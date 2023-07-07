package repos

import (
	"context"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/ent"
	"github.com/tauki/invoiceexchange/ent/investor"
	"github.com/tauki/invoiceexchange/internal/balance"
	"github.com/tauki/invoiceexchange/internal/errors"
	domain "github.com/tauki/invoiceexchange/internal/investor"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

type InvestorRepo struct {
	client *ent.Client
}

var _ domain.Repository = (*InvestorRepo)(nil)

func NewInvestorRepo(client *ent.Client) *InvestorRepo {
	return &InvestorRepo{
		client: client,
	}
}

func (r *InvestorRepo) CreateInvestor(ctx context.Context, inv *domain.Investor, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	_, err := client.Balance.
		Create().
		SetID(inv.Balance.ID).
		SetEntityID(inv.ID).
		SetTotalAmount(inv.Balance.TotalAmount).
		SetAvailableAmount(inv.Balance.AvailableAmount).
		SetCreatedAt(inv.Balance.CreatedAt).
		SetUpdatedAt(inv.Balance.UpdatedAt).
		Save(ctx)

	if err != nil {
		return err
	}

	_, err = client.Investor.
		Create().
		SetID(inv.ID).
		SetName(inv.Name).
		SetBalanceID(inv.Balance.ID).
		SetJoinedAt(inv.JoinedAt).
		Save(ctx)

	if err != nil {
		return err
	}
	return nil
}

func (r *InvestorRepo) GetInvestorByID(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) (*domain.Investor, error) {
	client := defaultClient(opts, r.client)
	inv, err := client.Investor.Query().WithBalance().Where(investor.ID(id)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if inv == nil {
		return nil, nil
	}

	return toDomainInvestor(inv), nil
}

func (r *InvestorRepo) UpdateInvestor(ctx context.Context, inv *domain.Investor, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	_, err := client.Investor.
		UpdateOneID(inv.ID).
		SetName(inv.Name).
		Save(ctx)

	if err != nil {
		return errors.NewInfrastructureError("failed to update investor", err)
	}

	if inv.Balance != nil {
		_, err = client.Balance.
			UpdateOneID(inv.Balance.ID).
			SetTotalAmount(inv.Balance.TotalAmount).
			SetAvailableAmount(inv.Balance.AvailableAmount).
			SetUpdatedAt(inv.Balance.UpdatedAt).
			Save(ctx)
		if err != nil {
			return errors.NewInfrastructureError("failed to update investor", err)
		}
	}

	return nil
}

func (r *InvestorRepo) DeleteInvestor(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	err := client.Investor.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return errors.NewInfrastructureError("failed to delete investor", err)
	}

	return nil
}

func (r *InvestorRepo) ListInvestors(ctx context.Context) ([]*domain.Investor, error) {
	investors, err := r.client.Investor.Query().WithBalance().All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if investors == nil {
		return nil, nil
	}

	var domainInvestors []*domain.Investor
	for _, inv := range investors {
		domainInvestors = append(domainInvestors, toDomainInvestor(inv))
	}

	return domainInvestors, nil
}

func toDomainInvestor(inv *ent.Investor) *domain.Investor {
	var bal *balance.Balance
	if inv.Edges.Balance != nil {
		bal = toDomainBalance(inv.Edges.Balance)
	}
	return &domain.Investor{
		ID:       inv.ID,
		Name:     inv.Name,
		JoinedAt: inv.JoinedAt,
		Balance:  bal,
	}
}
