package tableapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

// RecordElement represents an element within a record.
type RecordElement interface {
	GetDisplayValue() (ElementValue, error)
	GetValue() (ElementValue, error)
	GetLink() (*string, error)
	SetDisplayValue(interface{}) error
	SetValue(interface{}) error
	setLink(*string) error
	value() (interface{}, error)
	serialization.Parsable
	store.BackedModel
}

func NewRecordElement() RecordElement {
	return &recordElement{
		backingStore:     store.BackingStoreFactoryInstance(),
		displayValueOnly: false,
	}
}

// recordElement is an implementation of RecordElement.
type recordElement struct {
	backingStore     store.BackingStore
	displayValueOnly bool
}

func CreateRecordElementFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	tableRecord := NewRecordElement()

	return tableRecord, nil
}

func (rE *recordElement) GetBackingStore() store.BackingStore {
	if internal.IsNil(rE) {
		return nil
	}

	if internal.IsNil(rE.backingStore) {
		rE.backingStore = store.NewInMemoryBackingStore()
	}

	return rE.backingStore
}

// Serialize writes the objects properties to the current writer.
func (rE *recordElement) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(rE) {
		return nil
	}

	return errors.New("Serialize not implemented")
}

func (rE *recordElement) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{}
}

// GetDisplayValue returns the display value.
func (rE *recordElement) GetDisplayValue() (ElementValue, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get("display_value")
	if err != nil {
		return nil, err
	}

	return &elementValue{val: val}, nil
}

// SetDisplayValue returns the display value.
func (rE *recordElement) SetDisplayValue(val interface{}) error {
	if internal.IsNil(rE) {
		return nil
	}

	err := rE.GetBackingStore().Set("display_value", val)
	if err != nil {
		return err
	}

	return nil
}

// GetValue returns the raw value.
func (rE *recordElement) GetValue() (ElementValue, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get("value")
	if err != nil {
		return nil, err
	}

	return &elementValue{val: val}, nil
}

// SetValue returns the display value.
func (rE *recordElement) SetValue(val interface{}) error {
	if internal.IsNil(rE) {
		return nil
	}

	err := rE.GetBackingStore().Set("value", val)
	if err != nil {
		return err
	}

	return nil
}

// GetLink returns the link.
func (rE *recordElement) GetLink() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get("link")
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("link is not a *string")
	}

	return typedVal, nil
}

// setLink returns the display value.
func (rE *recordElement) setLink(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	err := rE.GetBackingStore().Set("link", val)
	if err != nil {
		return err
	}

	return nil
}

func (rE *recordElement) value() (interface{}, error) {
	if rE.displayValueOnly {
		return rE.GetDisplayValue()
	}
	eV, err := rE.GetValue()
	if err != nil {
		return nil, err
	}
	return eV.GetRawValue()
}
