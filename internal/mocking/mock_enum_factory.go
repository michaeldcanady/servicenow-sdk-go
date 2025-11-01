package mocking

import (
	"github.com/stretchr/testify/mock"
)

type MockEnumFactory struct {
	mock.Mock
}

func NewMockEnumFactory() *MockEnumFactory {
	return &MockEnumFactory{
		Mock: mock.Mock{},
	}
}

func (m *MockEnumFactory) Factory(input string) (interface{}, error) {
	args := m.Called(input)
	return args.Get(0), args.Error(1)
}
