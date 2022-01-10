package entity

import (
	"time"
)

type RepositoryID string

type GitHubID string

type Repository struct {
	ID             RepositoryID `json:"id"`
	GitHubID       GitHubID     `json:"github_id"`
	Owner          string       `json:"owner"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	TotalPR        int64        `json:"total_pr"`
	TotalIssue     int64        `json:"total_issue"`
	GetPullRequest bool         `json:"get_pull_request"`
	GetIssue       bool         `json:"get_issue"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	PushedAt       *time.Time   `json:"pushed_at"`
}

type Repositories []*Repository
