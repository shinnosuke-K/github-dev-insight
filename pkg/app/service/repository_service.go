package service

import (
	"context"
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter"
	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/github"
	"github.com/shinnosuke-K/github-dev-insight/pkg/env"
)

type RepositoryService interface {
	ImportRepositories(ctx context.Context) error
}

type repositoryService struct {
	*adapter.Adapter
}

func NewRepositoryService(ad *adapter.Adapter) *repositoryService {
	return &repositoryService{
		ad,
	}
}

func (s *repositoryService) ImportRepositories(ctx context.Context) error {
	repos, err := s.GitHub().GetRepositories(ctx, &github.GetRepositoriesParams{UserName: env.Get().GitHub.LoginUser})
	if err != nil {
		return fmt.Errorf("failed to get repositories. %w", err)
	}
	if err := s.DataStore().Repository().Create(ctx, repos...); err != nil {
		return fmt.Errorf("failed to create repositories. %w", err)
	}
	return nil
}
