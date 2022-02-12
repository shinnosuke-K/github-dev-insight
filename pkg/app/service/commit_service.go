package service

import (
	"context"
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter"
	datastoreParams "github.com/shinnosuke-K/github-dev-insight/pkg/adapter/datastore/params"
	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/github/params"
	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
)

type CommitService interface {
	ImportCommit(ctx context.Context) error
}

type commitService struct {
	*adapter.Adapter
}

func NewCommitService(ad *adapter.Adapter) *commitService {
	return &commitService{
		ad,
	}
}

func (s *commitService) ImportCommit(ctx context.Context) error {
	for offset := 0; offset < 100000; offset += 100 {
		prs, err := s.DataStore().PullRequest().GetByStatusWithPaging(ctx, &datastoreParams.GetByStatusWithPaging{
			Status: false,
			Limit:  100,
			Offset: offset,
		})
		if err != nil {
			return fmt.Errorf("failed to get pull request. %w", err)
		}
		if len(prs) == 0 {
			break
		}
		if err := s.store(ctx, prs); err != nil {
			return fmt.Errorf("failed to store. %w", err)
		}
	}
	return nil
}

func (s *commitService) store(ctx context.Context, prs []*entity.PullRequest) error {
	for _, p := range prs {
		if p.TotalCommits == 0 {
			continue
		}
		p := p
		cm, err := s.GitHub().GetCommitsByGitHubID(ctx, &params.GetCommitsByGitHubID{
			GitHubID:      p.GitHubID,
			PullRequestID: p.ID,
		})
		if err != nil {
			return fmt.Errorf("failed to get commit by id:[%s]. %w", p.GitHubID, err)
		}
		if err := s.DataStore().Transaction(ctx, func(ctx context.Context) error {
			if err := s.DataStore().Commit().Create(ctx, cm...); err != nil {
				return fmt.Errorf("failed to create commit. %w", err)
			}
			if err := s.DataStore().PullRequest().UpdateStatusByID(ctx, p.ID, true); err != nil {
				return fmt.Errorf("failed to update pull request by id:[%s]. %w", p.ID, err)
			}
			return nil
		}); err != nil {
			return fmt.Errorf("error occurred in transaction. %w", err)
		}
		fmt.Printf("[info]: pr[%s] done\n", p.Title)
	}
	return nil
}
