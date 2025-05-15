package internal

import (
	"errors"
	"fmt"

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

type ServiceNowCollectionResponse[T serialization.Parsable] interface {
	GetResult() ([]T, error)
	GetNextLink() (*string, error)
	GetPreviousLink() (*string, error)
	GetFirstLink() (*string, error)
	GetLastLink() (*string, error)
	serialization.Parsable
	store.BackedModel
}

type BaseServiceNowCollectionResponse[T serialization.Parsable] struct {
	store store.BackingStore
}

func NewBaseServiceNowCollectionResponse[T serialization.Parsable]() *BaseServiceNowCollectionResponse[T] {
	return &BaseServiceNowCollectionResponse[T]{
		store: store.NewInMemoryBackingStore(),
	}
}

func (r *BaseServiceNowCollectionResponse[T]) GetBackingStore() store.BackingStore {
	return r.store
}

func (r *BaseServiceNowCollectionResponse[T]) GetResult() ([]T, error) {
	if IsNil(r) {
		return nil, nil
	}

	store := r.GetBackingStore()
	if store == nil {
		return nil, errors.New("store is nil")
	}

	val, err := store.Get(resultKey)
	if err != nil {
		return nil, err
	}

	unknownSlice, ok := val.([]interface{})
	if !ok {
		return nil, errors.New("val is not slice")
	}

	results := make([]T, len(unknownSlice), 0)

	for index, value := range unknownSlice {
		result, ok := value.(T)
		if !ok {
			return nil, fmt.Errorf("value is not %T", new(T))
		}

		results[index] = result
	}

	return results, nil
}

func (r *BaseServiceNowCollectionResponse[T]) GetNextLink() (*string, error) {
	if IsNil(r) {
		return nil, nil
	}

	store := r.GetBackingStore()
	if store == nil {
		return nil, errors.New("store is nil")
	}

	val, err := store.Get(nextKey)
	if err != nil {
		return nil, err
	}

	link, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return link, nil
}

func (r *BaseServiceNowCollectionResponse[T]) GetPreviousLink() (*string, error) {
	if IsNil(r) {
		return nil, nil
	}

	store := r.GetBackingStore()
	if store == nil {
		return nil, errors.New("store is nil")
	}

	val, err := store.Get(previousKey)
	if err != nil {
		return nil, err
	}

	link, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return link, nil
}

func (r *BaseServiceNowCollectionResponse[T]) GetFirstLink() (*string, error) {
	if IsNil(r) {
		return nil, nil
	}

	store := r.GetBackingStore()
	if store == nil {
		return nil, errors.New("store is nil")
	}

	val, err := store.Get(firstKey)
	if err != nil {
		return nil, err
	}

	link, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return link, nil
}

func (r *BaseServiceNowCollectionResponse[T]) GetLastLink() (*string, error) {
	if IsNil(r) {
		return nil, nil
	}

	store := r.GetBackingStore()
	if store == nil {
		return nil, errors.New("store is nil")
	}

	val, err := store.Get(lastKey)
	if err != nil {
		return nil, err
	}

	link, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return link, nil
}
