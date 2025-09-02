package internal

import (
	"errors"
	"strings"

	"github.com/microsoft/kiota-abstractions-go/store"
)

// StoreAccessorFunc[S,T] defines a generic function signature for retrieving a value from a backing store
// using a specified key and converting it to a desired type.
type StoreAccessorFunc[S store.BackingStore, T any] func(S, string) (T, error)

// DefaultStoreAccessorFunc[S, T] is a generic implementation of StoreAccessorFunc that retrieves a value
// from a backing store and attempts to convert it to the specified type.
func DefaultStoreAccessorFunc[S store.BackingStore, T any](store S, key string) (T, error) {
	var result T

	if IsNil(store) {
		return result, errors.New("store is nil")
	}
	if strings.TrimSpace(key) == "" {
		return result, errors.New("key is empty")
	}

	val, err := store.Get(key)
	if err != nil {
		return result, err
	}

	if err := As2(val, &result, true); err != nil {
		return result, err
	}

	return result, nil
}
