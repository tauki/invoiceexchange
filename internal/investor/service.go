package investor

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	CreateInvestor(ctx context.Context, investor *Investor) (*Investor, error)
	GetInvestorByID(ctx context.Context, id uuid.UUID) (*Investor, error)
	UpdateInvestor(ctx context.Context, investor *Investor) (*Investor, error)
	DeleteInvestor(ctx context.Context, id uuid.UUID) error
	ListInvestors(ctx context.Context) ([]*Investor, error)
}
