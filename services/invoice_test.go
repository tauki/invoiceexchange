package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tauki/invoiceexchange/internal/invoice"
	"github.com/tauki/invoiceexchange/internal/issuer"
	"github.com/tauki/invoiceexchange/internal/services/eventbus"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
	"testing"
)

func TestCreateInvoice(t *testing.T) {
	mockRepo := new(invoice.MockRepository)
	uf := new(unitofwork.MockFactory)
	w := new(unitofwork.MockWork)
	bus := &eventbus.NoOpPublisher{}
	mockIssuerRepo := new(issuer.MockRepository)

	service := NewInvoiceService(mockRepo, mockIssuerRepo, uf, bus)

	ctx := context.Background()

	invoiceData := &invoice.Invoice{
		ID:          uuid.New(),
		IssuerID:    uuid.New(),
		TotalAmount: 5000,
		IsLocked:    true,
	}

	uf.On("New", mock.AnythingOfType("*context.emptyCtx")).Return(w, nil)
	mockIssuerRepo.On("GetIssuerByID", ctx, invoiceData.IssuerID).Return(&issuer.Issuer{}, nil)
	mockRepo.On("CreateInvoice", ctx, invoiceData, mock.AnythingOfType("unitofwork.Option")).
		Return(nil)
	w.On("Commit").Return(nil)
	w.On("RollbackUnlessCommitted").Return(nil)

	_, err := service.CreateInvoice(ctx, invoiceData)
	require.NoError(t, err)

	mockRepo.AssertExpectations(t)
	uf.AssertExpectations(t)
	mockIssuerRepo.AssertExpectations(t)
}

func TestGetInvoiceByID(t *testing.T) {
	mockRepo := new(invoice.MockRepository)
	uf := new(unitofwork.MockFactory)

	service := NewInvoiceService(mockRepo, nil, uf, nil)

	inv := &invoice.Invoice{
		ID:            uuid.New(),
		AskingPrice:   100,
		IsLocked:      false,
		IsApproved:    false,
		InvoiceNumber: "INF-001",
	}

	mockRepo.On("GetInvoiceByID", mock.AnythingOfType("*context.emptyCtx"), inv.ID).
		Return(inv, nil)
	mockRepo.On("GetInvoiceItems", mock.AnythingOfType("*context.emptyCtx"), inv.ID).
		Return([]*invoice.Item{}, nil)

	returnedInv, err := service.GetInvoiceByID(context.Background(), inv.ID)
	require.NoError(t, err)

	require.EqualValues(t, inv, returnedInv)
}

func TestListInvoices(t *testing.T) {
	mockRepo := new(invoice.MockRepository)
	uf := new(unitofwork.MockFactory)

	inv := []*invoice.Invoice{
		{
			ID:            uuid.New(),
			AskingPrice:   100,
			IsLocked:      false,
			IsApproved:    false,
			InvoiceNumber: "INF-001",
		},
		{
			ID:            uuid.New(),
			AskingPrice:   100,
			IsLocked:      false,
			IsApproved:    false,
			InvoiceNumber: "INF-001",
		},
	}

	service := NewInvoiceService(mockRepo, nil, uf, nil)

	mockRepo.On("ListInvoices", mock.Anything).Return(inv, nil)

	returnedInvs, err := service.ListInvoices(context.Background())
	require.NoError(t, err)

	require.EqualValues(t, inv, returnedInvs)
}
