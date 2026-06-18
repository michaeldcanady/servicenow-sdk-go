package core

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type ModelOption = internal.Option[Model]

type backingStoreSettableOption[T BackingStoreFactorySetter] func(T) error

// WithBackingStoreFactory
func WithBackingStoreFactory[T BackingStoreFactorySetter](factory store.BackingStoreFactory) backingStoreSettableOption[T] {
	return func(config T) error {
		if conversion.IsNil(config) {
			return errors.New("config is nil")
		}

		return config.SetBackingStoreFactory(factory)
	}
}
