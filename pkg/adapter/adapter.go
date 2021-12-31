package adapter

import (
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/datastore"
	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/github"
	"github.com/shinnosuke-K/github-dev-insight/pkg/env"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/graphql"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb"
)

type Adapter struct {
	gitHub    github.GitHub
	dataStore datastore.DataStore
}

func NewAdapter() (*Adapter, error) {
	var (
		_github = env.Get().GitHub
		_rdb    = env.Get().RDB
	)
	graph, err := graphql.NewGraphQL(graphql.Config{
		LoginUser: _github.LoginUser,
		URL:       _github.GraphQLURL,
		Token:     _github.AccessToken,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create graphql. %w", err)
	}

	d, err := datastore.NewDataStore(rdb.Config{
		Driver: _rdb.Driver,
		User:   _rdb.User,
		Pass:   _rdb.Password,
		Host:   _rdb.Host,
		Port:   _rdb.Port,
		Name:   _rdb.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init datastore. %w", err)
	}

	return &Adapter{
		gitHub:    graph,
		dataStore: d,
	}, nil
}

func (a *Adapter) GitHub() github.GitHub {
	return a.gitHub
}

func (a *Adapter) DataStore() datastore.DataStore {
	return a.dataStore
}
