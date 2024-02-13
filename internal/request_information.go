package internal

import (
	"context"
	"net/http"
	"net/url"
)

type RequestInformation interface {
	SetStreamContent(content []byte)
	GetContent() []byte
	GetMethod() string
	AddQueryParameters(source interface{}) error
	SetUri(url *url.URL)
	Url() (string, error)
	ToRequest() (*http.Request, error)
	ToRequestWithContext(ctx context.Context) (*http.Request, error)
	AddHeaders(rawHeaders interface{}) error
	GetHeaders() RequestHeader
}
