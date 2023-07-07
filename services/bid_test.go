package services_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tauki/invoiceexchange/internal/balance"
	"github.com/tauki/invoiceexchange/internal/bid"
	"github.com/tauki/invoiceexchange/internal/investor"
	"github.com/tauki/invoiceexchange/internal/invoice"
	"github.com/tauki/invoiceexchange/internal/services/eventbus"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
	"github.com/tauki/invoiceexchange/services"
	"testing"
)

func TestCreateBid(t *testing.T) {
	mockRepo := new(bid.MockRepository)
	uf := new(unitofwork.MockFactory)
	w := new(unitofwork.MockWork)
	mockInvoiceRepo := new(invoice.MockRepository)
	mockInvestorRepo := new(investor.MockRepository)
	mockBalanceRepo := new(balance.MockRepository)

	service := services.NewBidService(
		mockRepo,
		uf,
		mockInvoiceRepo,
		mockInvestorRepo,
		mockBalanceRepo,
		&eventbus.NoOpPublisher{},
	)

	b := &bid.Bid{Amount: 500, InvestorID: uuid.New(), InvoiceID: uuid.New()}
	ctx := context.Background()

	uf.On("New", mock.AnythingOfType("*context.emptyCtx")).Return(w, nil)

	mockInvestorRepo.On("GetInvestorByID", mock.AnythingOfType("*context.emptyCtx"), b.InvestorID).Return(&investor.Investor{Balance: &balance.Balance{AvailableAmount: 1000}}, nil)
	mockInvoiceRepo.On("GetInvoiceByID", mock.AnythingOfType("*context.emptyCtx"), b.InvoiceID).Return(&invoice.Invoice{}, nil)
	mockRepo.On("CreateBid", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*bid.Bid"), mock.AnythingOfType("unitofwork.Option")).Return(nil)
	mockBalanceRepo.On("UpdateBalance", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*balance.Balance"), mock.AnythingOfType("unitofwork.Option")).Return(nil)

	w.On("Commit").Return(nil)
	w.On("RollbackUnlessCommitted").Return(nil)

	_, err := service.CreateBid(ctx, b)

	require.NoError(t, err)

	mockRepo.AssertExpectations(t)
	uf.AssertExpectations(t)
	mockInvoiceRepo.AssertExpectations(t)
	mockInvestorRepo.AssertExpectations(t)
	mockBalanceRepo.AssertExpectations(t)
}

func TestCreateBid_InvalidAmount(t *testing.T) {
	mockRepo := new(bid.MockRepository)
	mockInvoiceRepo := new(invoice.MockRepository)
	mockInvestorRepo := new(investor.MockRepository)
	mockBalanceRepo := new(balance.MockRepository)

	service := services.NewBidService(
		mockRepo,
		nil,
		mockInvoiceRepo,
		mockInvestorRepo,
		mockBalanceRepo,
		&eventbus.NoOpPublisher{},
	)

	b := &bid.Bid{Amount: 500, InvestorID: uuid.New(), InvoiceID: uuid.New()}
	ctx := context.Background()

	mockInvestorRepo.On("GetInvestorByID", mock.AnythingOfType("*context.emptyCtx"), b.InvestorID).Return(&investor.Investor{Balance: &balance.Balance{AvailableAmount: 100}}, nil)

	_, err := service.CreateBid(ctx, b)

	require.EqualError(t, err, "Error: insufficient balance")

	mockRepo.AssertExpectations(t)
	mockInvoiceRepo.AssertExpectations(t)
	mockInvestorRepo.AssertExpectations(t)
	mockBalanceRepo.AssertExpectations(t)
}

func TestCreateBid_InvalidInvoice(t *testing.T) {
	mockRepo := new(bid.MockRepository)
	mockInvoiceRepo := new(invoice.MockRepository)
	mockInvestorRepo := new(investor.MockRepository)
	mockBalanceRepo := new(balance.MockRepository)

	service := services.NewBidService(
		mockRepo,
		nil,
		mockInvoiceRepo,
		mockInvestorRepo,
		mockBalanceRepo,
		&eventbus.NoOpPublisher{},
	)

	b := &bid.Bid{Amount: 500, InvestorID: uuid.New(), InvoiceID: uuid.New()}
	ctx := context.Background()

	mockInvestorRepo.On("GetInvestorByID", mock.AnythingOfType("*context.emptyCtx"), b.InvestorID).Return(&investor.Investor{Balance: &balance.Balance{AvailableAmount: 1000}}, nil)
	mockInvoiceRepo.On("GetInvoiceByID", mock.AnythingOfType("*context.emptyCtx"), b.InvoiceID).Return(nil, nil)

	_, err := service.CreateBid(ctx, b)

	require.EqualError(t, err, "Error: invoice not found")

	mockRepo.AssertExpectations(t)
	mockInvoiceRepo.AssertExpectations(t)
	mockInvestorRepo.AssertExpectations(t)
	mockBalanceRepo.AssertExpectations(t)
}

func TestCreateBid_InvalidInvestor(t *testing.T) {
	mockRepo := new(bid.MockRepository)
	mockInvoiceRepo := new(invoice.MockRepository)
	mockInvestorRepo := new(investor.MockRepository)
	mockBalanceRepo := new(balance.MockRepository)

	service := services.NewBidService(
		mockRepo,
		nil,
		mockInvoiceRepo,
		mockInvestorRepo,
		mockBalanceRepo,
		&eventbus.NoOpPublisher{},
	)

	b := &bid.Bid{Amount: 500, InvestorID: uuid.New(), InvoiceID: uuid.New()}
	ctx := context.Background()

	mockInvestorRepo.On("GetInvestorByID", mock.AnythingOfType("*context.emptyCtx"), b.InvestorID).Return(nil, nil)

	_, err := service.CreateBid(ctx, b)

	require.EqualError(t, err, "Error: investor not found")

	mockRepo.AssertExpectations(t)
	mockInvoiceRepo.AssertExpectations(t)
	mockInvestorRepo.AssertExpectations(t)
	mockBalanceRepo.AssertExpectations(t)
}

func TestGetBidByID(t *testing.T) {
	mockRepo := new(bid.MockRepository)
	uf := new(unitofwork.MockFactory)

	service := services.NewBidService(mockRepo, uf, nil, nil, nil, nil)

	ctx := context.Background()

	b := &bid.Bid{Amount: 500, InvestorID: uuid.New(), InvoiceID: uuid.New()}
	mockRepo.On("GetBidByID", mock.AnythingOfType("*context.emptyCtx"), b.ID).Return(b, nil)

	returnedBid, err := service.GetBidByID(ctx, b.ID)

	require.NoError(t, err)
	require.Equal(t, b, returnedBid)

	mockRepo.AssertExpectations(t)
}

func TestGetBidByID_NotFound(t *testing.T) {
	mockRepo := new(bid.MockRepository)
	uf := new(unitofwork.MockFactory)

	service := services.NewBidService(mockRepo, uf, nil, nil, nil, nil)

	ctx := context.Background()

	b := &bid.Bid{Amount: 500, InvestorID: uuid.New(), InvoiceID: uuid.New()}
	mockRepo.On("GetBidByID", mock.AnythingOfType("*context.emptyCtx"), b.ID).Return(nil, nil)

	_, err := service.GetBidByID(ctx, b.ID)

	require.EqualError(t, err, "Error: bid not found")

	mockRepo.AssertExpectations(t)
}

func TestBidService_ListBidsByInvoiceID(t *testing.T) {
	mockRepo := new(bid.MockRepository)
	uf := new(unitofwork.MockFactory)

	service := services.NewBidService(mockRepo, uf, nil, nil, nil, nil)

	ctx := context.Background()

	b := &bid.Bid{Amount: 500, InvestorID: uuid.New(), InvoiceID: uuid.New()}
	bids := []*bid.Bid{b}
	mockRepo.On("ListBidsByInvoiceID", mock.AnythingOfType("*context.emptyCtx"), b.InvoiceID).Return(bids, nil)

	returnedBids, err := service.ListBidsByInvoiceID(ctx, b.InvoiceID)

	require.NoError(t, err)
	require.Equal(t, bids, returnedBids)

	mockRepo.AssertExpectations(t)
}

func TestBidService_ListBidsByInvoiceID_NotFound(t *testing.T) {
	mockRepo := new(bid.MockRepository)
	uf := new(unitofwork.MockFactory)

	service := services.NewBidService(mockRepo, uf, nil, nil, nil, nil)

	ctx := context.Background()

	b := &bid.Bid{Amount: 500, InvestorID: uuid.New(), InvoiceID: uuid.New()}
	mockRepo.On("ListBidsByInvoiceID", mock.AnythingOfType("*context.emptyCtx"), b.InvoiceID).Return(nil, nil)

	_, err := service.ListBidsByInvoiceID(ctx, b.InvoiceID)

	require.EqualError(t, err, "Error: no bids found for this invoice")

	mockRepo.AssertExpectations(t)
}

func TestBidService_ListBidsByInvestorID(t *testing.T) {
	mockRepo := new(bid.MockRepository)
	uf := new(unitofwork.MockFactory)

	service := services.NewBidService(mockRepo, uf, nil, nil, nil, nil)

	ctx := context.Background()

	b := &bid.Bid{Amount: 500, InvestorID: uuid.New(), InvoiceID: uuid.New()}
	bids := []*bid.Bid{b}
	mockRepo.On("ListBidsByInvestorID", mock.AnythingOfType("*context.emptyCtx"), b.InvestorID).Return(bids, nil)

	returnedBids, err := service.ListBidsByInvestorID(ctx, b.InvestorID)

	require.NoError(t, err)
	require.Equal(t, bids, returnedBids)

	mockRepo.AssertExpectations(t)
}

func TestBidService_ListBidsByInvestorID_NotFound(t *testing.T) {
	mockRepo := new(bid.MockRepository)
	uf := new(unitofwork.MockFactory)

	service := services.NewBidService(mockRepo, uf, nil, nil, nil, nil)

	ctx := context.Background()

	b := &bid.Bid{Amount: 500, InvestorID: uuid.New(), InvoiceID: uuid.New()}
	mockRepo.On("ListBidsByInvestorID", mock.AnythingOfType("*context.emptyCtx"), b.InvestorID).Return(nil, nil)

	_, err := service.ListBidsByInvestorID(ctx, b.InvestorID)

	require.EqualError(t, err, "Error: no bids found for this investor")

	mockRepo.AssertExpectations(t)
}
