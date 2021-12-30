// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shinnosuke-K/github-dev-insight/ent/issue"
	"github.com/shinnosuke-K/github-dev-insight/ent/repository"
)

// IssueCreate is the builder for creating a Issue entity.
type IssueCreate struct {
	config
	mutation *IssueMutation
	hooks    []Hook
}

// SetRepositoryID sets the "repository_id" field.
func (ic *IssueCreate) SetRepositoryID(s string) *IssueCreate {
	ic.mutation.SetRepositoryID(s)
	return ic
}

// SetGithubID sets the "github_id" field.
func (ic *IssueCreate) SetGithubID(s string) *IssueCreate {
	ic.mutation.SetGithubID(s)
	return ic
}

// SetTitle sets the "title" field.
func (ic *IssueCreate) SetTitle(s string) *IssueCreate {
	ic.mutation.SetTitle(s)
	return ic
}

// SetCreatedAt sets the "created_at" field.
func (ic *IssueCreate) SetCreatedAt(t time.Time) *IssueCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *IssueCreate) SetNillableCreatedAt(t *time.Time) *IssueCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *IssueCreate) SetUpdatedAt(t time.Time) *IssueCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *IssueCreate) SetNillableUpdatedAt(t *time.Time) *IssueCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetLastEditedAt sets the "last_edited_at" field.
func (ic *IssueCreate) SetLastEditedAt(t time.Time) *IssueCreate {
	ic.mutation.SetLastEditedAt(t)
	return ic
}

// SetNillableLastEditedAt sets the "last_edited_at" field if the given value is not nil.
func (ic *IssueCreate) SetNillableLastEditedAt(t *time.Time) *IssueCreate {
	if t != nil {
		ic.SetLastEditedAt(*t)
	}
	return ic
}

// SetClosedAt sets the "closed_at" field.
func (ic *IssueCreate) SetClosedAt(t time.Time) *IssueCreate {
	ic.mutation.SetClosedAt(t)
	return ic
}

// SetNillableClosedAt sets the "closed_at" field if the given value is not nil.
func (ic *IssueCreate) SetNillableClosedAt(t *time.Time) *IssueCreate {
	if t != nil {
		ic.SetClosedAt(*t)
	}
	return ic
}

// SetRepositoryID sets the "repository" edge to the Repository entity by ID.
func (ic *IssueCreate) SetRepositoryID(id int) *IssueCreate {
	ic.mutation.SetRepositoryID(id)
	return ic
}

// SetNillableRepositoryID sets the "repository" edge to the Repository entity by ID if the given value is not nil.
func (ic *IssueCreate) SetNillableRepositoryID(id *int) *IssueCreate {
	if id != nil {
		ic = ic.SetRepositoryID(*id)
	}
	return ic
}

// SetRepository sets the "repository" edge to the Repository entity.
func (ic *IssueCreate) SetRepository(r *Repository) *IssueCreate {
	return ic.SetRepositoryID(r.ID)
}

// Mutation returns the IssueMutation object of the builder.
func (ic *IssueCreate) Mutation() *IssueMutation {
	return ic.mutation
}

// Save creates the Issue in the database.
func (ic *IssueCreate) Save(ctx context.Context) (*Issue, error) {
	var (
		err  error
		node *Issue
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IssueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IssueCreate) SaveX(ctx context.Context) *Issue {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IssueCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IssueCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *IssueCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := issue.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := issue.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.LastEditedAt(); !ok {
		v := issue.DefaultLastEditedAt()
		ic.mutation.SetLastEditedAt(v)
	}
	if _, ok := ic.mutation.ClosedAt(); !ok {
		v := issue.DefaultClosedAt()
		ic.mutation.SetClosedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IssueCreate) check() error {
	if _, ok := ic.mutation.RepositoryID(); !ok {
		return &ValidationError{Name: "repository_id", err: errors.New(`ent: missing required field "repository_id"`)}
	}
	if v, ok := ic.mutation.RepositoryID(); ok {
		if err := issue.RepositoryIDValidator(v); err != nil {
			return &ValidationError{Name: "repository_id", err: fmt.Errorf(`ent: validator failed for field "repository_id": %w`, err)}
		}
	}
	if _, ok := ic.mutation.GithubID(); !ok {
		return &ValidationError{Name: "github_id", err: errors.New(`ent: missing required field "github_id"`)}
	}
	if v, ok := ic.mutation.GithubID(); ok {
		if err := issue.GithubIDValidator(v); err != nil {
			return &ValidationError{Name: "github_id", err: fmt.Errorf(`ent: validator failed for field "github_id": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := ic.mutation.Title(); ok {
		if err := issue.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := ic.mutation.LastEditedAt(); !ok {
		return &ValidationError{Name: "last_edited_at", err: errors.New(`ent: missing required field "last_edited_at"`)}
	}
	if _, ok := ic.mutation.ClosedAt(); !ok {
		return &ValidationError{Name: "closed_at", err: errors.New(`ent: missing required field "closed_at"`)}
	}
	return nil
}

func (ic *IssueCreate) sqlSave(ctx context.Context) (*Issue, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ic *IssueCreate) createSpec() (*Issue, *sqlgraph.CreateSpec) {
	var (
		_node = &Issue{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: issue.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: issue.FieldID,
			},
		}
	)
	if value, ok := ic.mutation.RepositoryID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: issue.FieldRepositoryID,
		})
		_node.RepositoryID = value
	}
	if value, ok := ic.mutation.GithubID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: issue.FieldGithubID,
		})
		_node.GithubID = value
	}
	if value, ok := ic.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: issue.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: issue.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: issue.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.LastEditedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: issue.FieldLastEditedAt,
		})
		_node.LastEditedAt = value
	}
	if value, ok := ic.mutation.ClosedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: issue.FieldClosedAt,
		})
		_node.ClosedAt = value
	}
	if nodes := ic.mutation.RepositoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   issue.RepositoryTable,
			Columns: []string{issue.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.repository_issues = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IssueCreateBulk is the builder for creating many Issue entities in bulk.
type IssueCreateBulk struct {
	config
	builders []*IssueCreate
}

// Save creates the Issue entities in the database.
func (icb *IssueCreateBulk) Save(ctx context.Context) ([]*Issue, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Issue, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IssueMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *IssueCreateBulk) SaveX(ctx context.Context) []*Issue {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IssueCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IssueCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
