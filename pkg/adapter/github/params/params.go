package params

import "github.com/shinnosuke-K/github-dev-insight/pkg/entity"

type GetRepositories struct {
	UserName string
}

type GetPullRequestsByGitHubID struct {
	GitHubID     entity.GitHubID
	RepositoryID entity.RepositoryID
}

type GetCommitsByGitHubID struct {
	GitHubID      entity.GitHubID
	PullRequestID entity.PullRequestID
}

type GetIssuesByGitHubID struct {
	GitHubID     entity.GitHubID
	RepositoryID entity.RepositoryID
}
