package errors

import "fmt"

// PermissionDeniedError represents a permission denied error.
// It indicates the caller does not have permission to execute the specified operation
type PermissionDeniedError struct {
	Message string
	Cause   error
}

func (e *PermissionDeniedError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("Error: %s. Cause: %s", e.Message, e.Cause)
	}
	return fmt.Sprintf("Error: %s", e.Message)
}

func NewPermissionDeniedError(message string, cause error) *PermissionDeniedError {
	return &PermissionDeniedError{Message: message, Cause: cause}
}
