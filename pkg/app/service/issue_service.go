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

	issues, err := s.GitHub().GetIssuesByGitHubID(ctx, &params.GetIssuesByGitHubID{
		GitHubID:     repos[0].GitHubID,
		RepositoryID: repos[0].ID,
	})
	if err != nil {
		return fmt.Errorf("failed to get issue by id:[%s]. %w", repos[0].GitHubID, err)
	}
	fmt.Println(issues)
	return nil
}
