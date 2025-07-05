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
	return NilAllowed[store.BackingStore](args, 0)
}

func (mock *MockModel) SetBackingStoreFactory(factory store.BackingStoreFactory) error {
	args := mock.Called(factory)
	return args.Error(0)
}
