// Code generated by ent, DO NOT EDIT.

package balance

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldLTE(FieldID, id))
}

// TotalAmount applies equality check predicate on the "total_amount" field. It's identical to TotalAmountEQ.
func TotalAmount(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldTotalAmount, v))
}

// AvailableAmount applies equality check predicate on the "available_amount" field. It's identical to AvailableAmountEQ.
func AvailableAmount(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldAvailableAmount, v))
}

// EntityID applies equality check predicate on the "entity_id" field. It's identical to EntityIDEQ.
func EntityID(v uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldEntityID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldUpdatedAt, v))
}

// TotalAmountEQ applies the EQ predicate on the "total_amount" field.
func TotalAmountEQ(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldTotalAmount, v))
}

// TotalAmountNEQ applies the NEQ predicate on the "total_amount" field.
func TotalAmountNEQ(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldNEQ(FieldTotalAmount, v))
}

// TotalAmountIn applies the In predicate on the "total_amount" field.
func TotalAmountIn(vs ...float64) predicate.Balance {
	return predicate.Balance(sql.FieldIn(FieldTotalAmount, vs...))
}

// TotalAmountNotIn applies the NotIn predicate on the "total_amount" field.
func TotalAmountNotIn(vs ...float64) predicate.Balance {
	return predicate.Balance(sql.FieldNotIn(FieldTotalAmount, vs...))
}

// TotalAmountGT applies the GT predicate on the "total_amount" field.
func TotalAmountGT(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldGT(FieldTotalAmount, v))
}

// TotalAmountGTE applies the GTE predicate on the "total_amount" field.
func TotalAmountGTE(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldGTE(FieldTotalAmount, v))
}

// TotalAmountLT applies the LT predicate on the "total_amount" field.
func TotalAmountLT(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldLT(FieldTotalAmount, v))
}

// TotalAmountLTE applies the LTE predicate on the "total_amount" field.
func TotalAmountLTE(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldLTE(FieldTotalAmount, v))
}

// AvailableAmountEQ applies the EQ predicate on the "available_amount" field.
func AvailableAmountEQ(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldAvailableAmount, v))
}

// AvailableAmountNEQ applies the NEQ predicate on the "available_amount" field.
func AvailableAmountNEQ(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldNEQ(FieldAvailableAmount, v))
}

// AvailableAmountIn applies the In predicate on the "available_amount" field.
func AvailableAmountIn(vs ...float64) predicate.Balance {
	return predicate.Balance(sql.FieldIn(FieldAvailableAmount, vs...))
}

// AvailableAmountNotIn applies the NotIn predicate on the "available_amount" field.
func AvailableAmountNotIn(vs ...float64) predicate.Balance {
	return predicate.Balance(sql.FieldNotIn(FieldAvailableAmount, vs...))
}

// AvailableAmountGT applies the GT predicate on the "available_amount" field.
func AvailableAmountGT(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldGT(FieldAvailableAmount, v))
}

// AvailableAmountGTE applies the GTE predicate on the "available_amount" field.
func AvailableAmountGTE(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldGTE(FieldAvailableAmount, v))
}

// AvailableAmountLT applies the LT predicate on the "available_amount" field.
func AvailableAmountLT(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldLT(FieldAvailableAmount, v))
}

// AvailableAmountLTE applies the LTE predicate on the "available_amount" field.
func AvailableAmountLTE(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldLTE(FieldAvailableAmount, v))
}

// EntityIDEQ applies the EQ predicate on the "entity_id" field.
func EntityIDEQ(v uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldEntityID, v))
}

// EntityIDNEQ applies the NEQ predicate on the "entity_id" field.
func EntityIDNEQ(v uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldNEQ(FieldEntityID, v))
}

// EntityIDIn applies the In predicate on the "entity_id" field.
func EntityIDIn(vs ...uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldIn(FieldEntityID, vs...))
}

// EntityIDNotIn applies the NotIn predicate on the "entity_id" field.
func EntityIDNotIn(vs ...uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldNotIn(FieldEntityID, vs...))
}

// EntityIDGT applies the GT predicate on the "entity_id" field.
func EntityIDGT(v uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldGT(FieldEntityID, v))
}

// EntityIDGTE applies the GTE predicate on the "entity_id" field.
func EntityIDGTE(v uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldGTE(FieldEntityID, v))
}

// EntityIDLT applies the LT predicate on the "entity_id" field.
func EntityIDLT(v uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldLT(FieldEntityID, v))
}

// EntityIDLTE applies the LTE predicate on the "entity_id" field.
func EntityIDLTE(v uuid.UUID) predicate.Balance {
	return predicate.Balance(sql.FieldLTE(FieldEntityID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasInvestor applies the HasEdge predicate on the "investor" edge.
func HasInvestor() predicate.Balance {
	return predicate.Balance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, InvestorTable, InvestorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInvestorWith applies the HasEdge predicate on the "investor" edge with a given conditions (other predicates).
func HasInvestorWith(preds ...predicate.Investor) predicate.Balance {
	return predicate.Balance(func(s *sql.Selector) {
		step := newInvestorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasIssuer applies the HasEdge predicate on the "issuer" edge.
func HasIssuer() predicate.Balance {
	return predicate.Balance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, IssuerTable, IssuerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIssuerWith applies the HasEdge predicate on the "issuer" edge with a given conditions (other predicates).
func HasIssuerWith(preds ...predicate.Issuer) predicate.Balance {
	return predicate.Balance(func(s *sql.Selector) {
		step := newIssuerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Balance) predicate.Balance {
	return predicate.Balance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Balance) predicate.Balance {
	return predicate.Balance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Balance) predicate.Balance {
	return predicate.Balance(func(s *sql.Selector) {
		p(s.Not())
	})
}
