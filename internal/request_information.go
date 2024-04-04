package internal

import (
	"context"
	"net/http"
	"net/url"
)

type RequestInformation interface {
	SetStreamContent(content []byte)
	AddQueryParameters(source interface{}) error
	SetUri(url *url.URL)  //nolint:stylecheck
	Url() (string, error) //nolint:stylecheck
	ToRequest() (*http.Request, error)
	ToRequestWithContext(ctx context.Context) (*http.Request, error)
	AddHeaders(rawHeaders interface{}) error
}
