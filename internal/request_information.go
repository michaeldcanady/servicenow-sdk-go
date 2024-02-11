package internal

import (
	"context"
	"net/http"
	"net/url"
)

type RequestInformation interface {
	AddRequestOptions([]RequestOption)
	GetRequestOptions() []RequestOption
	SetStreamContent([]byte)
	AddQueryParameters(interface{}) error
	SetURI(*url.URL)
	URL() (string, error)
	ToRequest() (*http.Request, error)
	ToRequestWithContext(context.Context) (*http.Request, error)
	AddHeaders(interface{}) error
}
