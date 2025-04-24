package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/mock"
)

type MockBackingStore struct {
	mock.Mock
}

func NewMockBackingStore() *MockBackingStore {
	return &MockBackingStore{
		Mock: mock.Mock{},
	}
}

// Get return a value from the backing store based on its key.
// Returns null if the value hasn't changed and "ReturnOnlyChangedValues" is true
func (bS *MockBackingStore) Get(key string) (interface{}, error) {
	args := bS.Called(key)
	return args.Get(0), args.Error(1)
}

// Set or updates the stored value for the given key.
// Will trigger subscriptions callbacks.
func (bS *MockBackingStore) Set(key string, value interface{}) error {
	args := bS.Called(key, value)
	return args.Error(0)
}

// Enumerate returns all the values stored in the backing store. Values will be filtered if "ReturnOnlyChangedValues" is true.
func (bS *MockBackingStore) Enumerate() map[string]interface{} {
	args := bS.Called()
	return args.Get(0).(map[string]interface{})
}

// EnumerateKeysForValuesChangedToNil returns the keys for all values that changed to null
func (bS *MockBackingStore) EnumerateKeysForValuesChangedToNil() []string {
	args := bS.Called()
	return args.Get(0).([]string)
}

// Subscribe registers a listener to any data change happening.
// returns a subscriptionId which can be used to reference the current subscription
func (bS *MockBackingStore) Subscribe(callback store.BackingStoreSubscriber) string {
	args := bS.Called(callback)
	return args.String(0)
}

// SubscribeWithId registers a listener to any data change happening and assigns the given id
func (bS *MockBackingStore) SubscribeWithId(callback store.BackingStoreSubscriber, subscriptionId string) error { // nolint: stylecheck // needs to match interface
	args := bS.Called(callback, subscriptionId)
	return args.Error(0)
}

// Unsubscribe Removes a subscription from the store based on its subscription id.
func (bS *MockBackingStore) Unsubscribe(subscriptionId string) error { // nolint: stylecheck // needs to match interface
	args := bS.Called(subscriptionId)
	return args.Error(0)
}

// Clear Removes the data stored in the backing store. Doesn't trigger any subscription.
func (bS *MockBackingStore) Clear() {
	_ = bS.Called()
}

// GetInitializationCompleted Track's status of object during serialization and deserialization.
// this property is used to initialize subscriber notifications
func (bS *MockBackingStore) GetInitializationCompleted() bool {
	args := bS.Called()
	return args.Bool(0)
}

// SetInitializationCompleted sets whether the initialization of the object and/or
// the initial deserialization has been completed to track whether objects have changed
func (bS *MockBackingStore) SetInitializationCompleted(val bool) {
	_ = bS.Called(val)
}

// GetReturnOnlyChangedValues is a flag that defines whether subscriber notifications should be sent only when
// data has been changed
func (bS *MockBackingStore) GetReturnOnlyChangedValues() bool {
	args := bS.Called()
	return args.Bool(0)
}

// SetReturnOnlyChangedValues Sets whether to return only values that have changed
// since the initialization of the object when calling the Get and Enumerate method
func (bS *MockBackingStore) SetReturnOnlyChangedValues(val bool) {
	_ = bS.Called(val)
}
