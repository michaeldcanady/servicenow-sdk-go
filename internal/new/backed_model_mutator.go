package internal

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/store"
)

// BackedModelMutatorFunc[S, T] defines a generic function signature for setting the value of a backed model
// using a specified key.
type BackedModelMutatorFunc[M BackedModel, T any] func(M, string, T) error

// DefaultStoreMutatorFunc[S, T] is a generic implementation of StoreMutatorFunc[S, T] that sets the value
// of a backed model.
func DefaultBackedModelMutatorFunc[M store.BackedModel, T any](model M, key string, value T) error {
	if IsNil(model) {
		return errors.New("model is nil")
	}

	return DefaultStoreMutatorFunc(model.GetBackingStore(), key, value)
}
