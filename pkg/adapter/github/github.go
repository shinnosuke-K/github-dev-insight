package github

import (
	"context"

	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
)

type GitHub interface {
	GetRepositories(ctx context.Context, params *GetRepositoriesParams) ([]*entity.Repository, error)
	GetPullRequestsByGitHubID(ctx context.Context, params *GetPullRequestsByGitHubIDParams) ([]*entity.PullRequest, error)
	GetCommitsByGitHubID(ctx context.Context, params *GetCommitsByGitHubIDParams) ([]*entity.Commit, error)
}

type GetRepositoriesParams struct {
	UserName string
}

type GetPullRequestsByGitHubIDParams struct {
	GitHubID     entity.GitHubID
	RepositoryID entity.RepositoryID
}

type GetCommitsByGitHubIDParams struct {
	GitHubID      entity.GitHubID
	PullRequestID entity.PullRequestID
}
