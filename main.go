package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/omeid/uconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
	"modernc.org/sqlite"

	"github.com/tauki/invoiceexchange/config"
	"github.com/tauki/invoiceexchange/ent"
	entbid "github.com/tauki/invoiceexchange/ent/bid"
	"github.com/tauki/invoiceexchange/eventhandler"
	"github.com/tauki/invoiceexchange/internal/services/eventbus"
	"github.com/tauki/invoiceexchange/repos"
	"github.com/tauki/invoiceexchange/router"
	"github.com/tauki/invoiceexchange/services"
)

func main() {
	cfg := &config.Config{}
	files := uconfig.Files{
		{"local.yaml", yaml.Unmarshal},
	}

	c, err := uconfig.Classic(cfg, files)
	if err != nil {
		c.Usage()
		log.Fatalf("Failed to parse configuration: %v", err)
	}

	var zaplog *zap.Logger
	if cfg.Environment == "dev" {
		zapConfig := zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zaplog, _ = zapConfig.Build()
	} else {
		zaplog, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("Failed to create logger: %v", err)
		}
	}
	_ = zap.RedirectStdLog(zaplog)
	_ = zap.ReplaceGlobals(zaplog)
	log.Println("Standard logger test")

	if cfg.Environment == "dev" {
		c.Usage()
		zaplog.Info("Server config", zap.Any("config", cfg))
	}

	ctx := context.Background()
	var servicesToClose []func()

	var entClient *ent.Client
	//var sqlDrv *entsql.Driver
	switch cfg.EntDriver {
	case config.SqliteEntDriver:
		entClient, _ = initEntSqlite(cfg, zaplog)
	case config.PostgresEntDriver:
		entClient, _ = initEntClient(cfg, zaplog)
	default:
		log.Fatalf("EntDriver expected to be (postgres,sqlite) got (%s)", cfg.EntDriver)
	}
	servicesToClose = append(servicesToClose, func() { _ = entClient.Close() })

	eventBus := eventbus.NewEventBus(1024)
	go eventBus.Run()
	servicesToClose = append(servicesToClose, func() { eventBus.Close() })

	// repos
	bidRepo := repos.NewBidRepo(entClient)
	investorRepo := repos.NewInvestorRepo(entClient)
	invoiceRepo := repos.NewInvoiceRepo(entClient)
	issuerRepo := repos.NewIssuerRepo(entClient)
	balanceRepo := repos.NewBalanceRepo(entClient)
	ledgerRepo := repos.NewLedgerRepo(entClient)
	uowFactory := repos.NewEntUOWFactory(entClient)

	// services
	bidService := services.NewBidService(
		bidRepo,
		uowFactory,
		invoiceRepo,
		investorRepo,
		balanceRepo,
		eventBus,
	)
	investorService := services.NewInvestorService(
		investorRepo,
		uowFactory,
	)
	invoiceService := services.NewInvoiceService(
		invoiceRepo,
		issuerRepo,
		uowFactory,
		eventBus,
	)
	issuerService := services.NewIssuerService(
		issuerRepo,
		uowFactory,
	)

	bidEventHandlers := eventhandler.NewBidEventHandlers(
		bidRepo, investorRepo, invoiceRepo, ledgerRepo, uowFactory)
	invoiceEventHandlers := eventhandler.NewInvoiceEventHandlers(
		investorRepo, issuerRepo, balanceRepo, ledgerRepo, invoiceRepo, uowFactory)
	eventBus.Subscribe(eventhandler.BidCreatedEvent, bidEventHandlers.BidCreatedHandler)
	eventBus.Subscribe(eventhandler.InvoiceApprovedEvent, invoiceEventHandlers.InvoiceApprovedHandler)

	s := &http.Server{
		Handler: router.NewRouter(
			invoiceService,
			bidService,
			issuerService,
			investorService,
			zaplog,
		),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Create channel to listen for signals.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	eventBus.Running()
	log.Println("Event bus running")
	if cfg.Server.TLS {
		go func() {
			log.Println("Serving on: ", cfg.Server.TLSPort)
			s.Addr = fmt.Sprintf(":%s", cfg.Server.TLSPort)
			err := s.ListenAndServeTLS(
				cfg.Server.CertPath,
				cfg.Server.CertPrivateKey)
			if err != nil && err != http.ErrServerClosed {
				log.Fatalf("failed to start server… %s", err.Error())
			}
		}()
	} else {
		go func() {
			log.Println("Serving on: ", cfg.Server.Port)
			s.Addr = fmt.Sprintf(":%s", cfg.Server.Port)
			err := s.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				log.Fatalf("failed to start server :%s", err.Error())
			}
		}()
	}

	// Receive output from signalChan.
	<-signalChan
	zaplog.Info("terminating server...")

	// Timeout if waiting for connections to return idle.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shutdown the server by waiting on existing requests.
	if err := s.Shutdown(ctx); err != nil {
		zaplog.Error("server shutdown failed", zap.Error(err))
	}

	// Close all services
	zaplog.Info("closing services...")
	for _, closeFn := range servicesToClose {
		closeFn()
	}
	zaplog.Info("server exited")
	_ = zaplog.Sync()
}

// withMiddleware is a helper function that applies middleware to a handler.
func withMiddleware(h http.HandlerFunc, middleware ...func(handlerFunc http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func initEntClient(cfg *config.Config, log *zap.Logger) (*ent.Client, *entsql.Driver) {
	conCfg, err := pgx.ParseConfig(cfg.PostgresDSN)
	if err != nil {
		log.Fatal("error parsing dsn", zap.Error(err))
	}

	log.Info(conCfg.ConnString())

	db, err := sql.Open(cfg.EntDriver, conCfg.ConnString())
	if err != nil {
		log.Fatal("error opening db", zap.Error(err))
	}

	log.Info("Loading Ent Driver")
	drv := entsql.OpenDB(dialect.Postgres, db)
	entClient := ent.NewClient(ent.Driver(drv) /*ent.Debug(), ent.Log(t.Log) */)
	return entClient, drv
}

func initEntSqlite(cfg *config.Config, log *zap.Logger) (*ent.Client, *entsql.Driver) {
	log.Info("Loading Ent Driver")

	// Default registration for "modernc.org/sqlite" is "sqlite".
	// Ent expects "sqlite3", so we register the driver here.
	sql.Register("sqlite3", &sqlite.Driver{})

	open, err := sql.Open("sqlite3", "file:local.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatal("error connecting database", zap.Error(err))
	}
	drv := entsql.OpenDB("sqlite3", open)

	// Enable Foreign Keys capability required by Ent Sqlite dialect.
	// modernc.org driver doesn't parse the _fk=1 dsn argument
	if _, err := drv.DB().Exec(pragmaFKOn()); err != nil {
		log.Fatal("failed setting foreign_keys pragma", zap.Error(err))
	}
	var entOptions = []ent.Option{ent.Driver(drv)}
	entClient := ent.NewClient(entOptions...)
	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatal("failed creating localSqlite schema", zap.Error(err))
	}

	preloadData(entClient, log)

	return entClient, drv
}

func pragmaFKOn() string {
	return "PRAGMA foreign_keys = ON"
}

func preloadData(entClient *ent.Client, log *zap.Logger) {
	IssuerID := uuid.MustParse("ac538328-f800-44bc-927f-6bbba9ffce28")
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
			log.Error("Bootstrap: ent.Balance.Create()", zap.Error(err))
		}

		_, err := entClient.Issuer.Create().
			SetID(IssuerID).
			SetName("John Doe").
			SetBalanceID(balanceID).
			Save(context.Background())
		if err != nil {
			log.Error("Bootstrap: ent.Issuer.Create()", zap.Error(err))
		}
	}

	InvestorID := uuid.MustParse("f743d4ab-3c86-4b31-b5e5-c48fe7f2f1b6")
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
			log.Error("Bootstrap: ent.Balance.Create()", zap.Error(err))
		}

		_, err := entClient.Investor.Create().
			SetID(InvestorID).
			SetName("Jhon Doe").
			SetBalanceID(balanceID).
			Save(context.Background())
		if err != nil {
			log.Error("Bootstrap: ent.Investor.Create()", zap.Error(err))
		}
	}

	invoiceID := uuid.MustParse("a33add00-900e-4d03-996d-e72efcfdb274")
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
			log.Error("Bootstrap: ent.Invoice.Create()", zap.Error(err))
		}
	}

	invoiceID2 := uuid.MustParse("d7edb506-b22a-466f-ae28-ded8b8018ad4")
	if _, err := entClient.Invoice.Get(context.Background(), invoiceID2); ent.IsNotFound(err) {
		_, err := entClient.Invoice.Create().
			SetID(invoiceID2).
			SetIsLocked(false).
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
			log.Error("Bootstrap: ent.Invoice.Create()", zap.Error(err))
		}
	}

	bidID := uuid.MustParse("bd880f51-d427-4a84-b239-db46af455d5c")
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
			log.Error("Bootstrap: ent.Bid.Create()", zap.Error(err))
		}
	}
}
