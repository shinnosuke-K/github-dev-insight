// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shinnosuke-K/github-dev-insight/ent/predicate"
	"github.com/shinnosuke-K/github-dev-insight/ent/repository"
)

// RepositoryDelete is the builder for deleting a Repository entity.
type RepositoryDelete struct {
	config
	hooks    []Hook
	mutation *RepositoryMutation
}

// Where appends a list predicates to the RepositoryDelete builder.
func (rd *RepositoryDelete) Where(ps ...predicate.Repository) *RepositoryDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *RepositoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, RepositoryMutation](ctx, rd.sqlExec, rd.mutation, rd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *RepositoryDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *RepositoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(repository.Table, sqlgraph.NewFieldSpec(repository.FieldID, field.TypeUUID))
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rd.mutation.done = true
	return affected, err
}

// RepositoryDeleteOne is the builder for deleting a single Repository entity.
type RepositoryDeleteOne struct {
	rd *RepositoryDelete
}

// Where appends a list predicates to the RepositoryDelete builder.
func (rdo *RepositoryDeleteOne) Where(ps ...predicate.Repository) *RepositoryDeleteOne {
	rdo.rd.mutation.Where(ps...)
	return rdo
}

// Exec executes the deletion query.
func (rdo *RepositoryDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{repository.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *RepositoryDeleteOne) ExecX(ctx context.Context) {
	if err := rdo.Exec(ctx); err != nil {
		panic(err)
	}
}
