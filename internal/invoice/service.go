package invoice

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	CreateInvoice(ctx context.Context, inv *Invoice) (*Invoice, error)
	GetInvoiceByID(ctx context.Context, id uuid.UUID) (*Invoice, error)
	UpdateInvoice(ctx context.Context, inv *Invoice) (*Invoice, error)
	DeleteInvoice(ctx context.Context, id uuid.UUID) error
	ListInvoices(ctx context.Context) ([]*Invoice, error)
	ApproveTrade(ctx context.Context, id uuid.UUID) (*Invoice, error)
}
