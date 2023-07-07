package repos

import (
	"context"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/ent"
	"github.com/tauki/invoiceexchange/ent/balance"
	domain "github.com/tauki/invoiceexchange/internal/balance"
	"github.com/tauki/invoiceexchange/internal/errors"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

type BalanceRepo struct {
	client *ent.Client
}

var _ domain.Repository = (*BalanceRepo)(nil)

func NewBalanceRepo(client *ent.Client) *BalanceRepo {
	return &BalanceRepo{
		client: client,
	}
}

func (r *BalanceRepo) GetBalance(ctx context.Context, entityID uuid.UUID, opts ...unitofwork.Option) (*domain.Balance, error) {
	client := defaultClient(opts, r.client)
	b, err := client.Balance.Query().Where(balance.EntityIDEQ(entityID)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if b == nil {
		return nil, nil
	}

	return toDomainBalance(b), nil
}

func (r *BalanceRepo) UpdateBalance(ctx context.Context, balance *domain.Balance, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	_, err := client.Balance.
		UpdateOneID(balance.ID).
		SetAvailableAmount(balance.AvailableAmount).
		SetTotalAmount(balance.TotalAmount).
		SetUpdatedAt(balance.UpdatedAt).
		Save(ctx)
	if err != nil {
		return errors.NewInfrastructureError("failed to update balance", err)
	}
	return nil
}

func (r *BalanceRepo) CreateBalance(ctx context.Context, balance *domain.Balance, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	_, err := client.Balance.
		Create().
		SetID(balance.ID).
		SetAvailableAmount(balance.AvailableAmount).
		SetTotalAmount(balance.TotalAmount).
		SetEntityID(balance.EntityID).
		SetCreatedAt(balance.CreatedAt).
		SetUpdatedAt(balance.UpdatedAt).
		Save(ctx)
	if err != nil {
		return errors.NewInfrastructureError("failed to create balance", err)
	}
	return nil
}

func toDomainBalance(b *ent.Balance) *domain.Balance {
	return &domain.Balance{
		ID:              b.ID,
		EntityID:        b.EntityID,
		AvailableAmount: b.AvailableAmount,
		TotalAmount:     b.TotalAmount,
		CreatedAt:       b.CreatedAt,
		UpdatedAt:       b.UpdatedAt,
	}
}
