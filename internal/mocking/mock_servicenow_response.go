package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/mock"
)

type MockServiceNowResponse[T any] struct {
	mock.Mock
}

func NewMockServiceNowResponse[T any]() *MockServiceNowResponse[T] {
	return &MockServiceNowResponse[T]{
		mock.Mock{},
	}
}

func (sR *MockServiceNowResponse[T]) GetResult() (T, error) {
	args := sR.Called()
	return args.Get(0).(T), args.Error(1)
}

func (sR *MockServiceNowResponse[T]) SetResult(result T) error {
	args := sR.Called(result)
	return args.Error(0)
}

func (sR *MockServiceNowResponse[T]) Serialize(writer serialization.SerializationWriter) error {
	args := sR.Called(writer)
	return args.Error(0)
}

func (sR *MockServiceNowResponse[T]) GetBackingStore() store.BackingStore {
	args := sR.Called()
	return args.Get(0).(store.BackingStore)
}

func (sR *MockServiceNowResponse[T]) GetFactory() (serialization.ParsableFactory, error) {
	args := sR.Called()
	return args.Get(0).(serialization.ParsableFactory), args.Error(1)
}
