package client

import (
	"context"
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/ent"
)

type Client struct {
	tx *ent.Tx
	db *ent.Client
}

type Config struct {
	Driver string
	User   string
	Pass   string
	Host   string
	Port   string
	Name   string
}

func NewClient(db *ent.Client) *Client {
	return &Client{
		db: db,
	}
}

func (c *Client) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if err := c.begin(ctx); err != nil {
		return fmt.Errorf("failed to begin. %w", err)
	}
	if err := fn(ctx); err != nil {
		if rollBackErr := c.rollBack(); rollBackErr != nil {
			return fmt.Errorf("failed to rollback in fn. %w", rollBackErr)
		}
		return err
	}
	if err := c.commit(); err != nil {
		if rollBackErr := c.rollBack(); rollBackErr != nil {
			return fmt.Errorf("failed to rollback in commit. %w", rollBackErr)
		}
		return err
	}
	return nil
}

func (c *Client) begin(ctx context.Context) error {
	t, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction. %w", err)
	}
	c.tx = t
	return nil
}

func (c *Client) commit() error {
	if c.tx == nil {
		return nil
	}
	if err := c.tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit tx. %w", err)
	}
	c.tx = nil
	return nil
}

func (c *Client) rollBack() error {
	if c.tx == nil {
		return nil
	}
	if rollBackErr := c.tx.Rollback(); rollBackErr != nil {
		return fmt.Errorf("failed to roll back transaction. %w", rollBackErr)
	}
	return nil
}

func (c *Client) DB() *ent.Client {
	if c.tx != nil {
		return c.tx.Client()
	}
	return c.db
}
