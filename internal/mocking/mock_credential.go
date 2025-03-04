package mocking

import "github.com/stretchr/testify/mock"

type MockCredential struct {
	mock.Mock
}

func NewMockCredential() *MockCredential {
	return &MockCredential{
		mock.Mock{},
	}
}

func (c *MockCredential) GetAuthentication() (string, error) {
	args := c.Called()
	return args.String(0), args.Error(1)
}
