package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"time"

	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Balance struct {
	ent.Schema
}

func (Balance) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable().
			Default(uuid.New),
		field.Float("total_amount"),
		field.Float("available_amount"),
		field.UUID("entity_id", uuid.UUID{}).
			Immutable().
			Unique(),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Balance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("investor", Investor.Type).Ref("balance"),
		edge.From("issuer", Issuer.Type).Ref("balance"),
	}
}
