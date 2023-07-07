//go:build e2e
// +build e2e

package handlers_test

import (
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tauki/invoiceexchange/ent"
	entbid "github.com/tauki/invoiceexchange/ent/bid"
	"github.com/tauki/invoiceexchange/internal/bid"
	"github.com/tauki/invoiceexchange/internal/investor"
	"github.com/tauki/invoiceexchange/internal/invoice"
	"github.com/tauki/invoiceexchange/internal/issuer"
	"github.com/tauki/invoiceexchange/internal/services/eventbus"
	"github.com/tauki/invoiceexchange/repos"
	"github.com/tauki/invoiceexchange/router/middlewares/logging"
	"github.com/tauki/invoiceexchange/services"
	"go.uber.org/zap"
	"modernc.org/sqlite"
	"net/http"
	"testing"
	"time"
)

var entClient *ent.Client

var preloadedIDs = map[string]uuid.UUID{
	"issuer":             uuid.MustParse("ac538328-f800-44bc-927f-6bbba9ffce28"),
	"investor":           uuid.MustParse("f743d4ab-3c86-4b31-b5e5-c48fe7f2f1b6"),
	"invoice":            uuid.MustParse("a33add00-900e-4d03-996d-e72efcfdb274"),
	"invoice-locked":     uuid.MustParse("cbb3cd1f-0bb1-4fd8-89b9-2bf64f6391a6"),
	"invoice-locked-bid": uuid.MustParse("ae58cc6d-6408-463a-8a0c-19965a1d2696"),
	"bid":                uuid.MustParse("bd880f51-d427-4a84-b239-db46af455d5c"),
}

func GetTestHandleFunc(controller func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//z, _ := zap.NewDevelopment()
		ctx := logging.WithContext(r.Context(), zap.NewNop())
		r = r.WithContext(ctx)
		controller(w, r)
	}
}

func getTestEntClient(t *testing.T) *ent.Client {
	if entClient != nil {
		return entClient
	}
	sql.Register("sqlite3", &sqlite.Driver{})

	open, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal("error connecting database", err)
	}
	drv := entsql.OpenDB("sqlite3", open)

	if _, err := drv.DB().Exec(pragmaFKOn()); err != nil {
		t.Fatal("failed setting foreign_keys pragma", err)
	}
	var entOptions = []ent.Option{ent.Driver(drv)}
	entClient = ent.NewClient(entOptions...)
	if err := entClient.Schema.Create(context.Background()); err != nil {
		t.Fatal("failed creating localSqlite schema", err)
	}

	preloadData(t, entClient)
	return entClient
}

func preloadData(t *testing.T, entClient *ent.Client) {
	IssuerID := preloadedIDs["issuer"]
	if _, err := entClient.Issuer.Get(context.Background(), IssuerID); ent.IsNotFound(err) {
		balanceID := uuid.New()
		_, err = entClient.Balance.Create().
			SetID(balanceID).
			SetEntityID(IssuerID).
			SetTotalAmount(10000).
			SetAvailableAmount(10000).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if err != nil {
			t.Error("Bootstrap: ent.Balance.Create()", zap.Error(err))
		}

		_, err := entClient.Issuer.Create().
			SetID(IssuerID).
			SetName("John Doe").
			SetBalanceID(balanceID).
			Save(context.Background())
		if err != nil {
			t.Error("Bootstrap: ent.Issuer.Create()", zap.Error(err))
		}
	}

	InvestorID := preloadedIDs["investor"]
	if _, err := entClient.Investor.Get(context.Background(), InvestorID); ent.IsNotFound(err) {
		balanceID := uuid.New()
		_, err = entClient.Balance.Create().
			SetID(balanceID).
			SetEntityID(InvestorID).
			SetTotalAmount(10000).
			SetAvailableAmount(10000).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if err != nil {
			t.Error("Bootstrap: ent.Balance.Create()", zap.Error(err))
		}

		_, err := entClient.Investor.Create().
			SetID(InvestorID).
			SetName("Jhon Doe").
			SetBalanceID(balanceID).
			Save(context.Background())
		if err != nil {
			t.Error("Bootstrap: ent.Investor.Create()", zap.Error(err))
		}
	}

	invoiceID := preloadedIDs["invoice"]
	if _, err := entClient.Invoice.Get(context.Background(), invoiceID); ent.IsNotFound(err) {
		_, err := entClient.Invoice.Create().
			SetID(invoiceID).
			SetIsLocked(false).
			SetStatus("pending").
			SetInvoiceNumber("INV-01").
			SetInvoiceDate(time.Now()).
			SetDueDate(time.Now().Add(time.Hour * 24 * 30)).
			SetCustomerName("Customer").
			SetCompanyName("Company").
			SetReference("REF-01").
			SetTotalVat(20.0).
			SetTotalAmount(120.0).
			SetAskingPrice(110.0).
			SetAmountDue(110.0).
			SetCurrency("GBP").
			SetIssuerID(IssuerID).
			Save(context.Background())
		if err != nil {
			t.Error("Bootstrap: ent.Invoice.Create()", zap.Error(err))
		}
	}

	invoiceID2 := preloadedIDs["invoice-locked"]
	if _, err := entClient.Invoice.Get(context.Background(), invoiceID2); ent.IsNotFound(err) {
		_, err := entClient.Invoice.Create().
			SetID(invoiceID2).
			SetIsLocked(true).
			SetStatus("pending").
			SetInvoiceNumber("INV-02").
			SetInvoiceDate(time.Now()).
			SetDueDate(time.Now().Add(time.Hour * 24 * 30)).
			SetCustomerName("Customer2").
			SetCompanyName("Company").
			SetReference("REF-02").
			SetTotalVat(20.0).
			SetTotalAmount(120.).
			SetAskingPrice(110.0).
			SetAmountDue(110.0).
			SetCurrency("GBP").
			SetIssuerID(IssuerID).
			AddInvestorIDs(InvestorID).
			Save(context.Background())
		if err != nil {
			t.Error("Bootstrap: ent.Invoice.Create()", zap.Error(err))
		}
	}

	bidID := preloadedIDs["bid"]
	if _, err := entClient.Bid.Get(context.Background(), bidID); ent.IsNotFound(err) {
		_, err := entClient.Bid.Create().
			SetID(bidID).
			SetStatus(entbid.DefaultStatus).
			SetAmount(10).
			SetInvoiceID(invoiceID2).
			SetInvestorID(InvestorID).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if err != nil {
			t.Error("Bootstrap: ent.Bid.Create()", zap.Error(err))
		}
	}

	lockedInvBid := preloadedIDs["bid"]
	if _, err := entClient.Bid.Get(context.Background(), lockedInvBid); ent.IsNotFound(err) {
		_, err := entClient.Bid.Create().
			SetID(lockedInvBid).
			SetStatus(entbid.StatusACCEPTED).
			SetAmount(10).
			SetInvoiceID(invoiceID2).
			SetInvestorID(InvestorID).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if err != nil {
			t.Error("Bootstrap: ent.Bid.Create()", zap.Error(err))
		}
	}
}

func GetTestBidService(t *testing.T) (bid.Service, []func()) {
	entClient := getTestEntClient(t)
	bidRepo := repos.NewBidRepo(entClient)
	uowFactory := repos.NewEntUOWFactory(entClient)
	investorRepo := repos.NewInvestorRepo(entClient)
	invoiceRepo := repos.NewInvoiceRepo(entClient)
	balanceRepo := repos.NewBalanceRepo(entClient)
	eventBus := eventbus.NewEventBus(5)
	go eventBus.Run()
	eventBus.Running()
	var servicesToClose []func()
	servicesToClose = append(servicesToClose, func() { eventBus.Close() })
	servicesToClose = append(servicesToClose, func() { _ = entClient.Close() })
	return services.NewBidService(bidRepo, uowFactory, invoiceRepo, investorRepo, balanceRepo, eventBus), servicesToClose
}

func GetTestInvoiceService(t *testing.T) (invoice.Service, []func()) {
	entClient := getTestEntClient(t)
	invoiceRepo := repos.NewInvoiceRepo(entClient)
	issuerRepo := repos.NewIssuerRepo(entClient)
	uowFactory := repos.NewEntUOWFactory(entClient)
	eventBus := eventbus.NewEventBus(5)
	go eventBus.Run()
	eventBus.Running()
	var servicesToClose []func()
	servicesToClose = append(servicesToClose, func() { _ = entClient.Close() })
	servicesToClose = append(servicesToClose, func() { eventBus.Close() })
	return services.NewInvoiceService(invoiceRepo, issuerRepo, uowFactory, eventBus), servicesToClose
}

func GetTestIssuerService(t *testing.T) (issuer.Service, []func()) {
	entClient := getTestEntClient(t)
	issuerRepo := repos.NewIssuerRepo(entClient)
	uowFactory := repos.NewEntUOWFactory(entClient)
	var servicesToClose []func()
	servicesToClose = append(servicesToClose, func() { _ = entClient.Close() })
	return services.NewIssuerService(issuerRepo, uowFactory), servicesToClose
}

func GetTestInvestorService(t *testing.T) (investor.Service, []func()) {
	entClient := getTestEntClient(t)
	investorRepo := repos.NewInvestorRepo(entClient)
	uowFactory := repos.NewEntUOWFactory(entClient)
	var servicesToClose []func()
	servicesToClose = append(servicesToClose, func() { _ = entClient.Close() })
	return services.NewInvestorService(investorRepo, uowFactory), servicesToClose
}

func pragmaFKOn() string {
	return "PRAGMA foreign_keys = ON"
}
