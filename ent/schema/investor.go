package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

type Investor struct {
	ent.Schema
}

func (Investor) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable().
			Default(uuid.New),
		field.String("name"),
		field.Time("joined_at").
			Immutable().
			Default(time.Now),
	}
}

func (Investor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bids", Bid.Type),
		edge.To("invoices", Invoice.Type),
		edge.To("balance", Balance.Type).Unique().Required(),
	}
}
