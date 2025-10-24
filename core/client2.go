package core

import (
	"context"
	"net/http"
)

type Client2 interface {
	Client
	SendWithContext(context.Context, IRequestInformation, ErrorMapping) (*http.Response, error)
}
