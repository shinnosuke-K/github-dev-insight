package service

import (
	"context"
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter"
	datastoreParams "github.com/shinnosuke-K/github-dev-insight/pkg/adapter/datastore/params"
	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/github/params"
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
	offset := 0
	prs, err := s.DataStore().PullRequest().GetByStatusWithPaging(ctx, &datastoreParams.GetByStatusWithPaging{
		Status: false,
		Limit:  100,
		Offset: offset,
	})
	if err != nil {
		return fmt.Errorf("failed to get pull request. %w", err)
	}
	if len(prs) == 0 {
		return nil
	}

	for _, p := range prs {
		if p.TotalCommits == 0 {
			continue
		}
		cm, err := s.GitHub().GetCommitsByGitHubID(ctx, &params.GetCommitsByGitHubID{
			GitHubID:      p.GitHubID,
			PullRequestID: p.ID,
		})
		if err != nil {
			return fmt.Errorf("failed to get commit by id:[%s]. %w", prs[0].GitHubID, err)
		}
		fmt.Println(cm)
	}
	return nil
}
