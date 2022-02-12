package github

import (
	"context"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/github/params"
	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
)

type GitHub interface {
	GetRepositories(ctx context.Context, params *params.GetRepositories) ([]*entity.Repository, error)
	GetPullRequestsByGitHubID(ctx context.Context, params *params.GetPullRequestsByGitHubID) ([]*entity.PullRequest, error)
	GetCommitsByGitHubID(ctx context.Context, params *params.GetCommitsByGitHubID) ([]*entity.Commit, error)
	GetIssuesByGitHubID(ctx context.Context, params *params.GetIssuesByGitHubID) ([]*entity.Issue, error)
}
