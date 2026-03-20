package credentials

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type mockCredential struct {
	mock.Mock
}

func newMockCredential() *mockCredential {
	return &mockCredential{
		Mock: mock.Mock{},
	}
}

func (mock *mockCredential) GetAuthentication() (string, error) {
	args := mock.Called()

	return args.String(0), args.Error(1)
}

type mockAuthorizationCodeServer struct {
	mock.Mock
}

func (m *mockAuthorizationCodeServer) GetAddr() string {
	args := m.Called()
	return args.String(0)
}

func (m *mockAuthorizationCodeServer) Result(ctx context.Context) (string, string, error) {
	args := m.Called(ctx)
	return args.String(0), args.String(1), args.Error(2)
}

func (m *mockAuthorizationCodeServer) Shutdown(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
