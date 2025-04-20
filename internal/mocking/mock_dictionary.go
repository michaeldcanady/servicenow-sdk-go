package mocking

import "github.com/stretchr/testify/mock"

type MockDictionary[K comparable, V any] struct {
	mock.Mock
}

func NewMockDictionary[K comparable, V any]() *MockDictionary[K, V] {
	return &MockDictionary[K, V]{
		mock.Mock{},
	}
}

func (d *MockDictionary[K, V]) Get(key K) (V, error) {
	args := d.Called(key)

	return args.Get(0).(V), args.Error(1)
}

func (d *MockDictionary[K, V]) Add(key K, value V) error {
	args := d.Called(key, value)

	return args.Error(0)
}

func (d *MockDictionary[K, V]) Update(key K, value V) error {
	args := d.Called(key, value)

	return args.Error(0)
}

func (d *MockDictionary[K, V]) Contains(key K) bool {
	args := d.Called(key)

	return args.Bool(0)
}

func (d *MockDictionary[K, V]) Remove(key K) error {
	args := d.Called(key)

	return args.Error(0)
}

func (d *MockDictionary[K, V]) Pop(key K) (V, error) {
	args := d.Called(key)

	return args.Get(0).(V), args.Error(1)
}
