package rdb

import (
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/ent"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb/client"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Driver string
	User   string
	Pass   string
	Host   string
	Port   string
	Name   string
}

func (c *Config) dataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", c.User, c.Pass, c.Host, c.Port, c.Name)
}

func NewRDB(cfg Config) (*client.Client, error) {
	db, err := ent.Open(cfg.Driver, cfg.dataSourceName(), ent.Debug())
	if err != nil {
		return nil, fmt.Errorf("failed to open rdb. %w", err)
	}
	return &client.Client{Client: db}, nil
}
