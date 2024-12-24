package http

import (
	"context"
	"net/http"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

type RequestInformation interface {
	SetStreamContent(content []byte)
	GetContent() []byte
	GetMethod() string
	GetHeaders() core.RequestHeader //nolint: staticcheck
	AddQueryParameters(source interface{}) error
	SetUri(url *url.URL)
	Url() (string, error)
	ToRequest() (*http.Request, error)
	ToRequestWithContext(ctx context.Context) (*http.Request, error)
	AddHeaders(rawHeaders interface{}) error
}
