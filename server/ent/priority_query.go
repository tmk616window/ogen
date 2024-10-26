// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"server/ent/predicate"
	"server/ent/priority"
	"server/ent/todo"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PriorityQuery is the builder for querying Priority entities.
type PriorityQuery struct {
	config
	ctx        *QueryContext
	order      []priority.OrderOption
	inters     []Interceptor
	predicates []predicate.Priority
	withTodo   *TodoQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PriorityQuery builder.
func (pq *PriorityQuery) Where(ps ...predicate.Priority) *PriorityQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PriorityQuery) Limit(limit int) *PriorityQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PriorityQuery) Offset(offset int) *PriorityQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PriorityQuery) Unique(unique bool) *PriorityQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PriorityQuery) Order(o ...priority.OrderOption) *PriorityQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryTodo chains the current query on the "todo" edge.
func (pq *PriorityQuery) QueryTodo() *TodoQuery {
	query := (&TodoClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(priority.Table, priority.FieldID, selector),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, priority.TodoTable, priority.TodoColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Priority entity from the query.
// Returns a *NotFoundError when no Priority was found.
func (pq *PriorityQuery) First(ctx context.Context) (*Priority, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{priority.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PriorityQuery) FirstX(ctx context.Context) *Priority {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Priority ID from the query.
// Returns a *NotFoundError when no Priority ID was found.
func (pq *PriorityQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{priority.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PriorityQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Priority entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Priority entity is found.
// Returns a *NotFoundError when no Priority entities are found.
func (pq *PriorityQuery) Only(ctx context.Context) (*Priority, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{priority.Label}
	default:
		return nil, &NotSingularError{priority.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PriorityQuery) OnlyX(ctx context.Context) *Priority {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Priority ID in the query.
// Returns a *NotSingularError when more than one Priority ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PriorityQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{priority.Label}
	default:
		err = &NotSingularError{priority.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PriorityQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Priorities.
func (pq *PriorityQuery) All(ctx context.Context) ([]*Priority, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Priority, *PriorityQuery]()
	return withInterceptors[[]*Priority](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PriorityQuery) AllX(ctx context.Context) []*Priority {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Priority IDs.
func (pq *PriorityQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryIDs)
	if err = pq.Select(priority.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PriorityQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PriorityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PriorityQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PriorityQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PriorityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PriorityQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PriorityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PriorityQuery) Clone() *PriorityQuery {
	if pq == nil {
		return nil
	}
	return &PriorityQuery{
		config:     pq.config,
		ctx:        pq.ctx.Clone(),
		order:      append([]priority.OrderOption{}, pq.order...),
		inters:     append([]Interceptor{}, pq.inters...),
		predicates: append([]predicate.Priority{}, pq.predicates...),
		withTodo:   pq.withTodo.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithTodo tells the query-builder to eager-load the nodes that are connected to
// the "todo" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PriorityQuery) WithTodo(opts ...func(*TodoQuery)) *PriorityQuery {
	query := (&TodoClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withTodo = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Priority.Query().
//		GroupBy(priority.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PriorityQuery) GroupBy(field string, fields ...string) *PriorityGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PriorityGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = priority.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Priority.Query().
//		Select(priority.FieldName).
//		Scan(ctx, &v)
func (pq *PriorityQuery) Select(fields ...string) *PrioritySelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PrioritySelect{PriorityQuery: pq}
	sbuild.label = priority.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PrioritySelect configured with the given aggregations.
func (pq *PriorityQuery) Aggregate(fns ...AggregateFunc) *PrioritySelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PriorityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !priority.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PriorityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Priority, error) {
	var (
		nodes       = []*Priority{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withTodo != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Priority).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Priority{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withTodo; query != nil {
		if err := pq.loadTodo(ctx, query, nodes,
			func(n *Priority) { n.Edges.Todo = []*Todo{} },
			func(n *Priority, e *Todo) { n.Edges.Todo = append(n.Edges.Todo, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PriorityQuery) loadTodo(ctx context.Context, query *TodoQuery, nodes []*Priority, init func(*Priority), assign func(*Priority, *Todo)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Priority)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(todo.FieldPriorityID)
	}
	query.Where(predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(priority.TodoColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PriorityID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "priority_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PriorityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PriorityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(priority.Table, priority.Columns, sqlgraph.NewFieldSpec(priority.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, priority.FieldID)
		for i := range fields {
			if fields[i] != priority.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PriorityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(priority.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = priority.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PriorityGroupBy is the group-by builder for Priority entities.
type PriorityGroupBy struct {
	selector
	build *PriorityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PriorityGroupBy) Aggregate(fns ...AggregateFunc) *PriorityGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PriorityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PriorityQuery, *PriorityGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PriorityGroupBy) sqlScan(ctx context.Context, root *PriorityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PrioritySelect is the builder for selecting fields of Priority entities.
type PrioritySelect struct {
	*PriorityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PrioritySelect) Aggregate(fns ...AggregateFunc) *PrioritySelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PrioritySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PriorityQuery, *PrioritySelect](ctx, ps.PriorityQuery, ps, ps.inters, v)
}

func (ps *PrioritySelect) sqlScan(ctx context.Context, root *PriorityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
