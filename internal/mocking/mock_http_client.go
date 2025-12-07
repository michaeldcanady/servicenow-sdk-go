package mocking

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

// MockHTTPClient implements oauth2.HTTPClient using testify/mock.
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	if resp, ok := args.Get(0).(*http.Response); ok {
		return resp, args.Error(1)
	}
	return nil, args.Error(1)
}
