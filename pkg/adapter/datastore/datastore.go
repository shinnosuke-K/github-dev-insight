package datastore

import (
	"context"
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb/client"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb/table"
)

type DataStore interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
	Repository() Repository
	PullRequest() PullRequest
}

type Repository interface {
	Create(ctx context.Context, ents ...*entity.Repository) error
	UpdateStatusByID(ctx context.Context, id entity.RepositoryID, targetType entity.TargetType, status bool) error
	GetAll(ctx context.Context) ([]*entity.Repository, error)
	GetByTargetTypeAndStatus(ctx context.Context, targetType entity.TargetType, status bool) ([]*entity.Repository, error)
}

type PullRequest interface {
	Create(ctx context.Context, ents ...*entity.PullRequest) error
	Update(ctx context.Context, ent *entity.PullRequest) error
}

type dataStore struct {
	db          *client.Client
	repository  *table.Repository
	pullRequest *table.PullRequest
}

func NewDataStore(cfg rdb.Config) (DataStore, error) {
	db, err := rdb.NewRDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init dyanamodb. %w", err)
	}
	return &dataStore{
		db:          db,
		repository:  &table.Repository{Client: db},
		pullRequest: &table.PullRequest{Client: db},
	}, nil
}

func (d *dataStore) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.Transaction(ctx, fn)
}

func (d *dataStore) Repository() Repository {
	return d.repository
}

func (d *dataStore) PullRequest() PullRequest {
	return d.pullRequest
}
