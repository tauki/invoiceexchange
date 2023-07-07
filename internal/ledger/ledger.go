package ledger

import (
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/internal/errors"
	"time"
)

type Status string
type Entity string

const (
	StatusPending  Status = "pending"
	StatusResolved Status = "resolved"

	EntityTypeIssuer   Entity = "issuer"
	EntityTypeInvestor Entity = "investor"
)

type Ledger struct {
	ID uuid.UUID `json:"id"`
	// Status is the status of the ledger
	Status Status `json:"status"`
	// InvoiceID is the ID of the invoice that this ledger is associated with
	InvoiceID uuid.UUID `json:"invoice_id"`
	// Entity is the type of entity that this ledger is associated with
	Entity Entity `json:"entity"`
	// EntityID is the ID of the entity that this ledger is associated with
	EntityID uuid.UUID `json:"entity_id"`
	// Amount is the amount of money that this ledger is associated with
	Amount float64 `json:"amount"`
	// CreatedAt is the time that this ledger was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time that this ledger was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

func New(invoiceID uuid.UUID, entity Entity, entityID uuid.UUID, amount float64) (*Ledger, error) {
	if invoiceID == uuid.Nil {
		return nil, errors.NewValidationError("invoice_id cannot be empty", "invoice_id")
	}

	if entityID == uuid.Nil {
		return nil, errors.NewValidationError("entity_id cannot be empty", "entity_id")
	}

	return &Ledger{
		ID:        uuid.New(),
		Status:    StatusPending,
		InvoiceID: invoiceID,
		Entity:    entity,
		EntityID:  entityID,
		Amount:    amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
