// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/ent/balance"
	"github.com/tauki/invoiceexchange/ent/investor"
	"github.com/tauki/invoiceexchange/ent/issuer"
	"github.com/tauki/invoiceexchange/ent/predicate"
)

// BalanceQuery is the builder for querying Balance entities.
type BalanceQuery struct {
	config
	ctx          *QueryContext
	order        []balance.OrderOption
	inters       []Interceptor
	predicates   []predicate.Balance
	withInvestor *InvestorQuery
	withIssuer   *IssuerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BalanceQuery builder.
func (bq *BalanceQuery) Where(ps ...predicate.Balance) *BalanceQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BalanceQuery) Limit(limit int) *BalanceQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BalanceQuery) Offset(offset int) *BalanceQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BalanceQuery) Unique(unique bool) *BalanceQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BalanceQuery) Order(o ...balance.OrderOption) *BalanceQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryInvestor chains the current query on the "investor" edge.
func (bq *BalanceQuery) QueryInvestor() *InvestorQuery {
	query := (&InvestorClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(balance.Table, balance.FieldID, selector),
			sqlgraph.To(investor.Table, investor.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, balance.InvestorTable, balance.InvestorColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIssuer chains the current query on the "issuer" edge.
func (bq *BalanceQuery) QueryIssuer() *IssuerQuery {
	query := (&IssuerClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(balance.Table, balance.FieldID, selector),
			sqlgraph.To(issuer.Table, issuer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, balance.IssuerTable, balance.IssuerColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Balance entity from the query.
// Returns a *NotFoundError when no Balance was found.
func (bq *BalanceQuery) First(ctx context.Context) (*Balance, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{balance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BalanceQuery) FirstX(ctx context.Context) *Balance {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Balance ID from the query.
// Returns a *NotFoundError when no Balance ID was found.
func (bq *BalanceQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{balance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BalanceQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Balance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Balance entity is found.
// Returns a *NotFoundError when no Balance entities are found.
func (bq *BalanceQuery) Only(ctx context.Context) (*Balance, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{balance.Label}
	default:
		return nil, &NotSingularError{balance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BalanceQuery) OnlyX(ctx context.Context) *Balance {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Balance ID in the query.
// Returns a *NotSingularError when more than one Balance ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BalanceQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{balance.Label}
	default:
		err = &NotSingularError{balance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BalanceQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Balances.
func (bq *BalanceQuery) All(ctx context.Context) ([]*Balance, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Balance, *BalanceQuery]()
	return withInterceptors[[]*Balance](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BalanceQuery) AllX(ctx context.Context) []*Balance {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Balance IDs.
func (bq *BalanceQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(balance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BalanceQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BalanceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BalanceQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BalanceQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BalanceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BalanceQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BalanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BalanceQuery) Clone() *BalanceQuery {
	if bq == nil {
		return nil
	}
	return &BalanceQuery{
		config:       bq.config,
		ctx:          bq.ctx.Clone(),
		order:        append([]balance.OrderOption{}, bq.order...),
		inters:       append([]Interceptor{}, bq.inters...),
		predicates:   append([]predicate.Balance{}, bq.predicates...),
		withInvestor: bq.withInvestor.Clone(),
		withIssuer:   bq.withIssuer.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithInvestor tells the query-builder to eager-load the nodes that are connected to
// the "investor" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BalanceQuery) WithInvestor(opts ...func(*InvestorQuery)) *BalanceQuery {
	query := (&InvestorClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withInvestor = query
	return bq
}

// WithIssuer tells the query-builder to eager-load the nodes that are connected to
// the "issuer" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BalanceQuery) WithIssuer(opts ...func(*IssuerQuery)) *BalanceQuery {
	query := (&IssuerClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withIssuer = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TotalAmount float64 `json:"total_amount,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Balance.Query().
//		GroupBy(balance.FieldTotalAmount).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BalanceQuery) GroupBy(field string, fields ...string) *BalanceGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BalanceGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = balance.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TotalAmount float64 `json:"total_amount,omitempty"`
//	}
//
//	client.Balance.Query().
//		Select(balance.FieldTotalAmount).
//		Scan(ctx, &v)
func (bq *BalanceQuery) Select(fields ...string) *BalanceSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BalanceSelect{BalanceQuery: bq}
	sbuild.label = balance.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BalanceSelect configured with the given aggregations.
func (bq *BalanceQuery) Aggregate(fns ...AggregateFunc) *BalanceSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BalanceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !balance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BalanceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Balance, error) {
	var (
		nodes       = []*Balance{}
		_spec       = bq.querySpec()
		loadedTypes = [2]bool{
			bq.withInvestor != nil,
			bq.withIssuer != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Balance).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Balance{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withInvestor; query != nil {
		if err := bq.loadInvestor(ctx, query, nodes,
			func(n *Balance) { n.Edges.Investor = []*Investor{} },
			func(n *Balance, e *Investor) { n.Edges.Investor = append(n.Edges.Investor, e) }); err != nil {
			return nil, err
		}
	}
	if query := bq.withIssuer; query != nil {
		if err := bq.loadIssuer(ctx, query, nodes,
			func(n *Balance) { n.Edges.Issuer = []*Issuer{} },
			func(n *Balance, e *Issuer) { n.Edges.Issuer = append(n.Edges.Issuer, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BalanceQuery) loadInvestor(ctx context.Context, query *InvestorQuery, nodes []*Balance, init func(*Balance), assign func(*Balance, *Investor)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Balance)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Investor(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(balance.InvestorColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.investor_balance
		if fk == nil {
			return fmt.Errorf(`foreign-key "investor_balance" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "investor_balance" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (bq *BalanceQuery) loadIssuer(ctx context.Context, query *IssuerQuery, nodes []*Balance, init func(*Balance), assign func(*Balance, *Issuer)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Balance)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Issuer(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(balance.IssuerColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.issuer_balance
		if fk == nil {
			return fmt.Errorf(`foreign-key "issuer_balance" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "issuer_balance" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bq *BalanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BalanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(balance.Table, balance.Columns, sqlgraph.NewFieldSpec(balance.FieldID, field.TypeUUID))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, balance.FieldID)
		for i := range fields {
			if fields[i] != balance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BalanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(balance.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = balance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BalanceGroupBy is the group-by builder for Balance entities.
type BalanceGroupBy struct {
	selector
	build *BalanceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BalanceGroupBy) Aggregate(fns ...AggregateFunc) *BalanceGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BalanceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BalanceQuery, *BalanceGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BalanceGroupBy) sqlScan(ctx context.Context, root *BalanceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BalanceSelect is the builder for selecting fields of Balance entities.
type BalanceSelect struct {
	*BalanceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BalanceSelect) Aggregate(fns ...AggregateFunc) *BalanceSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BalanceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BalanceQuery, *BalanceSelect](ctx, bs.BalanceQuery, bs, bs.inters, v)
}

func (bs *BalanceSelect) sqlScan(ctx context.Context, root *BalanceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
