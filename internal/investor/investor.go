package investor

import (
	"github.com/tauki/invoiceexchange/internal/balance"
	"github.com/tauki/invoiceexchange/internal/errors"
	"time"

	"github.com/google/uuid"
)

type Investor struct {
	ID       uuid.UUID        `json:"id"`
	Name     string           `json:"name"`
	JoinedAt time.Time        `json:"joined_at"`
	Balance  *balance.Balance `json:"balance"`
}

func New(name string, balanceAmount float64) (*Investor, error) {
	if name == "" {
		return nil, errors.NewValidationError("name cannot be empty", "name")
	}

	investorID := uuid.New()

	bal, err := balance.New(investorID, balanceAmount)
	if err != nil {
		return nil, err
	}

	return &Investor{
		ID:       uuid.New(),
		Name:     name,
		Balance:  bal,
		JoinedAt: time.Now(),
	}, nil
}
