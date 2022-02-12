package graphql

import (
	"context"
	"errors"
	"fmt"
	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/github"
	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/github/params"
	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
	"github.com/shinnosuke-K/github-dev-insight/pkg/util/strconvert"
	"net/http"
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
func (c *Client) GetRepositories(ctx context.Context, params *params.GetRepositories) ([]*entity.Repository, error) {
	var (
		res   []*entity.Repository
		after *string
		total = 100
	)
	for i := 0; i < total; i++ {
		repos, err := c.Repositories(ctx, 100, params.UserName, after)
		if err != nil {
			return nil, fmt.Errorf("failed to get repositories. %w", err)
		}
		for _, n := range repos.Organization.Repositories.Nodes {
			res = append(res, &entity.Repository{
				GitHubID:    entity.GitHubID(n.ID),
				Owner:       params.UserName,
				Name:        n.Name,
				TotalIssue:  n.Issues.TotalCount,
				TotalPR:     n.PullRequests.TotalCount,
				Description: strconvert.StringOrDefault(n.Description),
				CreatedAt:   strconvert.ToTime(n.CreatedAt),
				UpdatedAt:   strconvert.ToTime(n.UpdatedAt),
				PushedAt:    strconvert.ToTimePtr(strconvert.StringOrDefault(n.PushedAt)),
			})
		}
		if !repos.Organization.Repositories.PageInfo.HasNextPage {
			break
		}
		after = repos.Organization.Repositories.PageInfo.EndCursor
	}
	return res, nil
}

func (c *Client) GetPullRequestsByGitHubID(ctx context.Context, params *params.GetPullRequestsByGitHubID) ([]*entity.PullRequest, error) {
	var (
		res   []*entity.PullRequest
		after *string
		total = 100
	)
	for i := 0; i < total; i++ {
		prs, err := c.PullRequests(ctx, 100, string(params.GitHubID), after)
		if err != nil {
			return nil, fmt.Errorf("failed to get pull requests. %w", err)
		}
		for _, n := range prs.Node.PullRequests.Nodes {
			res = append(res, &entity.PullRequest{
				GitHubID:     entity.GitHubID(n.ID),
				RepositoryID: params.RepositoryID,
				Title:        n.Title,
				TotalCommits: n.Commits.TotalCount,
				CreatedAt:    strconvert.ToTime(n.CreatedAt),
				UpdatedAt:    strconvert.ToTime(n.UpdatedAt),
				ClosedAt:     strconvert.ToTimePtr(strconvert.StringOrDefault(n.ClosedAt)),
				MergedAt:     strconvert.ToTimePtr(strconvert.StringOrDefault(n.MergedAt)),
			})
		}
		if !prs.Node.PullRequests.PageInfo.HasNextPage {
			break
		}
		after = prs.Node.PullRequests.PageInfo.EndCursor
	}
	return res, nil
}

func (c *Client) GetCommitsByGitHubID(ctx context.Context, params *params.GetCommitsByGitHubID) ([]*entity.Commit, error) {
	var (
		res   []*entity.Commit
		after *string
		total = 100
	)
	for i := 0; i < total; i++ {
		commits, err := c.Commits(ctx, 100, string(params.GitHubID), after)
		if err != nil {
			return nil, fmt.Errorf("failed to get commits. %w", err)
		}
		for _, n := range commits.Node.Commits.Nodes {
			res = append(res, &entity.Commit{
				GitHubID:      entity.GitHubID(n.Commit.ID),
				Message:       n.Commit.Message,
				Additions:     n.Commit.Additions,
				Deletions:     n.Commit.Deletions,
				ChangeFiles:   n.Commit.ChangedFiles,
				CommittedAt:   strconvert.ToTime(n.Commit.CommittedDate),
				PushedAt:      strconvert.ToTimePtr(strconvert.StringOrDefault(n.Commit.PushedDate)),
				PullRequestID: params.PullRequestID,
			})
		}
		if !commits.Node.Commits.PageInfo.HasNextPage {
			break
		}
		after = commits.Node.Commits.PageInfo.EndCursor
	}
	return res, nil
}

func (c *Client) GetIssuesByGitHubID(ctx context.Context, params *params.GetIssuesByGitHubID) ([]*entity.Issue, error) {
	var (
		res   []*entity.Issue
		after *string
		total = 100
	)
	for i := 0; i < total; i++ {
		issues, err := c.Issues(ctx, 100, string(params.GitHubID), after)
		if err != nil {
			return nil, fmt.Errorf("failed to get issues. %w", err)
		}
		for _, n := range issues.Node.Issues.Nodes {
			res = append(res, &entity.Issue{
				GitHubID:     entity.GitHubID(n.ID),
				Title:        n.Title,
				CreatedAt:    strconvert.ToTime(n.CreatedAt),
				UpdatedAt:    strconvert.ToTime(n.UpdatedAt),
				LastEditedAt: strconvert.ToTime(strconvert.StringOrDefault(n.LastEditedAt)),
				ClosedAt:     strconvert.ToTime(strconvert.StringOrDefault(n.ClosedAt)),
				RepositoryID: params.RepositoryID,
			})
		}
	}
	return res, nil
}
