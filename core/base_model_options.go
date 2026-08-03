package core

import (
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/store"
)

// ModelOption configures a Model at construction time.
type ModelOption = internal.Option[Model]

// BackingStoreSettableOption[T] configures a T that supports setting its backing store factory.
type BackingStoreSettableOption[T BackingStoreFactorySetter] func(T) error

// WithBackingStoreFactory
func WithBackingStoreFactory[T BackingStoreFactorySetter](factory store.BackingStoreFactory) BackingStoreSettableOption[T] {
	return func(config T) error {
		if conversion.IsNil(config) {
			return snerrors.ErrNilConfig
		}

		return config.SetBackingStoreFactory(factory)
	}
}
