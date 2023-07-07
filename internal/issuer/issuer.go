package issuer

import (
	"github.com/tauki/invoiceexchange/internal/balance"
	"time"

	"github.com/google/uuid"

	"github.com/tauki/invoiceexchange/internal/errors"
)

type Issuer struct {
	ID       uuid.UUID        `json:"id"`
	Name     string           `json:"name"`
	JoinedAt time.Time        `json:"joined_at"`
	Balance  *balance.Balance `json:"balance"`
}

func New(name string, balanceAmount float64) (*Issuer, error) {
	if name == "" {
		return nil, errors.NewValidationError("name is required", "name")
	}

	issuerID := uuid.New()

	bal, err := balance.New(issuerID, balanceAmount)
	if err != nil {
		return nil, err
	}

	return &Issuer{
		ID:       uuid.New(),
		Name:     name,
		Balance:  bal,
		JoinedAt: time.Now(),
	}, nil
}
