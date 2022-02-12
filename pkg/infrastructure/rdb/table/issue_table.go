package table

import (
	"context"
	"github.com/google/uuid"
	"github.com/shinnosuke-K/github-dev-insight/ent"
	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb/client"
)

type Issue struct {
	Client *client.Client
}

func (t *Issue) Create(ctx context.Context, ents ...*entity.Issue) error {
	var (
		bulk = make([]*ent.IssueCreate, len(ents))
	)
	for i, e := range ents {
		bulk[i] = t.Client.DB().Issue.Create().
			SetRepositoryID(uuid.MustParse(string(e.RepositoryID))).
			SetGithubID(string(e.GitHubID)).
			SetTitle(e.Title).
			SetCreatedAt(e.CreatedAt).
			SetUpdatedAt(e.UpdatedAt).
			SetNillableLastEditedAt(e.LastEditedAt).
			SetNillableClosedAt(e.ClosedAt)
	}
	if _, err := t.Client.DB().Issue.CreateBulk(bulk...).Save(ctx); err != nil {
		return err
	}
	return nil
}
