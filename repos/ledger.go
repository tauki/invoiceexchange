package repos

import (
	"context"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/ent"
	"github.com/tauki/invoiceexchange/ent/ledger"
	"github.com/tauki/invoiceexchange/internal/errors"
	domain "github.com/tauki/invoiceexchange/internal/ledger"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

type LedgerRepo struct {
	client *ent.Client
}

var _ domain.Repository = (*LedgerRepo)(nil)

func NewLedgerRepo(client *ent.Client) *LedgerRepo {
	return &LedgerRepo{
		client: client,
	}
}

func (r *LedgerRepo) AddLedgerEntry(ctx context.Context, l *domain.Ledger, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	_, err := client.Ledger.
		Create().
		SetID(l.ID).
		SetInvoiceID(l.InvoiceID).
		SetEntity(ledger.Entity(l.Entity)).
		SetEntityID(l.EntityID).
		SetAmount(l.Amount).
		SetCreatedAt(l.CreatedAt).
		SetUpdatedAt(l.UpdatedAt).
		Save(ctx)
	if err != nil {
		return errors.NewInfrastructureError("failed to create ledger entry", err)
	}
	return nil
}

func (r *LedgerRepo) GetLedgerEntriesByInvoiceID(ctx context.Context, invoiceID uuid.UUID, opts ...unitofwork.Option) ([]*domain.Ledger, error) {
	client := defaultClient(opts, r.client)
	entries, err := client.Ledger.
		Query().
		Where(ledger.InvoiceIDEQ(invoiceID)).
		All(ctx)
	if err != nil {
		return nil, errors.NewInfrastructureError("failed to get ledger entries", err)
	}

	return toDomainLedger(entries), nil
}

func toDomainLedger(entries []*ent.Ledger) []*domain.Ledger {
	var domainEntries []*domain.Ledger
	for _, e := range entries {
		domainEntries = append(domainEntries, toDomainLedgerEntry(e))
	}
	return domainEntries
}

func toDomainLedgerEntry(e *ent.Ledger) *domain.Ledger {
	return &domain.Ledger{
		ID:        e.ID,
		InvoiceID: e.InvoiceID,
		Entity:    domain.Entity(e.Entity),
		EntityID:  e.EntityID,
		Amount:    e.Amount,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
