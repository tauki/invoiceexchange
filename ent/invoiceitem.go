// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/ent/invoice"
	"github.com/tauki/invoiceexchange/ent/invoiceitem"
)

// InvoiceItem is the model entity for the InvoiceItem schema.
type InvoiceItem struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// UnitPrice holds the value of the "unit_price" field.
	UnitPrice float64 `json:"unit_price,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount,omitempty"`
	// VatRate holds the value of the "vat_rate" field.
	VatRate float64 `json:"vat_rate,omitempty"`
	// VatAmount holds the value of the "vat_amount" field.
	VatAmount float64 `json:"vat_amount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InvoiceItemQuery when eager-loading is set.
	Edges         InvoiceItemEdges `json:"edges"`
	invoice_items *uuid.UUID
	selectValues  sql.SelectValues
}

// InvoiceItemEdges holds the relations/edges for other nodes in the graph.
type InvoiceItemEdges struct {
	// Invoice holds the value of the invoice edge.
	Invoice *Invoice `json:"invoice,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// InvoiceOrErr returns the Invoice value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvoiceItemEdges) InvoiceOrErr() (*Invoice, error) {
	if e.loadedTypes[0] {
		if e.Invoice == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: invoice.Label}
		}
		return e.Invoice, nil
	}
	return nil, &NotLoadedError{edge: "invoice"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InvoiceItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case invoiceitem.FieldUnitPrice, invoiceitem.FieldAmount, invoiceitem.FieldVatRate, invoiceitem.FieldVatAmount:
			values[i] = new(sql.NullFloat64)
		case invoiceitem.FieldQuantity:
			values[i] = new(sql.NullInt64)
		case invoiceitem.FieldDescription:
			values[i] = new(sql.NullString)
		case invoiceitem.FieldID:
			values[i] = new(uuid.UUID)
		case invoiceitem.ForeignKeys[0]: // invoice_items
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InvoiceItem fields.
func (ii *InvoiceItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case invoiceitem.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ii.ID = *value
			}
		case invoiceitem.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ii.Description = value.String
			}
		case invoiceitem.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				ii.Quantity = int(value.Int64)
			}
		case invoiceitem.FieldUnitPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field unit_price", values[i])
			} else if value.Valid {
				ii.UnitPrice = value.Float64
			}
		case invoiceitem.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				ii.Amount = value.Float64
			}
		case invoiceitem.FieldVatRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field vat_rate", values[i])
			} else if value.Valid {
				ii.VatRate = value.Float64
			}
		case invoiceitem.FieldVatAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field vat_amount", values[i])
			} else if value.Valid {
				ii.VatAmount = value.Float64
			}
		case invoiceitem.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field invoice_items", values[i])
			} else if value.Valid {
				ii.invoice_items = new(uuid.UUID)
				*ii.invoice_items = *value.S.(*uuid.UUID)
			}
		default:
			ii.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the InvoiceItem.
// This includes values selected through modifiers, order, etc.
func (ii *InvoiceItem) Value(name string) (ent.Value, error) {
	return ii.selectValues.Get(name)
}

// QueryInvoice queries the "invoice" edge of the InvoiceItem entity.
func (ii *InvoiceItem) QueryInvoice() *InvoiceQuery {
	return NewInvoiceItemClient(ii.config).QueryInvoice(ii)
}

// Update returns a builder for updating this InvoiceItem.
// Note that you need to call InvoiceItem.Unwrap() before calling this method if this InvoiceItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (ii *InvoiceItem) Update() *InvoiceItemUpdateOne {
	return NewInvoiceItemClient(ii.config).UpdateOne(ii)
}

// Unwrap unwraps the InvoiceItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ii *InvoiceItem) Unwrap() *InvoiceItem {
	_tx, ok := ii.config.driver.(*txDriver)
	if !ok {
		panic("ent: InvoiceItem is not a transactional entity")
	}
	ii.config.driver = _tx.drv
	return ii
}

// String implements the fmt.Stringer.
func (ii *InvoiceItem) String() string {
	var builder strings.Builder
	builder.WriteString("InvoiceItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ii.ID))
	builder.WriteString("description=")
	builder.WriteString(ii.Description)
	builder.WriteString(", ")
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", ii.Quantity))
	builder.WriteString(", ")
	builder.WriteString("unit_price=")
	builder.WriteString(fmt.Sprintf("%v", ii.UnitPrice))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", ii.Amount))
	builder.WriteString(", ")
	builder.WriteString("vat_rate=")
	builder.WriteString(fmt.Sprintf("%v", ii.VatRate))
	builder.WriteString(", ")
	builder.WriteString("vat_amount=")
	builder.WriteString(fmt.Sprintf("%v", ii.VatAmount))
	builder.WriteByte(')')
	return builder.String()
}

// InvoiceItems is a parsable slice of InvoiceItem.
type InvoiceItems []*InvoiceItem
