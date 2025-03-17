package internal

import (
	"fmt"
	"sync"
)

type ConcurrentDictionary[K comparable, V any] struct {
	mapping map[K]V
	lock    sync.RWMutex
}

func NewConcurrentDictionary[K comparable, V any]() *ConcurrentDictionary[K, V] {
	return &ConcurrentDictionary[K, V]{
		mapping: make(map[K]V),
		lock:    sync.RWMutex{},
	}
}

// Get retrieves the value associated with the given key
func (cd *ConcurrentDictionary[K, V]) Get(key K) (V, error) {
	if IsNil(cd) {
		var empty V
		return empty, nil
	}

	cd.lock.RLock()
	defer cd.lock.RUnlock()

	value, exists := cd.mapping[key]
	if !exists {
		var zero V
		return zero, fmt.Errorf("key '%v' does not exist", key)
	}
	return value, nil
}

// Add inserts a key-value pair into the dictionary. Returns an error if the key already exists.
func (cd *ConcurrentDictionary[K, V]) Add(key K, value V) error {
	if IsNil(cd) {
		return nil
	}

	cd.lock.Lock()
	defer cd.lock.Unlock()

	if _, exists := cd.mapping[key]; exists {
		return fmt.Errorf("key '%v' already exists", key)
	}
	cd.mapping[key] = value
	return nil
}

// Update modifies the value for an existing key in the dictionary.
func (cd *ConcurrentDictionary[K, V]) Update(key K, value V) error {
	if IsNil(cd) {
		return nil
	}

	cd.lock.Lock()
	defer cd.lock.Unlock()

	if _, exists := cd.mapping[key]; !exists {
		return fmt.Errorf("key '%v' does not exist", key)
	}
	cd.mapping[key] = value
	return nil
}

// Contains checks if the given key exists in the dictionary
func (cd *ConcurrentDictionary[K, V]) Contains(key K) bool {
	if IsNil(cd) {
		return false
	}

	cd.lock.RLock()
	defer cd.lock.RUnlock()

	_, exists := cd.mapping[key]
	return exists
}

// Remove deletes the key-value pair associated with the given key.
func (cd *ConcurrentDictionary[K, V]) Remove(key K) error {
	if IsNil(cd) {
		return nil
	}

	cd.lock.Lock()
	defer cd.lock.Unlock()

	if _, exists := cd.mapping[key]; !exists {
		return fmt.Errorf("key '%v' does not exist", key)
	}
	delete(cd.mapping, key)
	return nil
}

// Pop retrieves and removes the value associated with the given key.
func (cd *ConcurrentDictionary[K, V]) Pop(key K) (V, error) {
	if IsNil(cd) {
		var empty V
		return empty, nil
	}

	cd.lock.Lock()
	defer cd.lock.Unlock()

	value, exists := cd.mapping[key]
	if !exists {
		var zero V
		return zero, fmt.Errorf("key '%v' does not exist", key)
	}
	delete(cd.mapping, key)
	return value, nil
}
