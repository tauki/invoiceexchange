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

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"github.com/tauki/invoiceexchange/handlers"
	"github.com/tauki/invoiceexchange/internal/investor"
)

func TestIssuerController_CreateIssuer(t *testing.T) {
	issService, _ := GetTestIssuerService(t)
	controller := handlers.NewIssuerController(issService)
	router := mux.NewRouter()
	router.HandleFunc("/issuer", GetTestHandleFunc(controller.CreateIssuer)).Methods("POST")

	testServer := httptest.NewServer(router)

	reqBody, _ := json.Marshal(handlers.CreateInvestorRequest{
		Name:    "Test Issuer",
		Balance: 10000.0,
	})

	response, err := http.Post(testServer.URL+"/issuer", "application/json", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	require.Equal(t, http.StatusCreated, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var createdInvestor investor.Investor
	require.NoError(t, json.NewDecoder(bytes.NewReader(body)).Decode(&createdInvestor))
}

func TestIssuerController_GetIssuer(t *testing.T) {
	issService, _ := GetTestIssuerService(t)
	controller := handlers.NewIssuerController(issService)
	router := mux.NewRouter()
	router.HandleFunc("/issuer/{id}", GetTestHandleFunc(controller.GetIssuer)).Methods("GET")
	testServer := httptest.NewServer(router)

	response, err := http.Get(testServer.URL + "/issuer/" + preloadedIDs["issuer"].String())
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var retrievedInvestor investor.Investor
	require.NoError(t, json.NewDecoder(bytes.NewReader(body)).Decode(&retrievedInvestor))
}

func TestIssuerController_GetIssuer_InvalidID(t *testing.T) {
	issService, _ := GetTestIssuerService(t)
	controller := handlers.NewIssuerController(issService)
	router := mux.NewRouter()
	router.HandleFunc("/issuer/{id}", GetTestHandleFunc(controller.GetIssuer)).Methods("GET")
	testServer := httptest.NewServer(router)

	response, err := http.Get(testServer.URL + "/issuer/invalidID")
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestIssuerController_GetIssuer_NotFound(t *testing.T) {
	issService, _ := GetTestIssuerService(t)
	controller := handlers.NewIssuerController(issService)
	router := mux.NewRouter()
	router.HandleFunc("/issuer/{id}", GetTestHandleFunc(controller.GetIssuer)).Methods("GET")
	testServer := httptest.NewServer(router)

	response, err := http.Get(testServer.URL + "/issuer/" + uuid.New().String())
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestIssuerController_ListIssuer(t *testing.T) {
	issService, _ := GetTestIssuerService(t)
	controller := handlers.NewIssuerController(issService)
	router := mux.NewRouter()
	router.HandleFunc("/issuers", GetTestHandleFunc(controller.ListIssuer)).Methods("GET")
	testServer := httptest.NewServer(router)

	response, err := http.Get(testServer.URL + "/issuers")
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var investors []investor.Investor
	require.NoError(t, json.NewDecoder(bytes.NewReader(body)).Decode(&investors))

	// We should have at least one investor
	require.Greater(t, len(investors), 0)
}
