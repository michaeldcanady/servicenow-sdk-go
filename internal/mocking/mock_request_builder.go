package mocking

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/mock"
)

type MockRequestBuilder struct {
	mock.Mock
}

func NewMockRequestBuilder() *MockRequestBuilder {
	return &MockRequestBuilder{
		mock.Mock{},
	}
}

func (m *MockRequestBuilder) GetPathParameters() map[string]string {
	args := m.Called()
	return args.Get(0).(map[string]string)
}

func (m *MockRequestBuilder) SetPathParameters(pathParameters map[string]string) error {
	args := m.Called(pathParameters)
	return args.Error(0)
}

func (m *MockRequestBuilder) GetRequestAdapter() abstractions.RequestAdapter {
	args := m.Called()
	return args.Get(0).(abstractions.RequestAdapter)
}

func (m *MockRequestBuilder) SetRequestAdapter(requestAdapter abstractions.RequestAdapter) error {
	args := m.Called(requestAdapter)
	return args.Error(0)
}

func (m *MockRequestBuilder) GetURLTemplate() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockRequestBuilder) SetURLTemplate(urlTemplate string) error {
	args := m.Called(urlTemplate)
	return args.Error(0)
}
