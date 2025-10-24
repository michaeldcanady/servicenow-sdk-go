package mocking

import (
	"context"
	"net/http"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/mock"
)

type MockCoreClient2 struct {
	mock.Mock
}

func NewMockCoreClient2() *MockCoreClient2 {
	return &MockCoreClient2{mock.Mock{}}
}

func (mock *MockCoreClient2) Send(requestInfo core.IRequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {
	args := mock.Called(requestInfo, errorMapping)

	return args.Get(0).(*http.Response), args.Error(1)
}

func (mock *MockCoreClient2) SendWithContext(ctx context.Context, requestInfo core.IRequestInformation, mapping core.ErrorMapping) (*http.Response, error) {
	args := mock.Called(ctx, requestInfo, mapping)

	return args.Get(0).(*http.Response), args.Error(1)
}
