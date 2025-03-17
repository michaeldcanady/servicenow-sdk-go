package mocking

import "github.com/microsoft/kiota-abstractions-go/store"

type MockMainError struct {
	*MockParsable
}

func NewMockMainError() *MockMainError {
	return &MockMainError{
		NewMockParsable(),
	}
}

func (mE *MockMainError) GetBackingStore() store.BackingStore {
	args := mE.Called()
	return args.Get(0).(store.BackingStore)
}

func (mE *MockMainError) SetBackingStoreFactory(factory store.BackingStoreFactory) error {
	args := mE.Called(factory)
	return args.Error(0)
}

func (mE *MockMainError) GetDetail() (*string, error) {
	args := mE.Called()
	return args.Get(0).(*string), args.Error(1)
}

func (mE *MockMainError) GetMessage() (*string, error) {
	args := mE.Called()
	return args.Get(0).(*string), args.Error(1)
}

func (mE *MockMainError) GetStatus() (*string, error) {
	args := mE.Called()
	return args.Get(0).(*string), args.Error(1)
}
