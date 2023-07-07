// Code generated by ent, DO NOT EDIT.

package balance

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the balance type in the database.
	Label = "balance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTotalAmount holds the string denoting the total_amount field in the database.
	FieldTotalAmount = "total_amount"
	// FieldAvailableAmount holds the string denoting the available_amount field in the database.
	FieldAvailableAmount = "available_amount"
	// FieldEntityID holds the string denoting the entity_id field in the database.
	FieldEntityID = "entity_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeInvestor holds the string denoting the investor edge name in mutations.
	EdgeInvestor = "investor"
	// EdgeIssuer holds the string denoting the issuer edge name in mutations.
	EdgeIssuer = "issuer"
	// Table holds the table name of the balance in the database.
	Table = "balances"
	// InvestorTable is the table that holds the investor relation/edge.
	InvestorTable = "investors"
	// InvestorInverseTable is the table name for the Investor entity.
	// It exists in this package in order to avoid circular dependency with the "investor" package.
	InvestorInverseTable = "investors"
	// InvestorColumn is the table column denoting the investor relation/edge.
	InvestorColumn = "investor_balance"
	// IssuerTable is the table that holds the issuer relation/edge.
	IssuerTable = "issuers"
	// IssuerInverseTable is the table name for the Issuer entity.
	// It exists in this package in order to avoid circular dependency with the "issuer" package.
	IssuerInverseTable = "issuers"
	// IssuerColumn is the table column denoting the issuer relation/edge.
	IssuerColumn = "issuer_balance"
)

// Columns holds all SQL columns for balance fields.
var Columns = []string{
	FieldID,
	FieldTotalAmount,
	FieldAvailableAmount,
	FieldEntityID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Balance queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTotalAmount orders the results by the total_amount field.
func ByTotalAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalAmount, opts...).ToFunc()
}

// ByAvailableAmount orders the results by the available_amount field.
func ByAvailableAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvailableAmount, opts...).ToFunc()
}

// ByEntityID orders the results by the entity_id field.
func ByEntityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntityID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByInvestorCount orders the results by investor count.
func ByInvestorCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInvestorStep(), opts...)
	}
}

// ByInvestor orders the results by investor terms.
func ByInvestor(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvestorStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByIssuerCount orders the results by issuer count.
func ByIssuerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIssuerStep(), opts...)
	}
}

// ByIssuer orders the results by issuer terms.
func ByIssuer(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIssuerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newInvestorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvestorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, InvestorTable, InvestorColumn),
	)
}
func newIssuerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IssuerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, IssuerTable, IssuerColumn),
	)
}
