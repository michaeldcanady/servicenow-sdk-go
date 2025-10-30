package store

import (
	"errors"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type StoreMutatorFunc[S store.BackingStore, T any] func(S, string, value T) error

// DefaultStoreMutatorFunc[T] sets the store at the provided key to the provided value.
func DefaultStoreMutatorFunc[T any](store store.BackingStore, key string, value T) error {
	if internal.IsNil(store) {
		return errors.New("store is nil")
	}

	if strings.TrimSpace(key) == "" {
		return errors.New("key is empty")
	}

	return store.Set(key, value)
}
