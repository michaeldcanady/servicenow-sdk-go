package core

import (
	"context"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type Client2 interface {
	Send(RequestInformation, ErrorMapping) (*http.Response, error)
	SendWithContext(context.Context, RequestInformation, ErrorMapping) (*http.Response, error)
	GetBaseURL() string
}

type MockClient2 struct {
	mock.Mock
}

func (c *MockClient2) Send(info RequestInformation, mapping ErrorMapping) (*http.Response, error) {
	args := c.Called(info, mapping)
	return args.Get(0).(*http.Response), args.Error(1)
}
func (c *MockClient2) SendWithContext(ctx context.Context, info RequestInformation, mapping ErrorMapping) (*http.Response, error) {
	args := c.Called(ctx, info, mapping)
	return args.Get(0).(*http.Response), args.Error(1)
}
func (c *MockClient2) GetBaseURL() string {
	args := c.Called()
	return args.String(0)
}

var _ Client2 = (*MockClient2)(nil)
