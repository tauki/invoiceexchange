package bid

import (
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/internal/errors"
	"time"
)

type Status string

const (
	// Pending declares a bid has been created but has not been approved or rejected yet
	Pending Status = "PENDING"
	// Accepted declares a bid has been approved for transaction
	Accepted Status = "ACCEPTED"
	// Rejected declares the bid has not been approved
	Rejected Status = "REJECTED"
)

type Bid struct {
	ID             uuid.UUID `json:"id"`
	InvoiceID      uuid.UUID `json:"invoice_id"`
	InvestorID     uuid.UUID `json:"investor_id"`
	Amount         float64   `json:"amount"`
	AcceptedAmount float64   `json:"accepted_amount"`
	Status         Status    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func New(invoiceID, investorID uuid.UUID, amount float64) (*Bid, error) {
	if invoiceID == uuid.Nil {
		return nil, errors.NewValidationError("invoice id cannot be nil", "invoiceID")
	}
	if investorID == uuid.Nil {
		return nil, errors.NewValidationError("investor id cannot be nil", "investorID")
	}
	if amount <= 0 {
		return nil, errors.NewValidationError("amount cannot be less than or equal to 0", "amount")
	}

	return &Bid{
		ID:         uuid.New(),
		InvoiceID:  invoiceID,
		InvestorID: investorID,
		Amount:     amount,
		Status:     Pending,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
