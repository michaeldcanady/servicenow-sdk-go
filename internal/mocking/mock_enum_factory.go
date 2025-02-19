package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/mock"
)

var _ serialization.EnumFactory = ((*MockEnumFactory)(nil)).EnumFactory

type MockEnumFactory struct {
	mock.Mock
}

func NewMockEnumFactory() *MockEnumFactory {
	return &MockEnumFactory{Mock: mock.Mock{}}
}

func (eF *MockEnumFactory) EnumFactory(input string) (interface{}, error) {
	args := eF.Called(input)
	return args.Get(0), args.Error(1)
}
