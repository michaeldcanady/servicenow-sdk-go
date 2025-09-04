package internal

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
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
	GetPreviousLink() (*string, error)
	GetFirstLink() (*string, error)
	GetLastLink() (*string, error)
	serialization.Parsable
	BackedModel
}

// BaseServiceNowCollectionResponse[T] Basic implementation of the ServiceNowCollectionResponse[T]
type BaseServiceNowCollectionResponse[T serialization.Parsable] struct {
	// factory The ParsableFactory for serializing the result type.
	factory serialization.ParsableFactory
	// backingStoreFactory The BackingStoreFactory for creating the BackingStore to use.
	backingStoreFactory store.BackingStoreFactory
	// backingStore The BackingStore
	backingStore store.BackingStore
}

// NewBaseServiceNowCollectionResponse[T] Creates a new instance of the BaseServiceNowCollectionResponse[T].
func NewBaseServiceNowCollectionResponse[T serialization.Parsable](factory serialization.ParsableFactory) *BaseServiceNowCollectionResponse[T] {
	return &BaseServiceNowCollectionResponse[T]{
		factory:             factory,
		backingStoreFactory: store.NewInMemoryBackingStore,
		backingStore:        store.NewInMemoryBackingStore(),
	}
}

// Serialize writes the objects properties to the current writer
func (bR *BaseServiceNowCollectionResponse[T]) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(bR) {
		return nil
	}

	return errors.New("Serialize not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object
func (bR *BaseServiceNowCollectionResponse[T]) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		resultKey: DeserializeObjectArrayFunc(bR.SetResult, bR.factory),
	}
}

// GetBackingStore returns the backing store, if store is nil it instantiates a new store.
func (r *BaseServiceNowCollectionResponse[T]) GetBackingStore() store.BackingStore {
	return r.backingStore
}

func (r *BaseServiceNowCollectionResponse[T]) SetResult(result []T) error {
	return DefaultBackedModelMutatorFunc(r, resultKey, result)
}

// GetResult Returns the result values of the response.
func (r *BaseServiceNowCollectionResponse[T]) GetResult() ([]T, error) {
	return DefaultBackedModelAccessorFunc[*BaseServiceNowCollectionResponse[T], []T](r, resultKey)
}

// GetNextLink Returns the url to the next page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetNextLink() (*string, error) {
	return DefaultBackedModelAccessorFunc[*BaseServiceNowCollectionResponse[T], *string](r, nextKey)
}

// SetNextLink Sets the url to the next page of results.
func (r *BaseServiceNowCollectionResponse[T]) SetNextLink(link *string) error {
	return DefaultBackedModelMutatorFunc(r, nextKey, link)
}

// GetPreviousLink Returns the url to the previous page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetPreviousLink() (*string, error) {
	return DefaultBackedModelAccessorFunc[*BaseServiceNowCollectionResponse[T], *string](r, previousKey)
}

// SetPreviousLink Sets the url to the previous page of results.
func (r *BaseServiceNowCollectionResponse[T]) SetPreviousLink(link *string) error {
	return DefaultBackedModelMutatorFunc(r, previousKey, link)
}

// GetFirstLink Returns the url to the first page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetFirstLink() (*string, error) {
	return DefaultBackedModelAccessorFunc[*BaseServiceNowCollectionResponse[T], *string](r, firstKey)
}

// SetFirstLink Sets the url to the first page of results.
func (r *BaseServiceNowCollectionResponse[T]) SetFirstLink(link *string) error {
	return DefaultBackedModelMutatorFunc(r, firstKey, link)
}

// GetLastLink Returns the url to the last page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetLastLink() (*string, error) {
	return DefaultBackedModelAccessorFunc[*BaseServiceNowCollectionResponse[T], *string](r, lastKey)
}

// SetLastLink Sets the url to the last page of results.
func (r *BaseServiceNowCollectionResponse[T]) SetLastLink(link *string) error {
	return DefaultBackedModelMutatorFunc(r, lastKey, link)
}
