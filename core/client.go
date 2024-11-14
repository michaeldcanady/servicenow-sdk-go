package core

import (
	"context"
	"net/http"
)

type Client interface {
	Send(ctx context.Context, requestInfo IRequestInformation, errorMapping ErrorMapping) (*http.Response, error)
}
