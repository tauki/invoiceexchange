package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	domain "github.com/tauki/invoiceexchange/internal/bid"
	"time"
)

type Bid struct {
	ent.Schema
}

func (Bid) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable().
			Default(uuid.New),
		field.Enum("status").Values(
			string(domain.Pending),
			string(domain.Accepted),
			string(domain.Rejected),
		).Default(string(domain.Pending)),
		field.Float("amount"),
		field.Float("accepted_amount").
			Default(0),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Bid) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("invoice", Invoice.Type).
			Ref("bids").Unique().Required().Immutable(),
		edge.From("investor", Investor.Type).
			Ref("bids").Unique().Required().Immutable(),
	}
}
