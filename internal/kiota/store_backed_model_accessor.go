package kiota

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// BackedModelAccessorFunc[S,T] defines a generic function signature for retrieving a value from a backing store
// using a specified key and converting it to a desired type.
type BackedModelAccessorFunc[S kiotaStore.BackingStore, T any] func(S, string) (T, error)

// ModelAccessor represents a function for getting and typing a property for a store backed model.
type ModelAccessor[S kiotaStore.BackingStore, T any] BackedModelAccessorFunc[S, T]

// DefaultBackedModelAccessorFunc[S, T] is a generic implementation of BackedModelAccessorFunc that retrieves a value
// from a backing store and attempts to convert it to the specified type.
func DefaultBackedModelAccessorFunc[S kiotaStore.BackingStore, T any](backingStore S, key string) (T, error) {
	return conversion.DefaultBackedModelAccessorFunc[S, T](backingStore, key)
}
