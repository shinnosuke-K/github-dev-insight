package table

import (
	"context"

	"github.com/google/uuid"
	"github.com/shinnosuke-K/github-dev-insight/ent"
	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb/client"
)

type PullRequest struct {
	Client *client.Client
}

func (t *PullRequest) Create(ctx context.Context, ents ...*entity.PullRequest) error {
	var (
		bulk = make([]*ent.PullRequestCreate, len(ents))
	)
	for i, e := range ents {
		bulk[i] = t.Client.DB().PullRequest.Create().
			SetRepositoryID(uuid.MustParse(string(e.RepositoryID))).
			SetGithubID(string(e.GitHubID)).
			SetTitle(e.Title).
			SetTotalCommits(e.TotalCommit).
			SetCreatedAt(e.CreatedAt).
			SetUpdatedAt(e.UpdatedAt).
			SetNillableMergedAt(e.MergedAt).
			SetNillableClosedAt(e.ClosedAt)
	}
	if _, err := t.Client.DB().PullRequest.CreateBulk(bulk...).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (t *PullRequest) Update(ctx context.Context, ent *entity.PullRequest) error {
	if _, err := t.Client.DB().PullRequest.UpdateOneID(uuid.MustParse(string(ent.ID))).Save(ctx); err != nil {
		return err
	}
	return nil
}
