package table

import (
	"context"

	"github.com/google/uuid"
	"github.com/shinnosuke-K/github-dev-insight/ent"
	"github.com/shinnosuke-K/github-dev-insight/ent/pullrequest"
	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter/datastore/params"
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
			SetTotalCommits(e.TotalCommits).
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

func (t *PullRequest) UpdateStatusByID(ctx context.Context, id entity.PullRequestID, status bool) error {
	if _, err := t.Client.DB().PullRequest.UpdateOneID(uuid.MustParse(string(id))).SetGetCommit(status).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (t *PullRequest) GetByStatusWithPaging(ctx context.Context, params *params.GetByStatusWithPaging) ([]*entity.PullRequest, error) {
	var (
		res     []*entity.PullRequest
		builder = t.Client.DB().PullRequest.Query().Where(pullrequest.GetCommit(params.Status)).Order(ent.Asc(pullrequest.FieldTotalCommits))
	)
	if err := builder.Offset(params.Offset).Limit(params.Limit).Select(pullrequest.Columns...).Scan(ctx, &res); err != nil {
		return nil, err
	}
	return res, nil
}
