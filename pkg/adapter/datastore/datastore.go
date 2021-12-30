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
}

type Repository interface {
	Create(ctx context.Context, ent ...entity.Repository) error
}

type dataStore struct {
	repository *table.Repository
}

func NewDataStore(cfg rdb.Config) (DataStore, error) {
	db, err := rdb.NewRDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init dyanamodb. %w", err)
	}
	return &dataStore{
		repository: &table.Repository{Client: db},
	}, nil
}

func (d *dataStore) Repository() Repository {
	return d.repository
}
