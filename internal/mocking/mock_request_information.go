package mocking

import (
	"context"
	"net/http"
	"net/url"

	"github.com/RecoLabs/servicenow-sdk-go/internal/core"
	"github.com/stretchr/testify/mock"
)

type MockRequestInformation struct {
	mock.Mock
}

func (m *MockRequestInformation) Url() (string, error) { //nolint:stylecheck
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockRequestInformation) GetContent() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

func (m *MockRequestInformation) GetMethod() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockRequestInformation) GetHeaders() core.RequestHeader {
	args := m.Called()
	return args.Get(0).(core.RequestHeader)
}

func (m *MockRequestInformation) AddHeaders(headers interface{}) error {
	args := m.Called(headers)
	return args.Error(0)
}

func (m *MockRequestInformation) AddQueryParameters(params interface{}) error {
	args := m.Called(params)
	return args.Error(0)
}

func (m *MockRequestInformation) SetStreamContent(data []byte) {
	m.Called(data)
}

func (m *MockRequestInformation) SetUri(uri *url.URL) { //nolint:stylecheck
	m.Called(uri)
}

func (m *MockRequestInformation) ToRequest() (*http.Request, error) {
	args := m.Called()
	return args.Get(0).(*http.Request), args.Error(1)
}

func (m *MockRequestInformation) ToRequestWithContext(ctx context.Context) (*http.Request, error) {
	args := m.Called(ctx)
	return args.Get(0).(*http.Request), args.Error(1)
}

func (m *MockRequestInformation) SetContent(content []byte, contentType string) {
	_ = m.Called(content, contentType)
}
