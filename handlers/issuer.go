package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/tauki/invoiceexchange/internal/errors"
	"github.com/tauki/invoiceexchange/internal/issuer"
	"github.com/tauki/invoiceexchange/router/middlewares/logging"
)

type IssuerController struct {
	service issuer.Service
}

func NewIssuerController(service issuer.Service) *IssuerController {
	return &IssuerController{
		service: service,
	}
}

type CreateIssuerRequest struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (c *IssuerController) CreateIssuer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.GetZap(ctx)

	var req CreateIssuerRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		HTTPErrorResponse(w, errors.NewValidationError("invalid request body", "req.body"))
		return
	}

	inv, err := issuer.New(req.Name, req.Balance)
	if err != nil {
		HTTPErrorResponse(w, err)
		return
	}

	if inv, err = c.service.CreateIssuer(ctx, inv); err != nil {
		log.Error("error calling service.CreateIssuer", zap.Error(err))
		HTTPErrorResponse(w, err)
		return
	}

	response, _ := json.Marshal(inv)
	HTTPResponse(w, http.StatusCreated, response)
}

func (c *IssuerController) GetIssuer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.GetZap(ctx)

	id := mux.Vars(r)["id"]
	invID, err := uuid.Parse(id)
	if err != nil {
		HTTPErrorResponse(w, errors.NewValidationError("invalid issuer ID", "req.params.id"))
		return
	}

	inv, err := c.service.GetIssuerByID(ctx, invID)
	if err != nil {
		log.Error("error calling service.GetIssuerByID", zap.Error(err))
		HTTPErrorResponse(w, err)
		return
	}

	response, _ := json.Marshal(inv)
	HTTPResponse(w, http.StatusOK, response)
}

func (c *IssuerController) ListIssuer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.GetZap(ctx)

	invs, err := c.service.ListIssuers(ctx)
	if err != nil {
		log.Error("error calling service.ListIssuers", zap.Error(err))
		HTTPErrorResponse(w, err)
		return
	}

	response, _ := json.Marshal(invs)
	HTTPResponse(w, http.StatusOK, response)
}
