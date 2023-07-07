// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/ent/balance"
	"github.com/tauki/invoiceexchange/ent/bid"
	"github.com/tauki/invoiceexchange/ent/investor"
	"github.com/tauki/invoiceexchange/ent/invoice"
)

// InvestorCreate is the builder for creating a Investor entity.
type InvestorCreate struct {
	config
	mutation *InvestorMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ic *InvestorCreate) SetName(s string) *InvestorCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetJoinedAt sets the "joined_at" field.
func (ic *InvestorCreate) SetJoinedAt(t time.Time) *InvestorCreate {
	ic.mutation.SetJoinedAt(t)
	return ic
}

// SetNillableJoinedAt sets the "joined_at" field if the given value is not nil.
func (ic *InvestorCreate) SetNillableJoinedAt(t *time.Time) *InvestorCreate {
	if t != nil {
		ic.SetJoinedAt(*t)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *InvestorCreate) SetID(u uuid.UUID) *InvestorCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *InvestorCreate) SetNillableID(u *uuid.UUID) *InvestorCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// AddBidIDs adds the "bids" edge to the Bid entity by IDs.
func (ic *InvestorCreate) AddBidIDs(ids ...uuid.UUID) *InvestorCreate {
	ic.mutation.AddBidIDs(ids...)
	return ic
}

// AddBids adds the "bids" edges to the Bid entity.
func (ic *InvestorCreate) AddBids(b ...*Bid) *InvestorCreate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ic.AddBidIDs(ids...)
}

// AddInvoiceIDs adds the "invoices" edge to the Invoice entity by IDs.
func (ic *InvestorCreate) AddInvoiceIDs(ids ...uuid.UUID) *InvestorCreate {
	ic.mutation.AddInvoiceIDs(ids...)
	return ic
}

// AddInvoices adds the "invoices" edges to the Invoice entity.
func (ic *InvestorCreate) AddInvoices(i ...*Invoice) *InvestorCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ic.AddInvoiceIDs(ids...)
}

// SetBalanceID sets the "balance" edge to the Balance entity by ID.
func (ic *InvestorCreate) SetBalanceID(id uuid.UUID) *InvestorCreate {
	ic.mutation.SetBalanceID(id)
	return ic
}

// SetBalance sets the "balance" edge to the Balance entity.
func (ic *InvestorCreate) SetBalance(b *Balance) *InvestorCreate {
	return ic.SetBalanceID(b.ID)
}

// Mutation returns the InvestorMutation object of the builder.
func (ic *InvestorCreate) Mutation() *InvestorMutation {
	return ic.mutation
}

// Save creates the Investor in the database.
func (ic *InvestorCreate) Save(ctx context.Context) (*Investor, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InvestorCreate) SaveX(ctx context.Context) *Investor {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InvestorCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InvestorCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *InvestorCreate) defaults() {
	if _, ok := ic.mutation.JoinedAt(); !ok {
		v := investor.DefaultJoinedAt()
		ic.mutation.SetJoinedAt(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := investor.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InvestorCreate) check() error {
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Investor.name"`)}
	}
	if _, ok := ic.mutation.JoinedAt(); !ok {
		return &ValidationError{Name: "joined_at", err: errors.New(`ent: missing required field "Investor.joined_at"`)}
	}
	if _, ok := ic.mutation.BalanceID(); !ok {
		return &ValidationError{Name: "balance", err: errors.New(`ent: missing required edge "Investor.balance"`)}
	}
	return nil
}

func (ic *InvestorCreate) sqlSave(ctx context.Context) (*Investor, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *InvestorCreate) createSpec() (*Investor, *sqlgraph.CreateSpec) {
	var (
		_node = &Investor{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(investor.Table, sqlgraph.NewFieldSpec(investor.FieldID, field.TypeUUID))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(investor.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.JoinedAt(); ok {
		_spec.SetField(investor.FieldJoinedAt, field.TypeTime, value)
		_node.JoinedAt = value
	}
	if nodes := ic.mutation.BidsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   investor.BidsTable,
			Columns: []string{investor.BidsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bid.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.InvoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   investor.InvoicesTable,
			Columns: investor.InvoicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.BalanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   investor.BalanceTable,
			Columns: []string{investor.BalanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(balance.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.investor_balance = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InvestorCreateBulk is the builder for creating many Investor entities in bulk.
type InvestorCreateBulk struct {
	config
	builders []*InvestorCreate
}

// Save creates the Investor entities in the database.
func (icb *InvestorCreateBulk) Save(ctx context.Context) ([]*Investor, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Investor, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InvestorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InvestorCreateBulk) SaveX(ctx context.Context) []*Investor {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InvestorCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InvestorCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
