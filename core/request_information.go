package core

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"reflect"

	"github.com/RecoLabs/servicenow-sdk-go/internal"
	"github.com/RecoLabs/servicenow-sdk-go/internal/core"
	"github.com/mozillazg/go-httpheader"
)

// RequestInformation represents an abstract HTTP request.
type RequestInformation struct {
	// The HTTP method of the request.
	Method HttpMethod
	// The Request Headers.
	Headers http.Header
	// The Request Body.
	Content []byte
	options map[string]RequestOption
	uri     *UrlInformation
}

// NewRequestInformation creates a new RequestInformation object with default values.
func NewRequestInformation() *RequestInformation {
	return &RequestInformation{
		Headers: make(http.Header),
		options: make(map[string]RequestOption),
		uri:     NewURLInformation(),
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

func (rI *RequestInformation) GetContent() []byte {
	return rI.Content
}

func (rI *RequestInformation) GetHeaders() core.RequestHeader {
	headers := internal.NewRequestHeader()
	headers.SetAll(rI.Headers)
	return headers
}

func (rI *RequestInformation) GetMethod() string {
	return rI.Method.String()
}

// AddQueryParameters adds the query parameters to the request by reading the properties from the provided object.
func (rI *RequestInformation) AddQueryParameters(source interface{}) error {
	return rI.uri.AddQueryParameters(source)
}

func (rI *RequestInformation) getContentReader() *bytes.Reader {
	return bytes.NewReader(rI.Content)
}

func (rI *RequestInformation) SetUri(url *url.URL) { //nolint:stylecheck
	//TODO: Add validation that url is valid

	rI.uri.PathParameters = map[string]string{rawURLKey: url.String()}
}

func (rI *RequestInformation) Url() (string, error) { //nolint:stylecheck
	uri, err := rI.uri.ToUrl()
	if err != nil {
		return "", err
	}

	url := uri.String()

	return url, nil
}

// ToRequest converts the RequestInformation object into an HTTP request.
func (rI *RequestInformation) ToRequest() (*http.Request, error) {
	url, err := rI.Url()
	if err != nil {
		return nil, err
	}

	contentReader := rI.getContentReader()
	methodString := rI.Method.String()

	req, err := http.NewRequest(methodString, url, contentReader)
	if err != nil {
		return nil, err
	}

	for key, values := range rI.Headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	return req, nil
}

// ToRequestWithContext converts the RequestInformation object into an HTTP request with context.
func (rI *RequestInformation) ToRequestWithContext(ctx context.Context) (*http.Request, error) {
	url, err := rI.Url()
	if err != nil {
		return nil, err
	}
	contentReader := rI.getContentReader()
	methodString := rI.Method.String()

	req, err := http.NewRequestWithContext(ctx, methodString, url, contentReader)
	req.Header = rI.Headers
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (rI *RequestInformation) AddHeaders(rawHeaders interface{}) error {
	var headers http.Header
	var err error

	val := reflect.ValueOf(rawHeaders)

	if val.Kind() == reflect.Struct {
		// use the httpheader.Encode function from the httpheader package
		// to encode the pointer value into an http.Header map
		headers, err = httpheader.Encode(rawHeaders)
		if err != nil {
			return err
		}
	} else if val.Type() == reflect.TypeOf(http.Header{}) {
		// if the value is already an http.Header map, just assign it to headers
		headers = rawHeaders.(http.Header)
	} else {
		// otherwise, return an error
		return ErrInvalidHeaderType
	}

	// iterate over the headers map and add each key-value pair to rI.Headers
	for key, values := range headers {
		for _, value := range values {
			rI.Headers.Add(key, value)
		}
	}
	return nil
}
