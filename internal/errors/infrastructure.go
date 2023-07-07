package errors

import "fmt"

// InfrastructureError represents an Infrastructure error such as database connection, HTTP request errors etc
type InfrastructureError struct {
	Message string
	Cause   error
}

func (e *InfrastructureError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("Error: %s. Cause: %s", e.Message, e.Cause)
	}
	return fmt.Sprintf("Error: %s", e.Message)
}

func NewInfrastructureError(message string, cause error) *InfrastructureError {
	return &InfrastructureError{Message: message, Cause: cause}
}
