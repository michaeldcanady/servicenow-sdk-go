package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/mock"
)

type MockParsable struct {
	mock.Mock
	serialization.Parsable
}

func NewMockParsable() *MockParsable {
	return &MockParsable{
		Mock: mock.Mock{},
	}
}

func (m *MockParsable) Serialize(writer serialization.SerializationWriter) error {
	args := m.Called(writer)
	return args.Error(0)
}

func (m *MockParsable) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	args := m.Called()
	return args.Get(0).(map[string]func(serialization.ParseNode) error)
}
