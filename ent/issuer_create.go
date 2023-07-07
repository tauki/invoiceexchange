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
	"github.com/tauki/invoiceexchange/ent/invoice"
	"github.com/tauki/invoiceexchange/ent/issuer"
)

// IssuerCreate is the builder for creating a Issuer entity.
type IssuerCreate struct {
	config
	mutation *IssuerMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ic *IssuerCreate) SetName(s string) *IssuerCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetJoinedAt sets the "joined_at" field.
func (ic *IssuerCreate) SetJoinedAt(t time.Time) *IssuerCreate {
	ic.mutation.SetJoinedAt(t)
	return ic
}

// SetNillableJoinedAt sets the "joined_at" field if the given value is not nil.
func (ic *IssuerCreate) SetNillableJoinedAt(t *time.Time) *IssuerCreate {
	if t != nil {
		ic.SetJoinedAt(*t)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *IssuerCreate) SetID(u uuid.UUID) *IssuerCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *IssuerCreate) SetNillableID(u *uuid.UUID) *IssuerCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// AddInvoiceIDs adds the "invoices" edge to the Invoice entity by IDs.
func (ic *IssuerCreate) AddInvoiceIDs(ids ...uuid.UUID) *IssuerCreate {
	ic.mutation.AddInvoiceIDs(ids...)
	return ic
}

// AddInvoices adds the "invoices" edges to the Invoice entity.
func (ic *IssuerCreate) AddInvoices(i ...*Invoice) *IssuerCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ic.AddInvoiceIDs(ids...)
}

// SetBalanceID sets the "balance" edge to the Balance entity by ID.
func (ic *IssuerCreate) SetBalanceID(id uuid.UUID) *IssuerCreate {
	ic.mutation.SetBalanceID(id)
	return ic
}

// SetBalance sets the "balance" edge to the Balance entity.
func (ic *IssuerCreate) SetBalance(b *Balance) *IssuerCreate {
	return ic.SetBalanceID(b.ID)
}

// Mutation returns the IssuerMutation object of the builder.
func (ic *IssuerCreate) Mutation() *IssuerMutation {
	return ic.mutation
}

// Save creates the Issuer in the database.
func (ic *IssuerCreate) Save(ctx context.Context) (*Issuer, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IssuerCreate) SaveX(ctx context.Context) *Issuer {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IssuerCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IssuerCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *IssuerCreate) defaults() {
	if _, ok := ic.mutation.JoinedAt(); !ok {
		v := issuer.DefaultJoinedAt()
		ic.mutation.SetJoinedAt(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := issuer.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IssuerCreate) check() error {
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Issuer.name"`)}
	}
	if _, ok := ic.mutation.JoinedAt(); !ok {
		return &ValidationError{Name: "joined_at", err: errors.New(`ent: missing required field "Issuer.joined_at"`)}
	}
	if _, ok := ic.mutation.BalanceID(); !ok {
		return &ValidationError{Name: "balance", err: errors.New(`ent: missing required edge "Issuer.balance"`)}
	}
	return nil
}

func (ic *IssuerCreate) sqlSave(ctx context.Context) (*Issuer, error) {
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

func (ic *IssuerCreate) createSpec() (*Issuer, *sqlgraph.CreateSpec) {
	var (
		_node = &Issuer{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(issuer.Table, sqlgraph.NewFieldSpec(issuer.FieldID, field.TypeUUID))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(issuer.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.JoinedAt(); ok {
		_spec.SetField(issuer.FieldJoinedAt, field.TypeTime, value)
		_node.JoinedAt = value
	}
	if nodes := ic.mutation.InvoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   issuer.InvoicesTable,
			Columns: []string{issuer.InvoicesColumn},
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
			Table:   issuer.BalanceTable,
			Columns: []string{issuer.BalanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(balance.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.issuer_balance = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IssuerCreateBulk is the builder for creating many Issuer entities in bulk.
type IssuerCreateBulk struct {
	config
	builders []*IssuerCreate
}

// Save creates the Issuer entities in the database.
func (icb *IssuerCreateBulk) Save(ctx context.Context) ([]*Issuer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Issuer, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IssuerMutation)
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
func (icb *IssuerCreateBulk) SaveX(ctx context.Context) []*Issuer {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IssuerCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IssuerCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
