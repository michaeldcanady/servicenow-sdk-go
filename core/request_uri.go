package core

import (
	"errors"
	"net/url"
	"strings"

	"maps"

	t "github.com/yosida95/uritemplate/v3"
)

var (
	ErrEmptyUri          = errors.New("uri cannot be empty")
	ErrNilPathParameters = errors.New("uri template parameters cannot be nil")
	ErrNilQueryParamters = errors.New("uri query parameters cannot be nil")
)

const (
	contentTypeHeader = "Content-Type"
	binaryContentType = "application/octet-steam"
	rawUrlKey         = "request-raw-url"
)

// UrlInformation represents an abstract Url.
type UrlInformation struct {
	// The Query Parameters of the request.
	QueryParameters map[string]string
	// The path parameters to use for the URL template when generating the URI.
	PathParameters map[string]string
	// The Url template for the current request.
	UrlTemplate string
}

// NewUrlInformation creates a new RequestUri object.
func NewUrlInformation() *UrlInformation {
	return &UrlInformation{
		QueryParameters: make(map[string]string),
		PathParameters:  make(map[string]string),
	}
}

// validateUrlTemplate checks if the URL template is empty.
func (uI *UrlInformation) validateUrlTemplate() error {
	if uI.UrlTemplate == "" {
		return ErrEmptyUri
	}
	return nil
}

// validatePathParams checks if path parameters are nil.
func (uI *UrlInformation) validatePathParams() error {
	if uI.PathParameters == nil {
		return ErrNilPathParameters
	}
	return nil
}

// validateQueryParams checks if query parameters are nil.
func (uI *UrlInformation) validateQueryParams() error {
	if uI.QueryParameters == nil {
		return ErrNilQueryParamters
	}
	return nil
}

// validateParams checks all the required parameters.
func (uI *UrlInformation) validateParams() error {
	err := uI.validateUrlTemplate()
	if err != nil {
		return err
	}
	err = uI.validatePathParams()
	if err != nil {
		return err
	}
	err = uI.validateQueryParams()
	if err != nil {
		return err
	}
	return nil
}

// getUriFromRaw retrieves the URI from the raw URL.
func (uI *UrlInformation) getUriFromRaw() (*url.URL, error) {
	if rawURL := uI.PathParameters[rawUrlKey]; rawURL != "" {
		return uI.parseRawURL(rawURL)
	}
	return nil, nil
}

// buildValues builds the values for URI template expansion.
func (uI *UrlInformation) buildValues(normalizedNames map[string]string) t.Values {
	values := t.Values{}

	values = uI.buildPathParams(normalizedNames, values)
	values = uI.buildQueryParams(normalizedNames, values)
	return values
}

// normalizeVarNames normalizes variable names for URI template expansion.
func (uI *UrlInformation) normalizeVarNames(varNames []string) map[string]string {
	normalizedNames := make(map[string]string)
	for _, varName := range varNames {
		normalizedNames[strings.ToLower(varName)] = varName
	}
	return normalizedNames
}

// buildPathParams builds path parameters for URI template expansion.
func (uI *UrlInformation) buildPathParams(normalizedNames map[string]string, values t.Values) t.Values {
	for key, value := range uI.PathParameters {
		addParameterWithOriginalName(key, value, normalizedNames, values)
	}
	return values
}

// buildQueryParams builds query parameters for URI template expansion.
func (uI *UrlInformation) buildQueryParams(normalizedNames map[string]string, values t.Values) t.Values {
	for key, value := range uI.QueryParameters {
		addParameterWithOriginalName(key, value, normalizedNames, values)
	}
	return values
}

// buildUriFromTemplate builds the URI from the template.
func (uI *UrlInformation) buildUriFromTemplate() (string, error) {
	uriTemplate, err := t.New(uI.UrlTemplate)
	if err != nil {
		return "", err
	}

	normalizedNames := uI.normalizeVarNames(uriTemplate.Varnames())
	values := uI.buildValues(normalizedNames)

	url, err := uriTemplate.Expand(values)
	if err != nil {
		return "", err
	}

	return url, nil
}

// checkBaseUrlRequirement checks if the "baseurl" parameter is required.
func (uI *UrlInformation) checkBaseUrlRequirement() error {
	_, baseurlExists := uI.PathParameters["baseurl"]
	if !baseurlExists && strings.Contains(strings.ToLower(uI.UrlTemplate), "{+baseurl}") {
		return errors.New("pathParameters must contain a value for \"baseurl\" for the URL to be built")
	}
	return nil
}

// getUriFromTemplate retrieves the URI from the URL template.
func (uI *UrlInformation) getUriFromTemplate() (*url.URL, error) {

	if err := uI.checkBaseUrlRequirement(); err != nil {
		return nil, err
	}

	uri, err := uI.buildUriFromTemplate()
	if err != nil {
		return nil, err
	}

	return uI.parseRawURL(uri)
}

// parseRawURL parses a raw URL.
func (uI *UrlInformation) parseRawURL(rawURL string) (*url.URL, error) {
	uri, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return uri, nil
}

// ToUrl retrieves the URI, either from the raw URL or the URL template.
func (uI *UrlInformation) ToUrl() (*url.URL, error) {

	err := uI.validateParams()
	if err != nil {
		return nil, err
	}

	uri, err := uI.getUriFromRaw()
	if uri != nil || err != nil {
		return uri, err
	}

	return uI.getUriFromTemplate()
}

// AddQueryParameters adds the query parameters to the request by reading the properties from the provided object.
func (uI *UrlInformation) AddQueryParameters(source interface{}) error {
	if source == nil || uI == nil {
		return errors.New("source or request is nil")
	}

	queryMap, err := toQueryMap(source)
	if err != nil {
		return err
	}

	maps.Copy(uI.QueryParameters, queryMap)

	return nil
}
