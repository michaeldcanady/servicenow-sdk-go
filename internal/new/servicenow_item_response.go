package internal

import (
	"fmt"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

// ServiceNowItemResponse implementation of ServiceNowResponse
type ServiceNowItemResponse[T serialization.Parsable] struct {
	*ServiceNowResponseImpl[T]
}

// NewServiceNowItemResponse creates a new instance of a ServiceNowResponse from a parsable factory.
func NewServiceNowItemResponse[T serialization.Parsable](factory serialization.ParsableFactory, storeFactory store.BackingStoreFactory) *ServiceNowItemResponse[T] {
	return &ServiceNowItemResponse[T]{
		NewServiceNowResponse[T](ServiceNowResponseTypeItem, factory, storeFactory),
	}
}

// CreateServiceNowItemResponseFromDiscriminatorValue[T serialization.Parsable] is a factory for creating a ServiceNowResponse
func CreateServiceNowItemResponseFromDiscriminatorValue[T serialization.Parsable](factory serialization.ParsableFactory) serialization.ParsableFactory {
	return func(parseNode serialization.ParseNode) (serialization.Parsable, error) {
		return NewServiceNowItemResponse[T](factory, store.BackingStoreFactoryInstance), nil
	}
}

// Serialize writes the objects properties to the current writer
func (tE *ServiceNowItemResponse[T]) Serialize(writer serialization.SerializationWriter) error {
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object
func (tE *ServiceNowItemResponse[T]) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		resultKey: func(pn serialization.ParseNode) error {
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

			elem, ok := rawElem.(T)
			if !ok {
				return fmt.Errorf("elem is not %v", new(T))
			}

			if err := tE.SetResult(elem); err != nil {
				return err
			}

			return nil
		},
	}
}
