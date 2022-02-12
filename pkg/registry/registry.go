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

func RegisterPullRequestService() (service.PullRequestService, error) {
	ad, err := adapter.NewAdapter()
	if err != nil {
		return nil, fmt.Errorf("faild to create adapter. %w", err)
	}
	return service.NewPullRequestService(ad), nil
}

func RegisterCommitService() (service.CommitService, error) {
	ad, err := adapter.NewAdapter()
	if err != nil {
		return nil, fmt.Errorf("faild to create adapter. %w", err)
	}
	return service.NewCommitService(ad), nil
}

func RegisterIssueService() (service.IssueService, error) {
	ad, err := adapter.NewAdapter()
	if err != nil {
		return nil, fmt.Errorf("faild to create adapter. %w", err)
	}
	return service.NewIssueService(ad), nil
}
