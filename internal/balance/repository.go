package balance

import (
	"context"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

type Repository interface {
	CreateBalance(ctx context.Context, balance *Balance, opts ...unitofwork.Option) error
	GetBalance(ctx context.Context, entityID uuid.UUID, opts ...unitofwork.Option) (*Balance, error)
	UpdateBalance(ctx context.Context, balance *Balance, opts ...unitofwork.Option) error
}
