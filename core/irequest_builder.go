package core

import (
	"context"
	"net/http"
	"net/url"
)

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
}
