package main

import (
	"fmt"
	"os"

	"github.com/shinnosuke-K/github-dev-insight/pkg/env"
)

func _main() error {
	if err := env.Load(); err != nil {
		return fmt.Errorf("failed to load env config. %w", err)
	}
	return nil
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
