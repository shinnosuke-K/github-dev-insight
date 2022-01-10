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
)

// CommitsCreate is the builder for creating a Commits entity.
type CommitsCreate struct {
	config
	mutation *CommitsMutation
	hooks    []Hook
}

// SetGithubID sets the "github_id" field.
func (cc *CommitsCreate) SetGithubID(s string) *CommitsCreate {
	cc.mutation.SetGithubID(s)
	return cc
}

// SetMessage sets the "message" field.
func (cc *CommitsCreate) SetMessage(s string) *CommitsCreate {
	cc.mutation.SetMessage(s)
	return cc
}

// SetAdditions sets the "additions" field.
func (cc *CommitsCreate) SetAdditions(i int64) *CommitsCreate {
	cc.mutation.SetAdditions(i)
	return cc
}

// SetDeletions sets the "deletions" field.
func (cc *CommitsCreate) SetDeletions(i int64) *CommitsCreate {
	cc.mutation.SetDeletions(i)
	return cc
}

// SetChangeFiles sets the "change_files" field.
func (cc *CommitsCreate) SetChangeFiles(i int64) *CommitsCreate {
	cc.mutation.SetChangeFiles(i)
	return cc
}

// SetCommittedAt sets the "committed_at" field.
func (cc *CommitsCreate) SetCommittedAt(t time.Time) *CommitsCreate {
	cc.mutation.SetCommittedAt(t)
	return cc
}

// SetNillableCommittedAt sets the "committed_at" field if the given value is not nil.
func (cc *CommitsCreate) SetNillableCommittedAt(t *time.Time) *CommitsCreate {
	if t != nil {
		cc.SetCommittedAt(*t)
	}
	return cc
}

// SetPushedAt sets the "pushed_at" field.
func (cc *CommitsCreate) SetPushedAt(t time.Time) *CommitsCreate {
	cc.mutation.SetPushedAt(t)
	return cc
}

// SetNillablePushedAt sets the "pushed_at" field if the given value is not nil.
func (cc *CommitsCreate) SetNillablePushedAt(t *time.Time) *CommitsCreate {
	if t != nil {
		cc.SetPushedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CommitsCreate) SetID(u uuid.UUID) *CommitsCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetPullRequestID sets the "pull_request" edge to the PullRequest entity by ID.
func (cc *CommitsCreate) SetPullRequestID(id uuid.UUID) *CommitsCreate {
	cc.mutation.SetPullRequestID(id)
	return cc
}

// SetNillablePullRequestID sets the "pull_request" edge to the PullRequest entity by ID if the given value is not nil.
func (cc *CommitsCreate) SetNillablePullRequestID(id *uuid.UUID) *CommitsCreate {
	if id != nil {
		cc = cc.SetPullRequestID(*id)
	}
	return cc
}

// SetPullRequest sets the "pull_request" edge to the PullRequest entity.
func (cc *CommitsCreate) SetPullRequest(p *PullRequest) *CommitsCreate {
	return cc.SetPullRequestID(p.ID)
}

// Mutation returns the CommitsMutation object of the builder.
func (cc *CommitsCreate) Mutation() *CommitsMutation {
	return cc.mutation
}

// Save creates the Commits in the database.
func (cc *CommitsCreate) Save(ctx context.Context) (*Commits, error) {
	var (
		err  error
		node *Commits
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommitsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommitsCreate) SaveX(ctx context.Context) *Commits {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommitsCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommitsCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CommitsCreate) defaults() {
	if _, ok := cc.mutation.CommittedAt(); !ok {
		v := commits.DefaultCommittedAt()
		cc.mutation.SetCommittedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := commits.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommitsCreate) check() error {
	if _, ok := cc.mutation.GithubID(); !ok {
		return &ValidationError{Name: "github_id", err: errors.New(`ent: missing required field "github_id"`)}
	}
	if v, ok := cc.mutation.GithubID(); ok {
		if err := commits.GithubIDValidator(v); err != nil {
			return &ValidationError{Name: "github_id", err: fmt.Errorf(`ent: validator failed for field "github_id": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "message"`)}
	}
	if _, ok := cc.mutation.Additions(); !ok {
		return &ValidationError{Name: "additions", err: errors.New(`ent: missing required field "additions"`)}
	}
	if _, ok := cc.mutation.Deletions(); !ok {
		return &ValidationError{Name: "deletions", err: errors.New(`ent: missing required field "deletions"`)}
	}
	if _, ok := cc.mutation.ChangeFiles(); !ok {
		return &ValidationError{Name: "change_files", err: errors.New(`ent: missing required field "change_files"`)}
	}
	if _, ok := cc.mutation.CommittedAt(); !ok {
		return &ValidationError{Name: "committed_at", err: errors.New(`ent: missing required field "committed_at"`)}
	}
	return nil
}

func (cc *CommitsCreate) sqlSave(ctx context.Context) (*Commits, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *CommitsCreate) createSpec() (*Commits, *sqlgraph.CreateSpec) {
	var (
		_node = &Commits{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: commits.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: commits.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.GithubID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: commits.FieldGithubID,
		})
		_node.GithubID = value
	}
	if value, ok := cc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: commits.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := cc.mutation.Additions(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commits.FieldAdditions,
		})
		_node.Additions = value
	}
	if value, ok := cc.mutation.Deletions(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commits.FieldDeletions,
		})
		_node.Deletions = value
	}
	if value, ok := cc.mutation.ChangeFiles(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: commits.FieldChangeFiles,
		})
		_node.ChangeFiles = value
	}
	if value, ok := cc.mutation.CommittedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: commits.FieldCommittedAt,
		})
		_node.CommittedAt = value
	}
	if value, ok := cc.mutation.PushedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: commits.FieldPushedAt,
		})
		_node.PushedAt = value
	}
	if nodes := cc.mutation.PullRequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   commits.PullRequestTable,
			Columns: []string{commits.PullRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pullrequest.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.pull_request_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommitsCreateBulk is the builder for creating many Commits entities in bulk.
type CommitsCreateBulk struct {
	config
	builders []*CommitsCreate
}

// Save creates the Commits entities in the database.
func (ccb *CommitsCreateBulk) Save(ctx context.Context) ([]*Commits, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Commits, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommitsMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommitsCreateBulk) SaveX(ctx context.Context) []*Commits {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommitsCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommitsCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
