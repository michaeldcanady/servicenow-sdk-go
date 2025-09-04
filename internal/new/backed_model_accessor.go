package internal

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/store"
)

// BackedModelAccessorFunc[S,T] defines a generic function signature for retrieving a value from a backed model
// using a specified key and converting it to a desired type.
type BackedModelAccessorFunc[M BackedModel, T any] func(M, string) (T, error)

// DefaultBackedModelAccessorFunc[S, T] is a generic implementation of BackedModelAccessorFunc that retrieves a value
// from a backed model and attempts to convert it to the specified type.
func DefaultBackedModelAccessorFunc[M store.BackedModel, T any](model M, key string) (T, error) {
	var result T

	if IsNil(model) {
		return result, errors.New("model is nil")
	}

	return DefaultStoreAccessorFunc[store.BackingStore, T](model.GetBackingStore(), key)
}

func DefaultBackedModelMutatedAccessorFunc[M store.BackedModel, T, S any](model M, key string, mutator func(S) (T, error)) (T, error) {
	var result T
	rawValue, err := DefaultBackedModelAccessorFunc[M, S](model, key)
	if err != nil {
		return result, err
	}
	return mutator(rawValue)
}
