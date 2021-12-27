package adapter

import (
	"context"
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/graphql"
)

type Adapter struct {
	gitHub GitHub
}

func CreateAdapter(ctx context.Context) (*Adapter, error) {

	graph, err := graphql.NewGraphQL(graphql.Config{
		LoginUser: "",
		URL:       "",
		Token:     "",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init graphql. %w", err)
	}
	return &Adapter{
		gitHub: graph,
	}, nil
}

func (a *Adapter) Github() GitHub {
	return a.gitHub
}
