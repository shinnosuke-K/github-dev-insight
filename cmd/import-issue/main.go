package main

import (
	"context"
	"fmt"
	"github.com/shinnosuke-K/github-dev-insight/pkg/env"
	"github.com/shinnosuke-K/github-dev-insight/pkg/registry"
	"os"
)

func init() {
	if err := env.Load(); err != nil {
		panic(fmt.Sprintf("failed to load env confing. %s", err))
	}
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func _main() error {
	issueSvc, err := registry.RegisterIssueService()
	if err != nil {
		return fmt.Errorf("failed to register service. %w", err)
	}
	if err := issueSvc.ImportIssue(context.Background()); err != nil {
		return fmt.Errorf("failed to import isseu. %w", err)
	}
	return nil
}
