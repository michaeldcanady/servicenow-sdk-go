package core

import (
	"context"
	"net/http"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

type Client2 interface {
	Send(internal.RequestInformation, internal.ErrorMapping) (*http.Response, error)
	SendWithContext(context.Context, internal.RequestInformation, internal.ErrorMapping) (*http.Response, error)
	GetBaseURL() string
}
