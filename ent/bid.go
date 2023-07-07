// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/ent/bid"
	"github.com/tauki/invoiceexchange/ent/investor"
	"github.com/tauki/invoiceexchange/ent/invoice"
)

// Bid is the model entity for the Bid schema.
type Bid struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status bid.Status `json:"status,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount,omitempty"`
	// AcceptedAmount holds the value of the "accepted_amount" field.
	AcceptedAmount float64 `json:"accepted_amount,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BidQuery when eager-loading is set.
	Edges         BidEdges `json:"edges"`
	investor_bids *uuid.UUID
	invoice_bids  *uuid.UUID
	selectValues  sql.SelectValues
}

// BidEdges holds the relations/edges for other nodes in the graph.
type BidEdges struct {
	// Invoice holds the value of the invoice edge.
	Invoice *Invoice `json:"invoice,omitempty"`
	// Investor holds the value of the investor edge.
	Investor *Investor `json:"investor,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// InvoiceOrErr returns the Invoice value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BidEdges) InvoiceOrErr() (*Invoice, error) {
	if e.loadedTypes[0] {
		if e.Invoice == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: invoice.Label}
		}
		return e.Invoice, nil
	}
	return nil, &NotLoadedError{edge: "invoice"}
}

// InvestorOrErr returns the Investor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BidEdges) InvestorOrErr() (*Investor, error) {
	if e.loadedTypes[1] {
		if e.Investor == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: investor.Label}
		}
		return e.Investor, nil
	}
	return nil, &NotLoadedError{edge: "investor"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bid) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bid.FieldAmount, bid.FieldAcceptedAmount:
			values[i] = new(sql.NullFloat64)
		case bid.FieldStatus:
			values[i] = new(sql.NullString)
		case bid.FieldCreatedAt, bid.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case bid.FieldID:
			values[i] = new(uuid.UUID)
		case bid.ForeignKeys[0]: // investor_bids
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case bid.ForeignKeys[1]: // invoice_bids
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bid fields.
func (b *Bid) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bid.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case bid.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = bid.Status(value.String)
			}
		case bid.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				b.Amount = value.Float64
			}
		case bid.FieldAcceptedAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field accepted_amount", values[i])
			} else if value.Valid {
				b.AcceptedAmount = value.Float64
			}
		case bid.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case bid.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case bid.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field investor_bids", values[i])
			} else if value.Valid {
				b.investor_bids = new(uuid.UUID)
				*b.investor_bids = *value.S.(*uuid.UUID)
			}
		case bid.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field invoice_bids", values[i])
			} else if value.Valid {
				b.invoice_bids = new(uuid.UUID)
				*b.invoice_bids = *value.S.(*uuid.UUID)
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Bid.
// This includes values selected through modifiers, order, etc.
func (b *Bid) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryInvoice queries the "invoice" edge of the Bid entity.
func (b *Bid) QueryInvoice() *InvoiceQuery {
	return NewBidClient(b.config).QueryInvoice(b)
}

// QueryInvestor queries the "investor" edge of the Bid entity.
func (b *Bid) QueryInvestor() *InvestorQuery {
	return NewBidClient(b.config).QueryInvestor(b)
}

// Update returns a builder for updating this Bid.
// Note that you need to call Bid.Unwrap() before calling this method if this Bid
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bid) Update() *BidUpdateOne {
	return NewBidClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Bid entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Bid) Unwrap() *Bid {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bid is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bid) String() string {
	var builder strings.Builder
	builder.WriteString("Bid(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", b.Amount))
	builder.WriteString(", ")
	builder.WriteString("accepted_amount=")
	builder.WriteString(fmt.Sprintf("%v", b.AcceptedAmount))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Bids is a parsable slice of Bid.
type Bids []*Bid
