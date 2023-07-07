package eventhandler_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/tauki/invoiceexchange/eventhandler"
	"github.com/tauki/invoiceexchange/internal/balance"
	"github.com/tauki/invoiceexchange/internal/bid"
	"github.com/tauki/invoiceexchange/internal/investor"
	"github.com/tauki/invoiceexchange/internal/invoice"
	"github.com/tauki/invoiceexchange/internal/issuer"
	"github.com/tauki/invoiceexchange/internal/ledger"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

func TestBidCreatedHandler_acceptBid(t *testing.T) {
	ctx := context.Background()
	bidID := uuid.New()
	bidRepo := &bid.MockRepository{}
	investorRepo := &investor.MockRepository{}
	invoiceRepo := &invoice.MockRepository{}
	ledgerRepo := &ledger.MockRepository{}
	w := &unitofwork.MockWork{}
	uowFactory := &unitofwork.MockFactory{}

	handler := eventhandler.NewBidEventHandlers(bidRepo, investorRepo, invoiceRepo, ledgerRepo, uowFactory)

	iss := &issuer.Issuer{
		ID:   uuid.New(),
		Name: "iss name",
	}

	inv := &invoice.Invoice{
		ID:            uuid.New(),
		Status:        "",
		AskingPrice:   100,
		IsLocked:      false,
		IsApproved:    false,
		InvoiceNumber: "INV-001",
		AmountDue:     100,
		IssuerID:      iss.ID,
	}

	invs := &investor.Investor{
		ID:   uuid.New(),
		Name: "invs name",
		Balance: &balance.Balance{
			ID:              uuid.New(),
			EntityID:        inv.ID,
			TotalAmount:     1000,
			AvailableAmount: 900,
		},
	}

	createdBid := &bid.Bid{
		ID:         bidID,
		Amount:     100,
		Status:     bid.Pending,
		InvoiceID:  inv.ID,
		InvestorID: invs.ID,
	}

	bidRepo.On("GetBidByID", ctx, bidID).Return(createdBid, nil)
	invoiceRepo.On("GetInvoiceByID", ctx, createdBid.InvoiceID, mock.AnythingOfType("unitofwork.Option")).Return(inv, nil)
	invoiceRepo.On("UpdateInvoice", ctx, mock.AnythingOfType("*invoice.Invoice"), mock.AnythingOfType("unitofwork.Option")).Return(nil)
	bidRepo.On("UpdateBid", ctx, mock.AnythingOfType("*bid.Bid"), mock.AnythingOfType("unitofwork.Option")).Return(nil)
	invoiceRepo.On("UpdateInvoice", ctx, mock.AnythingOfType("*invoice.Invoice"), mock.AnythingOfType("unitofwork.Option")).Return(nil)
	ledgerRepo.On("AddLedgerEntry", ctx, mock.AnythingOfType("*ledger.Ledger"), mock.AnythingOfType("unitofwork.Option")).Return(nil).Times(2)

	uowFactory.On("New", ctx).Return(w, nil)
	w.On("Commit").Return(nil)
	w.On("RollbackUnlessCommitted").Return(nil)

	handler.BidCreatedHandler(bidID.String())

	mock.AssertExpectationsForObjects(t, bidRepo, investorRepo, invoiceRepo, w, uowFactory, ledgerRepo)
}

func TestBidCreatedHandler_rejectBid(t *testing.T) {
	ctx := context.Background()
	bidID := uuid.New()
	bidRepo := &bid.MockRepository{}
	investorRepo := &investor.MockRepository{}
	invoiceRepo := &invoice.MockRepository{}
	ledgerRepo := &ledger.MockRepository{}
	w := &unitofwork.MockWork{}
	uowFactory := &unitofwork.MockFactory{}

	handler := eventhandler.NewBidEventHandlers(bidRepo, investorRepo, invoiceRepo, ledgerRepo, uowFactory)

	iss := &issuer.Issuer{
		ID:   uuid.New(),
		Name: "iss name",
	}

	inv := &invoice.Invoice{
		ID:            uuid.New(),
		Status:        "",
		AskingPrice:   100,
		IsLocked:      true,
		IsApproved:    false,
		InvoiceNumber: "INV-001",
		AmountDue:     0,
		IssuerID:      iss.ID,
	}

	invs := &investor.Investor{
		ID:   uuid.New(),
		Name: "invs name",
		Balance: &balance.Balance{
			ID:              uuid.New(),
			EntityID:        inv.ID,
			TotalAmount:     1000,
			AvailableAmount: 900,
		},
	}

	createdBid := &bid.Bid{
		ID:         bidID,
		Amount:     100,
		Status:     bid.Pending,
		InvoiceID:  inv.ID,
		InvestorID: invs.ID,
	}

	bidRepo.On("GetBidByID", ctx, bidID).Return(createdBid, nil)
	invoiceRepo.On("GetInvoiceByID", ctx, createdBid.InvoiceID, mock.AnythingOfType("unitofwork.Option")).Return(inv, nil)
	investorRepo.On("GetInvestorByID", ctx, invs.ID).Return(invs, nil)
	investorRepo.On("UpdateInvestor", ctx, mock.AnythingOfType("*investor.Investor"), mock.AnythingOfType("unitofwork.Option")).Return(nil)
	bidRepo.On("UpdateBid", ctx, mock.AnythingOfType("*bid.Bid"), mock.AnythingOfType("unitofwork.Option")).Return(nil)

	uowFactory.On("New", ctx).Return(w, nil)
	w.On("Commit").Return(nil)
	w.On("RollbackUnlessCommitted").Return(nil)

	handler.BidCreatedHandler(bidID.String())

	mock.AssertExpectationsForObjects(t, bidRepo, investorRepo, invoiceRepo, w, uowFactory)
}
