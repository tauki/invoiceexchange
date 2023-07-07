package balance

import (
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/internal/errors"
	"time"
)

type Balance struct {
	ID uuid.UUID `json:"id"`
	// EntityID is the ID of the entity that owns this balance
	EntityID uuid.UUID `json:"entity_id"`
	// TotalAmount is the total amount of money in this balance
	TotalAmount float64 `json:"total_amount"`
	// AvailableAmount is the amount of money that is available for use
	AvailableAmount float64   `json:"available_amount"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func New(entityID uuid.UUID, totalAmount float64) (*Balance, error) {
	if entityID == uuid.Nil {
		return nil, errors.NewValidationError("entity_id cannot be empty", "entity_id")
	}

	if totalAmount < 0 {
		return nil, errors.NewValidationError("total_amount cannot be negative", "total_amount")
	}

	return &Balance{
		ID:              uuid.New(),
		EntityID:        entityID,
		TotalAmount:     totalAmount,
		AvailableAmount: totalAmount,
		CreatedAt:       time.Now(),
	}, nil
}
