package repos

import (
	"context"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/ent"
	"github.com/tauki/invoiceexchange/ent/issuer"
	"github.com/tauki/invoiceexchange/internal/balance"
	"github.com/tauki/invoiceexchange/internal/errors"
	domain "github.com/tauki/invoiceexchange/internal/issuer"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

type IssuerRepo struct {
	client *ent.Client
}

func NewIssuerRepo(client *ent.Client) *IssuerRepo {
	return &IssuerRepo{
		client: client,
	}
}

var _ domain.Repository = (*IssuerRepo)(nil)

func (r *IssuerRepo) CreateIssuer(ctx context.Context, issuer *domain.Issuer, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	_, err := client.Balance.Create().
		SetID(issuer.Balance.ID).
		SetEntityID(issuer.ID).
		SetTotalAmount(issuer.Balance.TotalAmount).
		SetAvailableAmount(issuer.Balance.AvailableAmount).
		SetCreatedAt(issuer.Balance.CreatedAt).
		SetUpdatedAt(issuer.Balance.UpdatedAt).
		Save(ctx)
	_, err = client.Issuer.
		Create().
		SetID(issuer.ID).
		SetJoinedAt(issuer.JoinedAt).
		SetBalanceID(issuer.Balance.ID).
		SetName(issuer.Name).
		Save(ctx)

	if err != nil {
		return errors.NewInfrastructureError("failed to create issuer", err)
	}
	return nil
}

func (r *IssuerRepo) GetIssuerByID(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) (*domain.Issuer, error) {
	client := defaultClient(opts, r.client)
	iss, err := client.Issuer.Query().WithBalance().Where(issuer.IDEQ(id)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.NewInfrastructureError("failed to get issuer", err)
	}

	if iss == nil {
		return nil, nil
	}

	return toDomainIssuer(iss), nil
}

func (r *IssuerRepo) UpdateIssuer(ctx context.Context, issuer *domain.Issuer, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	_, err := client.Issuer.
		UpdateOneID(issuer.ID).
		SetName(issuer.Name).
		Save(ctx)

	if err != nil {
		return errors.NewInfrastructureError("failed to update issuer", err)
	}

	if issuer.Balance != nil {
		_, err = client.Balance.
			UpdateOneID(issuer.Balance.ID).
			SetTotalAmount(issuer.Balance.TotalAmount).
			SetAvailableAmount(issuer.Balance.AvailableAmount).
			SetUpdatedAt(issuer.Balance.UpdatedAt).
			Save(ctx)

		if err != nil {
			return errors.NewInfrastructureError("failed to update issuer", err)
		}
	}

	return nil
}

func (r *IssuerRepo) DeleteIssuer(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) error {
	err := r.client.Issuer.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return errors.NewInfrastructureError("failed to delete issuer", err)
	}

	return nil
}

func (r *IssuerRepo) ListIssuers(ctx context.Context) ([]*domain.Issuer, error) {
	issuers, err := r.client.Issuer.Query().WithBalance().All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.NewInfrastructureError("failed to list issuers", err)
	}

	if issuers == nil {
		return nil, nil
	}

	var domainIssuers []*domain.Issuer
	for _, iss := range issuers {
		domainIssuers = append(domainIssuers, toDomainIssuer(iss))
	}

	return domainIssuers, nil
}

func toDomainIssuer(iss *ent.Issuer) *domain.Issuer {
	var bal *balance.Balance
	if iss.Edges.Balance != nil {
		bal = toDomainBalance(iss.Edges.Balance)
	}
	return &domain.Issuer{
		ID:       iss.ID,
		Name:     iss.Name,
		Balance:  bal,
		JoinedAt: iss.JoinedAt,
	}
}
