package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/tauki/invoiceexchange/internal/bid"
	"github.com/tauki/invoiceexchange/internal/errors"
	"github.com/tauki/invoiceexchange/router/middlewares/logging"
	"go.uber.org/zap"
	"net/http"
)

type BidController struct {
	service bid.Service
}

func NewBidController(service bid.Service) *BidController {
	return &BidController{
		service: service,
	}
}

type CreateBidRequest struct {
	InvoiceID  uuid.UUID `json:"invoice_id"`
	InvestorID uuid.UUID `json:"investor_id"`
	Amount     float64   `json:"amount"`
}

func (c *BidController) CreateBid(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.GetZap(ctx)
	var req CreateBidRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		HTTPErrorResponse(w, errors.NewValidationError("invalid request body", "req.body"))
		return
	}

	b, err := bid.New(req.InvoiceID, req.InvestorID, req.Amount)
	if err != nil {
		log.Error("error calling bid.New", zap.Error(err))
		HTTPErrorResponse(w, err)
		return
	}

	if b, err = c.service.CreateBid(ctx, b); err != nil {
		log.Error("error calling service.CreateBid", zap.Error(err))
		HTTPErrorResponse(w, err)
		return
	}

	response, _ := json.Marshal(b)
	HTTPResponse(w, http.StatusCreated, response)
}

func (c *BidController) GetBid(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.GetZap(ctx)

	id := mux.Vars(r)["id"]
	bidID, err := uuid.Parse(id)
	if err != nil {
		HTTPErrorResponse(w, errors.NewValidationError("invalid bid ID", "req.params.id"))
		return
	}
	b, err := c.service.GetBidByID(ctx, bidID)
	if err != nil {
		log.Error("error calling service.GetBidByID", zap.Error(err))
		HTTPErrorResponse(w, err)
		return
	}

	response, _ := json.Marshal(b)
	HTTPResponse(w, http.StatusOK, response)
}
