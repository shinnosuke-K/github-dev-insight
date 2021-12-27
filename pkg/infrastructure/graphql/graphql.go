package graphql

import (
	"context"
	"errors"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter"
	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
)

type Config struct {
	LoginUser string
	URL       string
	Token     string
}

func NewGraphQL(cfg Config) (adapter.GitHub, error) {
	switch {
	case cfg.LoginUser == "":
		return nil, errors.New("login_user is required")
	case cfg.URL == "":
		return nil, errors.New("url is required")
	case cfg.Token == "":
		return nil, errors.New("token is required")
	}
	return nil, nil
}

// GetRepositories は指定したユーザに紐づくリポジトリ一覧を取得する
func (c *Client) GetRepositories(ctx context.Context, params *adapter.GetRepositoriesParams) (entity.Repositories, error) {
	return nil, nil
}
