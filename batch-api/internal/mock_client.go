package internal

import (
	"context"
	"net/http"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (c *MockClient) Send(ctx context.Context, requestInfo core.IRequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {
	args := c.Called(ctx, requestInfo, errorMapping)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (c *MockClient) GetBaseURL() string {
	args := c.Called()
	return args.String(0)
}

func (c *MockClient) SendWithContext(ctx context.Context, requestInfo core.IRequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {
	args := c.Called(ctx, requestInfo, errorMapping)
	return args.Get(0).(*http.Response), args.Error(1)
}
