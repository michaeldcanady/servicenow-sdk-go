package mocking

import (
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

func (mM *MockModel) GetBackingStore() store.BackingStore {
	args := mM.Called()
	return args.Get(0).(store.BackingStore)
}

func (mM *MockModel) SetBackingStoreFactory(factory store.BackingStoreFactory) error {
	args := mM.Called(factory)
	return args.Error(0)
}
