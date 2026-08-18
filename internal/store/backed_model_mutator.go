package store

import (
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// BackedModelMutatorFunc[S, T] defines a generic function signature for setting the value of a [kiotaStore.BackedModel]
// using a specified key.
type BackedModelMutatorFunc[S kiotaStore.BackedModel, T any] func(S, string, T) error

// DefaultBackedModelMutatorFunc[S, T] is a generic implementation of [BackedModelMutatorFunc] that sets the value
// of a backed model.
func DefaultBackedModelMutatorFunc[S kiotaStore.BackedModel, T any](model S, key string, value T) error {
	if conversion.IsNil(model) {
		return snerrors.ErrNilModel
	}

	bs := model.GetBackingStore()
	if conversion.IsNil(bs) {
		return snerrors.ErrNilStore
	}

	return DefaultStoreMutatorFunc[kiotaStore.BackingStore](bs, key, value)
}
