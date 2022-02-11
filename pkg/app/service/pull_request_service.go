package service

import (
	"context"
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter"
	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/github/params"
	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
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
	repos, err := s.DataStore().Repository().GetByTargetTypeAndStatus(ctx, entity.TargetTypePullRequest, false)
	if err != nil {
		return fmt.Errorf("failed to get repository. %w", err)
	}

	if len(repos) == 0 {
		return nil
	}

	sem := semaphore.NewWeighted(5)
	eg, ctx := errgroup.WithContext(ctx)

	for _, r := range repos {
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Printf("failed to acquire semaphore: %v\n", err)
			break
		}

		ad, err := adapter.NewAdapter()
		if err != nil {
			return fmt.Errorf("failed to init adapter. %w", err)
		}
		r := r
		eg.Go(func() error {
			defer sem.Release(1)
			if err := ad.DataStore().Transaction(ctx, func(ctx context.Context) error {
				prs, err := ad.GitHub().GetPullRequestsByGitHubID(ctx, &params.GetPullRequestsByGitHubID{
					GitHubID:     r.GitHubID,
					RepositoryID: r.ID,
				})
				if err != nil {
					return fmt.Errorf("failed to get prs by id:[%s]. %w", r.GitHubID, err)
				}
				if err := ad.DataStore().PullRequest().Create(ctx, prs...); err != nil {
					return fmt.Errorf("failed to create prs. %w", err)
				}
				if err := ad.DataStore().Repository().UpdateStatusByID(ctx, r.ID, entity.TargetTypePullRequest, true); err != nil {
					return fmt.Errorf("failed to update repository by id:[%s]. %w", r.ID, err)
				}
				return nil
			}); err != nil {
				return fmt.Errorf("error occurred in transaction. %w", err)
			}
			fmt.Printf("[info] repo:[%s] done\n", r.Name)
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return fmt.Errorf("error occurred in goroutine. %w", err)
	}
	return nil
}
