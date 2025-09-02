package internal

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/store"
)

type BackedModelAccessorFunc[M store.BackedModel, T any] func(M, string) (T, error)

func BaseBackedModelAccessorFunc[M store.BackedModel, T any](model M, key string) (T, error) {
	var result T

	if IsNil(model) {
		return result, errors.New("model is nil")
	}

	return DefaultStoreAccessorFunc[store.BackingStore, T](model.GetBackingStore(), key)
}
