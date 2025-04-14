package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/mock"
)

type MockBackingStoreFactory struct {
	mock.Mock
}

func NewMockBackingStoreFactory() *MockBackingStoreFactory {
	return &MockBackingStoreFactory{
		Mock: mock.Mock{},
	}
}

func (bSF *MockBackingStoreFactory) MockBackingStoreFactory() store.BackingStore {
	args := bSF.Called()
	return args.Get(0).(store.BackingStore)
}
