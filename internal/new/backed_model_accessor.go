package internal

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/store"
)

// BackedModelAccessorFunc[S,T] defines a generic function signature for retrieving a value from a backed model
// using a specified key and converting it to a desired type.
type BackedModelAccessorFunc[M store.BackedModel, T any] func(M, string) (T, error)

// DefaultBackedModelAccessorFunc[S, T] is a generic implementation of BackedModelAccessorFunc that retrieves a value
// from a backed model and attempts to convert it to the specified type.
func DefaultBackedModelAccessorFunc[M store.BackedModel, T any](model M, key string) (T, error) {
	var result T

	if IsNil(model) {
		return result, errors.New("model is nil")
	}

	return DefaultStoreAccessorFunc[store.BackingStore, T](model.GetBackingStore(), key)
}
