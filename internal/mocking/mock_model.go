package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/mock"
)

type MockModel struct {
	mock.Mock
}

func NewMockModel() *MockModel {
	return &MockModel{
		Mock: mock.Mock{},
	}
}

func (mock *MockModel) GetBackingStore() store.BackingStore {
	args := mock.Called()
	return NilAllowed[store.BackingStore](args, 0)
}

func (mock *MockModel) SetBackingStoreFactory(factory store.BackingStoreFactory) error {
	args := mock.Called(factory)
	return args.Error(0)
}

func (mock *MockModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	args := mock.Called()
	return args.Get(0).(map[string]func(serialization.ParseNode) error)
}

func (mock *MockModel) Serialize(writer serialization.SerializationWriter) error {
	args := mock.Called(writer)
	return args.Error(0)
}
