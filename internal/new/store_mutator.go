package internal

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/store"
)

type StoreMutatorFunc[S store.BackingStore, T any] func(S, string, value T) error

func BaseStoreMutatorFunc[S store.BackingStore, T any](store S, key string, value T) error {
	if IsNil(store) {
		return errors.New("store is nil")
	}

	if key == "" {
		return errors.New("key is empty")
	}

	return store.Set(key, value)
}
