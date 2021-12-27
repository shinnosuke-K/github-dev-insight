package registry

import (
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter"
	"github.com/shinnosuke-K/github-dev-insight/pkg/app/service"
	"github.com/shinnosuke-K/github-dev-insight/pkg/env"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/graphql"
)

func newAdapter() (*adapter.Adapter, error) {
	var (
		github = env.Get().GitHub
	)
	graph, err := graphql.NewGraphQL(graphql.Config{
		LoginUser: github.LoginUser,
		URL:       github.GraphQLURL,
		Token:     github.AccessToken,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create graphql. %w", err)
	}

	return &adapter.Adapter{
		GitHub: graph,
	}, nil
}

func RegisterRepositoryService() (service.RepositoryService, error) {
	ad, err := newAdapter()
	if err != nil {
		return nil, fmt.Errorf("faild to create adapter. %w", err)
	}
	return service.NewRepositoryService(ad), nil
}
