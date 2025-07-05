package credentials

import "github.com/stretchr/testify/mock"

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
