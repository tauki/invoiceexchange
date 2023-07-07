package ledger

import (
	"context"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

type Repository interface {
	AddLedgerEntry(ctx context.Context, ledger *Ledger, opts ...unitofwork.Option) error
	GetLedgerEntriesByInvoiceID(ctx context.Context, invoiceID uuid.UUID, opts ...unitofwork.Option) ([]*Ledger, error)
}
