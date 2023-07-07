package bid

import (
	"context"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

type Repository interface {
	CreateBid(ctx context.Context, b *Bid, opts ...unitofwork.Option) error
	GetBidByID(ctx context.Context, id uuid.UUID) (*Bid, error)
	UpdateBid(ctx context.Context, b *Bid, opts ...unitofwork.Option) error
	DeleteBid(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) error
	ListBidsByInvoiceID(ctx context.Context, invoiceID uuid.UUID, opts ...unitofwork.Option) ([]*Bid, error)
	ListBidsByInvestorID(ctx context.Context, investorID uuid.UUID) ([]*Bid, error)
}
