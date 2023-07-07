package eventhandler

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/tauki/invoiceexchange/internal/bid"
	"github.com/tauki/invoiceexchange/internal/investor"
	"github.com/tauki/invoiceexchange/internal/invoice"
	"github.com/tauki/invoiceexchange/internal/ledger"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

const (
	BidCreatedEvent string = "bid.created"
)

type BidEventHandlers struct {
	bidRepo      bid.Repository
	investorRepo investor.Repository
	invoiceRepo  invoice.Repository
	ledgerRepo   ledger.Repository
	uowFactory   unitofwork.Factory
	lock         *invoiceLocks
}

func NewBidEventHandlers(
	bidRepo bid.Repository,
	investorRepo investor.Repository,
	invoiceRepo invoice.Repository,
	ledgerRepo ledger.Repository,
	uowFactory unitofwork.Factory,
) *BidEventHandlers {
	return &BidEventHandlers{
		bidRepo:      bidRepo,
		investorRepo: investorRepo,
		invoiceRepo:  invoiceRepo,
		ledgerRepo:   ledgerRepo,
		uowFactory:   uowFactory,
		lock:         newInvoiceLocks(),
	}
}

func (ev *BidEventHandlers) BidCreatedHandler(data interface{}) {
	ctx := context.Background()
	BidID, err := uuid.Parse(data.(string))
	if err != nil {
		log.Printf("Error %s: event data %+v is not an UUID.", BidCreatedEvent, data)
		return
	}

	createdBid, err := ev.bidRepo.GetBidByID(ctx, BidID)
	if err != nil {
		log.Println("Error calling GetBidByID", err)
		return
	}

	if createdBid.Status != bid.Pending {
		log.Println("bid not in pending status")
		return
	}

	// locking the invoice in-memory for processing
	ev.lock.Lock(createdBid.InvoiceID.String())
	defer ev.lock.Unlock(createdBid.InvoiceID.String())

	uow, err := ev.uowFactory.New(ctx)
	if err != nil {
		log.Println("unable to create unitofwork", err)
		return
	}
	defer uow.RollbackUnlessCommitted()

	inv, err := ev.invoiceRepo.GetInvoiceByID(ctx, createdBid.InvoiceID, unitofwork.With(uow))
	if err != nil {
		log.Println("Error calling GetInvoiceByID", err)
		return
	}

	if inv.IsLocked {
		err = ev.rejectBid(ctx, createdBid, uow)
		if err != nil {
			log.Println("Error calling rejectBid", err)
			return
		}
	} else {
		err = ev.acceptBid(ctx, createdBid, inv, uow)
		if err != nil {
			log.Println("Error calling acceptBid", err)
			return
		}
	}

	if err = uow.Commit(); err != nil {
		log.Println("unable to commit", err)
	}

	log.Println("bid handled successfully", createdBid)
}

func (ev *BidEventHandlers) acceptBid(ctx context.Context, b *bid.Bid, inv *invoice.Invoice, uow unitofwork.Work) error {
	var reverseAmount float64
	bidAmount := b.Amount
	if inv.AmountDue-bidAmount < 0 {
		reverseAmount = b.Amount - inv.AmountDue
		bidAmount = inv.AmountDue
	}
	inv.AmountDue = inv.AmountDue - bidAmount
	b.AcceptedAmount = bidAmount
	b.Status = bid.Accepted

	err := ev.invoiceRepo.UpdateInvoice(ctx, inv, unitofwork.With(uow))
	if err != nil {
		return err
	}

	err = ev.bidRepo.UpdateBid(ctx, b, unitofwork.With(uow))
	if err != nil {
		return err
	}

	if reverseAmount > 0 {
		err = ev.reverseBid(ctx, b, reverseAmount, uow)
		log.Println("reverse amount", reverseAmount)
		if err != nil {
			return err
		}
	}

	ledgerEntryInvestor, err := ledger.New(inv.ID, ledger.EntityTypeInvestor, b.InvestorID, -bidAmount)
	if err != nil {
		return err
	}
	ledgerEntryIssuer, err := ledger.New(inv.ID, ledger.EntityTypeIssuer, inv.IssuerID, bidAmount)
	if err != nil {
		return err
	}

	err = ev.ledgerRepo.AddLedgerEntry(ctx, ledgerEntryInvestor, unitofwork.With(uow))
	if err != nil {
		return err
	}

	err = ev.ledgerRepo.AddLedgerEntry(ctx, ledgerEntryIssuer, unitofwork.With(uow))
	if err != nil {
		return err
	}

	if inv.AmountDue == 0 {
		inv.IsLocked = true
		err = ev.invoiceRepo.UpdateInvoice(ctx, inv, unitofwork.With(uow))
		if err != nil {
			log.Println("Error calling UpdateInvoice", inv.ID)
			return err
		}
	}

	return nil
}

func (ev *BidEventHandlers) rejectBid(ctx context.Context, b *bid.Bid, uow unitofwork.Work) error {
	b.Status = bid.Rejected
	return ev.reverseBid(ctx, b, b.Amount, uow)
}

func (ev *BidEventHandlers) reverseBid(ctx context.Context, b *bid.Bid, amount float64, uow unitofwork.Work) error {
	inv, err := ev.investorRepo.GetInvestorByID(ctx, b.InvestorID)
	if err != nil {
		return err
	}

	inv.Balance.AvailableAmount = inv.Balance.AvailableAmount + amount
	err = ev.investorRepo.UpdateInvestor(ctx, inv, unitofwork.With(uow))
	if err != nil {
		return err
	}

	b.AcceptedAmount = b.Amount - amount
	return ev.bidRepo.UpdateBid(ctx, b, unitofwork.With(uow))
}
