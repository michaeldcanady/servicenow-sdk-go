package internal

import (
	"errors"
	"strings"

	"github.com/microsoft/kiota-abstractions-go/store"
)

// StoreAccessorFunc Function for accessing specific key, and converting it to the desired type, from backing store.
type StoreAccessorFunc[S store.BackingStore, T any] func(S, string) (T, error)

// DefaultStoreAccessorFunc Gets key from store, converting it to the desired type.
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
