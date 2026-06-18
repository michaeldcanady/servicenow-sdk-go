package errors

import "errors"

// Standard sentinel errors for consistent error handling.
var (
	ErrNilRequestAdapter = errors.New("requestAdapter cannot be nil")
	ErrNilResponse       = errors.New("response cannot be nil")
	ErrNilContext        = errors.New("context cannot be nil")
)

// NewValidationError creates a standardized validation error message.
func NewValidationError(parameter string) error {
	return errors.New(parameter + " cannot be nil")
}
