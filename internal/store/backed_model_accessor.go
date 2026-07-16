package store

import (
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// BackedModelAccessorFunc[S,T] defines a generic function signature for retrieving a value from a backing store
// using a specified key and converting it to a desired type.
type BackedModelAccessorFunc[S kiotaStore.BackedModel, T any] func(S, string) (T, error)

// ModelAccessor represents a function for getting and typing a property for a kiotaStore.backed model.
type ModelAccessor[S kiotaStore.BackedModel, T any] BackedModelAccessorFunc[S, T]

// DefaultBackedModelAccessorFunc[S, T] is a generic implementation of BackedModelAccessorFunc that retrieves a value
// from a backing kiotaStore.and attempts to convert it to the specified type.
// DefaultBackedModelAccessorFunc[S, T] is a generic implementation of BackedModelAccessorFunc that retrieves a value
// from a backing kiotaStore.and attempts to convert it to the specified type.
func DefaultBackedModelAccessorFunc[S kiotaStore.BackedModel, T any](model S, key string) (T, error) {
	var result T

	if conversion.IsNil(model) {
		return result, snerrors.ErrNilModel
	}

	return DefaultStoreAccessorFunc[kiotaStore.BackingStore, T](model.GetBackingStore(), key)
}
