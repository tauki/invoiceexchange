package services

import (
	"context"
	"github.com/tauki/invoiceexchange/internal/balance"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/tauki/invoiceexchange/eventhandler"
	domain "github.com/tauki/invoiceexchange/internal/bid"
	"github.com/tauki/invoiceexchange/internal/errors"
	"github.com/tauki/invoiceexchange/internal/investor"
	"github.com/tauki/invoiceexchange/internal/invoice"
	"github.com/tauki/invoiceexchange/internal/services/eventbus"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
	"github.com/tauki/invoiceexchange/router/middlewares/logging"
)

type bidService struct {
	repo         domain.Repository
	uowFactory   unitofwork.Factory
	invoiceRepo  invoice.Repository
	investorRepo investor.Repository
	balanceRepo  balance.Repository
	eventbus     eventbus.Publisher
}

var _ domain.Service = (*bidService)(nil)

func NewBidService(
	repo domain.Repository,
	uowFactory unitofwork.Factory,
	invoiceRepo invoice.Repository,
	investorRepo investor.Repository,
	balanceRepo balance.Repository,
	eventbus eventbus.Publisher,
) domain.Service {
	return &bidService{
		repo:         repo,
		uowFactory:   uowFactory,
		invoiceRepo:  invoiceRepo,
		investorRepo: investorRepo,
		balanceRepo:  balanceRepo,
		eventbus:     eventbus,
	}
}

func (s *bidService) CreateBid(ctx context.Context, bid *domain.Bid) (*domain.Bid, error) {
	log := logging.GetZap(ctx)

	invstr, err := s.investorRepo.GetInvestorByID(ctx, bid.InvestorID)
	if err != nil {
		//return nil, errors.Wrap(err, "unable to fetch investor")
		log.Error("error calling GetInvestorByID", zap.Error(err))
		return nil, err
	}

	if invstr == nil {
		return nil, errors.NewNotFoundError("investor not found")
	}

	if invstr.Balance.AvailableAmount < bid.Amount {
		return nil, errors.NewPermissionDeniedError("insufficient balance", nil)
	}

	invc, err := s.invoiceRepo.GetInvoiceByID(ctx, bid.InvoiceID)
	if err != nil {
		log.Error("error calling GetInvoiceByID", zap.Error(err))
		return nil, err
	}

	if invc == nil {
		return nil, errors.NewNotFoundError("invoice not found")
	}

	if invc.IsLocked {
		return nil, errors.NewPermissionDeniedError("invoice is locked", nil)
	}

	uow, err := s.uowFactory.New(ctx)
	if err != nil {
		return nil, err
	}
	defer uow.RollbackUnlessCommitted()

	if err = s.repo.CreateBid(ctx, bid, unitofwork.With(uow)); err != nil {
		log.Error("error calling CreateBid", zap.Error(err))
		return nil, err
	}

	invstr.Balance.AvailableAmount -= bid.Amount
	if err = s.balanceRepo.UpdateBalance(ctx, invstr.Balance, unitofwork.With(uow)); err != nil {
		log.Error("error calling UpdateBalance", zap.Error(err))
		return nil, err
	}

	if err = uow.Commit(); err != nil {
		log.Error("error calling Commit", zap.Error(err))
		return nil, err
	}

	log.Debug("Successfully created bid", zap.String("bidID", bid.ID.String()))
	s.eventbus.Publish(eventhandler.BidCreatedEvent, bid.ID.String())

	return bid, nil
}

func (s *bidService) GetBidByID(ctx context.Context, id uuid.UUID) (*domain.Bid, error) {
	log := logging.GetZap(ctx)
	bid, err := s.repo.GetBidByID(ctx, id)
	if err != nil {
		log.Error("error calling GetBidByID", zap.Error(err))
		return nil, err
	}

	if bid == nil {
		return nil, errors.NewNotFoundError("bid not found")
	}

	return bid, nil
}

func (s *bidService) UpdateBid(ctx context.Context, bid *domain.Bid) (*domain.Bid, error) {
	return nil, errors.NewInfrastructureError("not implemented", nil)
}

func (s *bidService) DeleteBid(ctx context.Context, id uuid.UUID) error {
	return errors.NewInfrastructureError("not implemented", nil)
}

func (s *bidService) ListBidsByInvoiceID(ctx context.Context, invoiceID uuid.UUID) ([]*domain.Bid, error) {
	log := logging.GetZap(ctx)
	bids, err := s.repo.ListBidsByInvoiceID(ctx, invoiceID)
	if err != nil {
		log.Error("error calling ListBidsByInvoiceID", zap.Error(err))
		return nil, err
	}

	if len(bids) == 0 {
		return nil, errors.NewNotFoundError("no bids found for this invoice")
	}
	return bids, nil
}

func (s *bidService) ListBidsByInvestorID(ctx context.Context, investorID uuid.UUID) ([]*domain.Bid, error) {
	log := logging.GetZap(ctx)
	bids, err := s.repo.ListBidsByInvestorID(ctx, investorID)
	if err != nil {
		log.Error("error calling ListBidsByInvestorID", zap.Error(err))
		return nil, err
	}

	if len(bids) == 0 {
		return nil, errors.NewNotFoundError("no bids found for this investor")
	}

	return bids, nil
}
