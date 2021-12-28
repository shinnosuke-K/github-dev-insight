package dynamo

import (
	"context"

	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/aws/dynamo/client"
)

type Repository struct {
	*client.Client
}

func (d *Repository) Create(ctx context.Context, ents ...*entity.Repository) error {
	//items, err := d.CreateItems(ents)
	//if err != nil {
	//	return fmt.Errorf("failed to create item. %w", err)
	//}
	_ = d.Client.SetTable(entity.Repository{})
	return nil
}
