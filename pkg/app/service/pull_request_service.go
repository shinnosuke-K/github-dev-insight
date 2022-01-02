package service

import (
	"context"
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter"
	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/github"
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
		prs, err := s.GitHub().GetPullRequestsByGitHubID(ctx, &github.GetPullRequestsByGitHubIDParams{
			GitHubID:     r.GitHubID,
			RepositoryID: r.ID,
		})
		if err != nil {
			return fmt.Errorf("failed to get prs by id:[%s]. %w", r.GitHubID, err)
		}
		if err := s.DataStore().PullRequest().Create(ctx, prs...); err != nil {
			return fmt.Errorf("failed to create prs. %w", err)
		}
	}
	return nil
}
