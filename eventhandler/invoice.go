package eventhandler

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"log"

	"github.com/tauki/invoiceexchange/internal/balance"
	"github.com/tauki/invoiceexchange/internal/investor"
	"github.com/tauki/invoiceexchange/internal/invoice"
	"github.com/tauki/invoiceexchange/internal/issuer"
	"github.com/tauki/invoiceexchange/internal/ledger"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

const (
	InvoiceApprovedEvent string = "invoice.approved"
)

type InvoiceEventHandlers struct {
	investorRepo investor.Repository
	issuerRepo   issuer.Repository
	balanceRepo  balance.Repository
	invoiceRepo  invoice.Repository
	uowFactory   unitofwork.Factory
	ledgerRepo   ledger.Repository
	lock         *invoiceLocks
}

func NewInvoiceEventHandlers(
	investorRepo investor.Repository,
	issuerRepo issuer.Repository,
	balanceRepo balance.Repository,
	ledgerRepo ledger.Repository,
	invoiceRepo invoice.Repository,
	uowFactory unitofwork.Factory,
) *InvoiceEventHandlers {
	return &InvoiceEventHandlers{
		investorRepo: investorRepo,
		issuerRepo:   issuerRepo,
		balanceRepo:  balanceRepo,
		uowFactory:   uowFactory,
		ledgerRepo:   ledgerRepo,
		invoiceRepo:  invoiceRepo,
		lock:         newInvoiceLocks(),
	}
}

func (ev *InvoiceEventHandlers) InvoiceApprovedHandler(data interface{}) {
	log.Printf("Event %s: %+v", InvoiceApprovedEvent, data)
	ctx := context.Background()
	invoiceID, err := uuid.Parse(data.(string))
	if err != nil {
		log.Printf("Error %s: %s", InvoiceApprovedEvent, err)
		return
	}

	inv, err := ev.invoiceRepo.GetInvoiceByID(ctx, invoiceID)
	if err != nil {
		log.Printf("Error %s: %s", InvoiceApprovedEvent, err)
		return
	}

	if inv == nil {
		log.Printf("Error %s: %s", InvoiceApprovedEvent, errors.New("invoice not found"))
		return
	}

	if inv.Status != invoice.StatusPending {
		log.Printf("Error %s: %s", InvoiceApprovedEvent, errors.New("invoice is not pending"))
		return
	}

	// locking the invoice in-memory for processing
	ev.lock.Lock(inv.ID.String())
	defer ev.lock.Unlock(inv.ID.String())

	issr, err := ev.issuerRepo.GetIssuerByID(ctx, inv.IssuerID)
	if err != nil {
		log.Printf("Error %s: %s", InvoiceApprovedEvent, err)
		return
	}

	if issr == nil {
		log.Printf("Error %s: %s", InvoiceApprovedEvent, errors.New("issuer not found"))
		return
	}

	ledgerEntries, err := ev.ledgerRepo.GetLedgerEntriesByInvoiceID(ctx, invoiceID)
	if err != nil {
		log.Printf("Error %s: %s", InvoiceApprovedEvent, err)
		return
	}

	uow, err := ev.uowFactory.New(ctx)
	if err != nil {
		log.Printf("Error %s: %s", InvoiceApprovedEvent, err)
		return
	}
	defer uow.RollbackUnlessCommitted()

	if err = ev.processLedgers(ledgerEntries, issr, inv, uow); err != nil {
		log.Printf("Error %s: %s", InvoiceApprovedEvent, err)
		return
	}

	inv.Status = invoice.StatusProcessed
	if err = ev.invoiceRepo.UpdateInvoice(ctx, inv, unitofwork.With(uow)); err != nil {
		log.Printf("Error %s: %s", InvoiceApprovedEvent, err)
		return
	}

	if err = uow.Commit(); err != nil {
		log.Printf("Error %s: %s", InvoiceApprovedEvent, err)
		return
	}
	log.Println("Event", InvoiceApprovedEvent, "processed successfully")
}

func (ev *InvoiceEventHandlers) checkLedgerSum(ledgerEntries []*ledger.Ledger) error {
	var sum float64
	for _, entry := range ledgerEntries {
		sum += entry.Amount
	}
	if sum != 0 {
		return errors.New("ledger entries do not sum to zero")
	}

	return nil
}

func (ev *InvoiceEventHandlers) processLedgers(ledgerEntries []*ledger.Ledger, issr *issuer.Issuer, inv *invoice.Invoice, uow unitofwork.Work) error {
	if err := ev.checkLedgerSum(ledgerEntries); err != nil {
		return err
	}

	for _, entry := range ledgerEntries {
		switch entry.Entity {
		case ledger.EntityTypeIssuer:
			log.Println(issr)
			issr.Balance.TotalAmount += entry.Amount
			issr.Balance.AvailableAmount += entry.Amount
			if err := ev.issuerRepo.UpdateIssuer(context.Background(), issr, unitofwork.With(uow)); err != nil {
				return err
			}
		case ledger.EntityTypeInvestor:
			invstr, err := ev.investorRepo.GetInvestorByID(context.Background(), entry.EntityID, unitofwork.With(uow))
			if err != nil {
				return err
			}

			invstr.Balance.TotalAmount += entry.Amount
			if err = ev.investorRepo.UpdateInvestor(context.Background(), invstr, unitofwork.With(uow)); err != nil {
				return err
			}

			err = inv.AddInvestor(invstr.ID)
			if err != nil {
				return err
			}

			if err = ev.invoiceRepo.UpdateInvoice(context.Background(), inv, unitofwork.With(uow)); err != nil {
				return err
			}

		}
	}
	return nil
}
