package internal

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/store"
)

type BackedModelMutatorFunc[M store.BackedModel, T any] func(M, string, T) error

func DefaultBackedModelMutatorFunc[M store.BackedModel, T any](model M, key string, value T) error {
	if IsNil(model) {
		return errors.New("model is nil")
	}

	return DefaultStoreMutatorFunc(model.GetBackingStore(), key, value)
}
