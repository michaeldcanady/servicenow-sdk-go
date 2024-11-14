package http

import (
	"errors"
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/internal/core"
	"github.com/RecoLabs/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthorizationProvider struct{}

func (m *MockAuthorizationProvider) AuthorizeRequest(request RequestInformation) error {
	// Mock implementation
	return nil
}

var (
	sharedAuthorizationHandler = &AuthorizationHandler{
		BaseHandler: &BaseHandler{},
		provider:    &MockAuthorizationProvider{},
	}
)

func TestNewAuthorizationHandler(t *testing.T) {
	tests := []mocking.Test[*AuthorizationHandler]{
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

type mockAuthorizationProvider struct {
	mock.Mock
}

func (m *mockAuthorizationProvider) AuthorizeRequest(request RequestInformation) error {
	args := m.Called(request)
	return args.Error(0)
}

func TestAuthorizationHandler_Handler(t *testing.T) {
	reqInfo := &mocking.MockRequestInformation{}

	authProvider := &mockAuthorizationProvider{}

	authHandler := &AuthorizationHandler{
		BaseHandler: &BaseHandler{},
		provider:    authProvider,
	}

	tests := []mocking.Test[*AuthorizationHandler]{
		{
			Title: "Successful",
			Setup: func() {
				authProvider.On("AuthorizeRequest", reqInfo).Return(nil)
			},
			Input:       reqInfo,
			ExpectedErr: nil,
			Cleanup: func() {
				authProvider.On("AuthorizeRequest", reqInfo).Unset()
			},
		},
		{
			Title: "Errored",
			Setup: func() {
				authProvider.On("AuthorizeRequest", reqInfo).Return(errors.New("Unable to authorize the request"))
			},
			Input:       reqInfo,
			ExpectedErr: errors.New("Unable to authorize the request"),
			Cleanup: func() {
				authProvider.On("AuthorizeRequest", reqInfo).Unset()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			err := authHandler.Handle(test.Input.(core.RequestInformation))

			assert.Equal(t, test.ExpectedErr, err)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}
