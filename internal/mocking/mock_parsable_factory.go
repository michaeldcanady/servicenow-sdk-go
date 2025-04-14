package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/mock"
)

type MockParsableFactory struct {
	mock.Mock
}

func NewMockParsableFactory() *MockParsableFactory {
	return &MockParsableFactory{
		Mock: mock.Mock{},
	}
}

func (mPF *MockParsableFactory) Factory(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	args := mPF.Called(parseNode)
	return args.Get(0).(serialization.Parsable), args.Error(1)
}
