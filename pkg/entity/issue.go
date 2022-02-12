package entity

import "time"

type IssueID string

type Issue struct {
	ID           IssueID      `json:"id"`
	GitHubID     GitHubID     `json:"git_hub_id"`
	Title        string       `json:"title"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	LastEditedAt *time.Time   `json:"last_edited_at"`
	ClosedAt     *time.Time   `json:"closed_at"`
	RepositoryID RepositoryID `json:"repository_id"`
}
