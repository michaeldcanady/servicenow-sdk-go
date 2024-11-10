package tableapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

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

type serviceNowCollectionResponse struct {
	factory      serialization.ParsableFactory
	backingStore store.BackingStore
}

func NewServiceNowCollectionResponse(factory serialization.ParsableFactory) ServiceNowCollectionResponse {
	return &serviceNowCollectionResponse{
		factory:      factory,
		backingStore: store.BackingStoreFactoryInstance(),
	}
}

func CreateServiceNowCollectionResponseFromDiscriminatorValue(factory serialization.ParsableFactory) serialization.ParsableFactory {
	return func(parseNode serialization.ParseNode) (serialization.Parsable, error) {
		return NewServiceNowCollectionResponse(factory), nil
	}
}

func (tE *serviceNowCollectionResponse) GetBackingStore() store.BackingStore {
	if internal.IsNil(tE) {
		return nil
	}

	if internal.IsNil(tE.backingStore) {
		tE.backingStore = store.BackingStoreFactoryInstance()
	}

	return tE.backingStore
}

// Serialize writes the objects properties to the current writer.
func (tE *serviceNowCollectionResponse) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(tE) {
		return nil
	}
	return nil
}

func (tE *serviceNowCollectionResponse) GetResult() ([]serialization.Parsable, error) {
	if internal.IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get("Result")
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

func (tE *serviceNowCollectionResponse) setResult(result []serialization.Parsable) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set("Result", result)
}

func (tE *serviceNowCollectionResponse) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(tE) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		"Result": func(pn serialization.ParseNode) error {
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

func (tE *serviceNowCollectionResponse) setNextLink(nextLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(nextLinkHeaderKey, nextLink)
}

func (tE *serviceNowCollectionResponse) setPreviousLink(previousLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(prevLinkHeaderKey, previousLink)
}

func (tE *serviceNowCollectionResponse) setFirstLink(firstLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(firstLinkHeaderKey, firstLink)
}

func (tE *serviceNowCollectionResponse) setLastLink(lastLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(lastLinkHeaderKey, lastLink)
}
