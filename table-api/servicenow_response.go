package tableapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type ServiceNowResponse interface {
	GetResult() (serialization.Parsable, error)
	setResult(serialization.Parsable) error
	serialization.Parsable
	store.BackedModel
}

type serviceNowResponse struct {
	factory      serialization.ParsableFactory
	backingStore store.BackingStore
}

func NewServiceNowResponse(factory serialization.ParsableFactory) ServiceNowResponse {
	return &serviceNowResponse{
		factory:      factory,
		backingStore: store.BackingStoreFactoryInstance(),
	}
}

func CreateServiceNowResponseFromDiscriminatorValue(factory serialization.ParsableFactory) serialization.ParsableFactory {
	return func(parseNode serialization.ParseNode) (serialization.Parsable, error) {
		return NewServiceNowResponse(factory), nil
	}
}

func (tE *serviceNowResponse) GetBackingStore() store.BackingStore {
	return tE.backingStore
}

// Serialize writes the objects properties to the current writer.
func (tE *serviceNowResponse) Serialize(writer serialization.SerializationWriter) error {
	return nil
}

func (tE *serviceNowResponse) GetResult() (serialization.Parsable, error) {
	if internal.IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get("Result")
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(serialization.Parsable)
	if !ok {
		return nil, errors.New("val is not serialization.Parsable")
	}

	return typedVal, nil
}

func (tE *serviceNowResponse) setResult(result serialization.Parsable) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set("Result", result)
}

func (tE *serviceNowResponse) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		"Result": func(pn serialization.ParseNode) error {
			val, err := pn.GetRawValue()
			if err != nil {
				return nil
			}

			var rawElem interface{}

			switch val.(type) {
			case map[string]interface{}:
				rawElem, err = pn.GetObjectValue(tE.factory)
			case []map[string]interface{}:
				rawElem, err = pn.GetCollectionOfObjectValues(tE.factory)
			}

			if err != nil {
				return err
			}

			elem, ok := rawElem.(serialization.Parsable)
			if !ok {
				return errors.New("elem is not serialization.Parsable")
			}

			if err := tE.setResult(elem); err != nil {
				return err
			}

			return nil
		},
	}
}
