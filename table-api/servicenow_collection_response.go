package tableapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	resultKey = "Result"
)

// ServiceNowCollectionResponse represents a Service-Now collection response
type ServiceNowCollectionResponse interface {
	GetResult() ([]serialization.Parsable, error)
	GetNextLink() (*string, error)
	GetPreviousLink() (*string, error)
	GetFirstLink() (*string, error)
	GetLastLink() (*string, error)
	setNextLink(*string) error
	setPreviousLink(*string) error
	setFirstLink(*string) error
	setLastLink(*string) error
	setResult([]serialization.Parsable) error
	serialization.Parsable
	store.BackedModel
}

// serviceNowCollectionResponse implementation of ServiceNowCollectionResponse
type serviceNowCollectionResponse struct {
	factory             serialization.ParsableFactory
	backingStoreFactory store.BackingStoreFactory
	backingStore        store.BackingStore
}

// NewServiceNowCollectionResponse creates a new instance of a ServiceNowCollectionResponse from a parsable factory.
func NewServiceNowCollectionResponse(factory serialization.ParsableFactory) ServiceNowCollectionResponse {
	return &serviceNowCollectionResponse{
		factory:             factory,
		backingStoreFactory: store.BackingStoreFactoryInstance,
		backingStore:        store.BackingStoreFactoryInstance(),
	}
}

// CreateServiceNowCollectionResponseFromDiscriminatorValue is a factory for creating a ServiceNowCollectionResponse
func CreateServiceNowCollectionResponseFromDiscriminatorValue(factory serialization.ParsableFactory) serialization.ParsableFactory {
	return func(parseNode serialization.ParseNode) (serialization.Parsable, error) {
		return NewServiceNowCollectionResponse(factory), nil
	}
}

// Serialize writes the objects properties to the current writer
func (tE *serviceNowCollectionResponse) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(tE) {
		return nil
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object
func (tE *serviceNowCollectionResponse) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(tE) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		resultKey: func(pn serialization.ParseNode) error {
			elem, err := pn.GetCollectionOfObjectValues(tE.factory)
			if err != nil {
				return err
			}

			if err := tE.setResult(elem); err != nil {
				return err
			}

			return nil
		},
	}
}

// GetBackingStore returns the backing store of the record
func (tE *serviceNowCollectionResponse) GetBackingStore() store.BackingStore {
	if internal.IsNil(tE) {
		return nil
	}

	if internal.IsNil(tE.backingStore) {
		tE.backingStore = tE.backingStoreFactory()
	}

	return tE.backingStore
}

// GetResult returns result slice from Service-Now Response
func (tE *serviceNowCollectionResponse) GetResult() ([]serialization.Parsable, error) {
	if internal.IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(resultKey)
	if err != nil {
		return nil, err
	}
	if internal.IsNil(val) {
		return []serialization.Parsable{}, nil
	}

	typedVal, ok := val.([]serialization.Parsable)
	if !ok {
		return nil, errors.New("val is not serialization.Parsable")
	}

	return typedVal, nil
}

// setResult sets the result slice for the Service-Now Response
func (tE *serviceNowCollectionResponse) setResult(result []serialization.Parsable) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(resultKey, result)
}

// GetNextLink returns next link, if it exists
func (tE *serviceNowCollectionResponse) GetNextLink() (*string, error) {
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

// GetPreviousLink returns previous link, if it exists
func (tE *serviceNowCollectionResponse) GetPreviousLink() (*string, error) {
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

// GetFirstLink returns first link, if it exists
func (tE *serviceNowCollectionResponse) GetFirstLink() (*string, error) {
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

// GetLastLink returns last link, if it exists
func (tE *serviceNowCollectionResponse) GetLastLink() (*string, error) {
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

// setNextLink sets next link
func (tE *serviceNowCollectionResponse) setNextLink(nextLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(nextLinkHeaderKey, nextLink)
}

// setPreviousLink sets previous link
func (tE *serviceNowCollectionResponse) setPreviousLink(previousLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(prevLinkHeaderKey, previousLink)
}

// setFirstLink sets first link
func (tE *serviceNowCollectionResponse) setFirstLink(firstLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(firstLinkHeaderKey, firstLink)
}

// setLastLink sets last link
func (tE *serviceNowCollectionResponse) setLastLink(lastLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(lastLinkHeaderKey, lastLink)
}
