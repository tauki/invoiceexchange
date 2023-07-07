package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/tauki/invoiceexchange/internal/bid"
	"github.com/tauki/invoiceexchange/internal/errors"
	"github.com/tauki/invoiceexchange/internal/invoice"
	"github.com/tauki/invoiceexchange/router/middlewares/logging"
)

type InvoiceController struct {
	service    invoice.Service
	bidService bid.Service
}

func NewInvoiceController(
	service invoice.Service,
	bidService bid.Service,
) *InvoiceController {
	return &InvoiceController{
		service:    service,
		bidService: bidService,
	}
}

type CreateInvoiceRequest struct {
	InvoiceDate   time.Time `json:"invoice_date"`
	InvoiceNumber string    `json:"invoice_number"`
	Reference     string    `json:"reference"`
	CompanyName   string    `json:"company_name"`
	CustomerName  string    `json:"customer_name"`
	Currency      string    `json:"currency"`
	TotalAmount   float64   `json:"total_amount"`
	TotalVat      float64   `json:"total_vat"`
	DueDate       time.Time `json:"due_date"`
	IssuerID      uuid.UUID `json:"issuer_id"`
	AskingPrice   float64   `json:"asking_price"`
	InvoiceItems  []struct {
		Description string `json:"description"`
		// quantity might be better as a float
		Quantity  int     `json:"quantity"`
		UnitPrice float64 `json:"unit_price"`
		Amount    float64 `json:"amount"`
		Vat       float64 `json:"vat"`
		Total     float64 `json:"total"`
	} `json:"items"`
}

func (c *InvoiceController) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.GetZap(ctx)

	var req CreateInvoiceRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		HTTPErrorResponse(w, errors.NewValidationError("invalid request body", "req.body"))
		return
	}

	inv, err := invoice.New(
		req.InvoiceDate,
		req.DueDate,
		req.InvoiceNumber,
		req.Reference,
		req.CompanyName,
		req.CustomerName,
		req.Currency,
		req.TotalAmount,
		req.TotalVat,
		req.AskingPrice,
		req.IssuerID,
	)
	if err != nil {
		HTTPErrorResponse(w, err)
		return
	}

	items := make([]*invoice.Item, len(req.InvoiceItems))
	for _, item := range req.InvoiceItems {
		invItem, err := invoice.NewItem(item.Description, item.Quantity, item.UnitPrice, item.Vat, item.Total)
		if err != nil {
			HTTPErrorResponse(w, err)
			return
		}
		items = append(items, invItem)
	}
	inv.Items = items

	createdInv, err := c.service.CreateInvoice(ctx, inv)
	if err != nil {
		log.Error("failed to create invoice", zap.Error(err))
		HTTPErrorResponse(w, err)
		return
	}

	response, _ := json.Marshal(createdInv)
	HTTPResponse(w, http.StatusCreated, response)
}

func (c *InvoiceController) GetInvoice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]
	invID, err := uuid.Parse(id)
	if err != nil {
		HTTPErrorResponse(w, errors.NewValidationError("invalid invoice ID", "req.params.id"))
		return
	}

	inv, err := c.service.GetInvoiceByID(ctx, invID)
	if err != nil {
		HTTPErrorResponse(w, err)
		return
	}

	response, _ := json.Marshal(inv)
	HTTPResponse(w, http.StatusOK, response)
}

func (c *InvoiceController) ListInvoices(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.GetZap(ctx)

	invs, err := c.service.ListInvoices(ctx)
	if err != nil {
		log.Error("error calling service.ListInvoices", zap.Error(err))
		HTTPErrorResponse(w, err)
		return
	}

	response, _ := json.Marshal(invs)
	HTTPResponse(w, http.StatusOK, response)
}

func (c *InvoiceController) GetBids(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.GetZap(ctx)

	id := mux.Vars(r)["id"]
	invID, err := uuid.Parse(id)
	if err != nil {
		HTTPErrorResponse(w, errors.NewValidationError("invalid invoice ID", "req.params.id"))
		return
	}

	bids, err := c.bidService.ListBidsByInvoiceID(ctx, invID)
	if err != nil {
		log.Error("error calling service.ListBidsByInvoiceID", zap.Error(err))
		HTTPErrorResponse(w, err)
		return
	}

	response, _ := json.Marshal(bids)
	HTTPResponse(w, http.StatusOK, response)
}

func (c *InvoiceController) ApproveTrade(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.GetZap(ctx)

	id := mux.Vars(r)["id"]
	invID, err := uuid.Parse(id)
	if err != nil {
		HTTPErrorResponse(w, errors.NewValidationError("invalid invoice ID", "req.params.id"))
		return
	}

	inv, err := c.service.ApproveTrade(ctx, invID)
	if err != nil {
		log.Warn("error calling service.ApproveTrade", zap.Error(err))
		HTTPErrorResponse(w, err)
		return
	}

	response, _ := json.Marshal(inv)
	HTTPResponse(w, http.StatusAccepted, response)
}
