package recovery

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleware_ErrorOnPanic(t *testing.T) {
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("something went wrong")
	})

	handler := Middleware(handlerFunc)
	req := httptest.NewRequest("GET", "http://testing", nil)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Result().Header.Get("content-type"))
	assert.JSONEq(t, `{"description":"Internal Server Error", "status":500}`, recorder.Body.String())
}

func TestMiddleware_SuccessOnNoPanic(t *testing.T) {
	message := "OK, no panic"

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})

	handler := Middleware(handlerFunc)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, nil)

	body, err := io.ReadAll(recorder.Result().Body)
	assert.Nil(t, err)
	defer recorder.Result().Body.Close()

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, message, string(body))
}
