package core

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yosida95/uritemplate/v3"
)

func TestNewUrlInformation(t *testing.T) {
	ui := NewURLInformation()
	assert.NotNil(t, ui)
	assert.Empty(t, ui.QueryParameters)
	assert.Empty(t, ui.PathParameters)
}

func TestUrlInformationValidateUrlTemplate(t *testing.T) {
	ui := NewURLInformation()
	err := ui.validateURLTemplate()
	assert.Equal(t, ErrEmptyURI, err)

	ui.UrlTemplate = "https://"
	err = ui.validateURLTemplate()
	assert.Equal(t, ErrMissingBasePathTemplate, err)

	ui.UrlTemplate = "https://{+baseurl}/endpoint"
	err = ui.validateURLTemplate()
	assert.NoError(t, err)
}

func TestUrlInformationValidatePathParams(t *testing.T) {
	ui := NewURLInformation()
	err := ui.validatePathParams()
	assert.Equal(t, ErrMissingBasePathParam, err)

	ui.PathParameters = make(map[string]string)
	err = ui.validatePathParams()
	assert.Equal(t, ErrMissingBasePathParam, err)

	ui.PathParameters["baseurl"] = "example.com"
	err = ui.validatePathParams()
	assert.NoError(t, err)
}

func TestUrlInformationValidateQueryParams(t *testing.T) {
	ui := NewURLInformation()
	ui.QueryParameters = nil
	err := ui.validateQueryParams()
	assert.Equal(t, ErrNilQueryParamters, err)

	ui.QueryParameters = make(map[string]string)
	err = ui.validateQueryParams()
	assert.NoError(t, err)
}

func TestUrlInformationValidateParams(t *testing.T) {
	ui := NewURLInformation()
	err := ui.validateParams()
	assert.Equal(t, ErrEmptyURI, err)

	ui.UrlTemplate = "https://{+baseurl}/endpoint"
	err = ui.validateParams()
	assert.Equal(t, ErrMissingBasePathParam, err)

	ui.PathParameters = nil
	err = ui.validateParams()
	assert.Equal(t, ErrNilPathParameters, err)

	ui.QueryParameters = nil
	ui.PathParameters = map[string]string{"baseurl": "example.com"}
	err = ui.validateParams()
	assert.Equal(t, ErrNilQueryParamters, err)

	ui.QueryParameters = make(map[string]string)
	ui.PathParameters["baseurl"] = "example.com"
	err = ui.validateParams()
	assert.NoError(t, err)

	ui.QueryParameters = make(map[string]string)
	err = ui.validateParams()
	assert.NoError(t, err)
}

func TestUrlInformationBuildValues(t *testing.T) {
	ui := NewURLInformation()
	ui.QueryParameters = map[string]string{"param1": "value1", "param2": "value2"}
	ui.PathParameters = map[string]string{"baseurl": "example.com"}

	normalizedNames := make(map[string]string)
	values := ui.buildValues(normalizedNames)

	assert.NotNil(t, values)
	assert.Equal(t, uritemplate.Value{
		T: uritemplate.ValueTypeString,
		V: []string{"value1"},
	}, values["param1"])
	assert.Equal(t, uritemplate.Value{
		T: uritemplate.ValueTypeString,
		V: []string{"value2"},
	}, values["param2"])
	assert.Equal(t, uritemplate.Value{
		T: uritemplate.ValueTypeString,
		V: []string{"example.com"},
	}, values["baseurl"])
}

func TestUrlInformationBuildUriFromTemplate(t *testing.T) {
	ui := NewURLInformation()
	ui.UrlTemplate = "https://{+baseurl}/endpoint{?param1}"
	ui.PathParameters = map[string]string{"baseurl": "example.com"}
	ui.QueryParameters = map[string]string{"param1": "value1"}

	url, err := ui.buildURIFromTemplate()
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com/endpoint?param1=value1", url)

	ui.UrlTemplate = "{+baseurl}/endpoint"
	_, err = ui.buildURIFromTemplate()
	assert.Nil(t, err)
}

func TestUrlInformationCheckBaseUrlRequirement(t *testing.T) {
	ui := NewURLInformation()
	ui.UrlTemplate = "{+baseurl}/endpoint"
	ui.PathParameters = map[string]string{"baseurl": "example.com"}

	err := ui.checkBaseURLRequirement()
	assert.NoError(t, err)

	ui.PathParameters = map[string]string{}
	err = ui.checkBaseURLRequirement()
	assert.Error(t, err)
}

func TestUrlInformationGetUriFromTemplate(t *testing.T) {
	ui := NewURLInformation()
	ui.UrlTemplate = "https://{+baseurl}/endpoint"
	ui.PathParameters = map[string]string{"baseurl": "example.com"}

	uri, err := ui.getUriFromTemplate()
	assert.NoError(t, err)
	assert.Equal(t, &url.URL{Scheme: "https", Opaque: "", User: (*url.Userinfo)(nil), Host: "example.com", Path: "/endpoint", RawPath: "", OmitHost: false, ForceQuery: false, RawQuery: "", Fragment: "", RawFragment: ""}, uri)

	ui.UrlTemplate = "{+baseurl}/endpoint"
	_, err = ui.getUriFromTemplate()
	assert.Error(t, err)

	ui.PathParameters = map[string]string{}
	_, err = ui.getUriFromTemplate()
	assert.Error(t, err)
}

func TestUrlInformationParseRawURL(t *testing.T) {
	ui := NewURLInformation()
	rawURL := "https://example.com/endpoint"

	uri, err := ui.parseRawURL(rawURL)
	assert.NoError(t, err)
	assert.NotNil(t, uri)
	assert.Equal(t, "https", uri.Scheme)
	assert.Equal(t, "example.com", uri.Host)
	assert.Equal(t, "/endpoint", uri.Path)

	invalidRawURL := ""
	_, err = ui.parseRawURL(invalidRawURL)
	assert.Error(t, err)
}

func TestUrlInformationToUrl(t *testing.T) {
	ui := NewURLInformation()
	ui.UrlTemplate = "https://{+baseurl}/endpoint"
	ui.PathParameters = map[string]string{"baseurl": "example.com"}

	url, err := ui.ToUrl()
	assert.NoError(t, err)
	assert.NotNil(t, url)
	assert.Equal(t, "https://example.com/endpoint", url.String())

	ui.PathParameters = map[string]string{}
	_, err = ui.ToUrl()
	assert.Error(t, err)
}

func TestUrlInformationAddQueryParameters(t *testing.T) {
	ui := NewURLInformation()
	ui.QueryParameters = make(map[string]string)
	source := struct {
		Param1 string `url:"param1"`
		Param2 string `url:"param2"`
	}{Param1: "value1", Param2: "value2"}

	err := ui.AddQueryParameters(source)
	assert.NoError(t, err)
	assert.Equal(t, "value1", ui.QueryParameters["param1"])
	assert.Equal(t, "value2", ui.QueryParameters["param2"])

	err = ui.AddQueryParameters(nil)
	assert.Error(t, err)
}
