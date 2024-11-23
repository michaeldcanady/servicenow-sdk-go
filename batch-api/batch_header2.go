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

type BatchHeaderable interface {
	GetName() (*string, error)
	SetName(*string) error
	GetValue() (*string, error)
	SetValue(*string) error
	serialization.Parsable
	store.BackedModel
}

type batchHeader2 struct {
	backingStore store.BackingStore
}

func NewBatchHeader2() BatchHeaderable {
	return &batchHeader2{
		backingStore: store.NewInMemoryBackingStore(),
	}
}

// CreateBatchHeader2FromDiscriminatorValue is a parsable factory for creating a BatchRequestable
func CreateBatchHeader2FromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewBatchHeader2(), nil
}

// GetBackingStore retrieves the backing store for the model.
func (rE *batchHeader2) GetBackingStore() store.BackingStore {
	if internal.IsNil(rE) {
		return nil
	}

	if internal.IsNil(rE.backingStore) {
		rE.backingStore = store.NewInMemoryBackingStore()
	}

	return rE.backingStore
}

// Serialize writes the objects properties to the current writer.
func (rE *batchHeader2) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(rE) {
		return nil
	}

	serializers := []func(serialization.SerializationWriter) error{
		func(sw serialization.SerializationWriter) error {
			name, err := rE.GetName()
			if err != nil {
				return err
			}
			return sw.WriteStringValue(nameKey, name)
		},
		func(sw serialization.SerializationWriter) error {
			value, err := rE.GetValue()
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
func (rE *batchHeader2) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(rE) {
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

func (rE *batchHeader2) GetName() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	name, err := rE.GetBackingStore().Get(nameKey)
	if err != nil {
		return nil, err
	}

	typedName, ok := name.(*string)
	if !ok {
		return nil, errors.New("name is not *string")
	}

	return typedName, nil
}
func (rE *batchHeader2) SetName(name *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(nameKey, name)
}

func (rE *batchHeader2) GetValue() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	value, err := rE.GetBackingStore().Get(valueKey)
	if err != nil {
		return nil, err
	}

	typedValue, ok := value.(*string)
	if !ok {
		return nil, errors.New("value is not *string")
	}

	return typedValue, nil
}
func (rE *batchHeader2) SetValue(value *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(valueKey, value)
}
