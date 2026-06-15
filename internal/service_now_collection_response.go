package internal

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
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
	BaseModel
	// factory The ParsableFactory for serializing the result type.
	factory serialization.ParsableFactory
}

// NewBaseServiceNowCollectionResponse[T] Creates a new instance of the BaseServiceNowCollectionResponse[T].
func NewBaseServiceNowCollectionResponse[T serialization.Parsable](factory serialization.ParsableFactory) *BaseServiceNowCollectionResponse[T] {
	res := &BaseServiceNowCollectionResponse[T]{
		BaseModel: *NewBaseModel(),
		factory:   factory,
	}
	res.GetBackingStore()
	return res
}

// Serialize writes the objects properties to the current writer
func (bR *BaseServiceNowCollectionResponse[T]) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(bR) {
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
		resultKey:   internalSerialization.DeserializeCollectionOfObjectValuesFunc[T](bR.factory)(bR.setResult),
		nextKey:     internalSerialization.DeserializeStringFunc()(bR.SetNextLink),
		previousKey: internalSerialization.DeserializeStringFunc()(bR.SetPreviousLink),
		firstKey:    internalSerialization.DeserializeStringFunc()(bR.SetFirstLink),
		lastKey:     internalSerialization.DeserializeStringFunc()(bR.SetLastLink),
	}
}

// setResult Sets the result values of the response.
func (r *BaseServiceNowCollectionResponse[T]) setResult(val []T) error {
	if conversion.IsNil(r) {
		return nil
	}

	anySlice, _ := conversion.CastCollection[T, any](val)

	return store.DefaultBackedModelMutatorFunc(r.GetBackingStore(), resultKey, anySlice)
}

// SetNextLink Sets the url to the next page of results.
func (r *BaseServiceNowCollectionResponse[T]) SetNextLink(val *string) error {
	if conversion.IsNil(r) {
		return nil
	}

	return store.DefaultBackedModelMutatorFunc(r.GetBackingStore(), nextKey, val)
}

// SetPreviousLink Sets the url to the previous page of results.
func (r *BaseServiceNowCollectionResponse[T]) SetPreviousLink(val *string) error {
	if conversion.IsNil(r) {
		return nil
	}

	return store.DefaultBackedModelMutatorFunc(r.GetBackingStore(), previousKey, val)
}

// SetFirstLink Sets the url to the first page of results.
func (r *BaseServiceNowCollectionResponse[T]) SetFirstLink(val *string) error {
	if conversion.IsNil(r) {
		return nil
	}

	return store.DefaultBackedModelMutatorFunc(r.GetBackingStore(), firstKey, val)
}

// SetLastLink Sets the url to the last page of results.
func (r *BaseServiceNowCollectionResponse[T]) SetLastLink(val *string) error {
	if conversion.IsNil(r) {
		return nil
	}

	return store.DefaultBackedModelMutatorFunc(r.GetBackingStore(), lastKey, val)
}

// GetResult Returns the result values of the response.
func (r *BaseServiceNowCollectionResponse[T]) GetResult() ([]T, error) {
	if conversion.IsNil(r) {
		return nil, nil
	}

	unknownSlice, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []any](r.GetBackingStore(), resultKey)
	if err != nil {
		return nil, err
	}

	return conversion.CastCollection[any, T](unknownSlice)
}

// GetNextLink Returns the url to the next page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetNextLink() (*string, error) {
	if conversion.IsNil(r) {
		return nil, nil
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](r.GetBackingStore(), nextKey)
}

// GetPreviousLink Returns the url to the previous page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetPreviousLink() (*string, error) {
	if conversion.IsNil(r) {
		return nil, nil
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](r.GetBackingStore(), previousKey)
}

// GetFirstLink Returns the url to the first page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetFirstLink() (*string, error) {
	if conversion.IsNil(r) {
		return nil, nil
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](r.GetBackingStore(), firstKey)
}

// GetLastLink Returns the url to the last page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetLastLink() (*string, error) {
	if conversion.IsNil(r) {
		return nil, nil
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](r.GetBackingStore(), lastKey)
}
