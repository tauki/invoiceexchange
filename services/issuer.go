package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/internal/errors"
	domain "github.com/tauki/invoiceexchange/internal/issuer"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
	"github.com/tauki/invoiceexchange/router/middlewares/logging"
	"go.uber.org/zap"
)

type issuerService struct {
	repo       domain.Repository
	uowFactory unitofwork.Factory
}

var _ domain.Service = (*issuerService)(nil)

func NewIssuerService(repo domain.Repository, uowFactory unitofwork.Factory) domain.Service {
	return &issuerService{
		repo:       repo,
		uowFactory: uowFactory,
	}
}

func (s *issuerService) CreateIssuer(ctx context.Context, issuer *domain.Issuer) (*domain.Issuer, error) {
	log := logging.GetZap(ctx)
	uow, err := s.uowFactory.New(ctx)
	defer uow.RollbackUnlessCommitted()
	if err != nil {
		return nil, err
	}

	if err = s.repo.CreateIssuer(ctx, issuer, unitofwork.With(uow)); err != nil {
		log.Error("error calling CreateBid", zap.Error(err))
		return nil, err
	}

	if err = uow.Commit(); err != nil {
		return nil, err
	}

	return issuer, nil
}

func (s *issuerService) GetIssuerByID(ctx context.Context, id uuid.UUID) (*domain.Issuer, error) {
	log := logging.GetZap(ctx)
	iss, err := s.repo.GetIssuerByID(ctx, id)
	if err != nil {
		log.Error("error calling GetIssuerByID", zap.Error(err))
		return nil, err
	}

	if iss == nil {
		return nil, errors.NewNotFoundError("issuer not found")
	}

	return iss, nil
}

func (s *issuerService) UpdateIssuer(ctx context.Context, issuer *domain.Issuer) (*domain.Issuer, error) {
	return issuer, s.repo.UpdateIssuer(ctx, issuer)
}

func (s *issuerService) DeleteIssuer(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteIssuer(ctx, id)
}

func (s *issuerService) ListIssuers(ctx context.Context) ([]*domain.Issuer, error) {
	log := logging.GetZap(ctx)
	invs, err := s.repo.ListIssuers(ctx)
	if err != nil {
		log.Error("error calling ListIssuers", zap.Error(err))
		return nil, err
	}

	if invs == nil {
		return nil, errors.NewNotFoundError("issuers not found")
	}

	return invs, nil
}
