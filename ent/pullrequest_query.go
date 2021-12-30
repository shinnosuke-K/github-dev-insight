// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shinnosuke-K/github-dev-insight/ent/commit"
	"github.com/shinnosuke-K/github-dev-insight/ent/predicate"
	"github.com/shinnosuke-K/github-dev-insight/ent/pullrequest"
	"github.com/shinnosuke-K/github-dev-insight/ent/repository"
)

// PullRequestQuery is the builder for querying PullRequest entities.
type PullRequestQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PullRequest
	// eager-loading edges.
	withCommits    *CommitQuery
	withRepository *RepositoryQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PullRequestQuery builder.
func (prq *PullRequestQuery) Where(ps ...predicate.PullRequest) *PullRequestQuery {
	prq.predicates = append(prq.predicates, ps...)
	return prq
}

// Limit adds a limit step to the query.
func (prq *PullRequestQuery) Limit(limit int) *PullRequestQuery {
	prq.limit = &limit
	return prq
}

// Offset adds an offset step to the query.
func (prq *PullRequestQuery) Offset(offset int) *PullRequestQuery {
	prq.offset = &offset
	return prq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (prq *PullRequestQuery) Unique(unique bool) *PullRequestQuery {
	prq.unique = &unique
	return prq
}

// Order adds an order step to the query.
func (prq *PullRequestQuery) Order(o ...OrderFunc) *PullRequestQuery {
	prq.order = append(prq.order, o...)
	return prq
}

// QueryCommits chains the current query on the "commits" edge.
func (prq *PullRequestQuery) QueryCommits() *CommitQuery {
	query := &CommitQuery{config: prq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := prq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pullrequest.Table, pullrequest.FieldID, selector),
			sqlgraph.To(commit.Table, commit.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, pullrequest.CommitsTable, pullrequest.CommitsColumn),
		)
		fromU = sqlgraph.SetNeighbors(prq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRepository chains the current query on the "repository" edge.
func (prq *PullRequestQuery) QueryRepository() *RepositoryQuery {
	query := &RepositoryQuery{config: prq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := prq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pullrequest.Table, pullrequest.FieldID, selector),
			sqlgraph.To(repository.Table, repository.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, pullrequest.RepositoryTable, pullrequest.RepositoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(prq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PullRequest entity from the query.
// Returns a *NotFoundError when no PullRequest was found.
func (prq *PullRequestQuery) First(ctx context.Context) (*PullRequest, error) {
	nodes, err := prq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pullrequest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (prq *PullRequestQuery) FirstX(ctx context.Context) *PullRequest {
	node, err := prq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PullRequest ID from the query.
// Returns a *NotFoundError when no PullRequest ID was found.
func (prq *PullRequestQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = prq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pullrequest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (prq *PullRequestQuery) FirstIDX(ctx context.Context) int {
	id, err := prq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PullRequest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one PullRequest entity is not found.
// Returns a *NotFoundError when no PullRequest entities are found.
func (prq *PullRequestQuery) Only(ctx context.Context) (*PullRequest, error) {
	nodes, err := prq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pullrequest.Label}
	default:
		return nil, &NotSingularError{pullrequest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (prq *PullRequestQuery) OnlyX(ctx context.Context) *PullRequest {
	node, err := prq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PullRequest ID in the query.
// Returns a *NotSingularError when exactly one PullRequest ID is not found.
// Returns a *NotFoundError when no entities are found.
func (prq *PullRequestQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = prq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pullrequest.Label}
	default:
		err = &NotSingularError{pullrequest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (prq *PullRequestQuery) OnlyIDX(ctx context.Context) int {
	id, err := prq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PullRequests.
func (prq *PullRequestQuery) All(ctx context.Context) ([]*PullRequest, error) {
	if err := prq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return prq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (prq *PullRequestQuery) AllX(ctx context.Context) []*PullRequest {
	nodes, err := prq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PullRequest IDs.
func (prq *PullRequestQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := prq.Select(pullrequest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (prq *PullRequestQuery) IDsX(ctx context.Context) []int {
	ids, err := prq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (prq *PullRequestQuery) Count(ctx context.Context) (int, error) {
	if err := prq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return prq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (prq *PullRequestQuery) CountX(ctx context.Context) int {
	count, err := prq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (prq *PullRequestQuery) Exist(ctx context.Context) (bool, error) {
	if err := prq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return prq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (prq *PullRequestQuery) ExistX(ctx context.Context) bool {
	exist, err := prq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PullRequestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (prq *PullRequestQuery) Clone() *PullRequestQuery {
	if prq == nil {
		return nil
	}
	return &PullRequestQuery{
		config:         prq.config,
		limit:          prq.limit,
		offset:         prq.offset,
		order:          append([]OrderFunc{}, prq.order...),
		predicates:     append([]predicate.PullRequest{}, prq.predicates...),
		withCommits:    prq.withCommits.Clone(),
		withRepository: prq.withRepository.Clone(),
		// clone intermediate query.
		sql:  prq.sql.Clone(),
		path: prq.path,
	}
}

// WithCommits tells the query-builder to eager-load the nodes that are connected to
// the "commits" edge. The optional arguments are used to configure the query builder of the edge.
func (prq *PullRequestQuery) WithCommits(opts ...func(*CommitQuery)) *PullRequestQuery {
	query := &CommitQuery{config: prq.config}
	for _, opt := range opts {
		opt(query)
	}
	prq.withCommits = query
	return prq
}

// WithRepository tells the query-builder to eager-load the nodes that are connected to
// the "repository" edge. The optional arguments are used to configure the query builder of the edge.
func (prq *PullRequestQuery) WithRepository(opts ...func(*RepositoryQuery)) *PullRequestQuery {
	query := &RepositoryQuery{config: prq.config}
	for _, opt := range opts {
		opt(query)
	}
	prq.withRepository = query
	return prq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RepositoryID string `json:"repository_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PullRequest.Query().
//		GroupBy(pullrequest.FieldRepositoryID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (prq *PullRequestQuery) GroupBy(field string, fields ...string) *PullRequestGroupBy {
	group := &PullRequestGroupBy{config: prq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return prq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RepositoryID string `json:"repository_id,omitempty"`
//	}
//
//	client.PullRequest.Query().
//		Select(pullrequest.FieldRepositoryID).
//		Scan(ctx, &v)
//
func (prq *PullRequestQuery) Select(fields ...string) *PullRequestSelect {
	prq.fields = append(prq.fields, fields...)
	return &PullRequestSelect{PullRequestQuery: prq}
}

func (prq *PullRequestQuery) prepareQuery(ctx context.Context) error {
	for _, f := range prq.fields {
		if !pullrequest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if prq.path != nil {
		prev, err := prq.path(ctx)
		if err != nil {
			return err
		}
		prq.sql = prev
	}
	return nil
}

func (prq *PullRequestQuery) sqlAll(ctx context.Context) ([]*PullRequest, error) {
	var (
		nodes       = []*PullRequest{}
		withFKs     = prq.withFKs
		_spec       = prq.querySpec()
		loadedTypes = [2]bool{
			prq.withCommits != nil,
			prq.withRepository != nil,
		}
	)
	if prq.withRepository != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, pullrequest.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PullRequest{config: prq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, prq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := prq.withCommits; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*PullRequest)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Commits = []*Commit{}
		}
		query.withFKs = true
		query.Where(predicate.Commit(func(s *sql.Selector) {
			s.Where(sql.InValues(pullrequest.CommitsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.pull_request_commits
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "pull_request_commits" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "pull_request_commits" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Commits = append(node.Edges.Commits, n)
		}
	}

	if query := prq.withRepository; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*PullRequest)
		for i := range nodes {
			if nodes[i].repository_pull_requests == nil {
				continue
			}
			fk := *nodes[i].repository_pull_requests
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(repository.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "repository_pull_requests" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Repository = n
			}
		}
	}

	return nodes, nil
}

func (prq *PullRequestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := prq.querySpec()
	return sqlgraph.CountNodes(ctx, prq.driver, _spec)
}

func (prq *PullRequestQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := prq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (prq *PullRequestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pullrequest.Table,
			Columns: pullrequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pullrequest.FieldID,
			},
		},
		From:   prq.sql,
		Unique: true,
	}
	if unique := prq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := prq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pullrequest.FieldID)
		for i := range fields {
			if fields[i] != pullrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := prq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := prq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := prq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := prq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (prq *PullRequestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(prq.driver.Dialect())
	t1 := builder.Table(pullrequest.Table)
	columns := prq.fields
	if len(columns) == 0 {
		columns = pullrequest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if prq.sql != nil {
		selector = prq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range prq.predicates {
		p(selector)
	}
	for _, p := range prq.order {
		p(selector)
	}
	if offset := prq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := prq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PullRequestGroupBy is the group-by builder for PullRequest entities.
type PullRequestGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (prgb *PullRequestGroupBy) Aggregate(fns ...AggregateFunc) *PullRequestGroupBy {
	prgb.fns = append(prgb.fns, fns...)
	return prgb
}

// Scan applies the group-by query and scans the result into the given value.
func (prgb *PullRequestGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := prgb.path(ctx)
	if err != nil {
		return err
	}
	prgb.sql = query
	return prgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (prgb *PullRequestGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := prgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (prgb *PullRequestGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(prgb.fields) > 1 {
		return nil, errors.New("ent: PullRequestGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := prgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (prgb *PullRequestGroupBy) StringsX(ctx context.Context) []string {
	v, err := prgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (prgb *PullRequestGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = prgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pullrequest.Label}
	default:
		err = fmt.Errorf("ent: PullRequestGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (prgb *PullRequestGroupBy) StringX(ctx context.Context) string {
	v, err := prgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (prgb *PullRequestGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(prgb.fields) > 1 {
		return nil, errors.New("ent: PullRequestGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := prgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (prgb *PullRequestGroupBy) IntsX(ctx context.Context) []int {
	v, err := prgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (prgb *PullRequestGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = prgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pullrequest.Label}
	default:
		err = fmt.Errorf("ent: PullRequestGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (prgb *PullRequestGroupBy) IntX(ctx context.Context) int {
	v, err := prgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (prgb *PullRequestGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(prgb.fields) > 1 {
		return nil, errors.New("ent: PullRequestGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := prgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (prgb *PullRequestGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := prgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (prgb *PullRequestGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = prgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pullrequest.Label}
	default:
		err = fmt.Errorf("ent: PullRequestGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (prgb *PullRequestGroupBy) Float64X(ctx context.Context) float64 {
	v, err := prgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (prgb *PullRequestGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(prgb.fields) > 1 {
		return nil, errors.New("ent: PullRequestGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := prgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (prgb *PullRequestGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := prgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (prgb *PullRequestGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = prgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pullrequest.Label}
	default:
		err = fmt.Errorf("ent: PullRequestGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (prgb *PullRequestGroupBy) BoolX(ctx context.Context) bool {
	v, err := prgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (prgb *PullRequestGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range prgb.fields {
		if !pullrequest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := prgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (prgb *PullRequestGroupBy) sqlQuery() *sql.Selector {
	selector := prgb.sql.Select()
	aggregation := make([]string, 0, len(prgb.fns))
	for _, fn := range prgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(prgb.fields)+len(prgb.fns))
		for _, f := range prgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(prgb.fields...)...)
}

// PullRequestSelect is the builder for selecting fields of PullRequest entities.
type PullRequestSelect struct {
	*PullRequestQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (prs *PullRequestSelect) Scan(ctx context.Context, v interface{}) error {
	if err := prs.prepareQuery(ctx); err != nil {
		return err
	}
	prs.sql = prs.PullRequestQuery.sqlQuery(ctx)
	return prs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (prs *PullRequestSelect) ScanX(ctx context.Context, v interface{}) {
	if err := prs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (prs *PullRequestSelect) Strings(ctx context.Context) ([]string, error) {
	if len(prs.fields) > 1 {
		return nil, errors.New("ent: PullRequestSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := prs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (prs *PullRequestSelect) StringsX(ctx context.Context) []string {
	v, err := prs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (prs *PullRequestSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = prs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pullrequest.Label}
	default:
		err = fmt.Errorf("ent: PullRequestSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (prs *PullRequestSelect) StringX(ctx context.Context) string {
	v, err := prs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (prs *PullRequestSelect) Ints(ctx context.Context) ([]int, error) {
	if len(prs.fields) > 1 {
		return nil, errors.New("ent: PullRequestSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := prs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (prs *PullRequestSelect) IntsX(ctx context.Context) []int {
	v, err := prs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (prs *PullRequestSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = prs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pullrequest.Label}
	default:
		err = fmt.Errorf("ent: PullRequestSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (prs *PullRequestSelect) IntX(ctx context.Context) int {
	v, err := prs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (prs *PullRequestSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(prs.fields) > 1 {
		return nil, errors.New("ent: PullRequestSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := prs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (prs *PullRequestSelect) Float64sX(ctx context.Context) []float64 {
	v, err := prs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (prs *PullRequestSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = prs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pullrequest.Label}
	default:
		err = fmt.Errorf("ent: PullRequestSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (prs *PullRequestSelect) Float64X(ctx context.Context) float64 {
	v, err := prs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (prs *PullRequestSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(prs.fields) > 1 {
		return nil, errors.New("ent: PullRequestSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := prs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (prs *PullRequestSelect) BoolsX(ctx context.Context) []bool {
	v, err := prs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (prs *PullRequestSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = prs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pullrequest.Label}
	default:
		err = fmt.Errorf("ent: PullRequestSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (prs *PullRequestSelect) BoolX(ctx context.Context) bool {
	v, err := prs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (prs *PullRequestSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := prs.sql.Query()
	if err := prs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
