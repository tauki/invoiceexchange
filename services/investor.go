package services

import (
	"context"
	"github.com/tauki/invoiceexchange/internal/errors"

	"github.com/google/uuid"
	"go.uber.org/zap"

	domain "github.com/tauki/invoiceexchange/internal/investor"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
	"github.com/tauki/invoiceexchange/router/middlewares/logging"
)

type investorService struct {
	repo       domain.Repository
	uowFactory unitofwork.Factory
}

var _ domain.Service = (*investorService)(nil)

func NewInvestorService(repo domain.Repository, uowFactory unitofwork.Factory) domain.Service {
	return &investorService{
		repo:       repo,
		uowFactory: uowFactory,
	}
}

func (s *investorService) CreateInvestor(ctx context.Context, investor *domain.Investor) (*domain.Investor, error) {
	log := logging.GetZap(ctx)
	uow, err := s.uowFactory.New(ctx)
	defer uow.RollbackUnlessCommitted()
	if err != nil {
		return nil, err
	}

	if err = s.repo.CreateInvestor(ctx, investor, unitofwork.With(uow)); err != nil {
		log.Error("error calling CreateInvestor", zap.Error(err))
		return nil, err
	}

	if err = uow.Commit(); err != nil {
		return nil, err
	}

	return investor, nil
}

func (s *investorService) GetInvestorByID(ctx context.Context, id uuid.UUID) (*domain.Investor, error) {
	log := logging.GetZap(ctx)
	inv, err := s.repo.GetInvestorByID(ctx, id)
	if err != nil {
		log.Error("error calling GetInvestorByID", zap.Error(err))
		return nil, err
	}

	if inv == nil {
		return nil, errors.NewNotFoundError("investor not found")
	}

	return inv, nil
}

func (s *investorService) UpdateInvestor(ctx context.Context, investor *domain.Investor) (*domain.Investor, error) {
	log := logging.GetZap(ctx)
	uow, err := s.uowFactory.New(ctx)
	defer uow.RollbackUnlessCommitted()
	if err != nil {
		return nil, err
	}

	if err = s.repo.UpdateInvestor(ctx, investor, unitofwork.With(uow)); err != nil {
		log.Error("error calling UpdateInvestor", zap.Error(err))
		return nil, err
	}

	if err = uow.Commit(); err != nil {
		return nil, err
	}

	return investor, nil
}

func (s *investorService) DeleteInvestor(ctx context.Context, id uuid.UUID) error {
	log := logging.GetZap(ctx)
	uow, err := s.uowFactory.New(ctx)
	defer uow.RollbackUnlessCommitted()
	if err != nil {
		return err
	}

	if err = s.repo.DeleteInvestor(ctx, id, unitofwork.With(uow)); err != nil {
		log.Error("error calling DeleteInvestor", zap.Error(err))
		return err
	}

	if err = uow.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *investorService) ListInvestors(ctx context.Context) ([]*domain.Investor, error) {
	log := logging.GetZap(ctx)
	invs, err := s.repo.ListInvestors(ctx)
	if err != nil {
		log.Error("error calling ListInvestors", zap.Error(err))
		return nil, err
	}

	if invs == nil {
		return nil, errors.NewNotFoundError("investors not found")
	}

	return invs, nil
}
