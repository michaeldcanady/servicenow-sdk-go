package core

import (
	"context"
	"net/http"
)

// Client2 interface designed to match core.Client2
type Client2 interface {
	ClientSendable
	Send(RequestInformation, ErrorMapping) (*http.Response, error)
	GetBaseURL() string
}

type ClientSendable interface {
	SendWithContext(context.Context, RequestInformation, ErrorMapping) (*http.Response, error)
}
