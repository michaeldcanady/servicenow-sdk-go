package tableapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// RecordElement represents a single field in a TableRecord.
//
// It contains the raw value, the display value, and an optional reference link.
type RecordElement struct {
	core.BackedModel
}

// NewRecordElement creates a new instance of RecordElement.
func NewRecordElement() *RecordElement {
	return &RecordElement{
		core.NewBaseModel(),
	}
}

const (
	recordDisplayValueKey = "display_value"
	recordValueKey        = "value"
	recordLinkKey         = "link"
)

// GetDisplayValue returns the display value of the element.
func (rE *RecordElement) GetDisplayValue() (ElementValue, error) {
	if conversion.IsNil(rE) || conversion.IsNil(rE.BackedModel) {
		return ElementValue{}, nil
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, ElementValue](rE.GetBackingStore(), recordDisplayValueKey)
}

// SetDisplayValue sets the display value of the element.
func (rE *RecordElement) SetDisplayValue(value any) error {
	val, err := NewElementValue(value)
	if err != nil {
		return err
	}
	return store.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), recordDisplayValueKey, *val)
}

// GetValue returns the raw value of the element.
func (rE *RecordElement) GetValue() (ElementValue, error) {
	if conversion.IsNil(rE) || conversion.IsNil(rE.BackedModel) {
		return ElementValue{}, nil
	}

	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, ElementValue](rE.GetBackingStore(), recordValueKey)
}

// SetValue sets the raw value of the element.
func (rE *RecordElement) SetValue(value any) error {
	if conversion.IsNil(rE) {
		return errors.New("model is nil")
	}

	val, err := NewElementValue(value)
	if err != nil {
		return err
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, recordValueKey, *val)
}

// GetLink returns the reference link of the element, if it is a reference field.
func (rE *RecordElement) GetLink() (string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, string](rE.GetBackingStore(), recordLinkKey)
}

// SetLink sets the reference link of the element.
func (rE *RecordElement) SetLink(link *string) error {
	var val string
	if link != nil {
		val = *link
	}
	return store.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), recordLinkKey, val)
}
