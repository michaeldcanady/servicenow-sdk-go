package batchapi

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicedRequest represents Service-Now Batch API response's serviced request.
type ServicedRequest interface {
	GetBodyAsParsable(serialization.ParsableFactory) (serialization.Parsable, error)
	GetBody() ([]byte, error)
	GetErrorMessage() (*string, error)
	GetExecutionTime() (*serialization.ISODuration, error)
	GetHeaders() ([]RestRequestHeader, error)
	GetID() (*string, error)
	GetRedirectURL() (*string, error)
	GetStatusCode() (*int64, error)
	GetStatusText() (*string, error)
	serialization.Parsable
}
