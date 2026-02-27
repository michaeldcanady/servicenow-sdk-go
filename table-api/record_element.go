package tableapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
)

const (
	displayValueKey = "display_value"
	valueKey        = "value"
	linkKey         = "link"
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
		model.NewBaseModel(),
	}
}

// GetDisplayValue retrieves the display value associated with the element.
func (rE *RecordElement) GetDisplayValue() (*ElementValue, error) {
	val, err := kiota.DefaultBackedModelAccessorFunc[*RecordElement, ElementValue](rE, displayValueKey)
	return &val, err
}

// SetDisplayValue updates the display value of the element.
func (rE *RecordElement) SetDisplayValue(value any) error {
	if _, ok := value.(*ElementValue); !ok {
		var err error
		if value, err = NewElementValue(value); err != nil {
			return err
		}
	}

	return kiota.DefaultBackedModelMutatorFunc(rE, displayValueKey, value)
}

// GetValue retrieves the raw stored value of the element.
func (rE *RecordElement) GetValue() (*ElementValue, error) {
	val, err := kiota.DefaultBackedModelAccessorFunc[*RecordElement, ElementValue](rE, valueKey)
	return &val, err
}

// SetValue updates the stored value of the element.
func (rE *RecordElement) SetValue(value any) error {
	if _, ok := value.(*ElementValue); !ok {
		var err error
		if value, err = NewElementValue(value); err != nil {
			return err
		}
	}

	return kiota.DefaultBackedModelMutatorFunc(rE, valueKey, value)
}

// GetLink retrieves the optional link associated with the element.
func (rE *RecordElement) GetLink() (*string, error) {
	val, err := kiota.DefaultBackedModelAccessorFunc[*RecordElement, string](rE, linkKey)
	return &val, err
}

// SetLink assigns an optional reference link to the element.
func (rE *RecordElement) SetLink(link *string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE, linkKey, link)
}
