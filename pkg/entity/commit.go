package entity

import "time"

type CommitID string

type Commit struct {
	ID            CommitID      `json:"id"`
	GitHubID      GitHubID      `json:"github_id"`
	Message       string        `json:"message"`
	Additions     int64         `json:"additions"`
	Deletions     int64         `json:"deletions"`
	ChangeFiles   int64         `json:"change_files"`
	CommittedAt   time.Time     `json:"committed_at"`
	PushedAt      *time.Time    `json:"pushed_at"`
	PullRequestID PullRequestID `json:"pull_request_id"`
}
