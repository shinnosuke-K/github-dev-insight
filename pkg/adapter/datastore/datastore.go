package datastore

import (
	"context"
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/aws/dynamo"
)

type DataStore interface {
	Repository() Repository
}

type Repository interface {
	Create(ctx context.Context, ent ...*entity.Repository) error
}

type dataStore struct {
	repository *dynamo.Repository
}

func NewDataStore(cfg dynamo.Config) (DataStore, error) {
	db, err := dynamo.NewDynamoDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init dyanamodb. %w", err)
	}
	return &dataStore{
		repository: &dynamo.Repository{Client: db},
	}, nil
}

func (d *dataStore) Repository() Repository {
	return d.repository
}
