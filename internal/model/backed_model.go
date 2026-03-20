package model

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/store"
)

// BackedModel Represents a model backed by a BackingStore
type BackedModel interface {
	store.BackedModel
}

// BackedModelAccessorFunc[S,T] defines a generic function signature for retrieving a value from a backing store
// using a specified key and converting it to a desired type.
type BackedModelAccessorFunc[S store.BackingStore, T any] conversion.BackedModelAccessorFunc[S, T]

// DefaultBackedModelAccessorFunc[S, T] is a generic implementation of BackedModelAccessorFunc that retrieves a value
// from a backing store and attempts to convert it to the specified type.
func DefaultBackedModelAccessorFunc[S store.BackingStore, T any](backingStore S, key string) (T, error) {
	return conversion.DefaultBackedModelAccessorFunc[S, T](backingStore, key)
}

// BackedModelMutatorFunc[S, T] defines a generic function signature for setting the value of a backing store
// using a specified key.
type BackedModelMutatorFunc[S store.BackingStore, T any] conversion.BackedModelMutatorFunc[S, T]

// DefaultBackedModelMutatorFunc[S, T] is a generic implementation of BackedModelMutatorFunc that sets the value
// of a backing store.
func DefaultBackedModelMutatorFunc[S store.BackingStore, T any](backingStore S, key string, value T) error {
	return conversion.DefaultBackedModelMutatorFunc[S, T](backingStore, key, value)
}
