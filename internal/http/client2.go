package http

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

type client2Adapter struct {
	adaptee RequestAdapter
}

func (a *client2Adapter) SendWithContext(ctx context.Context, info RequestInformation, mapping ErrorMapping) (*http.Response, error) {
	return a.adaptee.Send(ctx, info)
}

func (a *client2Adapter) Send(info RequestInformation, mapping ErrorMapping) (*http.Response, error) {
	return a.SendWithContext(context.Background(), info, mapping)
}

func (a *client2Adapter) GetBaseURL() string {
	return a.adaptee.GetBaseURL()
}
