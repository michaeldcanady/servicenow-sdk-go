package internal

import (
	"errors"
	"fmt"

	"github.com/microsoft/kiota-abstractions-go/store"
)

type StoreAccessorFunc[S store.BackingStore, T any] func(S, string) (T, error)

type BackedModelAccessorFunc[M store.BackedModel, T any] func(M, string) (T, error)

type BackedModelMutatorFunc[M store.BackedModel, T any] func(M, string, T) error

func BaseStoreAccessorFunc[S store.BackingStore, T any](store S, key string) (T, error) {
	var result T

	if IsNil(store) {
		return result, errors.New("store is nil")
	}
	if key == "" {
		return result, errors.New("key is empty")
	}

	val, err := store.Get(key)
	if err != nil {
		return result, err
	}

	var ok bool

	// TODO: use convert.As
	if result, ok = val.(T); !ok {
		return result, fmt.Errorf("val is not %T", val)
	}

	return result, nil
}

func BaseBackedModelAccessorFunc[M store.BackedModel, T any](model M, key string) (T, error) {
	var result T

	if IsNil(model) {
		return result, errors.New("model is nil")
	}

	return BaseStoreAccessorFunc[store.BackingStore, T](model.GetBackingStore(), key)
}

func BaseBackedModelMutatorFunc[M store.BackedModel, T any](model M, key string, value T) error {
	if IsNil(model) {
		return errors.New("model is nil")
	}

	return BaseStoreMutatorFunc[store.BackingStore, T](model.GetBackingStore(), key, value)
}
