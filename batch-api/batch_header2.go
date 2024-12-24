package batchapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	nameKey  = "name"
	valueKey = "value"
)

// BatchHeaderable represents a Service-Now batch header
type BatchHeaderable interface {
	GetName() (*string, error)
	SetName(*string) error
	GetValue() (*string, error)
	SetValue(*string) error
	serialization.Parsable
	store.BackedModel
}

// batchHeader2 implementation of BatchHeaderable
type batchHeader2 struct {
	// backingStoreFactory factory to create backingStore
	backingStoreFactory store.BackingStoreFactory
	// backingStore the store backing the model
	backingStore store.BackingStore
}

// NewBatchHeader2 creates new instance of BatchHeaderable
func NewBatchHeader2() BatchHeaderable {
	return &batchHeader2{
		backingStore:        store.NewInMemoryBackingStore(),
		backingStoreFactory: store.NewInMemoryBackingStore,
	}
}

// CreateBatchHeader2FromDiscriminatorValue is a parsable factory for creating a BatchRequestable
func CreateBatchHeader2FromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewBatchHeader2(), nil
}

// GetBackingStore retrieves the backing store for the model.
func (bH *batchHeader2) GetBackingStore() store.BackingStore {
	if internal.IsNil(bH) {
		return nil
	}

	if internal.IsNil(bH.backingStore) {
		bH.backingStore = bH.backingStoreFactory()
	}

	return bH.backingStore
}

// Serialize writes the objects properties to the current writer.
func (bH *batchHeader2) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(bH) {
		return nil
	}

	serializers := []func(serialization.SerializationWriter) error{
		func(sw serialization.SerializationWriter) error {
			name, err := bH.GetName()
			if err != nil {
				return err
			}
			return sw.WriteStringValue(nameKey, name)
		},
		func(sw serialization.SerializationWriter) error {
			value, err := bH.GetValue()
			if err != nil {
				return err
			}
			return sw.WriteStringValue(valueKey, value)
		},
	}

	for _, serializer := range serializers {
		if err := serializer(writer); err != nil {
			return err
		}
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (bH *batchHeader2) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(bH) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		nameKey: func(pn serialization.ParseNode) error {
			return errors.New("deserializer (nameKey) not implemented")
		},
		valueKey: func(pn serialization.ParseNode) error {
			return errors.New("deserializer (valueKey) not implemented")
		},
	}
}

// GetName returns the name of the header
func (bH *batchHeader2) GetName() (*string, error) {
	if internal.IsNil(bH) {
		return nil, nil
	}

	name, err := bH.GetBackingStore().Get(nameKey)
	if err != nil {
		return nil, err
	}

	typedName, ok := name.(*string)
	if !ok {
		return nil, errors.New("name is not *string")
	}

	return typedName, nil
}

// SetName sets name to provided value
func (bH *batchHeader2) SetName(name *string) error {
	if internal.IsNil(bH) {
		return nil
	}

	return bH.GetBackingStore().Set(nameKey, name)
}

// GetValue returns the value of the header
func (bH *batchHeader2) GetValue() (*string, error) {
	if internal.IsNil(bH) {
		return nil, nil
	}

	value, err := bH.GetBackingStore().Get(valueKey)
	if err != nil {
		return nil, err
	}

	typedValue, ok := value.(*string)
	if !ok {
		return nil, errors.New("value is not *string")
	}

	return typedValue, nil
}

// SetValue sets the value to the provided value
func (bH *batchHeader2) SetValue(value *string) error {
	if internal.IsNil(bH) {
		return nil
	}

	return bH.GetBackingStore().Set(valueKey, value)
}
