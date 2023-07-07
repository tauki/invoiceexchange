package eventhandler_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/tauki/invoiceexchange/eventhandler"
	"github.com/tauki/invoiceexchange/internal/balance"
	"github.com/tauki/invoiceexchange/internal/investor"
	"github.com/tauki/invoiceexchange/internal/invoice"
	"github.com/tauki/invoiceexchange/internal/issuer"
	"github.com/tauki/invoiceexchange/internal/ledger"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

func TestInvoiceApprovedHandler(t *testing.T) {
	ctx := context.Background()
	invoiceID := uuid.New()
	investorRepo := &investor.MockRepository{}
	issuerRepo := &issuer.MockRepository{}
	invoiceRepo := &invoice.MockRepository{}
	ledgerRepo := &ledger.MockRepository{}
	balanceRepo := &balance.MockRepository{}
	w := &unitofwork.MockWork{}
	uowFactory := &unitofwork.MockFactory{}

	handler := eventhandler.NewInvoiceEventHandlers(investorRepo, issuerRepo, balanceRepo, ledgerRepo, invoiceRepo, uowFactory)

	issID := uuid.New()
	iss := &issuer.Issuer{
		ID:   issID,
		Name: "iss name",
		Balance: &balance.Balance{
			ID:              uuid.New(),
			EntityID:        issID,
			TotalAmount:     1000,
			AvailableAmount: 900,
		},
	}

	inv := &invoice.Invoice{
		ID:       invoiceID,
		Status:   invoice.StatusPending,
		IssuerID: iss.ID,
	}

	invsID := uuid.New()
	invs := &investor.Investor{
		ID:   invsID,
		Name: "invs name",
		Balance: &balance.Balance{
			ID:              uuid.New(),
			EntityID:        invsID,
			TotalAmount:     1000,
			AvailableAmount: 900,
		},
	}

	ledgerEntries := []*ledger.Ledger{
		{
			ID:       uuid.New(),
			Amount:   100,
			Entity:   ledger.EntityTypeInvestor,
			EntityID: invs.ID,
		},
		{
			ID:       uuid.New(),
			Amount:   -100,
			Entity:   ledger.EntityTypeIssuer,
			EntityID: iss.ID,
		},
	}

	invoiceRepo.On("GetInvoiceByID", ctx, invoiceID).Return(inv, nil)
	issuerRepo.On("GetIssuerByID", ctx, inv.IssuerID).Return(iss, nil)
	ledgerRepo.On("GetLedgerEntriesByInvoiceID", ctx, invoiceID).Return(ledgerEntries, nil)

	uowFactory.On("New", ctx).Return(w, nil)
	w.On("Commit").Return(nil)
	w.On("RollbackUnlessCommitted").Return(nil)

	issuerRepo.On("UpdateIssuer", ctx, mock.AnythingOfType("*issuer.Issuer"), mock.AnythingOfType("unitofwork.Option")).Return(nil)
	investorRepo.On("GetInvestorByID", ctx, invs.ID, mock.AnythingOfType("unitofwork.Option")).Return(invs, nil)
	investorRepo.On("UpdateInvestor", ctx, mock.AnythingOfType("*investor.Investor"), mock.AnythingOfType("unitofwork.Option")).Return(nil)
	invoiceRepo.On("UpdateInvoice", ctx, mock.AnythingOfType("*invoice.Invoice"), mock.AnythingOfType("unitofwork.Option")).Return(nil).Times(2)

	handler.InvoiceApprovedHandler(invoiceID.String())

	mock.AssertExpectationsForObjects(t, investorRepo, issuerRepo, invoiceRepo, w, uowFactory, ledgerRepo)
}
