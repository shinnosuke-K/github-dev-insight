package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/shinnosuke-K/github-dev-insight/ent/migrate"
	"github.com/shinnosuke-K/github-dev-insight/pkg/env"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb"
)

var (
	fDeploy = flag.Bool("deploy", false, "マイグレーションを行う")
)

func init() {
	if err := env.Load(); err != nil {
		panic(fmt.Sprintf("failed to load env config. %s", err))
	}
	flag.Parse()
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func _main() error {
	db, err := rdb.NewRDB(rdb.Config{
		Driver: env.Get().RDB.Driver,
		User:   env.Get().RDB.User,
		Pass:   env.Get().RDB.Password,
		Host:   env.Get().RDB.Host,
		Port:   env.Get().RDB.Port,
		Name:   env.Get().RDB.Name,
	})
	if err != nil {
		return fmt.Errorf("failed to init rdb. %w", err)
	}
	defer db.DB().Close()
	if *fDeploy {
		if err := db.DB().Schema.Create(context.Background(), migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
			return fmt.Errorf("failed to migrate. %w", err)
		}
	} else {
		if err := db.DB().Schema.WriteTo(context.Background(), os.Stdin, migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
			return fmt.Errorf("failed to dry run. %w", err)
		}
	}
	return nil
}
