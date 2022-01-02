package datastore

import (
	"context"
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb/table"
)

type DataStore interface {
	Repository() Repository
	PullRequest() PullRequest
}

type Repository interface {
	Create(ctx context.Context, ents ...*entity.Repository) error
	GetAll(ctx context.Context) ([]*entity.Repository, error)
}

type PullRequest interface {
	Create(ctx context.Context, ents ...*entity.PullRequest) error
}

type dataStore struct {
	repository  *table.Repository
	pullRequest *table.PullRequest
}

func NewDataStore(cfg rdb.Config) (DataStore, error) {
	db, err := rdb.NewRDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init dyanamodb. %w", err)
	}
	return &dataStore{
		repository:  &table.Repository{Client: db},
		pullRequest: &table.PullRequest{Client: db},
	}, nil
}

func (d *dataStore) Repository() Repository {
	return d.repository
}

func (d *dataStore) PullRequest() PullRequest {
	return d.pullRequest
}
