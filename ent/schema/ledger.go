package schema

import (
	"entgo.io/ent"
	"time"

	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	domain "github.com/tauki/invoiceexchange/internal/ledger"
)

type Ledger struct {
	ent.Schema
}

func (Ledger) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable().
			Default(uuid.New),
		field.Enum("status").Values(
			string(domain.StatusPending),
			string(domain.StatusResolved),
		).Default(string(domain.StatusPending)),
		field.UUID("invoice_id", uuid.UUID{}).
			Immutable(),
		field.Enum("entity").Values(
			string(domain.EntityTypeIssuer),
			string(domain.EntityTypeInvestor),
		).Immutable(),
		field.UUID("entity_id", uuid.UUID{}).
			Immutable(),
		field.Float("amount"),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
