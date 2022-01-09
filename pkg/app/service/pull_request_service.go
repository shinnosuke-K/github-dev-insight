package service

import (
	"context"
	"fmt"
	"time"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter"
	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/github"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
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

	sem := semaphore.NewWeighted(10)
	eg, ctx := errgroup.WithContext(ctx)

	if err := s.DataStore().Transaction(ctx, func(ctx context.Context) error {
		for _, r := range repos {
			if err := sem.Acquire(ctx, 1); err != nil {
				fmt.Printf("failed to acquire semaphore: %v\n", err)
				break
			}
			r := r
			eg.Go(func() error {
				defer sem.Release(1)
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
				time.Sleep(500 * time.Millisecond)
				return nil
			})
		}
		if err := eg.Wait(); err != nil {
			return fmt.Errorf("error occurred in groutines. %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("error occurred in transaction. %w", err)
	}
	return nil
}
