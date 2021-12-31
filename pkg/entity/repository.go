package entity

import "time"

type RepositoryID string

type Repository struct {
	ID          RepositoryID
	GitHubID    string
	Owner       string
	Name        string
	Description string
	TotalPR     int64
	TotalIssue  int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PushedAt    time.Time
}

type Repositories []*Repository
