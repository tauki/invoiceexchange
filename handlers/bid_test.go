package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tauki/invoiceexchange/handlers"
	"github.com/tauki/invoiceexchange/internal/bid"
)

func TestCreateBid_CreateBid(t *testing.T) {
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewBidController(bidService)
	router := mux.NewRouter()
	router.HandleFunc("/bid", GetTestHandleFunc(controller.CreateBid))

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	reqBody, _ := json.Marshal(handlers.CreateBidRequest{
		InvoiceID:  preloadedIDs["invoice"],
		InvestorID: preloadedIDs["investor"],
		Amount:     100,
	})

	response, err := http.Post(testServer.URL+"/bid", "application/json", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	require.Equal(t, http.StatusCreated, response.StatusCode)

	var createdBid bid.Bid
	require.NoError(t, json.NewDecoder(response.Body).Decode(&createdBid))
}

func TestCreateBid_CreateBidWithInvalidInvoiceID(t *testing.T) {
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewBidController(bidService)
	router := mux.NewRouter()
	router.HandleFunc("/bid", GetTestHandleFunc(controller.CreateBid))

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	reqBody, _ := json.Marshal(handlers.CreateBidRequest{
		InvoiceID:  uuid.New(),
		InvestorID: preloadedIDs["investor"],
		Amount:     100,
	})

	response, err := http.Post(testServer.URL+"/bid", "application/json", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	require.Equal(t, http.StatusNotFound, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	require.Equal(t, `{"status":404,"description":"invoice not found"}`, string(body))
}

func TestCreateBid_CreateBidWithInvalidInvestorID(t *testing.T) {
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewBidController(bidService)
	router := mux.NewRouter()
	router.HandleFunc("/bid", GetTestHandleFunc(controller.CreateBid))

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	reqBody, _ := json.Marshal(handlers.CreateBidRequest{
		InvoiceID:  preloadedIDs["invoice"],
		InvestorID: uuid.New(),
		Amount:     100,
	})

	response, err := http.Post(testServer.URL+"/bid", "application/json", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	require.Equal(t, http.StatusNotFound, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	require.Equal(t, `{"status":404,"description":"investor not found"}`, string(body))
}

func TestBidController_GetBid(t *testing.T) {
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewBidController(bidService)
	router := mux.NewRouter()
	router.HandleFunc("/bid/{id}", GetTestHandleFunc(controller.GetBid))

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	response, err := http.Get(testServer.URL + "/bid/" + preloadedIDs["bid"].String())
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, response.StatusCode)

	var bid bid.Bid
	require.NoError(t, json.NewDecoder(response.Body).Decode(&bid))
	require.Equal(t, preloadedIDs["bid"], bid.ID)
}

func TestBidController_GetBidWithInvalidID(t *testing.T) {
	bidService, _ := GetTestBidService(t)
	controller := handlers.NewBidController(bidService)
	router := mux.NewRouter()
	router.HandleFunc("/bid/{id}", GetTestHandleFunc(controller.GetBid))

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	response, err := http.Get(testServer.URL + "/bid/" + uuid.New().String())
	require.NoError(t, err)

	require.Equal(t, http.StatusNotFound, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	require.Equal(t, `{"status":404,"description":"bid not found"}`, string(body))
}
