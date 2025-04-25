package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/mock"
)

type MockDeserializer struct {
	mock.Mock
}

func NewMockDeserializer() *MockDeserializer {
	return &MockDeserializer{
		mock.Mock{},
	}
}

func (m *MockDeserializer) Deserialize(contentType string, content []byte, parsableFactory serialization.ParsableFactory) (serialization.Parsable, error) {
	args := m.Called(contentType, content, parsableFactory)

	return args.Get(0).(serialization.Parsable), args.Error(1)
}
