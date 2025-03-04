package mocking

import "github.com/stretchr/testify/mock"

type MockOption[T any] struct {
	mock.Mock
}

func NewMockOption[T any]() *MockOption[T] {
	return &MockOption[T]{
		mock.Mock{},
	}
}

func (o *MockOption[T]) Option(config T) error {
	args := o.Called(config)
	return args.Error(0)
}
