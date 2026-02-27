package model

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type ModelOption = kiota.Option[Model]

type backingStoreSettableOption[T BackingStoreFactorySetter] func(T) error

// WithBackingStoreFactory
func WithBackingStoreFactory[T BackingStoreFactorySetter](factory store.BackingStoreFactory) backingStoreSettableOption[T] {
	return func(config T) error {
		if utils.IsNil(config) {
			return errors.New("config is nil")
		}

		return config.SetBackingStoreFactory(factory)
	}
}
