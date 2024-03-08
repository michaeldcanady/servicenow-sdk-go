package internal

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/mozillazg/go-httpheader"
	"github.com/stretchr/testify/assert"
)

type MockAuthorizationProvider struct{}

func (m *MockAuthorizationProvider) AuthorizeRequest(request RequestInformation) error {
	// Mock implementation
	return nil
}

type MockRequestInformation struct {
	Headers http.Header
}

func (m *MockRequestInformation) SetStreamContent(content []byte)             {}
func (m *MockRequestInformation) AddQueryParameters(source interface{}) error { return nil }
func (m *MockRequestInformation) SetUri(url *url.URL)                         {}                 //nolint:stylecheck
func (m *MockRequestInformation) Url() (string, error)                        { return "", nil } //nolint:stylecheck
func (m *MockRequestInformation) ToRequest() (*http.Request, error)           { return nil, nil }
func (m *MockRequestInformation) ToRequestWithContext(ctx context.Context) (*http.Request, error) {
	return nil, nil
}
func (m *MockRequestInformation) AddHeaders(rawHeaders interface{}) error {
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
			m.Headers.Add(key, value)
		}
	}
	return nil
}

var (
	sharedAuthorizationHandler = &AuthorizationHandler{
		BaseHandler: &BaseHandler{},
		provider:    &MockAuthorizationProvider{},
	}
)

func TestNewAuthorizationHandler(t *testing.T) {
	tests := []Test[*AuthorizationHandler]{
		{
			Title:    "Valid",
			Input:    &MockAuthorizationProvider{},
			Expected: sharedAuthorizationHandler,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			handler := NewAuthorizationHandler(test.Input.(AuthorizationProvider))
			assert.Equal(t, test.Expected, handler)
		})
	}
}

func TestAuthorizationHandler_Handler(t *testing.T) {
	tests := []Test[*AuthorizationHandler]{
		{
			Title:    "",
			Input:    &MockRequestInformation{},
			Expected: sharedAuthorizationHandler,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			err := sharedAuthorizationHandler.Handle(test.Input.(RequestInformation))

			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}
