package issuer

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	CreateIssuer(ctx context.Context, issuer *Issuer) (*Issuer, error)
	GetIssuerByID(ctx context.Context, id uuid.UUID) (*Issuer, error)
	UpdateIssuer(ctx context.Context, issuer *Issuer) (*Issuer, error)
	DeleteIssuer(ctx context.Context, id uuid.UUID) error
	ListIssuers(ctx context.Context) ([]*Issuer, error)
}
