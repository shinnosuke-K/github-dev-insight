// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/shinnosuke-K/github-dev-insight/ent/commits"
	"github.com/shinnosuke-K/github-dev-insight/ent/predicate"
	"github.com/shinnosuke-K/github-dev-insight/ent/pullrequest"
)

// CommitsQuery is the builder for querying Commits entities.
type CommitsQuery struct {
	config
	ctx             *QueryContext
	order           []OrderFunc
	inters          []Interceptor
	predicates      []predicate.Commits
	withPullRequest *PullRequestQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommitsQuery builder.
func (cq *CommitsQuery) Where(ps ...predicate.Commits) *CommitsQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CommitsQuery) Limit(limit int) *CommitsQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CommitsQuery) Offset(offset int) *CommitsQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CommitsQuery) Unique(unique bool) *CommitsQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CommitsQuery) Order(o ...OrderFunc) *CommitsQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryPullRequest chains the current query on the "pull_request" edge.
func (cq *CommitsQuery) QueryPullRequest() *PullRequestQuery {
	query := (&PullRequestClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(commits.Table, commits.FieldID, selector),
			sqlgraph.To(pullrequest.Table, pullrequest.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, commits.PullRequestTable, commits.PullRequestColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Commits entity from the query.
// Returns a *NotFoundError when no Commits was found.
func (cq *CommitsQuery) First(ctx context.Context) (*Commits, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{commits.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CommitsQuery) FirstX(ctx context.Context) *Commits {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Commits ID from the query.
// Returns a *NotFoundError when no Commits ID was found.
func (cq *CommitsQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{commits.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CommitsQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Commits entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Commits entity is found.
// Returns a *NotFoundError when no Commits entities are found.
func (cq *CommitsQuery) Only(ctx context.Context) (*Commits, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{commits.Label}
	default:
		return nil, &NotSingularError{commits.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CommitsQuery) OnlyX(ctx context.Context) *Commits {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Commits ID in the query.
// Returns a *NotSingularError when more than one Commits ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CommitsQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{commits.Label}
	default:
		err = &NotSingularError{commits.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CommitsQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CommitsSlice.
func (cq *CommitsQuery) All(ctx context.Context) ([]*Commits, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Commits, *CommitsQuery]()
	return withInterceptors[[]*Commits](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CommitsQuery) AllX(ctx context.Context) []*Commits {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Commits IDs.
func (cq *CommitsQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(commits.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CommitsQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CommitsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CommitsQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CommitsQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CommitsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CommitsQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommitsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CommitsQuery) Clone() *CommitsQuery {
	if cq == nil {
		return nil
	}
	return &CommitsQuery{
		config:          cq.config,
		ctx:             cq.ctx.Clone(),
		order:           append([]OrderFunc{}, cq.order...),
		inters:          append([]Interceptor{}, cq.inters...),
		predicates:      append([]predicate.Commits{}, cq.predicates...),
		withPullRequest: cq.withPullRequest.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithPullRequest tells the query-builder to eager-load the nodes that are connected to
// the "pull_request" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CommitsQuery) WithPullRequest(opts ...func(*PullRequestQuery)) *CommitsQuery {
	query := (&PullRequestClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withPullRequest = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GithubID string `json:"github_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Commits.Query().
//		GroupBy(commits.FieldGithubID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CommitsQuery) GroupBy(field string, fields ...string) *CommitsGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CommitsGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = commits.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GithubID string `json:"github_id,omitempty"`
//	}
//
//	client.Commits.Query().
//		Select(commits.FieldGithubID).
//		Scan(ctx, &v)
func (cq *CommitsQuery) Select(fields ...string) *CommitsSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CommitsSelect{CommitsQuery: cq}
	sbuild.label = commits.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CommitsSelect configured with the given aggregations.
func (cq *CommitsQuery) Aggregate(fns ...AggregateFunc) *CommitsSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CommitsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !commits.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CommitsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Commits, error) {
	var (
		nodes       = []*Commits{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withPullRequest != nil,
		}
	)
	if cq.withPullRequest != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, commits.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Commits).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Commits{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withPullRequest; query != nil {
		if err := cq.loadPullRequest(ctx, query, nodes, nil,
			func(n *Commits, e *PullRequest) { n.Edges.PullRequest = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CommitsQuery) loadPullRequest(ctx context.Context, query *PullRequestQuery, nodes []*Commits, init func(*Commits), assign func(*Commits, *PullRequest)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Commits)
	for i := range nodes {
		if nodes[i].pull_request_id == nil {
			continue
		}
		fk := *nodes[i].pull_request_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(pullrequest.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "pull_request_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cq *CommitsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CommitsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(commits.Table, commits.Columns, sqlgraph.NewFieldSpec(commits.FieldID, field.TypeUUID))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commits.FieldID)
		for i := range fields {
			if fields[i] != commits.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CommitsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(commits.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = commits.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommitsGroupBy is the group-by builder for Commits entities.
type CommitsGroupBy struct {
	selector
	build *CommitsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CommitsGroupBy) Aggregate(fns ...AggregateFunc) *CommitsGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CommitsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommitsQuery, *CommitsGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CommitsGroupBy) sqlScan(ctx context.Context, root *CommitsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CommitsSelect is the builder for selecting fields of Commits entities.
type CommitsSelect struct {
	*CommitsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CommitsSelect) Aggregate(fns ...AggregateFunc) *CommitsSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CommitsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommitsQuery, *CommitsSelect](ctx, cs.CommitsQuery, cs, cs.inters, v)
}

func (cs *CommitsSelect) sqlScan(ctx context.Context, root *CommitsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
