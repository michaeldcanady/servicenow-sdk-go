package core

import (
	"context"
	"net/http"
)

type Client2 interface {
	Send(IRequestInformation, ErrorMapping) (*http.Response, error)
	SendWithContext(context.Context, IRequestInformation, ErrorMapping) (*http.Response, error)
	GetBaseURL() string
}
