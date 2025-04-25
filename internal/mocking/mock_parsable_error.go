package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/mock"
)

type MockParsableError struct {
	mock.Mock
}

func NewMockParsableError() *MockParsableError {
	return &MockParsableError{
		Mock: mock.Mock{},
	}
}

// Serialize writes the objects properties to the current writer.
func (mP *MockParsableError) Serialize(writer serialization.SerializationWriter) error {
	args := mP.Called(writer)
	return args.Error(0)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (mP *MockParsableError) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	args := mP.Called()
	return args.Get(0).(map[string]func(serialization.ParseNode) error)
}

func (mP *MockParsableError) Error() string {
	args := mP.Called()
	return args.String(0)
}
