package kiota

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// BackedModelAccessorFunc[S,T] defines a generic function signature for retrieving a value from a backing store
// using a specified key and converting it to a desired type.
type BackedModelAccessorFunc[S kiotaStore.BackingStore, T any] func(S, string) (T, error)

// DefaultBackedModelAccessorFunc[S, T] is a generic implementation of BackedModelAccessorFunc that retrieves a value
// from a backing store and attempts to convert it to the specified type.
func DefaultBackedModelAccessorFunc[S kiotaStore.BackingStore, T any](backingStore S, key string) (T, error) {
	return utils.DefaultBackedModelAccessorFunc[S, T](backingStore, key)
}
