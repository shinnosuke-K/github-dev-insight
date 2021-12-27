package entity

import "time"

type RepositoryID string

type Repository struct {
	ID          RepositoryID
	GitHubID    string
	Owner       string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PushedAt    time.Time
}

type Repositories []*Repository
