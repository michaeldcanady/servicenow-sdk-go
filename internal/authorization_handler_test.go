package internal

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockAuthorizationProvider struct{}

func (m *MockAuthorizationProvider) AuthorizeRequest(request RequestInformation) error {
	// Mock implementation
	return nil
}

type MockRequestInformation struct{}

func (m *MockRequestInformation) SetStreamContent(content []byte)             {}
func (m *MockRequestInformation) AddQueryParameters(source interface{}) error { return nil }
func (m *MockRequestInformation) SetUri(url *url.URL)                         {}                 //nolint:stylecheck
func (m *MockRequestInformation) Url() (string, error)                        { return "", nil } //nolint:stylecheck
func (m *MockRequestInformation) ToRequest() (*http.Request, error)           { return nil, nil }
func (m *MockRequestInformation) ToRequestWithContext(ctx context.Context) (*http.Request, error) {
	return nil, nil
}
func (m *MockRequestInformation) AddHeaders(rawHeaders interface{}) error { return nil }

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

			assert.Equal(t, test.expectedErr, err)
		})
	}
}
