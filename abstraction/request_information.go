package abstraction

import (
	"bytes"
	"errors"

	"strings"

	"net/http"
	u "net/url"

	"maps"

	t "github.com/yosida95/uritemplate/v3"
)

// RequestInformation represents an abstract HTTP request.
type RequestInformation struct {
	// The HTTP method of the request.
	Method HttpMethod
	uri    *u.URL
	// The Request Headers.
	Headers *RequestHeaders
	// The Query Parameters of the request.
	QueryParameters map[string]string
	// The Request Body.
	Content []byte
	// The path parameters to use for the URL template when generating the URI.
	PathParameters map[string]string
	// The Url template for the current request.
	UrlTemplate string
	options     map[string]RequestOption
}

var (
	ErrEmptyUri          = errors.New("uri cannot be empty")
	ErrNilPathParameters = errors.New("uri template parameters cannot be nil")
	ErrNilQueryParamters = errors.New("uri query parameters cannot be nil")
)

const raw_url_key = "request-raw-url"

// NewRequestInformation creates a new RequestInformation object with default values.
func NewRequestInformation() *RequestInformation {
	return &RequestInformation{
		Headers:         NewRequestHeaders(),
		QueryParameters: make(map[string]string),
		options:         make(map[string]RequestOption),
		PathParameters:  make(map[string]string),
	}
}

// GetUri returns the URI of the request.
func (request *RequestInformation) GetUri() (*u.URL, error) {
	if request.uri != nil {
		return request.uri, nil
	}

	if err := request.validateParams(); err != nil {
		return nil, err
	}

	if rawURL := request.PathParameters[raw_url_key]; rawURL != "" {
		return request.parseRawURL(rawURL)
	}

	if err := request.checkBaseUrlRequirement(); err != nil {
		return nil, err
	}

	uri, err := request.buildUriFromTemplate()
	if err != nil {
		return nil, err
	}

	return u.Parse(uri)
}

func (request *RequestInformation) buildUriFromTemplate() (string, error) {
	uriTemplate, err := t.New(request.UrlTemplate)
	if err != nil {
		return "", err
	}

	normalizedNames := request.normalizeVarNames(uriTemplate.Varnames())
	values := request.buildValues(normalizedNames)

	url, err := uriTemplate.Expand(values)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (request *RequestInformation) validateUrlTemplate() error {
	if request.UrlTemplate == "" {
		return ErrEmptyUri
	}
	return nil
}

func (request *RequestInformation) validatePathParams() error {
	if request.PathParameters == nil {
		return ErrNilPathParameters
	}
	return nil
}

func (request *RequestInformation) validateQueryParams() error {
	if request.QueryParameters == nil {
		return ErrNilQueryParamters
	}
	return nil
}

func (request *RequestInformation) validateParams() error {
	err := request.validateUrlTemplate()
	if err != nil {
		return err
	}
	err = request.validatePathParams()
	if err != nil {
		return err
	}
	err = request.validateQueryParams()
	if err != nil {
		return err
	}
	return nil
}

func (request *RequestInformation) parseRawURL(rawURL string) (*u.URL, error) {
	uri, err := u.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	request.SetUri(*uri)
	return request.uri, nil
}

func (request *RequestInformation) checkBaseUrlRequirement() error {
	_, baseurlExists := request.PathParameters["baseurl"]
	if !baseurlExists && strings.Contains(strings.ToLower(request.UrlTemplate), "{+baseurl}") {
		return errors.New("pathParameters must contain a value for \"baseurl\" for the URL to be built")
	}
	return nil
}

func (request *RequestInformation) normalizeVarNames(varNames []string) map[string]string {
	normalizedNames := make(map[string]string)
	for _, varName := range varNames {
		normalizedNames[strings.ToLower(varName)] = varName
	}
	return normalizedNames
}

func (request *RequestInformation) buildPathParams(normalizedNames map[string]string, values t.Values) t.Values {
	for key, value := range request.PathParameters {
		addParameterWithOriginalName(key, value, normalizedNames, values)
	}
	return values
}

func (request *RequestInformation) buildQueryParams(normalizedNames map[string]string, values t.Values) t.Values {
	for key, value := range request.QueryParameters {
		addParameterWithOriginalName(key, value, normalizedNames, values)
	}
	return values
}

func (request *RequestInformation) buildValues(normalizedNames map[string]string) t.Values {
	values := t.Values{}

	values = request.buildPathParams(normalizedNames, values)
	values = request.buildQueryParams(normalizedNames, values)
	return values
}

// addParameterWithOriginalName adds the URI template parameter to the template using the right casing, because of go conventions, casing might have changed for the generated property
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

// SetUri updates the URI for the request from a raw URL.
func (request *RequestInformation) SetUri(url u.URL) {
	request.uri = &url
	for k := range request.PathParameters {
		delete(request.PathParameters, k)
	}
	for k := range request.QueryParameters {
		delete(request.QueryParameters, k)
	}
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

const contentTypeHeader = "Content-Type"
const binaryContentType = "application/octet-steam"

// SetStreamContent sets the request body to a binary stream.
func (request *RequestInformation) SetStreamContent(content []byte) {
	request.Content = content
	if request.Headers != nil {
		request.Headers.Add(contentTypeHeader, binaryContentType)
	}
}

// AddQueryParameters adds the query parameters to the request by reading the properties from the provided object.
func (request *RequestInformation) AddQueryParameters(source interface{}) {
	if source == nil || request == nil {
		return
	}

	params, err := uriParamValue(source)
	if err != nil {
		return
	}

	maps.Copy(request.QueryParameters, params)
}

func (request *RequestInformation) getContentReader() *bytes.Reader {
	return bytes.NewReader(request.Content)
}

func (request *RequestInformation) getUriString() (string, error) {
	uri, err := request.GetUri()
	if err != nil {
		return "", err
	}
	return uri.String(), err
}

func (request *RequestInformation) ToRequest() (*http.Request, error) {

	uri, err := request.getUriString()
	if err != nil {
		return nil, err
	}

	contentReader := request.getContentReader()
	methodString := request.Method.String()

	req, err := http.NewRequest(methodString, uri, contentReader)
	if err != nil {
		return nil, err
	}
	return req, nil
}
