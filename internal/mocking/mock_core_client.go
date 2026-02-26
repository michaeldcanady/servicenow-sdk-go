package mocking

import (
	"context"
	"net/http"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/mock"
)

type MockCoreClient struct {
	mock.Mock
}

func NewMockCoreClient() *MockCoreClient {
	return &MockCoreClient{mock.Mock{}}
}

func (mock *MockCoreClient) Send(requestInfo core.IRequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {
	args := mock.Called(requestInfo, errorMapping)

	return args.Get(0).(*http.Response), args.Error(1)
}

func (mock *MockCoreClient) SendWithContext(ctx context.Context, requestInfo core.IRequestInformation, mapping core.ErrorMapping) (*http.Response, error) {
	args := mock.Called(ctx, requestInfo, mapping)

	return args.Get(0).(*http.Response), args.Error(1)
}
