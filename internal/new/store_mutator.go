package internal

import (
	"errors"
	"strings"

	"github.com/microsoft/kiota-abstractions-go/store"
)

// StoreMutatorFunc[S, T] defines a generic function signature for setting the value of a backing store
// using a specified key.
type StoreMutatorFunc[S store.BackingStore, T any] func(S, string, value T) error

// DefaultStoreMutatorFunc[S, T] is a generic implementation of StoreMutatorFunc[S, T] that sets the value
// of a backing store.
func DefaultStoreMutatorFunc[S store.BackingStore, T any](store S, key string, value T) error {
	if IsNil(store) {
		return errors.New("store is nil")
	}

	if strings.TrimSpace(key) == "" {
		return errors.New("key is empty")
	}

	return store.Set(key, value)
}
