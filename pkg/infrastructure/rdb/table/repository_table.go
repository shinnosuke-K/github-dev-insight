package table

import (
	"context"

	"github.com/shinnosuke-K/github-dev-insight/ent"
	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb/client"
)

type Repository struct {
	Client *client.Client
}

func (t *Repository) Create(ctx context.Context, ents ...entity.Repository) error {
	var (
		repo = t.Client.DB().Repository.Create()
		bulk = make([]*ent.RepositoryCreate, len(ents))
	)
	for i, e := range ents {
		bulk[i] = repo.
			SetGithubID(e.GitHubID).
			SetName(e.Name).
			SetOwner(e.Owner).
			SetDescription(e.Description).
			SetTotalIssue(e.TotalIssue).
			SetTotalPr(e.TotalPR).
			SetCreatedAt(e.CreatedAt).
			SetUpdatedAt(e.UpdatedAt)
	}
	if _, err := t.Client.DB().Repository.CreateBulk(bulk...).Save(ctx); err != nil {
		return err
	}
	return nil
}
