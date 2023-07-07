package repos

import (
	"context"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/ent"
	"github.com/tauki/invoiceexchange/ent/bid"
	"github.com/tauki/invoiceexchange/ent/investor"
	"github.com/tauki/invoiceexchange/ent/invoice"
	domain "github.com/tauki/invoiceexchange/internal/bid"
	"github.com/tauki/invoiceexchange/internal/errors"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

type BidRepo struct {
	client *ent.Client
}

var _ domain.Repository = (*BidRepo)(nil)

func NewBidRepo(client *ent.Client) *BidRepo {
	return &BidRepo{
		client: client,
	}
}

func (r *BidRepo) CreateBid(ctx context.Context, b *domain.Bid, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	_, err := client.Bid.
		Create().
		SetID(b.ID).
		SetStatus(bid.Status(b.Status)).
		SetAmount(b.Amount).
		SetInvoiceID(b.InvoiceID).
		SetInvestorID(b.InvestorID).
		SetCreatedAt(b.CreatedAt).
		SetUpdatedAt(b.UpdatedAt).
		Save(ctx)
	if err != nil {
		return errors.NewInfrastructureError("failed to create bid", err)
	}
	return nil
}

func (r *BidRepo) GetBidByID(ctx context.Context, id uuid.UUID) (*domain.Bid, error) {
	e, err := r.client.Bid.
		Query().
		WithInvoice().
		WithInvestor().
		Where(bid.ID(id)).
		First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if e == nil {
		return nil, nil
	}

	return toDomainBid(e), nil
}

func (r *BidRepo) UpdateBid(ctx context.Context, b *domain.Bid, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	_, err := client.Bid.
		UpdateOneID(b.ID).
		SetStatus(bid.Status(b.Status)).
		SetAmount(b.Amount).
		SetAcceptedAmount(b.AcceptedAmount).
		SetUpdatedAt(b.UpdatedAt).
		Save(ctx)

	if err != nil {
		return errors.NewInfrastructureError("failed to update bid", err)
	}
	return nil
}

func (r *BidRepo) DeleteBid(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	err := client.Bid.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return errors.NewInfrastructureError("failed to delete bid", err)
	}
	return nil
}

func (r *BidRepo) ListBidsByInvoiceID(ctx context.Context, invoiceID uuid.UUID, opts ...unitofwork.Option) ([]*domain.Bid, error) {
	client := defaultClient(opts, r.client)
	eb, err := client.Bid.Query().
		Where(bid.HasInvoiceWith(invoice.IDEQ(invoiceID))).
		WithInvoice().
		WithInvestor().
		Order(bid.ByCreatedAt()).
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.NewInfrastructureError("failed to list bids", err)
	}

	if eb == nil {
		return nil, nil
	}

	bids := make([]*domain.Bid, len(eb))
	for i, e := range eb {
		bids[i] = toDomainBid(e)
	}

	return bids, nil
}

func (r *BidRepo) ListBidsByInvestorID(ctx context.Context, investorID uuid.UUID) ([]*domain.Bid, error) {
	eb, err := r.client.Bid.Query().
		Where(bid.HasInvestorWith(investor.ID(investorID))).
		Order(bid.ByCreatedAt()).
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.NewInfrastructureError("failed to list bids", err)
	}

	if eb == nil {
		return nil, nil
	}

	bids := make([]*domain.Bid, len(eb))
	for i, e := range eb {
		bids[i] = toDomainBid(e)
	}

	return bids, nil
}

func toDomainBid(e *ent.Bid) *domain.Bid {
	b := &domain.Bid{
		ID:             e.ID,
		Status:         domain.Status(e.Status),
		Amount:         e.Amount,
		AcceptedAmount: e.AcceptedAmount,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}

	if e.Edges.Invoice != nil {
		b.InvoiceID = e.Edges.Invoice.ID
	}

	if e.Edges.Investor != nil {
		b.InvestorID = e.Edges.Investor.ID
	}

	return b
}
