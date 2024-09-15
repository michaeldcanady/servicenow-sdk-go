package core

import (
	"context"
	"net/http"
)

type Client2 interface {
	ClientSendable
	Send(RequestInformation, ErrorMapping) (*http.Response, error)
	GetBaseURL() string
}

type ClientSendable interface {
	SendWithContext(context.Context, RequestInformation, ErrorMapping) (*http.Response, error)
}
