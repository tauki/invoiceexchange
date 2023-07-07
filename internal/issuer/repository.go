package issuer

import (
	"context"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

type Repository interface {
	CreateIssuer(ctx context.Context, issuer *Issuer, opts ...unitofwork.Option) error
	GetIssuerByID(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) (*Issuer, error)
	UpdateIssuer(ctx context.Context, issuer *Issuer, opts ...unitofwork.Option) error
	DeleteIssuer(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) error
	ListIssuers(ctx context.Context) ([]*Issuer, error)
}
