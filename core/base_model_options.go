package core

import (
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
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
			return snerrors.ErrNilConfig
		}

		return config.SetBackingStoreFactory(factory)
	}
}
