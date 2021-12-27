package main

import (
	"fmt"
	"os"

	"github.com/shinnosuke-K/github-dev-insight/pkg/env"
)

func init() {
	if err := env.Load(); err != nil {
		panic(fmt.Sprintf("failed to load env config. %s", err))
	}
}

func _main() error {
	return nil
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
