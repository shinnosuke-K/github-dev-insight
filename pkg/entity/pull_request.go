package entity

import (
	"time"
)

type PullRequestID string

type PullRequest struct {
	ID           PullRequestID
	RepositoryID RepositoryID
	GitHubID     GitHubID
	Title        string
	TotalCommit  int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ClosedAt     *time.Time
	MergedAt     *time.Time
}
