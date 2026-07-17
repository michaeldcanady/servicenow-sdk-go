package tableapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
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
	return store.DefaultBackedModelAccessorFunc[*RecordElement, ElementValue](rE, recordDisplayValueKey)
}

func (rE *RecordElement) set(key string, value any) error {
	val, err := NewElementValue(value)
	if err != nil {
		return err
	}
	return store.DefaultBackedModelMutatorFunc(rE, key, *val)
}

// SetDisplayValue sets the display value of the element.
func (rE *RecordElement) SetDisplayValue(value any) error {
	return rE.set(recordDisplayValueKey, value)
}

// GetValue returns the raw value of the element.
func (rE *RecordElement) GetValue() (ElementValue, error) {
	return store.DefaultBackedModelAccessorFunc[*RecordElement, ElementValue](rE, recordValueKey)
}

// SetValue sets the raw value of the element.
func (rE *RecordElement) SetValue(value any) error {
	return rE.set(recordValueKey, value)
}

// GetLink returns the reference link of the element, if it is a reference field.
func (rE *RecordElement) GetLink() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*RecordElement, *string](rE, recordLinkKey)
}

// SetLink sets the reference link of the element.
func (rE *RecordElement) SetLink(link *string) error {
	return store.DefaultBackedModelMutatorFunc(rE, recordLinkKey, link)
}
