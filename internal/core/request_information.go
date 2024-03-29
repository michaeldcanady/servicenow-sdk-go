package core

import (
	"context"
	"net/http"
	"net/url"
)

type RequestInformation interface {
	SetStreamContent(content []byte)
	GetContent() []byte
	GetMethod() string
	GetHeaders() RequestHeader
	AddQueryParameters(source interface{}) error
	SetUri(url *url.URL)
	Url() (string, error)
	ToRequest() (*http.Request, error)
	ToRequestWithContext(ctx context.Context) (*http.Request, error)
	AddHeaders(rawHeaders interface{}) error
}
