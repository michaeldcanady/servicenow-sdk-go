package core

import (
	"context"
	"net/http"
)

type Client2 interface {
	Send(RequestInformation, ErrorMapping) (*http.Response, error)
	SendWithContext(context.Context, RequestInformation, ErrorMapping) (*http.Response, error)
	GetBaseURL() string
}
