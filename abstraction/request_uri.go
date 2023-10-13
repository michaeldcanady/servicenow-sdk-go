package abstraction

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
func (request *UrlInformation) validateUrlTemplate() error {
	if request.UrlTemplate == "" {
		return ErrEmptyUri
	}
	return nil
}

// validatePathParams checks if path parameters are nil.
func (request *UrlInformation) validatePathParams() error {
	if request.PathParameters == nil {
		return ErrNilPathParameters
	}
	return nil
}

// validateQueryParams checks if query parameters are nil.
func (request *UrlInformation) validateQueryParams() error {
	if request.QueryParameters == nil {
		return ErrNilQueryParamters
	}
	return nil
}

// validateParams checks all the required parameters.
func (request *UrlInformation) validateParams() error {
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

// getUriFromRaw retrieves the URI from the raw URL.
func (r *UrlInformation) getUriFromRaw() (*url.URL, error) {
	if rawURL := r.PathParameters[rawUrlKey]; rawURL != "" {
		return r.parseRawURL(rawURL)
	}
	return nil, nil
}

// buildValues builds the values for URI template expansion.
func (request *UrlInformation) buildValues(normalizedNames map[string]string) t.Values {
	values := t.Values{}

	values = request.buildPathParams(normalizedNames, values)
	values = request.buildQueryParams(normalizedNames, values)
	return values
}

// normalizeVarNames normalizes variable names for URI template expansion.
func (request *UrlInformation) normalizeVarNames(varNames []string) map[string]string {
	normalizedNames := make(map[string]string)
	for _, varName := range varNames {
		normalizedNames[strings.ToLower(varName)] = varName
	}
	return normalizedNames
}

// buildPathParams builds path parameters for URI template expansion.
func (request *UrlInformation) buildPathParams(normalizedNames map[string]string, values t.Values) t.Values {
	for key, value := range request.PathParameters {
		addParameterWithOriginalName(key, value, normalizedNames, values)
	}
	return values
}

// buildQueryParams builds query parameters for URI template expansion.
func (request *UrlInformation) buildQueryParams(normalizedNames map[string]string, values t.Values) t.Values {
	for key, value := range request.QueryParameters {
		addParameterWithOriginalName(key, value, normalizedNames, values)
	}
	return values
}

// buildUriFromTemplate builds the URI from the template.
func (request *UrlInformation) buildUriFromTemplate() (string, error) {
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

// checkBaseUrlRequirement checks if the "baseurl" parameter is required.
func (request *UrlInformation) checkBaseUrlRequirement() error {
	_, baseurlExists := request.PathParameters["baseurl"]
	if !baseurlExists && strings.Contains(strings.ToLower(request.UrlTemplate), "{+baseurl}") {
		return errors.New("pathParameters must contain a value for \"baseurl\" for the URL to be built")
	}
	return nil
}

// getUriFromTemplate retrieves the URI from the URL template.
func (r *UrlInformation) getUriFromTemplate() (*url.URL, error) {

	if err := r.checkBaseUrlRequirement(); err != nil {
		return nil, err
	}

	uri, err := r.buildUriFromTemplate()
	if err != nil {
		return nil, err
	}

	return r.parseRawURL(uri)
}

// parseRawURL parses a raw URL.
func (r *UrlInformation) parseRawURL(rawURL string) (*url.URL, error) {
	uri, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return uri, nil
}

// ToUrl retrieves the URI, either from the raw URL or the URL template.
func (r *UrlInformation) ToUrl() (*url.URL, error) {

	err := r.validateParams()
	if err != nil {
		return nil, err
	}

	uri, err := r.getUriFromRaw()
	if uri != nil || err != nil {
		return uri, err
	}

	return r.getUriFromTemplate()
}

// AddQueryParameters adds the query parameters to the request by reading the properties from the provided object.
func (request *UrlInformation) AddQueryParameters(source interface{}) error {
	if source == nil || request == nil {
		return errors.New("source or request is nil")
	}

	queryMap, err := toQueryMap(source)
	if err != nil {
		return err
	}

	maps.Copy(request.QueryParameters, queryMap)

	return nil
}
