package dynamo

import (
	"context"
	"fmt"

	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/aws/dynamo/client"
)

type Repository struct {
	*client.Client
}

func (d *Repository) Create(ctx context.Context, ents ...entity.Repository) error {
	if len(ents) == 0 {
		return nil
	}
	items, err := d.Client.CreateBatchItems(ents, entity.Repository{})
	if err != nil {
		return fmt.Errorf("failed to create batch items. %w", err)
	}

	if err := d.Client.BatchCreate(ctx, items); err != nil {
		return err
	}
	return nil
}
