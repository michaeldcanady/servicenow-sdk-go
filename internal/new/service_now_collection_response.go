package internal

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
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
	ParseHeaders(*abstractions.ResponseHeaders)
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

// ParseHeaders parses the needed headers from the response.
func (bR *BaseServiceNowCollectionResponse[T]) ParseHeaders(headers *abstractions.ResponseHeaders) {
	linkHeaderRegex := regexp.MustCompile(`<([^>]+)>;rel="([^"]+)"`)

	hearderLinks := headers.Get("Link")

	for _, header := range hearderLinks {
		linkMatches := linkHeaderRegex.FindAllStringSubmatch(header, -1)

		for _, match := range linkMatches {
			link := match[1]
			rel := match[2]

			// Determine the type of link based on the 'rel' attribute
			switch rel {
			case "first":
				_ = bR.setFirstLink(&link)
			case "prev":
				_ = bR.setPreviousLink(&link)
			case "next":
				_ = bR.setNextLink(&link)
			case "last":
				_ = bR.setLastLink(&link)
			}
		}
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
		nextKey: func(pn serialization.ParseNode) error {
			val, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return bR.setNextLink(val)
		},
		previousKey: func(pn serialization.ParseNode) error {
			val, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return bR.setPreviousLink(val)
		},
		firstKey: func(pn serialization.ParseNode) error {
			val, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return bR.setFirstLink(val)
		},
		lastKey: func(pn serialization.ParseNode) error {
			val, err := pn.GetStringValue()
			if err != nil {
				return err
			}
			return bR.setLastLink(val)
		},
	}
}

// setResult Sets the result values of the response.
func (r *BaseServiceNowCollectionResponse[T]) setResult(val []any) error {
	if IsNil(r) {
		return nil
	}

	store, err := r.GetBackingStore()
	if err != nil {
		return err
	}

	return store.Set(resultKey, val)
}

// setNextLink Sets the url to the next page of results.
func (r *BaseServiceNowCollectionResponse[T]) setNextLink(val *string) error {
	if IsNil(r) {
		return nil
	}

	store, err := r.GetBackingStore()
	if err != nil {
		return err
	}

	return store.Set(nextKey, val)
}

// setPreviousLink Sets the url to the previous page of results.
func (r *BaseServiceNowCollectionResponse[T]) setPreviousLink(val *string) error {
	if IsNil(r) {
		return nil
	}

	store, err := r.GetBackingStore()
	if err != nil {
		return err
	}

	return store.Set(previousKey, val)
}

// setFirstLink Sets the url to the first page of results.
func (r *BaseServiceNowCollectionResponse[T]) setFirstLink(val *string) error {
	if IsNil(r) {
		return nil
	}

	store, err := r.GetBackingStore()
	if err != nil {
		return err
	}

	return store.Set(firstKey, val)
}

// setLastLink Sets the url to the last page of results.
func (r *BaseServiceNowCollectionResponse[T]) setLastLink(val *string) error {
	if IsNil(r) {
		return nil
	}

	store, err := r.GetBackingStore()
	if err != nil {
		return err
	}

	return store.Set(lastKey, val)
}

// GetBackingStore returns the backing store, if store is nil it instantiates a new store.
func (r *BaseServiceNowCollectionResponse[T]) GetBackingStore() (store.BackingStore, error) {
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

// GetResult Returns the result values of the response.
func (r *BaseServiceNowCollectionResponse[T]) GetResult() ([]T, error) {
	if IsNil(r) {
		return nil, nil
	}

	store, err := r.GetBackingStore()
	if err != nil {
		return nil, err
	}

	val, err := store.Get(resultKey)
	if err != nil {
		return nil, err
	}

	unknownSlice, ok := val.([]any)
	if !ok {
		return nil, errors.New("val is not slice")
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
	if IsNil(r) {
		return nil, nil
	}

	store, err := r.GetBackingStore()
	if err != nil {
		return nil, err
	}

	val, err := store.Get(nextKey)
	if err != nil {
		return nil, err
	}

	if IsNil(val) {
		return nil, nil
	}

	link, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return link, nil
}

// GetPreviousLink Returns the url to the previous page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetPreviousLink() (*string, error) {
	if IsNil(r) {
		return nil, nil
	}

	store, err := r.GetBackingStore()
	if err != nil {
		return nil, err
	}

	val, err := store.Get(previousKey)
	if err != nil {
		return nil, err
	}

	if IsNil(val) {
		return nil, nil
	}

	link, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return link, nil
}

// GetFirstLink Returns the url to the first page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetFirstLink() (*string, error) {
	if IsNil(r) {
		return nil, nil
	}

	store, err := r.GetBackingStore()
	if err != nil {
		return nil, err
	}

	val, err := store.Get(firstKey)
	if err != nil {
		return nil, err
	}

	if IsNil(val) {
		return nil, nil
	}

	link, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return link, nil
}

// GetLastLink Returns the url to the last page of results.
func (r *BaseServiceNowCollectionResponse[T]) GetLastLink() (*string, error) {
	if IsNil(r) {
		return nil, nil
	}

	store, err := r.GetBackingStore()
	if err != nil {
		return nil, err
	}

	val, err := store.Get(lastKey)
	if err != nil {
		return nil, err
	}

	if IsNil(val) {
		return nil, nil
	}

	link, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return link, nil
}
