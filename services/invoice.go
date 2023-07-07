package services

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/tauki/invoiceexchange/eventhandler"
	"github.com/tauki/invoiceexchange/internal/errors"
	domain "github.com/tauki/invoiceexchange/internal/invoice"
	"github.com/tauki/invoiceexchange/internal/issuer"
	"github.com/tauki/invoiceexchange/internal/services/eventbus"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
	"github.com/tauki/invoiceexchange/router/middlewares/logging"
)

type invoiceService struct {
	repo       domain.Repository
	issuerRepo issuer.Repository
	uowFactory unitofwork.Factory
	eventBus   eventbus.Publisher
}

func NewInvoiceService(
	repo domain.Repository,
	issuerRepo issuer.Repository,
	uowFactory unitofwork.Factory,
	publisher eventbus.Publisher,
) domain.Service {
	return &invoiceService{
		repo:       repo,
		issuerRepo: issuerRepo,
		uowFactory: uowFactory,
		eventBus:   publisher,
	}
}

// CreateInvoice TODO: save items
func (s *invoiceService) CreateInvoice(
	ctx context.Context,
	inv *domain.Invoice,
) (*domain.Invoice, error) {
	log := logging.GetZap(ctx)

	iss, err := s.issuerRepo.GetIssuerByID(ctx, inv.IssuerID)
	if err != nil {
		log.Error("error calling GetIssuerByID", zap.Error(err))
		return nil, err
	}

	if iss == nil {
		return nil, errors.NewValidationError("issuer does not exist", "invoice.IssuerID")
	}

	uow, err := s.uowFactory.New(ctx)
	if err != nil {
		return nil, err
	}
	defer uow.RollbackUnlessCommitted()

	err = s.repo.CreateInvoice(ctx, inv, unitofwork.With(uow))
	if err != nil {
		log.Error("error calling CreateInvoice", zap.Error(err))
		return nil, err
	}

	err = uow.Commit()
	if err != nil {
		return nil, err
	}

	return inv, nil
}

func (s *invoiceService) GetInvoiceByID(ctx context.Context, id uuid.UUID) (*domain.Invoice, error) {
	log := logging.GetZap(ctx)
	inv, err := s.repo.GetInvoiceByID(ctx, id)
	if err != nil {
		log.Error("error calling GetInvoiceByID", zap.Error(err))
		return nil, err
	}

	if inv == nil {
		return nil, errors.NewNotFoundError("invoice not found")
	}

	items, err := s.repo.GetInvoiceItems(ctx, id)
	if err != nil {
		log.Error("error calling GetInvoiceItems", zap.Error(err))
		return nil, err
	}
	inv.Items = items

	return inv, nil
}

func (s *invoiceService) UpdateInvoice(ctx context.Context, inv *domain.Invoice) (*domain.Invoice, error) {
	return nil, errors.NewInfrastructureError("not implemented", nil)
}

func (s *invoiceService) DeleteInvoice(ctx context.Context, id uuid.UUID) error {
	return errors.NewInfrastructureError("not implemented", nil)
}

func (s *invoiceService) ListInvoices(ctx context.Context) ([]*domain.Invoice, error) {
	invoices, err := s.repo.ListInvoices(ctx)
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (s *invoiceService) ApproveTrade(ctx context.Context, id uuid.UUID) (*domain.Invoice, error) {
	log := logging.GetZap(ctx)
	uow, err := s.uowFactory.New(ctx)
	if err != nil {
		log.Error("error calling New", zap.Error(err))
		return nil, err
	}
	defer uow.RollbackUnlessCommitted()

	invoice, err := s.repo.GetInvoiceByID(ctx, id, unitofwork.With(uow))
	if err != nil {
		log.Error("error calling GetInvoiceByID", zap.Error(err))
		return nil, err
	}

	if invoice == nil {
		return nil, errors.NewNotFoundError("invoice not found")
	}

	if !invoice.IsLocked {
		return nil, errors.NewPermissionDeniedError("trade cannot be approved, invoice is not locked", nil)
	}

	if invoice.IsApproved {
		return nil, errors.NewPermissionDeniedError("trade cannot be approved, invoice is already approved", nil)
	}

	invoice.IsApproved = true
	err = s.repo.UpdateInvoice(ctx, invoice, unitofwork.With(uow))
	if err != nil {
		log.Error("error calling UpdateInvoice", zap.Error(err))
		return nil, err
	}

	if err = uow.Commit(); err != nil {
		return nil, err
	}

	s.eventBus.Publish(eventhandler.InvoiceApprovedEvent, invoice.ID.String())

	return invoice, nil
}
