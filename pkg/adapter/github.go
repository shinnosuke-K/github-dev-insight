package adapter

import (
	"context"

	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
)

type GitHub interface {
	GetRepositories(ctx context.Context, params *GetRepositoriesParams) (entity.Repositories, error)
}

type GetRepositoriesParams struct {
	UserName string
}
