package tableapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	displayValueKey = "display_value"
	valueKey        = "value"
	linkKey         = "link"
)

// RecordElement represents an element within a record.
type RecordElement interface {
	GetDisplayValue() (ElementValue, error)
	GetValue() (ElementValue, error)
	GetLink() (*string, error)
	SetDisplayValue(ElementValue) error
	SetValue(ElementValue) error
	setLink(*string) error
	IsDisplayValueOnly() bool
	value() (interface{}, error)
	serialization.Parsable
	store.BackedModel
}

// recordElement is an implementation of RecordElement.
type recordElement struct {
	backingStore        store.BackingStore
	backingStoreFactory store.BackingStoreFactory
	displayValueOnly    bool
}

// NewRecordElement creates a new Record Element instance.
func NewRecordElement() RecordElement {
	return &recordElement{
		backingStore:        store.BackingStoreFactoryInstance(),
		backingStoreFactory: store.BackingStoreFactoryInstance,
		displayValueOnly:    false,
	}
}

// CreateRecordElementFromDiscriminatorValue creates a new Record Element from a parse node.
func CreateRecordElementFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	tableRecord := NewRecordElement()

	return tableRecord, nil
}

// IsDisplayValueOnly returns if record is display value only.
func (rE *recordElement) IsDisplayValueOnly() bool {
	if internal.IsNil(rE) {
		return false
	}

	return rE.displayValueOnly
}

// GetBackingStore returns the backing store of the record.
func (rE *recordElement) GetBackingStore() store.BackingStore {
	if internal.IsNil(rE) {
		return nil
	}

	if internal.IsNil(rE.backingStore) {
		rE.backingStore = rE.backingStoreFactory()
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

// GetFieldDeserializers returns the deserialization information for this object.
func (rE *recordElement) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		displayValueKey: func(pn serialization.ParseNode) error {
			displayValue, err := pn.GetRawValue()
			if err != nil {
				return err
			}

			return rE.SetDisplayValue(newElementValue(displayValue))
		},
		valueKey: func(pn serialization.ParseNode) error {
			value, err := pn.GetObjectValue(CreateElementValuetFromDiscriminatorValue)
			if err != nil {
				return err
			}

			return rE.SetValue(newElementValue(value))
		},
		linkKey: func(pn serialization.ParseNode) error {
			value, err := pn.GetStringValue()
			if err != nil {
				return err
			}

			return rE.setLink(value)
		},
	}
}

// GetDisplayValue returns the display value.
func (rE *recordElement) GetDisplayValue() (ElementValue, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(displayValueKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(ElementValue)
	if !ok {
		return nil, errors.New("val is not ElementValue")
	}

	return typedVal, nil
}

// SetDisplayValue returns the display value.
func (rE *recordElement) SetDisplayValue(val ElementValue) error {
	if internal.IsNil(rE) {
		return nil
	}

	err := rE.GetBackingStore().Set(displayValueKey, val)
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

	if rE.displayValueOnly {
		return nil, errors.New("record is display value only")
	}

	val, err := rE.GetBackingStore().Get(valueKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(ElementValue)
	if !ok {
		return nil, errors.New("val is not ElementValue")
	}

	return typedVal, nil
}

// SetValue returns the display value.
func (rE *recordElement) SetValue(val ElementValue) error {
	if internal.IsNil(rE) {
		return nil
	}

	err := rE.GetBackingStore().Set(valueKey, val)
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

	val, err := rE.GetBackingStore().Get(linkKey)
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

	err := rE.GetBackingStore().Set(linkKey, val)
	if err != nil {
		return err
	}

	return nil
}

// value returns the proper value.
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
