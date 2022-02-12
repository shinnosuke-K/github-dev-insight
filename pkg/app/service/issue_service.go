package service

import (
	"context"
	"fmt"
	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter"
	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/github/params"
	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
)

type IssueService interface {
	ImportIssue(ctx context.Context) error
}

type issueService struct {
	*adapter.Adapter
}

func NewIssueService(ad *adapter.Adapter) *issueService {
	return &issueService{ad}
}

func (s *issueService) ImportIssue(ctx context.Context) error {
	repos, err := s.DataStore().Repository().GetByTargetTypeAndStatus(ctx, entity.TargetTypeIssue, false)
	if err != nil {
		return fmt.Errorf("failed to get repository. %w", err)
	}
	if len(repos) == 0 {
		return nil
	}

	for _, repo := range repos {
		if err := s.DataStore().Transaction(ctx, func(ctx context.Context) error {
			if repo.TotalIssue != 0 {
				if err := s.store(ctx, repo); err != nil {
					return fmt.Errorf("failed to store. %w", err)
				}
			}
			if err := s.DataStore().Repository().UpdateStatusByID(ctx, repo.ID, entity.TargetTypeIssue, true); err != nil {
				return fmt.Errorf("failed to update repository by id:[%s]. %w", repo.ID, err)
			}
			return nil
		}); err != nil {
			return fmt.Errorf("error occurred in transaction. %w", err)
		}
		fmt.Printf("[info] repo:[%s] done\n", repo.Name)
	}
	return nil
}

func (s *issueService) store(ctx context.Context, ent *entity.Repository) error {
	issues, err := s.GitHub().GetIssuesByGitHubID(ctx, &params.GetIssuesByGitHubID{
		GitHubID:     ent.GitHubID,
		RepositoryID: ent.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to get issue by id:[%s]. %w", ent.GitHubID, err)
	}
	if err := s.DataStore().Issue().Create(ctx, issues...); err != nil {
		return fmt.Errorf("failed to create. %w", err)
	}
	return nil
}
