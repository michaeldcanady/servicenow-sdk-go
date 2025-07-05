package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/mock"
)

type MockSerializationWriterFactory struct {
	mock.Mock
}

func NewMockSerializationWriterFactory() *MockSerializationWriterFactory {
	return &MockSerializationWriterFactory{
		Mock: mock.Mock{},
	}
}

// GetValidContentType returns the valid content type for the SerializationWriterFactoryRegistry
func (m *MockSerializationWriterFactory) GetValidContentType() (string, error) {
	args := m.Called()

	return args.String(0), args.Error(1)
}

// GetSerializationWriter returns the relevant SerializationWriter instance for the given content type
func (m *MockSerializationWriterFactory) GetSerializationWriter(contentType string) (serialization.SerializationWriter, error) {
	args := m.Called(contentType)

	return args.Get(0).(serialization.SerializationWriter), args.Error(1)
}
