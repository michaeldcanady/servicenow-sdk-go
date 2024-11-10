package core

import (
	"errors"
	"net/url"
	"strings"

	"maps"

	"github.com/yosida95/uritemplate/v3"
)

// UrlInformation represents an abstract Url.
type UrlInformation struct { //nolint:stylecheck
	// The Query Parameters of the request.
	QueryParameters map[string]string
	// The path parameters to use for the URL template when generating the URI.
	PathParameters map[string]string
	// The Url template for the current request.
	UrlTemplate string //nolint:stylecheck
}

// Deprecated: deprecated as of v1.4.0, use `NewURLInformation` instead.
//
// NewUrlInformation creates a new RequestUri object.
func NewUrlInformation() *UrlInformation { //nolint:stylecheck
	return NewURLInformation()
}

// NewURLInformation creates a new RequestUri object.
func NewURLInformation() *UrlInformation {
	return &UrlInformation{
		QueryParameters: make(map[string]string),
		PathParameters:  make(map[string]string),
	}
}

// validateURLTemplate checks if the URL template is empty.
func (uI *UrlInformation) validateURLTemplate() error {
	if uI.UrlTemplate == "" {
		return ErrEmptyURI
	}

	if !strings.Contains(strings.ToLower(uI.UrlTemplate), "{+baseurl}") {
		return ErrMissingBasePathTemplate
	}

	return nil
}

// validatePathParams checks if path parameters are nil.
func (uI *UrlInformation) validatePathParams() error {
	if uI.PathParameters == nil {
		return ErrNilPathParameters
	}

	_, baseurlExists := uI.PathParameters["baseurl"]
	if !baseurlExists {
		return ErrMissingBasePathParam
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
	err := uI.validateURLTemplate()
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

// getURIFromRaw retrieves the URI from the raw URL.
func (uI *UrlInformation) getURIFromRaw() (*url.URL, error) {
	if rawURL := uI.PathParameters[rawURLKey]; rawURL != "" {
		return uI.parseRawURL(rawURL)
	}
	return nil, nil
}

// buildValues builds the values for URI template expansion.
func (uI *UrlInformation) buildValues(normalizedNames map[string]string) uritemplate.Values {
	values := addParametersWithOriginalNames(uI.QueryParameters, normalizedNames, nil)
	values = addParametersWithOriginalNames(uI.PathParameters, normalizedNames, values)

	return values
}

// buildURIFromTemplate builds the URI from the template.
func (uI *UrlInformation) buildURIFromTemplate() (string, error) {
	uriTemplate, err := uritemplate.New(uI.UrlTemplate)
	if err != nil {
		return "", err
	}

	normalizedNames := normalizeVarNames(uriTemplate.Varnames())
	values := uI.buildValues(normalizedNames)

	url, err := uriTemplate.Expand(values)
	if err != nil {
		return "", err
	}

	return url, nil
}

// checkBaseURLRequirement checks if the "baseurl" parameter is required.
func (uI *UrlInformation) checkBaseURLRequirement() error {
	_, baseurlExists := uI.PathParameters["baseurl"]
	if !baseurlExists && strings.Contains(strings.ToLower(uI.UrlTemplate), "{+baseurl}") {
		return errors.New("pathParameters must contain a value for \"baseurl\" for the URL to be built")
	}
	return nil
}

// getUriFromTemplate retrieves the URI from the URL template.
func (uI *UrlInformation) getUriFromTemplate() (*url.URL, error) { //nolint:stylecheck
	if err := uI.checkBaseURLRequirement(); err != nil {
		return nil, err
	}

	uri, err := uI.buildURIFromTemplate()
	if err != nil {
		return nil, err
	}

	return uI.parseRawURL(uri)
}

// parseRawURL parses a raw URL.
func (uI *UrlInformation) parseRawURL(rawURL string) (*url.URL, error) {
	if rawURL == "" {
		return nil, ErrEmptyRawUrl
	}

	uri, err := url.Parse(rawURL)

	if uri.Scheme == "" {
		return nil, ErrMissingSchema
	}

	if err != nil {
		return nil, err
	}
	return uri, nil
}

// ToUrl retrieves the URI, either from the raw URL or the URL template.
func (uI *UrlInformation) ToUrl() (*url.URL, error) { //nolint:stylecheck
	uri, err := uI.getURIFromRaw()
	if uri != nil || err != nil {
		return uri, err
	}

	err = uI.validateParams()
	if err != nil {
		return nil, err
	}

	return uI.getUriFromTemplate()
}

// AddQueryParameters adds the query parameters to the request by reading the properties from the provided object.
func (uI *UrlInformation) AddQueryParameters(source interface{}) error {
	if source == nil || uI == nil {
		return errors.New("source or request is nil")
	}

	queryMap, err := ToQueryMap(source)
	if err != nil {
		return err
	}

	maps.Copy(uI.QueryParameters, queryMap)

	return nil
}
