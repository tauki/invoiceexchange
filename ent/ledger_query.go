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
	"github.com/tauki/invoiceexchange/ent/ledger"
	"github.com/tauki/invoiceexchange/ent/predicate"
)

// LedgerQuery is the builder for querying Ledger entities.
type LedgerQuery struct {
	config
	ctx        *QueryContext
	order      []ledger.OrderOption
	inters     []Interceptor
	predicates []predicate.Ledger
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LedgerQuery builder.
func (lq *LedgerQuery) Where(ps ...predicate.Ledger) *LedgerQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LedgerQuery) Limit(limit int) *LedgerQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LedgerQuery) Offset(offset int) *LedgerQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LedgerQuery) Unique(unique bool) *LedgerQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LedgerQuery) Order(o ...ledger.OrderOption) *LedgerQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// First returns the first Ledger entity from the query.
// Returns a *NotFoundError when no Ledger was found.
func (lq *LedgerQuery) First(ctx context.Context) (*Ledger, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ledger.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LedgerQuery) FirstX(ctx context.Context) *Ledger {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Ledger ID from the query.
// Returns a *NotFoundError when no Ledger ID was found.
func (lq *LedgerQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ledger.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LedgerQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Ledger entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Ledger entity is found.
// Returns a *NotFoundError when no Ledger entities are found.
func (lq *LedgerQuery) Only(ctx context.Context) (*Ledger, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ledger.Label}
	default:
		return nil, &NotSingularError{ledger.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LedgerQuery) OnlyX(ctx context.Context) *Ledger {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Ledger ID in the query.
// Returns a *NotSingularError when more than one Ledger ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LedgerQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ledger.Label}
	default:
		err = &NotSingularError{ledger.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LedgerQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Ledgers.
func (lq *LedgerQuery) All(ctx context.Context) ([]*Ledger, error) {
	ctx = setContextOp(ctx, lq.ctx, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Ledger, *LedgerQuery]()
	return withInterceptors[[]*Ledger](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LedgerQuery) AllX(ctx context.Context) []*Ledger {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Ledger IDs.
func (lq *LedgerQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, "IDs")
	if err = lq.Select(ledger.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LedgerQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LedgerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LedgerQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LedgerQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LedgerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, "Exist")
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LedgerQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LedgerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LedgerQuery) Clone() *LedgerQuery {
	if lq == nil {
		return nil
	}
	return &LedgerQuery{
		config:     lq.config,
		ctx:        lq.ctx.Clone(),
		order:      append([]ledger.OrderOption{}, lq.order...),
		inters:     append([]Interceptor{}, lq.inters...),
		predicates: append([]predicate.Ledger{}, lq.predicates...),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status ledger.Status `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Ledger.Query().
//		GroupBy(ledger.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LedgerQuery) GroupBy(field string, fields ...string) *LedgerGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LedgerGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = ledger.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status ledger.Status `json:"status,omitempty"`
//	}
//
//	client.Ledger.Query().
//		Select(ledger.FieldStatus).
//		Scan(ctx, &v)
func (lq *LedgerQuery) Select(fields ...string) *LedgerSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LedgerSelect{LedgerQuery: lq}
	sbuild.label = ledger.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LedgerSelect configured with the given aggregations.
func (lq *LedgerQuery) Aggregate(fns ...AggregateFunc) *LedgerSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LedgerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.ctx.Fields {
		if !ledger.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LedgerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Ledger, error) {
	var (
		nodes = []*Ledger{}
		_spec = lq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Ledger).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Ledger{config: lq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lq *LedgerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Node.Columns = lq.ctx.Fields
	if len(lq.ctx.Fields) > 0 {
		_spec.Unique = lq.ctx.Unique != nil && *lq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LedgerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(ledger.Table, ledger.Columns, sqlgraph.NewFieldSpec(ledger.FieldID, field.TypeUUID))
	_spec.From = lq.sql
	if unique := lq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lq.path != nil {
		_spec.Unique = true
	}
	if fields := lq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ledger.FieldID)
		for i := range fields {
			if fields[i] != ledger.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LedgerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(ledger.Table)
	columns := lq.ctx.Fields
	if len(columns) == 0 {
		columns = ledger.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.ctx.Unique != nil && *lq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LedgerGroupBy is the group-by builder for Ledger entities.
type LedgerGroupBy struct {
	selector
	build *LedgerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LedgerGroupBy) Aggregate(fns ...AggregateFunc) *LedgerGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LedgerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LedgerQuery, *LedgerGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LedgerGroupBy) sqlScan(ctx context.Context, root *LedgerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LedgerSelect is the builder for selecting fields of Ledger entities.
type LedgerSelect struct {
	*LedgerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LedgerSelect) Aggregate(fns ...AggregateFunc) *LedgerSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LedgerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LedgerQuery, *LedgerSelect](ctx, ls.LedgerQuery, ls, ls.inters, v)
}

func (ls *LedgerSelect) sqlScan(ctx context.Context, root *LedgerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
