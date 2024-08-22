package core

import (
	"context"
	"net/http"
)

type Client2 interface {
	SendableWithContext
	Send(RequestInformation, ErrorMapping) (*http.Response, error)
	GetBaseURL() string
}

type SendableWithContext interface {
	SendWithContext(context.Context, RequestInformation, ErrorMapping) (*http.Response, error)
}
