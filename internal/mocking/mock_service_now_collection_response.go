package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/mock"
)

type MockServiceNowCollectionResponse[T serialization.Parsable] struct {
	mock.Mock
}

func (m *MockServiceNowCollectionResponse[T]) GetResult() ([]T, error) {
	args := m.Called()
	return args.Get(0).([]T), args.Error(1)
}

func (m *MockServiceNowCollectionResponse[T]) GetNextLink() (*string, error) {
	args := m.Called()
	return args.Get(0).(*string), args.Error(1)
}

func (m *MockServiceNowCollectionResponse[T]) GetPreviousLink() (*string, error) {
	args := m.Called()
	return args.Get(0).(*string), args.Error(1)
}

func (m *MockServiceNowCollectionResponse[T]) GetFirstLink() (*string, error) {
	args := m.Called()
	return args.Get(0).(*string), args.Error(1)
}

func (m *MockServiceNowCollectionResponse[T]) GetLastLink() (*string, error) {
	args := m.Called()
	return args.Get(0).(*string), args.Error(1)
}

// Serialize writes the objects properties to the current writer.
func (mP *MockServiceNowCollectionResponse[T]) Serialize(writer serialization.SerializationWriter) error {
	args := mP.Called(writer)
	return args.Error(0)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (mP *MockServiceNowCollectionResponse[T]) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	args := mP.Called()
	return args.Get(0).(map[string]func(serialization.ParseNode) error)
}

// GetBackingStore returns the BackingStore of the model.
func (mM *MockServiceNowCollectionResponse[T]) GetBackingStore() store.BackingStore {
	args := mM.Called()
	return args.Get(0).(store.BackingStore)
}
