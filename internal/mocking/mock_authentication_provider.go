package mocking

import (
	"context"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/mock"
)

type MockAuthenticationProvider struct {
	mock.Mock
}

func NewMockAuthenticationProvider() *MockAuthenticationProvider {
	return &MockAuthenticationProvider{
		mock.Mock{},
	}
}

func (aP *MockAuthenticationProvider) AuthenticateRequest(context context.Context, request *abstractions.RequestInformation, additionalAuthenticationContext map[string]interface{}) error {
	args := aP.Called(context, request, additionalAuthenticationContext)

	return args.Error(0)
}
