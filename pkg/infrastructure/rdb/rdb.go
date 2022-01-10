package rdb

import (
	"database/sql"
	"fmt"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shinnosuke-K/github-dev-insight/ent"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb/client"
)

type Config struct {
	Driver      string
	User        string
	Pass        string
	Host        string
	Port        string
	Name        string
	MaxIdleConn int
	MaxOpenConn int
	MaxIdleTime time.Duration
	MaxLifeTime time.Duration
}

func (c *Config) dataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", c.User, c.Pass, c.Host, c.Port, c.Name)
}

func NewRDB(cfg Config) (*client.Client, error) {
	db, err := sql.Open(cfg.Driver, cfg.dataSourceName())
	if err != nil {
		return nil, fmt.Errorf("failed to open rdb. %w", err)
	}
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetConnMaxIdleTime(cfg.MaxIdleTime * time.Millisecond)
	db.SetConnMaxLifetime(cfg.MaxLifeTime * time.Millisecond)
	drv := entsql.OpenDB(cfg.Driver, db)
	return client.NewClient(ent.NewClient(ent.Driver(drv))), nil
}
