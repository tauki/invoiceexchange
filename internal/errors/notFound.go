package errors

import "fmt"

// NotFoundError represents a not found error
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{Message: message}
}
