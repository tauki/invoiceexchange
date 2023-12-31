// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tauki/invoiceexchange/ent/issuer"
	"github.com/tauki/invoiceexchange/ent/predicate"
)

// IssuerDelete is the builder for deleting a Issuer entity.
type IssuerDelete struct {
	config
	hooks    []Hook
	mutation *IssuerMutation
}

// Where appends a list predicates to the IssuerDelete builder.
func (id *IssuerDelete) Where(ps ...predicate.Issuer) *IssuerDelete {
	id.mutation.Where(ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *IssuerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, id.sqlExec, id.mutation, id.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (id *IssuerDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *IssuerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(issuer.Table, sqlgraph.NewFieldSpec(issuer.FieldID, field.TypeUUID))
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, id.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	id.mutation.done = true
	return affected, err
}

// IssuerDeleteOne is the builder for deleting a single Issuer entity.
type IssuerDeleteOne struct {
	id *IssuerDelete
}

// Where appends a list predicates to the IssuerDelete builder.
func (ido *IssuerDeleteOne) Where(ps ...predicate.Issuer) *IssuerDeleteOne {
	ido.id.mutation.Where(ps...)
	return ido
}

// Exec executes the deletion query.
func (ido *IssuerDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{issuer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *IssuerDeleteOne) ExecX(ctx context.Context) {
	if err := ido.Exec(ctx); err != nil {
		panic(err)
	}
}
