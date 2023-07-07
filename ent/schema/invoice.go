package schema

import (
	domain "github.com/tauki/invoiceexchange/internal/invoice"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

type Invoice struct {
	ent.Schema
}

func (Invoice) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable().
			Default(uuid.New),
		field.Enum("status").
			Values(string(domain.StatusPending), string(domain.StatusProcessed)).
			Default(string(domain.StatusPending)),
		field.Float("asking_price").
			Default(0.0),
		field.Bool("is_locked").
			Default(false),
		field.Bool("is_approved").
			Default(false),
		field.String("invoice_number"),
		field.Time("invoice_date"),
		field.Time("due_date"),
		field.Float("amount_due"),
		field.String("customer_name"),
		field.String("reference").
			Optional(),
		field.String("company_name").
			Optional(),
		field.String("currency").
			Default("GBP"),
		field.Float("total_amount").
			Default(0.0).Optional(),
		field.Float("total_vat").
			Default(0.0).Optional(),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
	}
}

func (Invoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("items", InvoiceItem.Type),
		edge.To("bids", Bid.Type),
		edge.From("issuer", Issuer.Type).Ref("invoices").Required().Unique(),
		edge.From("investor", Investor.Type).Ref("invoices"),
	}
}

type InvoiceItem struct {
	ent.Schema
}

func (InvoiceItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable().
			Default(uuid.New),
		field.String("description"),
		field.Int("quantity"),
		field.Float("unit_price"),
		field.Float("amount"),
		field.Float("vat_rate").
			Default(0.0),
		field.Float("vat_amount").
			Default(0.0),
	}
}

func (InvoiceItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("invoice", Invoice.Type).Ref("items").Unique(),
	}
}
