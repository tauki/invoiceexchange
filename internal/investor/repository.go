package investor

import (
	"context"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

type Repository interface {
	CreateInvestor(ctx context.Context, investor *Investor, opts ...unitofwork.Option) error
	GetInvestorByID(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) (*Investor, error)
	UpdateInvestor(ctx context.Context, investor *Investor, opts ...unitofwork.Option) error
	DeleteInvestor(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) error
	ListInvestors(ctx context.Context) ([]*Investor, error)
}
