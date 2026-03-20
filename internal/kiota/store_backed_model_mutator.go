package kiota

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// BackedModelMutatorFunc[S, T] defines a generic function signature for setting the value of a backing store
// using a specified key.
type BackedModelMutatorFunc[S kiotaStore.BackingStore, T any] func(S, string, T) error

// ModelMutator represents a function for mutating (setting) the property for a store backed model.
type ModelMutator[S kiotaStore.BackingStore, T any] BackedModelMutatorFunc[S, T]

// DefaultBackedModelMutatorFunc[S, T] is a generic implementation of BackedModelMutatorFunc that sets the value
// of a backed model.
func DefaultBackedModelMutatorFunc[S kiotaStore.BackingStore, T any](backingStore S, key string, value T) error {
	return utils.DefaultBackedModelMutatorFunc(backingStore, key, value)
}
