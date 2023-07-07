package bid

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	CreateBid(ctx context.Context, bid *Bid) (*Bid, error)
	GetBidByID(ctx context.Context, id uuid.UUID) (*Bid, error)
	UpdateBid(ctx context.Context, bid *Bid) (*Bid, error)
	DeleteBid(ctx context.Context, id uuid.UUID) error
	ListBidsByInvoiceID(ctx context.Context, invoiceID uuid.UUID) ([]*Bid, error)
	ListBidsByInvestorID(ctx context.Context, investorID uuid.UUID) ([]*Bid, error)
}
