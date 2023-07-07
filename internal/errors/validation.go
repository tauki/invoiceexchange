package errors

import "fmt"

// ValidationError represents a data validation error
type ValidationError struct {
	Message   string
	FieldName string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Field: %s. Error: %s", e.FieldName, e.Message)
}

func NewValidationError(message, fieldName string) *ValidationError {
	return &ValidationError{Message: message, FieldName: fieldName}
}
