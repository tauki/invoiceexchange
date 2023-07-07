package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tauki/invoiceexchange/internal/errors"
)

func HTTPResponse(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func HTTPErrorResponse(w http.ResponseWriter, err error) {
	details := ToHTTPError(err)
	response, _ := json.Marshal(details)
	HTTPResponse(w, details.Status, response)
}

type ErrorDetails struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Field       string `json:"field,omitempty"`
}

func ToHTTPError(domainErr error) *ErrorDetails {
	switch err := domainErr.(type) {
	case *errors.NotFoundError:
		return &ErrorDetails{
			Status:      http.StatusNotFound,
			Description: err.Message,
		}
	case *errors.ValidationError:
		return &ErrorDetails{
			Status:      http.StatusBadRequest,
			Description: err.Message,
			Field:       err.FieldName,
		}
	case *errors.PermissionDeniedError:
		return &ErrorDetails{
			Status:      http.StatusForbidden,
			Description: err.Message,
		}
	default:
		return &ErrorDetails{
			Status:      http.StatusInternalServerError,
			Description: "Internal Server Error",
		}
	}
}
