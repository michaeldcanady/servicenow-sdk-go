package core

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/hetiansu5/urlquery"
)

// RequestInformation represents an abstract HTTP request.
type RequestInformation struct {
	// The HTTP method of the request.
	Method HttpMethod
	// The Request Headers.
	Headers *RequestHeaders
	// The Request Body.
	Content []byte
	options map[string]RequestOption
	uri     *UrlInformation
}

// NewRequestInformation creates a new RequestInformation object with default values.
func NewRequestInformation() *RequestInformation {
	return &RequestInformation{
		Headers: NewRequestHeaders(),
		options: make(map[string]RequestOption),
		uri:     NewUrlInformation(),
	}
}

// AddRequestOptions adds an option to the request to be read by the middleware infrastructure.
func (rI *RequestInformation) AddRequestOptions(options []RequestOption) {
	if options == nil {
		return
	}
	if rI.options == nil {
		rI.options = make(map[string]RequestOption, len(options))
	}
	for _, option := range options {
		rI.options[option.GetKey().Key] = option
	}
}

// GetRequestOptions returns the options for this request. Options are unique by type. If an option of the same type is added twice, the last one wins.
func (rI *RequestInformation) GetRequestOptions() []RequestOption {
	if rI.options == nil {
		return []RequestOption{}
	}
	result := make([]RequestOption, len(rI.options))
	idx := 0

	for _, option := range rI.options {
		result[idx] = option
		idx++
	}
	return result
}

// SetStreamContent sets the request body to a binary stream.
func (rI *RequestInformation) SetStreamContent(content []byte) {
	rI.Content = content
	if rI.Headers != nil {
		rI.Headers.Add(contentTypeHeader, binaryContentType)
	}
}

// AddQueryParameters adds the query parameters to the request by reading the properties from the provided object.
func (rI *RequestInformation) AddQueryParameters(source interface{}) error {
	return rI.uri.AddQueryParameters(source)
}

func toQueryMap(source interface{}) (map[string]string, error) {
	if source == nil {
		return nil, errors.New("source is nil")
	}

	queryBytes, err := urlquery.Marshal(source)
	if err != nil {
		return nil, err
	}

	var queryMap map[string]string
	err = urlquery.Unmarshal(queryBytes, &queryMap)
	if err != nil {
		return nil, err
	}

	return queryMap, nil
}

func (rI *RequestInformation) getContentReader() *bytes.Reader {
	return bytes.NewReader(rI.Content)
}

func (rI *RequestInformation) SetUri(url *url.URL) {
	rI.uri.PathParameters = map[string]string{"request-raw-url": url.String()}
}

// ToRequest converts the RequestInformation object into an HTTP request.
func (rI *RequestInformation) ToRequest() (*http.Request, error) {

	uri, err := rI.uri.ToUrl()
	if err != nil {
		return nil, err
	}

	url := uri.String()
	contentReader := rI.getContentReader()
	methodString := rI.Method.String()

	req, err := http.NewRequest(methodString, url, contentReader)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ToRequestWithContext converts the RequestInformation object into an HTTP request with context.
func (rI *RequestInformation) ToRequestWithContext(ctx context.Context) (*http.Request, error) {
	uri, err := rI.uri.ToUrl()
	if err != nil {
		return nil, err
	}

	url := uri.String()
	contentReader := rI.getContentReader()
	methodString := rI.Method.String()

	req, err := http.NewRequestWithContext(ctx, methodString, url, contentReader)
	if err != nil {
		return nil, err
	}
	return req, nil
}
