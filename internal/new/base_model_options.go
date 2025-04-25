package internal

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/store"
)

type ModelOption = Option[Model]

type backingStoreSettableOption[T BackingStoreFactorySetter] func(T) error

// WithBackingStoreFactory
func WithBackingStoreFactory[T BackingStoreFactorySetter](factory store.BackingStoreFactory) backingStoreSettableOption[T] {
	return func(config T) error {
		if IsNil(config) {
			return errors.New("config is nil")
		}

		return config.SetBackingStoreFactory(factory)
	}
}
