package tableapi

import (
	"errors"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	resultKey = "Result"
)

// TODO: move ServiceNowCollectionResponse to internal
// TODO: make ServiceNowCollectionResponse generic for parsable type

// ServiceNowCollectionResponse represents a collection of Service-Now items
type ServiceNowCollectionResponse[T serialization.Parsable] struct {
	ServiceNowResponse[[]T]
}

// NewServiceNowCollectionResponse creates a new instance of a ServiceNowCollectionResponse from a parsable factory.
func NewServiceNowCollectionResponse[T serialization.Parsable](factory serialization.ParsableFactory, storeFactory store.BackingStoreFactory) *ServiceNowCollectionResponse[T] {
	return &ServiceNowCollectionResponse[T]{
		NewServiceNowResponse[[]T](ServiceNowResponseTypeCollection, factory, storeFactory),
	}
}

// CreateServiceNowCollectionResponseFromDiscriminatorValue is a factory for creating a ServiceNowCollectionResponse
func CreateServiceNowCollectionResponseFromDiscriminatorValue(factory serialization.ParsableFactory) serialization.ParsableFactory {
	return func(parseNode serialization.ParseNode) (serialization.Parsable, error) {
		return NewServiceNowCollectionResponse[serialization.Parsable](factory, store.BackingStoreFactoryInstance), nil
	}
}

// GetFieldDeserializers returns the deserialization information for this object
func (tE *ServiceNowCollectionResponse[T]) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(tE) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		resultKey: func(pn serialization.ParseNode) error {
			factory, err := tE.GetFactory()
			if err != nil {
				return err
			}
			if internal.IsNil(factory) {
				return errors.New("factory is nil")
			}
			elems, err := pn.GetCollectionOfObjectValues(factory)
			if err != nil {
				return err
			}

			collection := make([]T, len(elems))

			for index, item := range elems {
				typedItem, ok := item.(T)
				if !ok {
					return fmt.Errorf("item at %v is not %v", index, new(T))
				}

				collection[index] = typedItem
			}

			if err := tE.SetResult(collection); err != nil {
				return err
			}

			return nil
		},
	}
}

// GetNextLink returns next link, if it exists
func (tE *ServiceNowCollectionResponse[T]) GetNextLink() (*string, error) {
	if internal.IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(nextLinkHeaderKey)
	if err != nil {
		return nil, err
	}
	if internal.IsNil(val) {
		return nil, nil
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setNextLink sets next link
func (tE *ServiceNowCollectionResponse[T]) setNextLink(nextLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(nextLinkHeaderKey, nextLink)
}

// GetPreviousLink returns previous link, if it exists
func (tE *ServiceNowCollectionResponse[T]) GetPreviousLink() (*string, error) {
	if internal.IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(prevLinkHeaderKey)
	if err != nil {
		return nil, err
	}
	if internal.IsNil(val) {
		return nil, nil
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setPreviousLink sets previous link
func (tE *ServiceNowCollectionResponse[T]) setPreviousLink(previousLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(prevLinkHeaderKey, previousLink)
}

// GetFirstLink returns first link, if it exists
func (tE *ServiceNowCollectionResponse[T]) GetFirstLink() (*string, error) {
	if internal.IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(firstLinkHeaderKey)
	if err != nil {
		return nil, err
	}
	if internal.IsNil(val) {
		return nil, nil
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setFirstLink sets first link
func (tE *ServiceNowCollectionResponse[T]) setFirstLink(firstLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(firstLinkHeaderKey, firstLink)
}

// GetLastLink returns last link, if it exists
func (tE *ServiceNowCollectionResponse[T]) GetLastLink() (*string, error) {
	if internal.IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(lastLinkHeaderKey)
	if err != nil {
		return nil, err
	}
	if internal.IsNil(val) {
		return nil, nil
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setLastLink sets last link
func (tE *ServiceNowCollectionResponse[T]) setLastLink(lastLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(lastLinkHeaderKey, lastLink)
}
