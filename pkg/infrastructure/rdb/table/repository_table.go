package table

import (
	"context"

	"github.com/google/uuid"
	"github.com/shinnosuke-K/github-dev-insight/ent"
	"github.com/shinnosuke-K/github-dev-insight/ent/repository"
	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb/client"
)

type Repository struct {
	Client *client.Client
}

func (t *Repository) Create(ctx context.Context, ents ...*entity.Repository) error {
	var (
		bulk = make([]*ent.RepositoryCreate, len(ents))
	)
	for i, e := range ents {
		bulk[i] = t.Client.DB().Repository.Create().
			SetGithubID(string(e.GitHubID)).
			SetName(e.Name).
			SetOwner(e.Owner).
			SetDescription(e.Description).
			SetTotalIssue(e.TotalIssue).
			SetTotalPr(e.TotalPR).
			SetCreatedAt(e.CreatedAt).
			SetUpdatedAt(e.UpdatedAt).
			SetNillablePushedAt(e.PushedAt)
	}
	if _, err := t.Client.DB().Repository.CreateBulk(bulk...).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (t *Repository) Update(ctx context.Context, ent *entity.Repository) error {
	if _, err := t.Client.DB().Repository.UpdateOneID(uuid.MustParse(string(ent.ID))).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (t *Repository) GetAll(ctx context.Context) ([]*entity.Repository, error) {
	res := []*entity.Repository{}
	if err := t.Client.DB().Repository.Query().Select(repository.Columns...).Scan(ctx, &res); err != nil {
		return nil, err
	}
	return res, nil
}
