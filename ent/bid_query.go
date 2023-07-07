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
	"github.com/tauki/invoiceexchange/ent/bid"
	"github.com/tauki/invoiceexchange/ent/investor"
	"github.com/tauki/invoiceexchange/ent/invoice"
	"github.com/tauki/invoiceexchange/ent/predicate"
)

// BidQuery is the builder for querying Bid entities.
type BidQuery struct {
	config
	ctx          *QueryContext
	order        []bid.OrderOption
	inters       []Interceptor
	predicates   []predicate.Bid
	withInvoice  *InvoiceQuery
	withInvestor *InvestorQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BidQuery builder.
func (bq *BidQuery) Where(ps ...predicate.Bid) *BidQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BidQuery) Limit(limit int) *BidQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BidQuery) Offset(offset int) *BidQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BidQuery) Unique(unique bool) *BidQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BidQuery) Order(o ...bid.OrderOption) *BidQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryInvoice chains the current query on the "invoice" edge.
func (bq *BidQuery) QueryInvoice() *InvoiceQuery {
	query := (&InvoiceClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bid.Table, bid.FieldID, selector),
			sqlgraph.To(invoice.Table, invoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bid.InvoiceTable, bid.InvoiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInvestor chains the current query on the "investor" edge.
func (bq *BidQuery) QueryInvestor() *InvestorQuery {
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
			sqlgraph.From(bid.Table, bid.FieldID, selector),
			sqlgraph.To(investor.Table, investor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bid.InvestorTable, bid.InvestorColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Bid entity from the query.
// Returns a *NotFoundError when no Bid was found.
func (bq *BidQuery) First(ctx context.Context) (*Bid, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bid.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BidQuery) FirstX(ctx context.Context) *Bid {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Bid ID from the query.
// Returns a *NotFoundError when no Bid ID was found.
func (bq *BidQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bid.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BidQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Bid entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Bid entity is found.
// Returns a *NotFoundError when no Bid entities are found.
func (bq *BidQuery) Only(ctx context.Context) (*Bid, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bid.Label}
	default:
		return nil, &NotSingularError{bid.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BidQuery) OnlyX(ctx context.Context) *Bid {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Bid ID in the query.
// Returns a *NotSingularError when more than one Bid ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BidQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bid.Label}
	default:
		err = &NotSingularError{bid.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BidQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Bids.
func (bq *BidQuery) All(ctx context.Context) ([]*Bid, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Bid, *BidQuery]()
	return withInterceptors[[]*Bid](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BidQuery) AllX(ctx context.Context) []*Bid {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Bid IDs.
func (bq *BidQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(bid.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BidQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BidQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BidQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BidQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BidQuery) Exist(ctx context.Context) (bool, error) {
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
func (bq *BidQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BidQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BidQuery) Clone() *BidQuery {
	if bq == nil {
		return nil
	}
	return &BidQuery{
		config:       bq.config,
		ctx:          bq.ctx.Clone(),
		order:        append([]bid.OrderOption{}, bq.order...),
		inters:       append([]Interceptor{}, bq.inters...),
		predicates:   append([]predicate.Bid{}, bq.predicates...),
		withInvoice:  bq.withInvoice.Clone(),
		withInvestor: bq.withInvestor.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithInvoice tells the query-builder to eager-load the nodes that are connected to
// the "invoice" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BidQuery) WithInvoice(opts ...func(*InvoiceQuery)) *BidQuery {
	query := (&InvoiceClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withInvoice = query
	return bq
}

// WithInvestor tells the query-builder to eager-load the nodes that are connected to
// the "investor" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BidQuery) WithInvestor(opts ...func(*InvestorQuery)) *BidQuery {
	query := (&InvestorClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withInvestor = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status bid.Status `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Bid.Query().
//		GroupBy(bid.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BidQuery) GroupBy(field string, fields ...string) *BidGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BidGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = bid.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status bid.Status `json:"status,omitempty"`
//	}
//
//	client.Bid.Query().
//		Select(bid.FieldStatus).
//		Scan(ctx, &v)
func (bq *BidQuery) Select(fields ...string) *BidSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BidSelect{BidQuery: bq}
	sbuild.label = bid.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BidSelect configured with the given aggregations.
func (bq *BidQuery) Aggregate(fns ...AggregateFunc) *BidSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BidQuery) prepareQuery(ctx context.Context) error {
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
		if !bid.ValidColumn(f) {
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

func (bq *BidQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Bid, error) {
	var (
		nodes       = []*Bid{}
		withFKs     = bq.withFKs
		_spec       = bq.querySpec()
		loadedTypes = [2]bool{
			bq.withInvoice != nil,
			bq.withInvestor != nil,
		}
	)
	if bq.withInvoice != nil || bq.withInvestor != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, bid.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Bid).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Bid{config: bq.config}
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
	if query := bq.withInvoice; query != nil {
		if err := bq.loadInvoice(ctx, query, nodes, nil,
			func(n *Bid, e *Invoice) { n.Edges.Invoice = e }); err != nil {
			return nil, err
		}
	}
	if query := bq.withInvestor; query != nil {
		if err := bq.loadInvestor(ctx, query, nodes, nil,
			func(n *Bid, e *Investor) { n.Edges.Investor = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BidQuery) loadInvoice(ctx context.Context, query *InvoiceQuery, nodes []*Bid, init func(*Bid), assign func(*Bid, *Invoice)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Bid)
	for i := range nodes {
		if nodes[i].invoice_bids == nil {
			continue
		}
		fk := *nodes[i].invoice_bids
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(invoice.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "invoice_bids" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bq *BidQuery) loadInvestor(ctx context.Context, query *InvestorQuery, nodes []*Bid, init func(*Bid), assign func(*Bid, *Investor)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Bid)
	for i := range nodes {
		if nodes[i].investor_bids == nil {
			continue
		}
		fk := *nodes[i].investor_bids
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(investor.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "investor_bids" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (bq *BidQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BidQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bid.Table, bid.Columns, sqlgraph.NewFieldSpec(bid.FieldID, field.TypeUUID))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bid.FieldID)
		for i := range fields {
			if fields[i] != bid.FieldID {
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

func (bq *BidQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(bid.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = bid.Columns
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

// BidGroupBy is the group-by builder for Bid entities.
type BidGroupBy struct {
	selector
	build *BidQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BidGroupBy) Aggregate(fns ...AggregateFunc) *BidGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BidGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BidQuery, *BidGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BidGroupBy) sqlScan(ctx context.Context, root *BidQuery, v any) error {
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

// BidSelect is the builder for selecting fields of Bid entities.
type BidSelect struct {
	*BidQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BidSelect) Aggregate(fns ...AggregateFunc) *BidSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BidSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BidQuery, *BidSelect](ctx, bs.BidQuery, bs, bs.inters, v)
}

func (bs *BidSelect) sqlScan(ctx context.Context, root *BidQuery, v any) error {
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