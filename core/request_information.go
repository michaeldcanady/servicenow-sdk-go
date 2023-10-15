package core

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/hetiansu5/urlquery"
	t "github.com/yosida95/uritemplate/v3"
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

// addParameterWithOriginalName adds the URI template parameter to the template using the right casing, because of Go conventions, casing might have changed for the generated property.
func addParameterWithOriginalName(key string, value string, normalizedNames map[string]string, values t.Values) {
	paramName := getOriginalParameterName(key, normalizedNames)
	values.Set(paramName, t.String(value))
}

func getOriginalParameterName(key string, normalizedNames map[string]string) string {
	lowercaseKey := strings.ToLower(key)
	if paramName, ok := normalizedNames[lowercaseKey]; ok {
		return paramName
	}
	return key
}

// AddRequestOptions adds an option to the request to be read by the middleware infrastructure.
func (request *RequestInformation) AddRequestOptions(options []RequestOption) {
	if options == nil {
		return
	}
	if request.options == nil {
		request.options = make(map[string]RequestOption, len(options))
	}
	for _, option := range options {
		request.options[option.GetKey().Key] = option
	}
}

// GetRequestOptions returns the options for this request. Options are unique by type. If an option of the same type is added twice, the last one wins.
func (request *RequestInformation) GetRequestOptions() []RequestOption {
	if request.options == nil {
		return []RequestOption{}
	}
	result := make([]RequestOption, len(request.options))
	idx := 0

	for _, option := range request.options {
		result[idx] = option
		idx++
	}
	return result
}

// SetStreamContent sets the request body to a binary stream.
func (request *RequestInformation) SetStreamContent(content []byte) {
	request.Content = content
	if request.Headers != nil {
		request.Headers.Add(contentTypeHeader, binaryContentType)
	}
}

// AddQueryParameters adds the query parameters to the request by reading the properties from the provided object.
func (request *RequestInformation) AddQueryParameters(source interface{}) error {
	return request.uri.AddQueryParameters(source)
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

func (request *RequestInformation) getContentReader() *bytes.Reader {
	return bytes.NewReader(request.Content)
}

func (rI *RequestInformation) SetUri(url *url.URL) {

	rI.uri.PathParameters = map[string]string{"request-raw-url": url.String()}

}

// ToRequest converts the RequestInformation object into an HTTP request.
func (request *RequestInformation) ToRequest() (*http.Request, error) {

	uri, err := request.uri.ToUrl()
	if err != nil {
		return nil, err
	}

	url := uri.String()
	contentReader := request.getContentReader()
	methodString := request.Method.String()

	req, err := http.NewRequest(methodString, url, contentReader)
	if err != nil {
		return nil, err
	}
	return req, nil
}
