package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/mock"
)

type MockBatchHeader struct {
	mock.Mock
}

func NewMockBatchHeader() *MockBatchHeader {
	return &MockBatchHeader{
		Mock: mock.Mock{},
	}
}

func (mBH *MockBatchHeader) GetName() (*string, error) {
	args := mBH.Called()
	return args.Get(0).(*string), args.Error(1)
}
func (mBH *MockBatchHeader) SetName(name *string) error {
	args := mBH.Called(name)
	return args.Error(0)
}
func (mBH *MockBatchHeader) GetValue() (*string, error) {
	args := mBH.Called()
	return args.Get(0).(*string), args.Error(1)
}
func (mBH *MockBatchHeader) SetValue(value *string) error {
	args := mBH.Called(value)
	return args.Error(0)
}

// Serialize writes the objects properties to the current writer.
func (mBH *MockBatchHeader) Serialize(writer serialization.SerializationWriter) error {
	args := mBH.Called(writer)
	return args.Error(0)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (mBH *MockBatchHeader) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	args := mBH.Called()
	return args.Get(0).(map[string]func(serialization.ParseNode) error)
}

func (mBH *MockBatchHeader) GetBackingStore() store.BackingStore {
	args := mBH.Called()
	return args.Get(0).(store.BackingStore)
}

func (mBH *MockBatchHeader) SetBackingStoreFactory(backingStoreFactory store.BackingStoreFactory) error {
	args := mBH.Called(backingStoreFactory)
	return args.Error(0)
}
