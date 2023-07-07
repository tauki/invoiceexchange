package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tauki/invoiceexchange/internal/investor"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
	"testing"
)

func TestInvestorService_CreateInvestor(t *testing.T) {
	mockRepo := new(investor.MockRepository)
	uf := new(unitofwork.MockFactory)
	w := new(unitofwork.MockWork)

	service := NewInvestorService(mockRepo, uf)

	ctx := context.Background()

	investorData := &investor.Investor{
		ID:   uuid.New(),
		Name: "Test Investor",
	}

	uf.On("New", mock.AnythingOfType("*context.emptyCtx")).Return(w, nil)
	mockRepo.On("CreateInvestor", ctx, investorData, mock.AnythingOfType("unitofwork.Option")).Return(nil)
	w.On("Commit").Return(nil)
	w.On("RollbackUnlessCommitted").Return(nil)

	_, err := service.CreateInvestor(ctx, investorData)
	require.NoError(t, err)

	mockRepo.AssertExpectations(t)
	uf.AssertExpectations(t)
	w.AssertExpectations(t)
}

func TestInvestorService_GetInvestorByID(t *testing.T) {
	mockRepo := new(investor.MockRepository)

	service := NewInvestorService(mockRepo, nil)

	ctx := context.Background()

	investorID := uuid.New()

	mockRepo.On("GetInvestorByID", ctx, investorID).Return(&investor.Investor{
		ID:   investorID,
		Name: "Test Investor",
	}, nil)

	_, err := service.GetInvestorByID(ctx, investorID)
	require.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestInvestorService_UpdateInvestor(t *testing.T) {
	mockRepo := new(investor.MockRepository)
	uf := new(unitofwork.MockFactory)
	w := new(unitofwork.MockWork)

	service := NewInvestorService(mockRepo, uf)

	ctx := context.Background()

	investorData := &investor.Investor{
		ID:   uuid.New(),
		Name: "Updated Investor",
	}

	uf.On("New", mock.AnythingOfType("*context.emptyCtx")).Return(w, nil)
	mockRepo.On("UpdateInvestor", ctx, investorData, mock.AnythingOfType("unitofwork.Option")).Return(nil)
	w.On("Commit").Return(nil)
	w.On("RollbackUnlessCommitted").Return(nil)

	_, err := service.UpdateInvestor(ctx, investorData)
	require.NoError(t, err)

	mockRepo.AssertExpectations(t)
	uf.AssertExpectations(t)
	w.AssertExpectations(t)
}

func TestInvestorService_DeleteInvestor(t *testing.T) {
	mockRepo := new(investor.MockRepository)
	uf := new(unitofwork.MockFactory)
	w := new(unitofwork.MockWork)

	service := NewInvestorService(mockRepo, uf)

	ctx := context.Background()

	investorID := uuid.New()

	uf.On("New", mock.AnythingOfType("*context.emptyCtx")).Return(w, nil)
	mockRepo.On("DeleteInvestor", ctx, investorID, mock.AnythingOfType("unitofwork.Option")).Return(nil)
	w.On("Commit").Return(nil)
	w.On("RollbackUnlessCommitted").Return(nil)

	err := service.DeleteInvestor(ctx, investorID)
	require.NoError(t, err)

	mockRepo.AssertExpectations(t)
	uf.AssertExpectations(t)
	w.AssertExpectations(t)
}

func TestInvestorService_ListInvestors(t *testing.T) {
	mockRepo := new(investor.MockRepository)

	service := NewInvestorService(mockRepo, nil)

	ctx := context.Background()

	mockRepo.On("ListInvestors", ctx).Return([]*investor.Investor{
		{
			ID:   uuid.New(),
			Name: "Investor 1",
		},
		{
			ID:   uuid.New(),
			Name: "Investor 2",
		},
	}, nil)

	_, err := service.ListInvestors(ctx)
	require.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
