package store

import (
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type StoreMutatorFunc[S store.BackingStore, T any] func(S, string, value T) error

// DefaultStoreMutatorFunc[T] sets the store at the provided key to the provided value.
func DefaultStoreMutatorFunc[S store.BackingStore, T any](store store.BackingStore, key string, value T) error {
	if conversion.IsNil(store) {
		return snerrors.ErrNilStore
	}

	if strings.TrimSpace(key) == "" {
		return snerrors.ErrEmptyKey
	}

	return store.Set(key, value)
}
