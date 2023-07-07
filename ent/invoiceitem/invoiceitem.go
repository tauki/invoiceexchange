// Code generated by ent, DO NOT EDIT.

package invoiceitem

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the invoiceitem type in the database.
	Label = "invoice_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldUnitPrice holds the string denoting the unit_price field in the database.
	FieldUnitPrice = "unit_price"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldVatRate holds the string denoting the vat_rate field in the database.
	FieldVatRate = "vat_rate"
	// FieldVatAmount holds the string denoting the vat_amount field in the database.
	FieldVatAmount = "vat_amount"
	// EdgeInvoice holds the string denoting the invoice edge name in mutations.
	EdgeInvoice = "invoice"
	// Table holds the table name of the invoiceitem in the database.
	Table = "invoice_items"
	// InvoiceTable is the table that holds the invoice relation/edge.
	InvoiceTable = "invoice_items"
	// InvoiceInverseTable is the table name for the Invoice entity.
	// It exists in this package in order to avoid circular dependency with the "invoice" package.
	InvoiceInverseTable = "invoices"
	// InvoiceColumn is the table column denoting the invoice relation/edge.
	InvoiceColumn = "invoice_items"
)

// Columns holds all SQL columns for invoiceitem fields.
var Columns = []string{
	FieldID,
	FieldDescription,
	FieldQuantity,
	FieldUnitPrice,
	FieldAmount,
	FieldVatRate,
	FieldVatAmount,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "invoice_items"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"invoice_items",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultVatRate holds the default value on creation for the "vat_rate" field.
	DefaultVatRate float64
	// DefaultVatAmount holds the default value on creation for the "vat_amount" field.
	DefaultVatAmount float64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the InvoiceItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByUnitPrice orders the results by the unit_price field.
func ByUnitPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnitPrice, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByVatRate orders the results by the vat_rate field.
func ByVatRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVatRate, opts...).ToFunc()
}

// ByVatAmount orders the results by the vat_amount field.
func ByVatAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVatAmount, opts...).ToFunc()
}

// ByInvoiceField orders the results by invoice field.
func ByInvoiceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvoiceStep(), sql.OrderByField(field, opts...))
	}
}
func newInvoiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvoiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, InvoiceTable, InvoiceColumn),
	)
}