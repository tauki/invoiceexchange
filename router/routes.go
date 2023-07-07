package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/tauki/invoiceexchange/handlers"
	"github.com/tauki/invoiceexchange/internal/bid"
	"github.com/tauki/invoiceexchange/internal/investor"
	"github.com/tauki/invoiceexchange/internal/invoice"
	"github.com/tauki/invoiceexchange/internal/issuer"
	"github.com/tauki/invoiceexchange/router/middlewares/logging"
	"github.com/tauki/invoiceexchange/router/middlewares/recovery"
)

func NewRouter(
	invoiceService invoice.Service,
	bidService bid.Service,
	issuerService issuer.Service,
	investorService investor.Service,
	zaplog *zap.Logger,
) *mux.Router {
	// tribute to gorilla mux
	router := mux.NewRouter()
	router.Use(recovery.Middleware)
	router.Use(logging.ZapMiddleware(zaplog))
	router.Use(logging.InterceptorMiddleware)

	// controllers
	invoiceHandler := handlers.NewInvoiceController(invoiceService, bidService)
	bidHandler := handlers.NewBidController(bidService)
	issuerHandler := handlers.NewIssuerController(issuerService)
	investorHandler := handlers.NewInvestorController(investorService)

	// routes

	// invoice
	router.HandleFunc("/invoice", invoiceHandler.CreateInvoice).
		Methods(http.MethodPost)
	router.HandleFunc("/invoice/{id}", invoiceHandler.GetInvoice).
		Methods(http.MethodGet)
	//router.HandleFunc("/invoice/{id}", invoiceHandler.UpdateInvoice).
	//	Methods(http.MethodPut)
	//router.HandleFunc("/invoice/{id}", invoiceHandler.DeleteInvoice).
	//	Methods(http.MethodDelete)
	router.HandleFunc("/invoice/{id}/approve", invoiceHandler.ApproveTrade).
		Methods(http.MethodPatch)
	router.HandleFunc("/invoice/{id}/bids", invoiceHandler.GetBids).
		Methods(http.MethodGet)
	router.HandleFunc("/invoices", invoiceHandler.ListInvoices).
		Methods(http.MethodGet)

	// bid
	router.HandleFunc("/bid", bidHandler.CreateBid).
		Methods(http.MethodPost)
	router.HandleFunc("/bid/{id}", bidHandler.GetBid).
		Methods(http.MethodGet)
	//router.HandleFunc("/bid/{id}", bidHandler.UpdateBid).
	//	Methods(http.MethodPut)
	//router.HandleFunc("/bid/{id}", bidHandler.DeleteBid).
	//	Methods(http.MethodDelete)

	// issuer
	router.HandleFunc("/issuer", issuerHandler.CreateIssuer).
		Methods(http.MethodPost)
	router.HandleFunc("/issuer/{id}", issuerHandler.GetIssuer).
		Methods(http.MethodGet)
	//router.HandleFunc("/issuer/{id}", issuerHandler.UpdateIssuer).
	//	Methods(http.MethodPut)
	//router.HandleFunc("/issuer/{id}", issuerHandler.DeleteIssuer).
	//	Methods(http.MethodDelete)
	router.HandleFunc("/issuers", issuerHandler.ListIssuer).
		Methods(http.MethodGet)

	// investor
	router.HandleFunc("/investor", investorHandler.CreateInvestor).
		Methods(http.MethodPost)
	router.HandleFunc("/investor/{id}", investorHandler.GetInvestor).
		Methods(http.MethodGet)
	//router.HandleFunc("/investor/{id}", investorHandler.UpdateInvestor).
	//	Methods(http.MethodPut)
	//router.HandleFunc("/investor/{id}", investorHandler.DeleteInvestor).
	//	Methods(http.MethodDelete)
	router.HandleFunc("/investors", investorHandler.ListInvestors).
		Methods(http.MethodGet)
	//router.HandleFunc("/investor/{id}/bids", investorHandler.GetBids).
	//	Methods(http.MethodGet)

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	return router
}
