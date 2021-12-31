package graphql

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/github"
	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
	"github.com/shinnosuke-K/github-dev-insight/pkg/util/id"
	"github.com/shinnosuke-K/github-dev-insight/pkg/util/strconvert"
)

type Config struct {
	LoginUser string
	URL       string
	Token     string
}

func NewGraphQL(cfg Config) (github.GitHub, error) {
	switch {
	case cfg.LoginUser == "":
		return nil, errors.New("login_user is required")
	case cfg.URL == "":
		return nil, errors.New("url is required")
	case cfg.Token == "":
		return nil, errors.New("token is required")
	}
	client := NewClient(http.DefaultClient, cfg.URL, func(req *http.Request) {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.Token))
		return
	})
	return client, nil
}

// GetRepositories は指定したユーザに紐づくリポジトリ一覧を取得する
func (c *Client) GetRepositories(ctx context.Context, params *github.GetRepositoriesParams) ([]entity.Repository, error) {
	var (
		res   []entity.Repository
		after *string
		total = 10
	)
	for i := 0; i < total; i++ {
		repos, err := c.Repositories(ctx, 100, params.UserName, after)
		if err != nil {
			return nil, fmt.Errorf("failed to get repositories. %w", err)
		}
		for _, n := range repos.Organization.Repositories.Nodes {
			res = append(res, entity.Repository{
				ID:          entity.RepositoryID(id.Generate()),
				GitHubID:    n.ID,
				Owner:       params.UserName,
				Name:        n.Name,
				TotalIssue:  n.Issues.TotalCount,
				TotalPR:     n.PullRequests.TotalCount,
				Description: strconvert.StringOrDefault(n.Description),
				CreatedAt:   strconvert.ToTime(n.CreatedAt),
				UpdatedAt:   strconvert.ToTime(n.UpdatedAt),
				PushedAt:    strconvert.ToTime(strconvert.StringOrDefault(n.PushedAt)),
			})
		}
		if !repos.Organization.Repositories.PageInfo.HasNextPage {
			break
		}
		after = repos.Organization.Repositories.PageInfo.EndCursor
	}
	return res, nil
}
