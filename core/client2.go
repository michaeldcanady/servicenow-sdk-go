package core

import (
	"context"
	"net/http"
)

// Deprecated: deprecated since v{unreleased}.
//
// Client2 ...
type Client2 interface {
	Send(IRequestInformation, ErrorMapping) (*http.Response, error)
	SendWithContext(context.Context, IRequestInformation, ErrorMapping) (*http.Response, error)
	GetBaseURL() string
}
