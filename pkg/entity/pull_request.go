package entity

import (
	"time"
)

type PullRequestID string

type PullRequest struct {
	ID           PullRequestID `json:"id"`
	RepositoryID RepositoryID  `json:"repository_id"`
	GitHubID     GitHubID      `json:"github_id"`
	Title        string        `json:"title"`
	TotalCommits int64         `json:"total_commits"`
	GetCommit    bool          `json:"get_commit"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	ClosedAt     *time.Time    `json:"closed_at"`
	MergedAt     *time.Time    `json:"merged_at"`
}
