package invoice

import (
	"context"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

type Repository interface {
	CreateInvoice(ctx context.Context, inv *Invoice, opts ...unitofwork.Option) error
	GetInvoiceByID(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) (*Invoice, error)
	UpdateInvoice(ctx context.Context, inv *Invoice, opts ...unitofwork.Option) error
	DeleteInvoice(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) error
	ListInvoices(ctx context.Context) ([]*Invoice, error)
	GetInvoiceItems(ctx context.Context, invoiceID uuid.UUID) ([]*Item, error)
}

type ItemRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*Item, error)
	List(ctx context.Context) ([]*Item, error)
	Create(ctx context.Context, invoiceItem *Item) (*Item, error)
	Update(ctx context.Context, invoiceItem *Item) (*Item, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
