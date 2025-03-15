package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/mock"
)

type MockParsable struct {
	mock.Mock
}

func NewMockParsable() *MockParsable {
	return &MockParsable{
		Mock: mock.Mock{},
	}
}

// Serialize writes the objects properties to the current writer.
func (mP *MockParsable) Serialize(writer serialization.SerializationWriter) error {
	args := mP.Called(writer)
	return args.Error(0)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (mP *MockParsable) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	args := mP.Called()
	return args.Get(0).(map[string]func(serialization.ParseNode) error)
}
