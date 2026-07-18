package errors

import "errors"

// Standard sentinel errors for consistent error handling.
var (
	ErrNilRequestAdapter       = errors.New("requestAdapter cannot be nil")
	ErrNilRequestBuilder       = errors.New("request builder is nil")
	ErrNilResponse             = errors.New("response cannot be nil")
	ErrNilContext              = errors.New("context cannot be nil")
	ErrNilConfig               = errors.New("config is nil")
	ErrNilBody                 = errors.New("body is nil")
	ErrNilInput                = errors.New("input is nil")
	ErrNilRequestConfiguration = errors.New("requestConfiguration is nil")
	ErrNilQueryParameters      = errors.New("requestConfiguration.QueryParameters is nil")
	ErrNilFactory              = errors.New("factory is nil")
	ErrNilStore                = errors.New("store is nil")
	ErrNilPathParameters       = errors.New("pathParameters is nil")
	ErrEmptyPathParameters     = errors.New("pathParameters is empty")
	ErrNilMutator              = errors.New("mutator is nil")
	ErrNilModel                = errors.New("model is nil")
	ErrEmptyMiddleware         = errors.New("middleware is empty")
	ErrEmptyKey                = errors.New("key is empty")
)

// NewValidationError creates a standardized validation error message.
func NewValidationError(parameter string) error {
	return errors.New(parameter + " cannot be nil")
}
