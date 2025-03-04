package mocking

import (
	nethttp "net/http"

	nethttplibrary "github.com/microsoft/kiota-http-go"
	"github.com/stretchr/testify/mock"
)

var _ nethttplibrary.Middleware = (*MockMiddleware)(nil)

type MockMiddleware struct {
	mock.Mock
}

func NewMockMiddleware() *MockMiddleware {
	return &MockMiddleware{
		mock.Mock{},
	}
}

func (m *MockMiddleware) Intercept(pipeline nethttplibrary.Pipeline, middlewareIndex int, request *nethttp.Request) (*nethttp.Response, error) {
	args := m.Mock.Called(pipeline, middlewareIndex, request)
	return args.Get(0).(*nethttp.Response), args.Error(1)
}
