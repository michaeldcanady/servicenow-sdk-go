package internal

import (
	"errors"
	"strings"

	"github.com/microsoft/kiota-abstractions-go/store"
)

// StoreMutatorFunc[S, T]
type StoreMutatorFunc[S store.BackingStore, T any] func(S, string, value T) error

// DefaultStoreMutatorFunc[S, T] Inserts value at key in store.
func DefaultStoreMutatorFunc[S store.BackingStore, T any](store S, key string, value T) error {
	if IsNil(store) {
		return errors.New("store is nil")
	}

	if strings.TrimSpace(key) == "" {
		return errors.New("key is empty")
	}

	return store.Set(key, value)
}
