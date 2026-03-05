package tableapi

import (
	"errors"

	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// RecordElement represents a single field in a TableRecord.
//
// It contains the raw value, the display value, and an optional reference link.
type RecordElement struct {
	internal.Model
}

// NewRecordElement creates a new instance of RecordElement.
func NewRecordElement() *RecordElement {
	return &RecordElement{
		internal.NewBaseModel(),
	}
}

const (
	recordDisplayValueKey = "display_value"
	recordValueKey        = "value"
	recordLinkKey         = "link"
)

// GetDisplayValue returns the display value of the element.
func (rE *RecordElement) GetDisplayValue() (ElementValue, error) {
	if internal.IsNil(rE) {
		return ElementValue{}, errors.New("model is nil")
	}

	backingStore := rE.GetBackingStore()
	if internal.IsNil(backingStore) {
		return ElementValue{}, errors.New("store is nil")
	}

	val, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, ElementValue](backingStore, recordDisplayValueKey)
	if err != nil {
		return ElementValue{}, err
	}

	return val, nil
}

// SetDisplayValue sets the display value of the element.
func (rE *RecordElement) SetDisplayValue(value any) error {
	if internal.IsNil(rE) {
		return errors.New("model is nil")
	}

	val, err := NewElementValue(value)
	if err != nil {
		return err
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, recordDisplayValueKey, *val)
}

// GetValue returns the raw value of the element.
func (rE *RecordElement) GetValue() (ElementValue, error) {
	if internal.IsNil(rE) {
		return ElementValue{}, errors.New("model is nil")
	}

	backingStore := rE.GetBackingStore()
	if internal.IsNil(backingStore) {
		return ElementValue{}, errors.New("store is nil")
	}

	val, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, ElementValue](backingStore, recordValueKey)
	if err != nil {
		return ElementValue{}, err
	}

	return val, nil
}

// SetValue sets the raw value of the element.
func (rE *RecordElement) SetValue(value any) error {
	if internal.IsNil(rE) {
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
	if internal.IsNil(rE) {
		return "", errors.New("model is nil")
	}

	backingStore := rE.GetBackingStore()
	if internal.IsNil(backingStore) {
		return "", errors.New("store is nil")
	}

	val, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, string](backingStore, recordLinkKey)
	if err != nil {
		return "", err
	}

	return val, nil
}

// SetLink sets the reference link of the element.
func (rE *RecordElement) SetLink(link *string) error {
	if internal.IsNil(rE) {
		return errors.New("model is nil")
	}

	var val string
	if link != nil {
		val = *link
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, recordLinkKey, val)
}
