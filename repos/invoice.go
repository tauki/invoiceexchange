package repos

import (
	"context"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/ent"
	"github.com/tauki/invoiceexchange/ent/invoice"
	"github.com/tauki/invoiceexchange/ent/invoiceitem"
	"github.com/tauki/invoiceexchange/internal/errors"
	domain "github.com/tauki/invoiceexchange/internal/invoice"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

type InvoiceRepo struct {
	client *ent.Client
}

func NewInvoiceRepo(client *ent.Client) *InvoiceRepo {
	return &InvoiceRepo{
		client: client,
	}
}

var _ domain.Repository = (*InvoiceRepo)(nil)

func (r *InvoiceRepo) CreateInvoice(ctx context.Context, inv *domain.Invoice, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	_, err := client.Invoice.
		Create().
		SetID(inv.ID).
		SetStatus(invoice.Status(inv.Status)).
		SetAskingPrice(inv.AskingPrice).
		SetIsLocked(inv.IsLocked).
		SetIsApproved(inv.IsApproved).
		SetInvoiceNumber(inv.InvoiceNumber).
		SetInvoiceDate(inv.InvoiceDate).
		SetDueDate(inv.DueDate).
		SetAmountDue(inv.AmountDue).
		SetCustomerName(inv.CustomerName).
		SetReference(inv.Reference).
		SetCompanyName(inv.CompanyName).
		SetCurrency(inv.Currency).
		SetTotalAmount(inv.TotalAmount).
		SetTotalVat(inv.TotalVAT).
		SetIssuerID(inv.IssuerID).
		Save(ctx)

	if err != nil {
		return errors.NewInfrastructureError("failed to create invoice", err)
	}
	return nil
}

func (r *InvoiceRepo) GetInvoiceByID(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) (*domain.Invoice, error) {
	client := defaultClient(opts, r.client)
	inv, err := client.Invoice.Query().
		WithIssuer().
		WithInvestor().
		WithItems().
		Where(invoice.ID(id)).
		First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.NewInfrastructureError("failed to get invoice", err)
	}

	if inv == nil {
		return nil, nil
	}

	return toDomainInvoice(inv), nil
}

func (r *InvoiceRepo) UpdateInvoice(ctx context.Context, inv *domain.Invoice, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	_, err := client.Invoice.
		UpdateOneID(inv.ID).
		SetAskingPrice(inv.AskingPrice).
		SetIsLocked(inv.IsLocked).
		SetIsApproved(inv.IsApproved).
		SetStatus(invoice.Status(inv.Status)).
		SetInvoiceNumber(inv.InvoiceNumber).
		SetInvoiceDate(inv.InvoiceDate).
		SetDueDate(inv.DueDate).
		SetAmountDue(inv.AmountDue).
		SetCustomerName(inv.CustomerName).
		SetReference(inv.Reference).
		SetCompanyName(inv.CompanyName).
		SetCurrency(inv.Currency).
		SetTotalAmount(inv.TotalAmount).
		AddInvestorIDs(inv.InvestorIDs...).
		SetTotalVat(inv.TotalVAT).
		Save(ctx)
	if err != nil {
		return errors.NewInfrastructureError("failed to update invoice", err)
	}

	return nil
}

func (r *InvoiceRepo) DeleteInvoice(ctx context.Context, id uuid.UUID, opts ...unitofwork.Option) error {
	client := defaultClient(opts, r.client)
	err := client.Invoice.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return errors.NewInfrastructureError("failed to delete invoice", err)
	}

	return nil
}

func (r *InvoiceRepo) ListInvoices(ctx context.Context) ([]*domain.Invoice, error) {
	invoices, err := r.client.Invoice.Query().
		WithIssuer().
		WithInvestor().
		WithItems().
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.NewInfrastructureError("failed to list invoices", err)
	}

	if invoices == nil {
		return nil, nil
	}

	domainInvoices := make([]*domain.Invoice, len(invoices))
	for i, inv := range invoices {
		domainInvoices[i] = toDomainInvoice(inv)
	}

	return domainInvoices, nil
}

func (r *InvoiceRepo) GetInvoiceItems(ctx context.Context, invoiceID uuid.UUID) ([]*domain.Item, error) {
	items, err := r.client.InvoiceItem.Query().Where(
		invoiceitem.HasInvoiceWith(invoice.ID(invoiceID))).All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.NewInfrastructureError("failed to get invoice items", err)
	}

	if items == nil {
		return nil, nil
	}

	domainItems := make([]*domain.Item, len(items))
	for i, item := range items {
		domainItems[i] = toDomainItem(item)
	}

	return domainItems, nil
}

func toDomainInvoice(inv *ent.Invoice) *domain.Invoice {
	var issuerID uuid.UUID
	investorIDs := make([]uuid.UUID, 0)
	items := make([]*domain.Item, 0)
	if inv.Edges.Issuer != nil {
		issuerID = inv.Edges.Issuer.ID
	}
	if inv.Edges.Investor != nil {
		for _, inv := range inv.Edges.Investor {
			investorIDs = append(investorIDs, inv.ID)
		}
	}

	if inv.Edges.Items != nil {
		for _, item := range inv.Edges.Items {
			items = append(items, toDomainItem(item))
		}
	}

	return &domain.Invoice{
		ID:            inv.ID,
		Status:        domain.Status(inv.Status),
		AskingPrice:   inv.AskingPrice,
		IsLocked:      inv.IsLocked,
		IsApproved:    inv.IsApproved,
		InvoiceNumber: inv.InvoiceNumber,
		InvoiceDate:   inv.InvoiceDate,
		DueDate:       inv.DueDate,
		AmountDue:     inv.AmountDue,
		CustomerName:  inv.CustomerName,
		Reference:     inv.Reference,
		CompanyName:   inv.CompanyName,
		Currency:      inv.Currency,
		TotalAmount:   inv.TotalAmount,
		TotalVAT:      inv.TotalVat,
		IssuerID:      issuerID,
		InvestorIDs:   investorIDs,
		Items:         items,
	}
}

func toDomainItem(item *ent.InvoiceItem) *domain.Item {
	return &domain.Item{
		ID:          item.ID,
		Description: item.Description,
		Quantity:    item.Quantity,
		UnitPrice:   item.UnitPrice,
		Amount:      item.Amount,
		VatRate:     item.VatRate,
		VatAmount:   item.VatAmount,
	}
}
