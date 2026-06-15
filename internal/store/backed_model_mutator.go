package store

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/store"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// BackedModelMutatorFunc[S, T] defines a generic function signature for setting the value of a backing store
// using a specified key.
type BackedModelMutatorFunc[S kiotaStore.BackingStore, T any] func(S, string, T) error

// ModelMutator represents a function for mutating (setting) the property for a store backed model.
type ModelMutator[S kiotaStore.BackingStore, T any] BackedModelMutatorFunc[S, T]

// DefaultBackedModelMutatorFunc[S, T] is a generic implementation of BackedModelMutatorFunc that sets the value
// of a backed model.
// DefaultBackedModelMutatorFunc[S, T] is a generic implementation of BackedModelMutatorFunc that sets the value
// of a backing store.
func DefaultBackedModelMutatorFunc[S store.BackingStore, T any](backingStore S, key string, value T) error {
	if conversion.IsNil(backingStore) {
		return errors.New("backingStore is nil")
	}

	return backingStore.Set(key, value)
}
