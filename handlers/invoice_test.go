//go:build e2e
// +build e2e

package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/tauki/invoiceexchange/handlers"
	"github.com/tauki/invoiceexchange/internal/bid"
	"github.com/tauki/invoiceexchange/internal/invoice"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestInvoiceController_CreateInvoice(t *testing.T) {
	invService, _ := GetTestInvoiceService(t)
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewInvoiceController(invService, bidService)
	router := mux.NewRouter()
	router.HandleFunc("/invoice", GetTestHandleFunc(controller.CreateInvoice)).Methods("POST")

	testServer := httptest.NewServer(router)

	reqBody, _ := json.Marshal(handlers.CreateInvoiceRequest{
		InvoiceDate:   time.Now(),
		InvoiceNumber: "1234",
		Reference:     "Test Reference",
		CompanyName:   "Test Company",
		CustomerName:  "Test Customer",
		Currency:      "USD",
		TotalAmount:   10000.0,
		TotalVat:      2000.0,
		DueDate:       time.Now().AddDate(0, 0, 30),
		IssuerID:      preloadedIDs["issuer"],
		AskingPrice:   9000.0,
	})

	response, err := http.Post(testServer.URL+"/invoice", "application/json", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	require.Equal(t, http.StatusCreated, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var createdInvoice invoice.Invoice
	require.NoError(t, json.NewDecoder(bytes.NewReader(body)).Decode(&createdInvoice))
}

func TestInvoiceController_CreateInvoice_WithInvalidIssuer(t *testing.T) {
	invService, _ := GetTestInvoiceService(t)
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewInvoiceController(invService, bidService)
	router := mux.NewRouter()
	router.HandleFunc("/invoice", GetTestHandleFunc(controller.CreateInvoice)).Methods("POST")

	testServer := httptest.NewServer(router)

	reqBody, _ := json.Marshal(handlers.CreateInvoiceRequest{
		InvoiceDate:   time.Now(),
		InvoiceNumber: "1234",
		Reference:     "Test Reference",
		CompanyName:   "Test Company",
		CustomerName:  "Test Customer",
		Currency:      "USD",
		TotalAmount:   10000.0,
		TotalVat:      2000.0,
		DueDate:       time.Now().AddDate(0, 0, 30),
		IssuerID:      uuid.New(),
		AskingPrice:   9000.0,
	})

	response, err := http.Post(testServer.URL+"/invoice", "application/json", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	require.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	require.Equal(t, `{"status":400,"description":"issuer does not exist","field":"invoice.IssuerID"}`, string(body))
}

func TestInvoiceController_GetInvoice(t *testing.T) {
	invService, _ := GetTestInvoiceService(t)
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewInvoiceController(invService, bidService)
	router := mux.NewRouter()
	router.HandleFunc("/invoice/{id}", GetTestHandleFunc(controller.GetInvoice)).Methods("GET")
	testServer := httptest.NewServer(router)

	response, err := http.Get(testServer.URL + "/invoice/" + preloadedIDs["invoice"].String())
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var retrievedInvoice invoice.Invoice
	require.NoError(t, json.NewDecoder(bytes.NewReader(body)).Decode(&retrievedInvoice))
}

func TestInvoiceController_GetInvoice_NotFound(t *testing.T) {
	invService, _ := GetTestInvoiceService(t)
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewInvoiceController(invService, bidService)
	router := mux.NewRouter()
	router.HandleFunc("/invoice/{id}", GetTestHandleFunc(controller.GetInvoice)).Methods("GET")
	testServer := httptest.NewServer(router)

	response, err := http.Get(testServer.URL + "/invoice/" + uuid.New().String())
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestInvoiceController_GetInvoice_WithInvalidID(t *testing.T) {
	invService, _ := GetTestInvoiceService(t)
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewInvoiceController(invService, bidService)
	router := mux.NewRouter()
	router.HandleFunc("/invoice/{id}", GetTestHandleFunc(controller.GetInvoice)).Methods("GET")
	testServer := httptest.NewServer(router)

	response, err := http.Get(testServer.URL + "/invoice/invalid-id")
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestInvoiceController_GetBids(t *testing.T) {
	invService, _ := GetTestInvoiceService(t)
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewInvoiceController(invService, bidService)
	router := mux.NewRouter()
	router.HandleFunc("/invoice/{id}/bids", GetTestHandleFunc(controller.GetBids)).Methods("GET")
	testServer := httptest.NewServer(router)

	response, err := http.Get(testServer.URL + "/invoice/" + preloadedIDs["invoice"].String() + "/bids")
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var bids []bid.Bid
	require.NoError(t, json.NewDecoder(bytes.NewReader(body)).Decode(&bids))

	// We should have at least one bid
	require.Greater(t, len(bids), 0)
}

func TestInvoiceController_GetBids_NotFound(t *testing.T) {
	invService, _ := GetTestInvoiceService(t)
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewInvoiceController(invService, bidService)
	router := mux.NewRouter()
	router.HandleFunc("/invoice/{id}/bids", GetTestHandleFunc(controller.GetBids)).Methods("GET")
	testServer := httptest.NewServer(router)

	response, err := http.Get(testServer.URL + "/invoice/" + uuid.New().String() + "/bids")
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestInvoiceController_ApproveTrade_PermissionDenied(t *testing.T) {
	invService, _ := GetTestInvoiceService(t)
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewInvoiceController(invService, bidService)
	router := mux.NewRouter()
	router.HandleFunc("/invoice/{id}/approve", GetTestHandleFunc(controller.ApproveTrade)).Methods("PATCH")
	testServer := httptest.NewServer(router)

	req, err := http.NewRequest("PATCH", testServer.URL+"/invoice/"+preloadedIDs["invoice"].String()+"/approve", nil)
	require.NoError(t, err)

	client := &http.Client{}
	response, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusForbidden, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	require.Equal(t, `{"status":403,"description":"trade cannot be approved, invoice is not locked"}`, string(body))
}

func TestInvoiceController_ApproveTrade_Approved(t *testing.T) {
	invService, _ := GetTestInvoiceService(t)
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewInvoiceController(invService, bidService)
	router := mux.NewRouter()
	router.HandleFunc("/invoice/{id}/approve", GetTestHandleFunc(controller.ApproveTrade)).Methods("PATCH")
	testServer := httptest.NewServer(router)

	req, err := http.NewRequest("PATCH", testServer.URL+"/invoice/"+preloadedIDs["invoice-locked"].String()+"/approve", nil)
	require.NoError(t, err)

	client := &http.Client{}
	response, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusAccepted, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var inv invoice.Invoice
	require.NoError(t, json.NewDecoder(bytes.NewReader(body)).Decode(&inv))
}
