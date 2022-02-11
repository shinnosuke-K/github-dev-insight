package table

import (
	"context"

	"github.com/google/uuid"
	"github.com/shinnosuke-K/github-dev-insight/ent"
	"github.com/shinnosuke-K/github-dev-insight/pkg/entity"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/rdb/client"
)

type Commit struct {
	Client *client.Client
}

func (t *Commit) Create(ctx context.Context, ents ...*entity.Commit) error {
	var (
		bulk = make([]*ent.CommitsCreate, len(ents))
	)
	for i, e := range ents {
		bulk[i] = t.Client.DB().Commits.Create().
			SetPullRequestID(uuid.MustParse(string(e.PullRequestID))).
			SetGithubID(string(e.GitHubID)).
			SetMessage(e.Message).
			SetCommittedAt(e.CommittedAt).
			SetAdditions(e.Additions).
			SetDeletions(e.Deletions).
			SetChangeFiles(e.ChangeFiles).
			SetNillablePushedAt(e.PushedAt)
	}
	if _, err := t.Client.DB().Commits.CreateBulk(bulk...).Save(ctx); err != nil {
		return err
	}
	return nil
}
