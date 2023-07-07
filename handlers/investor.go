package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/tauki/invoiceexchange/internal/errors"
	"github.com/tauki/invoiceexchange/internal/investor"
	"github.com/tauki/invoiceexchange/router/middlewares/logging"
	"go.uber.org/zap"
	"net/http"
)

type InvestorController struct {
	service investor.Service
}

func NewInvestorController(service investor.Service) *InvestorController {
	return &InvestorController{
		service: service,
	}
}

type CreateInvestorRequest struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (c *InvestorController) CreateInvestor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.GetZap(ctx)

	var req CreateInvestorRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		HTTPErrorResponse(w, errors.NewValidationError("invalid request body", "req.body"))
		return
	}

	inv, err := investor.New(req.Name, req.Balance)
	if err != nil {
		HTTPErrorResponse(w, err)
		return
	}

	if inv, err = c.service.CreateInvestor(ctx, inv); err != nil {
		log.Error("error calling service.CreateInvestor", zap.Error(err))
		HTTPErrorResponse(w, err)
		return
	}

	response, _ := json.Marshal(inv)
	HTTPResponse(w, http.StatusCreated, response)
}

func (c *InvestorController) GetInvestor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.GetZap(ctx)

	id := mux.Vars(r)["id"]
	invID, err := uuid.Parse(id)
	if err != nil {
		HTTPErrorResponse(w, errors.NewValidationError("invalid investor ID", "req.params.id"))
		return
	}

	inv, err := c.service.GetInvestorByID(ctx, invID)
	if err != nil {
		log.Error("error calling service.GetInvestorByID", zap.Error(err))
		HTTPErrorResponse(w, err)
		return
	}

	response, _ := json.Marshal(inv)
	HTTPResponse(w, http.StatusOK, response)
}

func (c *InvestorController) ListInvestors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.GetZap(ctx)

	invs, err := c.service.ListInvestors(ctx)
	if err != nil {
		log.Error("error calling service.ListInvestors", zap.Error(err))
		HTTPErrorResponse(w, err)
		return
	}

	response, _ := json.Marshal(invs)
	HTTPResponse(w, http.StatusOK, response)
}
