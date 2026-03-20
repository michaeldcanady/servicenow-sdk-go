package model

import (
	"fmt"

	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	resultKey   = "result"
	nextKey     = "next"
	previousKey = "previous"
	lastKey     = "last"
	firstKey    = "first"
)

// ServiceNowCollectionResponse[T] Represents a Service-Now API collection response.
type ServiceNowCollectionResponse[T serialization.Parsable] interface {
	// GetResult Returns the result values of the response.
	GetResult() ([]T, error)
	GetNextLink() (*string, error)
	SetNextLink(val *string) error
	GetPreviousLink() (*string, error)
	SetPreviousLink(val *string) error
	GetFirstLink() (*string, error)
	SetFirstLink(val *string) error
	GetLastLink() (*string, error)
	SetLastLink(val *string) error
	serialization.Parsable
	BackedModel
}

// BaseServiceNowCollectionResponse[T] Basic implementation of the ServiceNowCollectionResponse[T]
type BaseServiceNowCollectionResponse[T serialization.Parsable] struct {
	// factory The ParsableFactory for serializing the result type.
	factory serialization.ParsableFactory
	// backingStoreFactory The BackingStoreFactory for creating the BackingStore to use.
	backingStoreFactory kiotaStore.BackingStoreFactory
	// backingStore The BackingStore
	backingStore kiotaStore.BackingStore
}

// NewBaseServiceNowCollectionResponse[T] Creates a new instance of the BaseServiceNowCollectionResponse[T].
func NewBaseServiceNowCollectionResponse[T serialization.Parsable](factory serialization.ParsableFactory) *BaseServiceNowCollectionResponse[T] {
	return &BaseServiceNowCollectionResponse[T]{
		factory:             factory,
		backingStoreFactory: kiotaStore.NewInMemoryBackingStore,
		backingStore:        kiotaStore.NewInMemoryBackingStore(),
	}
}

// Serialize writes the objects properties to the current writer
func (bR *BaseServiceNowCollectionResponse[T]) Serialize(writer serialization.SerializationWriter) error {
	if utils.IsNil(bR) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeCollectionOfObjectValuesFunc[T](resultKey)(bR.GetResult),
		internalSerialization.SerializeStringFunc(nextKey)(bR.GetNextLink),
		internalSerialization.SerializeStringFunc(previousKey)(bR.GetPreviousLink),
		internalSerialization.SerializeStringFunc(firstKey)(bR.GetFirstLink),
		internalSerialization.SerializeStringFunc(lastKey)(bR.GetLastLink),
	)
}

// GetFieldDeserializers returns the deserialization information for this object
func (bR *BaseServiceNowCollectionResponse[T]) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		resultKey: func(pn serialization.ParseNode) error {
			val, err := pn.GetCollectionOfObjectValues(bR.factory)
			if err != nil {
				return err
			}

			results := make([]any, len(val))
			for i, v := range val {
				results[i] = v
			}

			return bR.setResult(results)
		},
		nextKey:     internalSerialization.DeserializeStringFunc()(bR.SetNextLink),
		previousKey: internalSerialization.DeserializeStringFunc()(bR.SetPreviousLink),
		firstKey:    internalSerialization.DeserializeStringFunc()(bR.SetFirstLink),
		lastKey:     internalSerialization.DeserializeStringFunc()(bR.SetLastLink),
	}
}

// setResult Sets the result values of the response.
func (r *BaseServiceNowCollectionResponse[T]) setResult(val []any) error {
	if IsNil(r) {
		return nil
	}

	backingStore := r.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, resultKey, val)
}

// SetNextLink Sets the url to the next page of results.
func (r *BaseServiceNowCollectionResponse[T]) SetNextLink(val *string) error {
	if IsNil(r) {
		return nil
	}

	backingStore := r.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, nextKey, val)
}

// SetPreviousLink Sets the url to the previous page of results.
func (r *BaseServiceNowCollectionResponse[T]) SetPreviousLink(val *string) error {
	if IsNil(r) {
		return nil
	}

	backingStore := r.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, previousKey, val)
}

// SetFirstLink Sets the url to the first page of results.
func (r *BaseServiceNowCollectionResponse[T]) SetFirstLink(val *string) error {
	if IsNil(r) {
		return nil
	}

	backingStore := r.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, firstKey, val)
}

// SetLastLink Sets the url to the last page of results.
func (r *BaseServiceNowCollectionResponse[T]) SetLastLink(val *string) error {
	if IsNil(r) {
		return nil
	}

	backingStore := r.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, lastKey, val)
}

// GetBackingStore returns the backing store, if store is nil it instantiates a new store.
func (r *BaseServiceNowCollectionResponse[T]) GetBackingStore() kiotaStore.BackingStore {
	if IsNil(r) {
		return nil
	}

	if IsNil(r.backingStore) {
		if IsNil(r.backingStoreFactory) {
			return nil
		}
		r.backingStore = r.backingStoreFactory()
	}

	return r.backingStore
}

// GetResult Returns the result values of the response.
func (r *BaseServiceNowCollectionResponse[T]) GetResult() ([]T, error) {
	if utils.IsNil(r) {
		return nil, nil
	}

	backingStore := r.GetBackingStore()
	unknownSlice, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []any](backingStore, resultKey)
	if err != nil {
		return nil, err
	}

	results := make([]T, len(unknownSlice))

	for index, value := range unknownSlice {
		result, ok := value.(T)
		if !ok {
			return nil, fmt.Errorf("value is not %T", new(T))
		}

		results[index] = result
	}

	return results, nil
}

// GetNextLink Returns the url to the next page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetNextLink() (*string, error) {
	if utils.IsNil(r) {
		return nil, nil
	}

	backingStore := r.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, nextKey)
}

// GetPreviousLink Returns the url to the previous page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetPreviousLink() (*string, error) {
	if utils.IsNil(r) {
		return nil, nil
	}

	backingStore := r.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, previousKey)
}

// GetFirstLink Returns the url to the first page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetFirstLink() (*string, error) {
	if utils.IsNil(r) {
		return nil, nil
	}

	backingStore := r.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, firstKey)
}

// GetLastLink Returns the url to the last page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetLastLink() (*string, error) {
	if utils.IsNil(r) {
		return nil, nil
	}

	backingStore := r.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, lastKey)
}
