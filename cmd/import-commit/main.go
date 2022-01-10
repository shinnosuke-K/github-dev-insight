package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shinnosuke-K/github-dev-insight/pkg/env"
	"github.com/shinnosuke-K/github-dev-insight/pkg/registry"
)

func init() {
	if err := env.Load(); err != nil {
		panic(fmt.Sprintf("failed to load env config. %s", err))
	}
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func _main() error {
	commitSvc, err := registry.RegisterCommitService()
	if err != nil {
		return fmt.Errorf("failed to register service. %w", err)
	}
	if err := commitSvc.ImportCommit(context.Background()); err != nil {
		return fmt.Errorf("failed to import commits. %w", err)
	}
	return nil
}
