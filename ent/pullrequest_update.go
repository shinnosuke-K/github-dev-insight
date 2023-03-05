// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/shinnosuke-K/github-dev-insight/ent/commits"
	"github.com/shinnosuke-K/github-dev-insight/ent/predicate"
	"github.com/shinnosuke-K/github-dev-insight/ent/pullrequest"
	"github.com/shinnosuke-K/github-dev-insight/ent/repository"
)

// PullRequestUpdate is the builder for updating PullRequest entities.
type PullRequestUpdate struct {
	config
	hooks    []Hook
	mutation *PullRequestMutation
}

// Where appends a list predicates to the PullRequestUpdate builder.
func (pru *PullRequestUpdate) Where(ps ...predicate.PullRequest) *PullRequestUpdate {
	pru.mutation.Where(ps...)
	return pru
}

// SetGithubID sets the "github_id" field.
func (pru *PullRequestUpdate) SetGithubID(s string) *PullRequestUpdate {
	pru.mutation.SetGithubID(s)
	return pru
}

// SetTitle sets the "title" field.
func (pru *PullRequestUpdate) SetTitle(s string) *PullRequestUpdate {
	pru.mutation.SetTitle(s)
	return pru
}

// SetTotalCommits sets the "total_commits" field.
func (pru *PullRequestUpdate) SetTotalCommits(i int64) *PullRequestUpdate {
	pru.mutation.ResetTotalCommits()
	pru.mutation.SetTotalCommits(i)
	return pru
}

// SetNillableTotalCommits sets the "total_commits" field if the given value is not nil.
func (pru *PullRequestUpdate) SetNillableTotalCommits(i *int64) *PullRequestUpdate {
	if i != nil {
		pru.SetTotalCommits(*i)
	}
	return pru
}

// AddTotalCommits adds i to the "total_commits" field.
func (pru *PullRequestUpdate) AddTotalCommits(i int64) *PullRequestUpdate {
	pru.mutation.AddTotalCommits(i)
	return pru
}

// SetGetCommit sets the "get_commit" field.
func (pru *PullRequestUpdate) SetGetCommit(b bool) *PullRequestUpdate {
	pru.mutation.SetGetCommit(b)
	return pru
}

// SetNillableGetCommit sets the "get_commit" field if the given value is not nil.
func (pru *PullRequestUpdate) SetNillableGetCommit(b *bool) *PullRequestUpdate {
	if b != nil {
		pru.SetGetCommit(*b)
	}
	return pru
}

// SetCreatedAt sets the "created_at" field.
func (pru *PullRequestUpdate) SetCreatedAt(t time.Time) *PullRequestUpdate {
	pru.mutation.SetCreatedAt(t)
	return pru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pru *PullRequestUpdate) SetNillableCreatedAt(t *time.Time) *PullRequestUpdate {
	if t != nil {
		pru.SetCreatedAt(*t)
	}
	return pru
}

// SetUpdatedAt sets the "updated_at" field.
func (pru *PullRequestUpdate) SetUpdatedAt(t time.Time) *PullRequestUpdate {
	pru.mutation.SetUpdatedAt(t)
	return pru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pru *PullRequestUpdate) SetNillableUpdatedAt(t *time.Time) *PullRequestUpdate {
	if t != nil {
		pru.SetUpdatedAt(*t)
	}
	return pru
}

// SetClosedAt sets the "closed_at" field.
func (pru *PullRequestUpdate) SetClosedAt(t time.Time) *PullRequestUpdate {
	pru.mutation.SetClosedAt(t)
	return pru
}

// SetNillableClosedAt sets the "closed_at" field if the given value is not nil.
func (pru *PullRequestUpdate) SetNillableClosedAt(t *time.Time) *PullRequestUpdate {
	if t != nil {
		pru.SetClosedAt(*t)
	}
	return pru
}

// ClearClosedAt clears the value of the "closed_at" field.
func (pru *PullRequestUpdate) ClearClosedAt() *PullRequestUpdate {
	pru.mutation.ClearClosedAt()
	return pru
}

// SetMergedAt sets the "merged_at" field.
func (pru *PullRequestUpdate) SetMergedAt(t time.Time) *PullRequestUpdate {
	pru.mutation.SetMergedAt(t)
	return pru
}

// SetNillableMergedAt sets the "merged_at" field if the given value is not nil.
func (pru *PullRequestUpdate) SetNillableMergedAt(t *time.Time) *PullRequestUpdate {
	if t != nil {
		pru.SetMergedAt(*t)
	}
	return pru
}

// ClearMergedAt clears the value of the "merged_at" field.
func (pru *PullRequestUpdate) ClearMergedAt() *PullRequestUpdate {
	pru.mutation.ClearMergedAt()
	return pru
}

// AddCommitIDs adds the "commits" edge to the Commits entity by IDs.
func (pru *PullRequestUpdate) AddCommitIDs(ids ...uuid.UUID) *PullRequestUpdate {
	pru.mutation.AddCommitIDs(ids...)
	return pru
}

// AddCommits adds the "commits" edges to the Commits entity.
func (pru *PullRequestUpdate) AddCommits(c ...*Commits) *PullRequestUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pru.AddCommitIDs(ids...)
}

// SetRepositoryID sets the "repository" edge to the Repository entity by ID.
func (pru *PullRequestUpdate) SetRepositoryID(id uuid.UUID) *PullRequestUpdate {
	pru.mutation.SetRepositoryID(id)
	return pru
}

// SetNillableRepositoryID sets the "repository" edge to the Repository entity by ID if the given value is not nil.
func (pru *PullRequestUpdate) SetNillableRepositoryID(id *uuid.UUID) *PullRequestUpdate {
	if id != nil {
		pru = pru.SetRepositoryID(*id)
	}
	return pru
}

// SetRepository sets the "repository" edge to the Repository entity.
func (pru *PullRequestUpdate) SetRepository(r *Repository) *PullRequestUpdate {
	return pru.SetRepositoryID(r.ID)
}

// Mutation returns the PullRequestMutation object of the builder.
func (pru *PullRequestUpdate) Mutation() *PullRequestMutation {
	return pru.mutation
}

// ClearCommits clears all "commits" edges to the Commits entity.
func (pru *PullRequestUpdate) ClearCommits() *PullRequestUpdate {
	pru.mutation.ClearCommits()
	return pru
}

// RemoveCommitIDs removes the "commits" edge to Commits entities by IDs.
func (pru *PullRequestUpdate) RemoveCommitIDs(ids ...uuid.UUID) *PullRequestUpdate {
	pru.mutation.RemoveCommitIDs(ids...)
	return pru
}

// RemoveCommits removes "commits" edges to Commits entities.
func (pru *PullRequestUpdate) RemoveCommits(c ...*Commits) *PullRequestUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pru.RemoveCommitIDs(ids...)
}

// ClearRepository clears the "repository" edge to the Repository entity.
func (pru *PullRequestUpdate) ClearRepository() *PullRequestUpdate {
	pru.mutation.ClearRepository()
	return pru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pru *PullRequestUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PullRequestMutation](ctx, pru.sqlSave, pru.mutation, pru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pru *PullRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := pru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pru *PullRequestUpdate) Exec(ctx context.Context) error {
	_, err := pru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pru *PullRequestUpdate) ExecX(ctx context.Context) {
	if err := pru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pru *PullRequestUpdate) check() error {
	if v, ok := pru.mutation.GithubID(); ok {
		if err := pullrequest.GithubIDValidator(v); err != nil {
			return &ValidationError{Name: "github_id", err: fmt.Errorf(`ent: validator failed for field "PullRequest.github_id": %w`, err)}
		}
	}
	if v, ok := pru.mutation.Title(); ok {
		if err := pullrequest.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "PullRequest.title": %w`, err)}
		}
	}
	if v, ok := pru.mutation.TotalCommits(); ok {
		if err := pullrequest.TotalCommitsValidator(v); err != nil {
			return &ValidationError{Name: "total_commits", err: fmt.Errorf(`ent: validator failed for field "PullRequest.total_commits": %w`, err)}
		}
	}
	return nil
}

func (pru *PullRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(pullrequest.Table, pullrequest.Columns, sqlgraph.NewFieldSpec(pullrequest.FieldID, field.TypeUUID))
	if ps := pru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pru.mutation.GithubID(); ok {
		_spec.SetField(pullrequest.FieldGithubID, field.TypeString, value)
	}
	if value, ok := pru.mutation.Title(); ok {
		_spec.SetField(pullrequest.FieldTitle, field.TypeString, value)
	}
	if value, ok := pru.mutation.TotalCommits(); ok {
		_spec.SetField(pullrequest.FieldTotalCommits, field.TypeInt64, value)
	}
	if value, ok := pru.mutation.AddedTotalCommits(); ok {
		_spec.AddField(pullrequest.FieldTotalCommits, field.TypeInt64, value)
	}
	if value, ok := pru.mutation.GetCommit(); ok {
		_spec.SetField(pullrequest.FieldGetCommit, field.TypeBool, value)
	}
	if value, ok := pru.mutation.CreatedAt(); ok {
		_spec.SetField(pullrequest.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pru.mutation.UpdatedAt(); ok {
		_spec.SetField(pullrequest.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pru.mutation.ClosedAt(); ok {
		_spec.SetField(pullrequest.FieldClosedAt, field.TypeTime, value)
	}
	if pru.mutation.ClosedAtCleared() {
		_spec.ClearField(pullrequest.FieldClosedAt, field.TypeTime)
	}
	if value, ok := pru.mutation.MergedAt(); ok {
		_spec.SetField(pullrequest.FieldMergedAt, field.TypeTime, value)
	}
	if pru.mutation.MergedAtCleared() {
		_spec.ClearField(pullrequest.FieldMergedAt, field.TypeTime)
	}
	if pru.mutation.CommitsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.RemovedCommitsIDs(); len(nodes) > 0 && !pru.mutation.CommitsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.CommitsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pru.mutation.RepositoryCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.RepositoryIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pullrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pru.mutation.done = true
	return n, nil
}

// PullRequestUpdateOne is the builder for updating a single PullRequest entity.
type PullRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PullRequestMutation
}

// SetGithubID sets the "github_id" field.
func (pruo *PullRequestUpdateOne) SetGithubID(s string) *PullRequestUpdateOne {
	pruo.mutation.SetGithubID(s)
	return pruo
}

// SetTitle sets the "title" field.
func (pruo *PullRequestUpdateOne) SetTitle(s string) *PullRequestUpdateOne {
	pruo.mutation.SetTitle(s)
	return pruo
}

// SetTotalCommits sets the "total_commits" field.
func (pruo *PullRequestUpdateOne) SetTotalCommits(i int64) *PullRequestUpdateOne {
	pruo.mutation.ResetTotalCommits()
	pruo.mutation.SetTotalCommits(i)
	return pruo
}

// SetNillableTotalCommits sets the "total_commits" field if the given value is not nil.
func (pruo *PullRequestUpdateOne) SetNillableTotalCommits(i *int64) *PullRequestUpdateOne {
	if i != nil {
		pruo.SetTotalCommits(*i)
	}
	return pruo
}

// AddTotalCommits adds i to the "total_commits" field.
func (pruo *PullRequestUpdateOne) AddTotalCommits(i int64) *PullRequestUpdateOne {
	pruo.mutation.AddTotalCommits(i)
	return pruo
}

// SetGetCommit sets the "get_commit" field.
func (pruo *PullRequestUpdateOne) SetGetCommit(b bool) *PullRequestUpdateOne {
	pruo.mutation.SetGetCommit(b)
	return pruo
}

// SetNillableGetCommit sets the "get_commit" field if the given value is not nil.
func (pruo *PullRequestUpdateOne) SetNillableGetCommit(b *bool) *PullRequestUpdateOne {
	if b != nil {
		pruo.SetGetCommit(*b)
	}
	return pruo
}

// SetCreatedAt sets the "created_at" field.
func (pruo *PullRequestUpdateOne) SetCreatedAt(t time.Time) *PullRequestUpdateOne {
	pruo.mutation.SetCreatedAt(t)
	return pruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pruo *PullRequestUpdateOne) SetNillableCreatedAt(t *time.Time) *PullRequestUpdateOne {
	if t != nil {
		pruo.SetCreatedAt(*t)
	}
	return pruo
}

// SetUpdatedAt sets the "updated_at" field.
func (pruo *PullRequestUpdateOne) SetUpdatedAt(t time.Time) *PullRequestUpdateOne {
	pruo.mutation.SetUpdatedAt(t)
	return pruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pruo *PullRequestUpdateOne) SetNillableUpdatedAt(t *time.Time) *PullRequestUpdateOne {
	if t != nil {
		pruo.SetUpdatedAt(*t)
	}
	return pruo
}

// SetClosedAt sets the "closed_at" field.
func (pruo *PullRequestUpdateOne) SetClosedAt(t time.Time) *PullRequestUpdateOne {
	pruo.mutation.SetClosedAt(t)
	return pruo
}

// SetNillableClosedAt sets the "closed_at" field if the given value is not nil.
func (pruo *PullRequestUpdateOne) SetNillableClosedAt(t *time.Time) *PullRequestUpdateOne {
	if t != nil {
		pruo.SetClosedAt(*t)
	}
	return pruo
}

// ClearClosedAt clears the value of the "closed_at" field.
func (pruo *PullRequestUpdateOne) ClearClosedAt() *PullRequestUpdateOne {
	pruo.mutation.ClearClosedAt()
	return pruo
}

// SetMergedAt sets the "merged_at" field.
func (pruo *PullRequestUpdateOne) SetMergedAt(t time.Time) *PullRequestUpdateOne {
	pruo.mutation.SetMergedAt(t)
	return pruo
}

// SetNillableMergedAt sets the "merged_at" field if the given value is not nil.
func (pruo *PullRequestUpdateOne) SetNillableMergedAt(t *time.Time) *PullRequestUpdateOne {
	if t != nil {
		pruo.SetMergedAt(*t)
	}
	return pruo
}

// ClearMergedAt clears the value of the "merged_at" field.
func (pruo *PullRequestUpdateOne) ClearMergedAt() *PullRequestUpdateOne {
	pruo.mutation.ClearMergedAt()
	return pruo
}

// AddCommitIDs adds the "commits" edge to the Commits entity by IDs.
func (pruo *PullRequestUpdateOne) AddCommitIDs(ids ...uuid.UUID) *PullRequestUpdateOne {
	pruo.mutation.AddCommitIDs(ids...)
	return pruo
}

// AddCommits adds the "commits" edges to the Commits entity.
func (pruo *PullRequestUpdateOne) AddCommits(c ...*Commits) *PullRequestUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pruo.AddCommitIDs(ids...)
}

// SetRepositoryID sets the "repository" edge to the Repository entity by ID.
func (pruo *PullRequestUpdateOne) SetRepositoryID(id uuid.UUID) *PullRequestUpdateOne {
	pruo.mutation.SetRepositoryID(id)
	return pruo
}

// SetNillableRepositoryID sets the "repository" edge to the Repository entity by ID if the given value is not nil.
func (pruo *PullRequestUpdateOne) SetNillableRepositoryID(id *uuid.UUID) *PullRequestUpdateOne {
	if id != nil {
		pruo = pruo.SetRepositoryID(*id)
	}
	return pruo
}

// SetRepository sets the "repository" edge to the Repository entity.
func (pruo *PullRequestUpdateOne) SetRepository(r *Repository) *PullRequestUpdateOne {
	return pruo.SetRepositoryID(r.ID)
}

// Mutation returns the PullRequestMutation object of the builder.
func (pruo *PullRequestUpdateOne) Mutation() *PullRequestMutation {
	return pruo.mutation
}

// ClearCommits clears all "commits" edges to the Commits entity.
func (pruo *PullRequestUpdateOne) ClearCommits() *PullRequestUpdateOne {
	pruo.mutation.ClearCommits()
	return pruo
}

// RemoveCommitIDs removes the "commits" edge to Commits entities by IDs.
func (pruo *PullRequestUpdateOne) RemoveCommitIDs(ids ...uuid.UUID) *PullRequestUpdateOne {
	pruo.mutation.RemoveCommitIDs(ids...)
	return pruo
}

// RemoveCommits removes "commits" edges to Commits entities.
func (pruo *PullRequestUpdateOne) RemoveCommits(c ...*Commits) *PullRequestUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pruo.RemoveCommitIDs(ids...)
}

// ClearRepository clears the "repository" edge to the Repository entity.
func (pruo *PullRequestUpdateOne) ClearRepository() *PullRequestUpdateOne {
	pruo.mutation.ClearRepository()
	return pruo
}

// Where appends a list predicates to the PullRequestUpdate builder.
func (pruo *PullRequestUpdateOne) Where(ps ...predicate.PullRequest) *PullRequestUpdateOne {
	pruo.mutation.Where(ps...)
	return pruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pruo *PullRequestUpdateOne) Select(field string, fields ...string) *PullRequestUpdateOne {
	pruo.fields = append([]string{field}, fields...)
	return pruo
}

// Save executes the query and returns the updated PullRequest entity.
func (pruo *PullRequestUpdateOne) Save(ctx context.Context) (*PullRequest, error) {
	return withHooks[*PullRequest, PullRequestMutation](ctx, pruo.sqlSave, pruo.mutation, pruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pruo *PullRequestUpdateOne) SaveX(ctx context.Context) *PullRequest {
	node, err := pruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pruo *PullRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := pruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pruo *PullRequestUpdateOne) ExecX(ctx context.Context) {
	if err := pruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pruo *PullRequestUpdateOne) check() error {
	if v, ok := pruo.mutation.GithubID(); ok {
		if err := pullrequest.GithubIDValidator(v); err != nil {
			return &ValidationError{Name: "github_id", err: fmt.Errorf(`ent: validator failed for field "PullRequest.github_id": %w`, err)}
		}
	}
	if v, ok := pruo.mutation.Title(); ok {
		if err := pullrequest.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "PullRequest.title": %w`, err)}
		}
	}
	if v, ok := pruo.mutation.TotalCommits(); ok {
		if err := pullrequest.TotalCommitsValidator(v); err != nil {
			return &ValidationError{Name: "total_commits", err: fmt.Errorf(`ent: validator failed for field "PullRequest.total_commits": %w`, err)}
		}
	}
	return nil
}

func (pruo *PullRequestUpdateOne) sqlSave(ctx context.Context) (_node *PullRequest, err error) {
	if err := pruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(pullrequest.Table, pullrequest.Columns, sqlgraph.NewFieldSpec(pullrequest.FieldID, field.TypeUUID))
	id, ok := pruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PullRequest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pullrequest.FieldID)
		for _, f := range fields {
			if !pullrequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pullrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pruo.mutation.GithubID(); ok {
		_spec.SetField(pullrequest.FieldGithubID, field.TypeString, value)
	}
	if value, ok := pruo.mutation.Title(); ok {
		_spec.SetField(pullrequest.FieldTitle, field.TypeString, value)
	}
	if value, ok := pruo.mutation.TotalCommits(); ok {
		_spec.SetField(pullrequest.FieldTotalCommits, field.TypeInt64, value)
	}
	if value, ok := pruo.mutation.AddedTotalCommits(); ok {
		_spec.AddField(pullrequest.FieldTotalCommits, field.TypeInt64, value)
	}
	if value, ok := pruo.mutation.GetCommit(); ok {
		_spec.SetField(pullrequest.FieldGetCommit, field.TypeBool, value)
	}
	if value, ok := pruo.mutation.CreatedAt(); ok {
		_spec.SetField(pullrequest.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pruo.mutation.UpdatedAt(); ok {
		_spec.SetField(pullrequest.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pruo.mutation.ClosedAt(); ok {
		_spec.SetField(pullrequest.FieldClosedAt, field.TypeTime, value)
	}
	if pruo.mutation.ClosedAtCleared() {
		_spec.ClearField(pullrequest.FieldClosedAt, field.TypeTime)
	}
	if value, ok := pruo.mutation.MergedAt(); ok {
		_spec.SetField(pullrequest.FieldMergedAt, field.TypeTime, value)
	}
	if pruo.mutation.MergedAtCleared() {
		_spec.ClearField(pullrequest.FieldMergedAt, field.TypeTime)
	}
	if pruo.mutation.CommitsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.RemovedCommitsIDs(); len(nodes) > 0 && !pruo.mutation.CommitsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.CommitsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pruo.mutation.RepositoryCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.RepositoryIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PullRequest{config: pruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pullrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pruo.mutation.done = true
	return _node, nil
}
