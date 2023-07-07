package ledger

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	AddLedgerEntry(ctx context.Context, invoiceID uuid.UUID, entity Entity, amount float64) (*Ledger, error)
}
