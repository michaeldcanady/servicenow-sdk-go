package model

import (
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
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
	backingStoreFactory kiotaStore.BackingStoreFactory
	// backingStore The BackingStore
	backingStore kiotaStore.BackingStore
}

// NewBaseServiceNowItemResponse Creates a new instance of the BaseServiceNowItemResponse.
func NewBaseServiceNowItemResponse[T serialization.Parsable](factory serialization.ParsableFactory) *BaseServiceNowItemResponse[T] {
	return &BaseServiceNowItemResponse[T]{
		factory:             factory,
		backingStoreFactory: kiotaStore.NewInMemoryBackingStore,
		backingStore:        kiotaStore.NewInMemoryBackingStore(),
	}
}

// Serialize Writes the objects properties to the current writer.
func (bR *BaseServiceNowItemResponse[T]) Serialize(writer serialization.SerializationWriter) error {
	if utils.IsNil(bR) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[T](resultKey)(bR.GetResult),
	)
}

// GetFieldDeserializers Returns the deserialization information for this object.
func (bR *BaseServiceNowItemResponse[T]) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		resultKey: internalSerialization.DeserializeObjectValueFunc[T](bR.factory)(bR.setResult),
	}
}

// GetBackingStore Returns the backing store, if store is nil it instantiates a new store.
func (r *BaseServiceNowItemResponse[T]) GetBackingStore() kiotaStore.BackingStore {
	if utils.IsNil(r) {
		return nil
	}

	if utils.IsNil(r.backingStore) {
		if utils.IsNil(r.backingStoreFactory) {
			return nil
		}
		r.backingStore = r.backingStoreFactory()
	}

	return r.backingStore
}

// GetResult Returns the result value of the response.
func (bR *BaseServiceNowItemResponse[T]) GetResult() (T, error) {
	if utils.IsNil(bR) {
		var typedVal T
		return typedVal, nil
	}

	backingStore := bR.GetBackingStore()
	return DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, T](backingStore, resultKey)
}

// setResult Sets the result value of the response.
func (bR *BaseServiceNowItemResponse[T]) setResult(result T) error {
	if utils.IsNil(bR) {
		return nil
	}

	backingStore := bR.GetBackingStore()
	return DefaultBackedModelMutatorFunc(backingStore, resultKey, result)
}
