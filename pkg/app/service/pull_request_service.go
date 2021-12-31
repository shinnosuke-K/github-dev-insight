package service

import (
	"context"
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter"
)

type PullRequestService interface {
	ImportPullRequest(ctx context.Context) error
}

type pullRequestService struct {
	*adapter.Adapter
}

func NewPullRequestService(ad *adapter.Adapter) *pullRequestService {
	return &pullRequestService{
		ad,
	}
}

func (s *pullRequestService) ImportPullRequest(ctx context.Context) error {
	repos, err := s.DataStore().Repository().GetAll(ctx)
	if err != nil {
		return fmt.Errorf("failed to get repository. %w", err)
	}
	for _, r := range repos {
		fmt.Println(r)
	}
	return nil
}
