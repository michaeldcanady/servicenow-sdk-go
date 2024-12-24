package core

import (
	"context"
	"net/http"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

// Deprecated: deprecated since v{unreleased}.
type IRequestInformation interface {
	AddRequestOptions(options []RequestOption)
	SetStreamContent(content []byte)
	AddQueryParameters(source interface{}) error
	SetUri(url *url.URL)
	Url() (string, error)
	ToRequest() (*http.Request, error)
	ToRequestWithContext(ctx context.Context) (*http.Request, error)
	AddHeaders(rawHeaders interface{}) error
	GetRequestOptions() []RequestOption
	GetContent() []byte
	GetMethod() string
	GetHeaders() core.RequestHeader //nolint: staticcheck
}
