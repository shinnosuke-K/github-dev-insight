package datastore

import (
	"context"
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/datastore/params"
	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb/client"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb/table"
)

type DataStore interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
	Repository() Repository
	PullRequest() PullRequest
	Commit() Commit
	Issue() Issue
}

type Repository interface {
	Create(ctx context.Context, ents ...*entity.Repository) error
	UpdateStatusByID(ctx context.Context, id entity.RepositoryID, targetType entity.TargetType, status bool) error
	GetAll(ctx context.Context) ([]*entity.Repository, error)
	GetByTargetTypeAndStatus(ctx context.Context, targetType entity.TargetType, status bool) ([]*entity.Repository, error)
}

type PullRequest interface {
	Create(ctx context.Context, ents ...*entity.PullRequest) error
	UpdateStatusByID(ctx context.Context, id entity.PullRequestID, status bool) error
	GetByStatusWithPaging(ctx context.Context, params *params.GetByStatusWithPaging) ([]*entity.PullRequest, error)
}

type Commit interface {
	Create(ctx context.Context, ents ...*entity.Commit) error
}

type Issue interface {
	Create(ctx context.Context, ents ...*entity.Issue) error
}

type dataStore struct {
	db          *client.Client
	repository  *table.Repository
	pullRequest *table.PullRequest
	commit      *table.Commit
	issue       *table.Issue
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
		commit:      &table.Commit{Client: db},
		issue:       &table.Issue{Client: db},
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

func (d *dataStore) Commit() Commit {
	return d.commit
}

func (d *dataStore) Issue() Issue {
	return d.issue
}
