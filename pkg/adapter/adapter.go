package adapter

import (
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/datastore"
	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/github"
	"github.com/shinnosuke-K/github-dev-insight/pkg/env"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/aws/dynamo"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/graphql"
)

type Adapter struct {
	gitHub    github.GitHub
	dataStore datastore.DataStore
}

func NewAdapter() (*Adapter, error) {
	var (
		_github = env.Get().GitHub
		_aws    = env.Get().AWS
	)
	graph, err := graphql.NewGraphQL(graphql.Config{
		LoginUser: _github.LoginUser,
		URL:       _github.GraphQLURL,
		Token:     _github.AccessToken,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create graphql. %w", err)
	}

	d, err := datastore.NewDataStore(dynamo.Config{
		SecretAccessKey: _aws.SecretAccessKey,
		AccessToken:     _aws.AccessToken,
		Region:          _aws.Region,
		EndPoint:        _aws.EndPoint,
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
