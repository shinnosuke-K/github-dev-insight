// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/shinnosuke-K/github-dev-insight/ent/commits"
	"github.com/shinnosuke-K/github-dev-insight/ent/pullrequest"
	"github.com/shinnosuke-K/github-dev-insight/ent/repository"
)

// PullRequestCreate is the builder for creating a PullRequest entity.
type PullRequestCreate struct {
	config
	mutation *PullRequestMutation
	hooks    []Hook
}

// SetGithubID sets the "github_id" field.
func (prc *PullRequestCreate) SetGithubID(s string) *PullRequestCreate {
	prc.mutation.SetGithubID(s)
	return prc
}

// SetTitle sets the "title" field.
func (prc *PullRequestCreate) SetTitle(s string) *PullRequestCreate {
	prc.mutation.SetTitle(s)
	return prc
}

// SetTotalCommits sets the "total_commits" field.
func (prc *PullRequestCreate) SetTotalCommits(i int64) *PullRequestCreate {
	prc.mutation.SetTotalCommits(i)
	return prc
}

// SetNillableTotalCommits sets the "total_commits" field if the given value is not nil.
func (prc *PullRequestCreate) SetNillableTotalCommits(i *int64) *PullRequestCreate {
	if i != nil {
		prc.SetTotalCommits(*i)
	}
	return prc
}

// SetGetCommit sets the "get_commit" field.
func (prc *PullRequestCreate) SetGetCommit(b bool) *PullRequestCreate {
	prc.mutation.SetGetCommit(b)
	return prc
}

// SetNillableGetCommit sets the "get_commit" field if the given value is not nil.
func (prc *PullRequestCreate) SetNillableGetCommit(b *bool) *PullRequestCreate {
	if b != nil {
		prc.SetGetCommit(*b)
	}
	return prc
}

// SetCreatedAt sets the "created_at" field.
func (prc *PullRequestCreate) SetCreatedAt(t time.Time) *PullRequestCreate {
	prc.mutation.SetCreatedAt(t)
	return prc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (prc *PullRequestCreate) SetNillableCreatedAt(t *time.Time) *PullRequestCreate {
	if t != nil {
		prc.SetCreatedAt(*t)
	}
	return prc
}

// SetUpdatedAt sets the "updated_at" field.
func (prc *PullRequestCreate) SetUpdatedAt(t time.Time) *PullRequestCreate {
	prc.mutation.SetUpdatedAt(t)
	return prc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (prc *PullRequestCreate) SetNillableUpdatedAt(t *time.Time) *PullRequestCreate {
	if t != nil {
		prc.SetUpdatedAt(*t)
	}
	return prc
}

// SetClosedAt sets the "closed_at" field.
func (prc *PullRequestCreate) SetClosedAt(t time.Time) *PullRequestCreate {
	prc.mutation.SetClosedAt(t)
	return prc
}

// SetNillableClosedAt sets the "closed_at" field if the given value is not nil.
func (prc *PullRequestCreate) SetNillableClosedAt(t *time.Time) *PullRequestCreate {
	if t != nil {
		prc.SetClosedAt(*t)
	}
	return prc
}

// SetMergedAt sets the "merged_at" field.
func (prc *PullRequestCreate) SetMergedAt(t time.Time) *PullRequestCreate {
	prc.mutation.SetMergedAt(t)
	return prc
}

// SetNillableMergedAt sets the "merged_at" field if the given value is not nil.
func (prc *PullRequestCreate) SetNillableMergedAt(t *time.Time) *PullRequestCreate {
	if t != nil {
		prc.SetMergedAt(*t)
	}
	return prc
}

// SetID sets the "id" field.
func (prc *PullRequestCreate) SetID(u uuid.UUID) *PullRequestCreate {
	prc.mutation.SetID(u)
	return prc
}

// AddCommitIDs adds the "commits" edge to the Commits entity by IDs.
func (prc *PullRequestCreate) AddCommitIDs(ids ...uuid.UUID) *PullRequestCreate {
	prc.mutation.AddCommitIDs(ids...)
	return prc
}

// AddCommits adds the "commits" edges to the Commits entity.
func (prc *PullRequestCreate) AddCommits(c ...*Commits) *PullRequestCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return prc.AddCommitIDs(ids...)
}

// SetRepositoryID sets the "repository" edge to the Repository entity by ID.
func (prc *PullRequestCreate) SetRepositoryID(id uuid.UUID) *PullRequestCreate {
	prc.mutation.SetRepositoryID(id)
	return prc
}

// SetNillableRepositoryID sets the "repository" edge to the Repository entity by ID if the given value is not nil.
func (prc *PullRequestCreate) SetNillableRepositoryID(id *uuid.UUID) *PullRequestCreate {
	if id != nil {
		prc = prc.SetRepositoryID(*id)
	}
	return prc
}

// SetRepository sets the "repository" edge to the Repository entity.
func (prc *PullRequestCreate) SetRepository(r *Repository) *PullRequestCreate {
	return prc.SetRepositoryID(r.ID)
}

// Mutation returns the PullRequestMutation object of the builder.
func (prc *PullRequestCreate) Mutation() *PullRequestMutation {
	return prc.mutation
}

// Save creates the PullRequest in the database.
func (prc *PullRequestCreate) Save(ctx context.Context) (*PullRequest, error) {
	var (
		err  error
		node *PullRequest
	)
	prc.defaults()
	if len(prc.hooks) == 0 {
		if err = prc.check(); err != nil {
			return nil, err
		}
		node, err = prc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PullRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = prc.check(); err != nil {
				return nil, err
			}
			prc.mutation = mutation
			if node, err = prc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(prc.hooks) - 1; i >= 0; i-- {
			if prc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = prc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, prc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (prc *PullRequestCreate) SaveX(ctx context.Context) *PullRequest {
	v, err := prc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (prc *PullRequestCreate) Exec(ctx context.Context) error {
	_, err := prc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prc *PullRequestCreate) ExecX(ctx context.Context) {
	if err := prc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (prc *PullRequestCreate) defaults() {
	if _, ok := prc.mutation.TotalCommits(); !ok {
		v := pullrequest.DefaultTotalCommits
		prc.mutation.SetTotalCommits(v)
	}
	if _, ok := prc.mutation.GetCommit(); !ok {
		v := pullrequest.DefaultGetCommit
		prc.mutation.SetGetCommit(v)
	}
	if _, ok := prc.mutation.CreatedAt(); !ok {
		v := pullrequest.DefaultCreatedAt()
		prc.mutation.SetCreatedAt(v)
	}
	if _, ok := prc.mutation.UpdatedAt(); !ok {
		v := pullrequest.DefaultUpdatedAt()
		prc.mutation.SetUpdatedAt(v)
	}
	if _, ok := prc.mutation.ClosedAt(); !ok {
		v := pullrequest.DefaultClosedAt()
		prc.mutation.SetClosedAt(v)
	}
	if _, ok := prc.mutation.MergedAt(); !ok {
		v := pullrequest.DefaultMergedAt()
		prc.mutation.SetMergedAt(v)
	}
	if _, ok := prc.mutation.ID(); !ok {
		v := pullrequest.DefaultID()
		prc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (prc *PullRequestCreate) check() error {
	if _, ok := prc.mutation.GithubID(); !ok {
		return &ValidationError{Name: "github_id", err: errors.New(`ent: missing required field "github_id"`)}
	}
	if v, ok := prc.mutation.GithubID(); ok {
		if err := pullrequest.GithubIDValidator(v); err != nil {
			return &ValidationError{Name: "github_id", err: fmt.Errorf(`ent: validator failed for field "github_id": %w`, err)}
		}
	}
	if _, ok := prc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := prc.mutation.Title(); ok {
		if err := pullrequest.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if _, ok := prc.mutation.TotalCommits(); !ok {
		return &ValidationError{Name: "total_commits", err: errors.New(`ent: missing required field "total_commits"`)}
	}
	if v, ok := prc.mutation.TotalCommits(); ok {
		if err := pullrequest.TotalCommitsValidator(v); err != nil {
			return &ValidationError{Name: "total_commits", err: fmt.Errorf(`ent: validator failed for field "total_commits": %w`, err)}
		}
	}
	if _, ok := prc.mutation.GetCommit(); !ok {
		return &ValidationError{Name: "get_commit", err: errors.New(`ent: missing required field "get_commit"`)}
	}
	if _, ok := prc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := prc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (prc *PullRequestCreate) sqlSave(ctx context.Context) (*PullRequest, error) {
	_node, _spec := prc.createSpec()
	if err := sqlgraph.CreateNode(ctx, prc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (prc *PullRequestCreate) createSpec() (*PullRequest, *sqlgraph.CreateSpec) {
	var (
		_node = &PullRequest{config: prc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pullrequest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pullrequest.FieldID,
			},
		}
	)
	if id, ok := prc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := prc.mutation.GithubID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pullrequest.FieldGithubID,
		})
		_node.GithubID = value
	}
	if value, ok := prc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pullrequest.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := prc.mutation.TotalCommits(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: pullrequest.FieldTotalCommits,
		})
		_node.TotalCommits = value
	}
	if value, ok := prc.mutation.GetCommit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: pullrequest.FieldGetCommit,
		})
		_node.GetCommit = value
	}
	if value, ok := prc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pullrequest.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := prc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pullrequest.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := prc.mutation.ClosedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pullrequest.FieldClosedAt,
		})
		_node.ClosedAt = value
	}
	if value, ok := prc.mutation.MergedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pullrequest.FieldMergedAt,
		})
		_node.MergedAt = value
	}
	if nodes := prc.mutation.CommitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pullrequest.CommitsTable,
			Columns: []string{pullrequest.CommitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: commits.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := prc.mutation.RepositoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pullrequest.RepositoryTable,
			Columns: []string{pullrequest.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.repository_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PullRequestCreateBulk is the builder for creating many PullRequest entities in bulk.
type PullRequestCreateBulk struct {
	config
	builders []*PullRequestCreate
}

// Save creates the PullRequest entities in the database.
func (prcb *PullRequestCreateBulk) Save(ctx context.Context) ([]*PullRequest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(prcb.builders))
	nodes := make([]*PullRequest, len(prcb.builders))
	mutators := make([]Mutator, len(prcb.builders))
	for i := range prcb.builders {
		func(i int, root context.Context) {
			builder := prcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PullRequestMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, prcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, prcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, prcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (prcb *PullRequestCreateBulk) SaveX(ctx context.Context) []*PullRequest {
	v, err := prcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (prcb *PullRequestCreateBulk) Exec(ctx context.Context) error {
	_, err := prcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prcb *PullRequestCreateBulk) ExecX(ctx context.Context) {
	if err := prcb.Exec(ctx); err != nil {
		panic(err)
	}
}
