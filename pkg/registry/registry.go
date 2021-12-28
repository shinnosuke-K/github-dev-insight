package registry

import (
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter"
	"github.com/shinnosuke-K/github-dev-insight/pkg/app/service"
)

func RegisterRepositoryService() (service.RepositoryService, error) {
	ad, err := adapter.NewAdapter()
	if err != nil {
		return nil, fmt.Errorf("faild to create adapter. %w", err)
	}
	return service.NewRepositoryService(ad), nil
}
