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

type requestInformation struct {
}

func (rI *requestInformation) SetStreamContent(content []byte) {

}
func (rI *requestInformation) AddQueryParameters(source interface{}) error {
	return nil
}
func (rI *requestInformation) SetUri(url *url.URL) {

}
func (rI *requestInformation) Url() (string, error) {
	return "", nil
}
func (rI *requestInformation) ToRequest() (*http.Request, error) {
	return rI.ToRequestWithContext(context.Background())
}
func (rI *requestInformation) ToRequestWithContext(ctx context.Context) (*http.Request, error) {
	return nil, nil
}
func (rI *requestInformation) AddHeaders(rawHeaders interface{}) error {
	return nil
}
func (rI *requestInformation) GetHeaders() RequestHeader {
	return NewRequestHeader()
}
