package internal

import (
	"errors"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

// ServiceNowItemResponse[T] Represents a Service-Now API response containing a single result
type ServiceNowItemResponse[T serialization.Parsable] interface {
	// GetResult returns the result value from the response.
	GetResult() (T, error)
	serialization.Parsable
	BackedModel
}

// BaseServiceNowItemResponse[T] Basic implementation of the ServiceNowItemResponse[T].
type BaseServiceNowItemResponse[T serialization.Parsable] struct {
	// factory The ParsableFactory for serializing the result type.
	factory serialization.ParsableFactory
	// backingStoreFactory The BackingStoreFactory for creating the BackingStore to use.
	backingStoreFactory store.BackingStoreFactory
	// backingStore The BackingStore
	backingStore store.BackingStore
}

// NewBaseServiceNowItemResponse Creates a new instance of the BaseServiceNowItemResponse.
func NewBaseServiceNowItemResponse[T serialization.Parsable](factory serialization.ParsableFactory) *BaseServiceNowItemResponse[T] {
	return &BaseServiceNowItemResponse[T]{
		factory:             factory,
		backingStoreFactory: store.NewInMemoryBackingStore,
		backingStore:        store.NewInMemoryBackingStore(),
	}
}

// Serialize Writes the objects properties to the current writer.
func (bR *BaseServiceNowItemResponse[T]) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(bR) {
		return errors.New("serialization is not supported")
	}

	return errors.New("serialization is not supported")
}

// GetFieldDeserializers Returns the deserialization information for this object.
func (bR *BaseServiceNowItemResponse[T]) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		resultKey: func(pn serialization.ParseNode) error {
			var emptyVal T
			if IsNil(bR.factory) {
				return errors.New("factory is nil")
			}

			val, err := pn.GetObjectValue(bR.factory)
			if err != nil {
				return err
			}
			typedVal, ok := val.(T)
			if !ok {
				return fmt.Errorf("val is not %T", emptyVal)
			}
			return bR.setResult(typedVal)
		},
	}
}

// GetBackingStore Returns the backing store, if store is nil it instantiates a new store.
func (r *BaseServiceNowItemResponse[T]) GetBackingStore() (store.BackingStore, error) {
	if IsNil(r) {
		return nil, nil
	}

	if IsNil(r.backingStore) {
		if IsNil(r.backingStoreFactory) {
			return nil, errors.New("store is nil")
		}
		r.backingStore = r.backingStoreFactory()
	}

	return r.backingStore, nil
}

// GetResult Returns the result value of the response.
func (bR *BaseServiceNowItemResponse[T]) GetResult() (T, error) {
	var typedVal T

	if IsNil(bR) {
		return typedVal, nil
	}

	store, err := bR.GetBackingStore()
	if err != nil {
		return typedVal, err
	}

	val, err := store.Get(resultKey)
	if err != nil {
		return typedVal, err
	}

	typedVal, ok := val.(T)
	if !ok {
		return typedVal, fmt.Errorf("value is not %T", typedVal)
	}

	return typedVal, nil
}

// setResult Sets the result value of the response.
func (bR *BaseServiceNowItemResponse[T]) setResult(result T) error {
	if IsNil(bR) {
		return nil
	}

	store, err := bR.GetBackingStore()
	if err != nil {
		return err
	}

	return store.Set(resultKey, result)
}
