//go:build e2e
// +build e2e

package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"github.com/tauki/invoiceexchange/handlers"
	"github.com/tauki/invoiceexchange/internal/investor"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvestorController_CreateInvestor(t *testing.T) {
	invService, _ := GetTestInvestorService(t)
	controller := handlers.NewInvestorController(invService)
	router := mux.NewRouter()
	router.HandleFunc("/investor", GetTestHandleFunc(controller.CreateInvestor)).Methods("POST")

	testServer := httptest.NewServer(router)

	reqBody, _ := json.Marshal(handlers.CreateInvestorRequest{
		Name:    "Test Investor",
		Balance: 10000.0,
	})

	response, err := http.Post(testServer.URL+"/investor", "application/json", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	require.Equal(t, http.StatusCreated, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var createdInvestor investor.Investor
	require.NoError(t, json.NewDecoder(bytes.NewReader(body)).Decode(&createdInvestor))
}

func TestInvestorController_GetInvestor(t *testing.T) {
	invService, _ := GetTestInvestorService(t)
	controller := handlers.NewInvestorController(invService)
	router := mux.NewRouter()
	router.HandleFunc("/investor/{id}", GetTestHandleFunc(controller.GetInvestor)).Methods("GET")
	testServer := httptest.NewServer(router)

	response, err := http.Get(testServer.URL + "/investor/" + preloadedIDs["investor"].String())
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var retrievedInvestor investor.Investor
	require.NoError(t, json.NewDecoder(bytes.NewReader(body)).Decode(&retrievedInvestor))
}

func TestInvestorController_GetInvestor_InvalidID(t *testing.T) {
	invService, _ := GetTestInvestorService(t)
	controller := handlers.NewInvestorController(invService)
	router := mux.NewRouter()
	router.HandleFunc("/investor/{id}", GetTestHandleFunc(controller.GetInvestor)).Methods("GET")
	testServer := httptest.NewServer(router)

	response, err := http.Get(testServer.URL + "/investor/invalidID")
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestIssuerController_GetInvestor_NotFound(t *testing.T) {
	invService, _ := GetTestInvestorService(t)
	controller := handlers.NewInvestorController(invService)
	router := mux.NewRouter()
	router.HandleFunc("/investor/{id}", GetTestHandleFunc(controller.GetInvestor)).Methods("GET")
	testServer := httptest.NewServer(router)

	response, err := http.Get(testServer.URL + "/investor/" + uuid.New().String())
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestInvestorController_ListInvestors(t *testing.T) {
	invService, _ := GetTestInvestorService(t)
	controller := handlers.NewInvestorController(invService)
	router := mux.NewRouter()
	router.HandleFunc("/investors", GetTestHandleFunc(controller.ListInvestors)).Methods("GET")
	testServer := httptest.NewServer(router)

	response, err := http.Get(testServer.URL + "/investors")
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var investors []investor.Investor
	require.NoError(t, json.NewDecoder(bytes.NewReader(body)).Decode(&investors))

	// We should have at least one investor
	require.Greater(t, len(investors), 0)
}
