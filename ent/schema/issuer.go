package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Issuer struct {
	ent.Schema
}

func (Issuer) Fields() []ent.Field {
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

func (Issuer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invoices", Invoice.Type),
		edge.To("balance", Balance.Type).Unique().Required(),
	}
}
