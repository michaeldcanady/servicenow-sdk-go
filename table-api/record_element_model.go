//go:build preview.tableApiV2

package tableapi

import (
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
)

// RecordElement implements the RecordElement interface.
//
// This model encapsulates structured data storage for table entries and provides
// methods for retrieving and updating values related to a specific record.
//
// Example usage:
//
//	element := NewRecordElement()
//	err := element.SetValue("active")
//	if err != nil {
//	    log.Fatal(err)
//	}
type RecordElement struct {
	internal.Model
}

// NewRecordElement creates a new instance of RecordElementModel.
//
// This function initializes a RecordElementModel with a new backing store.
//
// Example:
//
//	element := NewRecordElement()
func NewRecordElement() *RecordElement {
	return &RecordElement{
		internal.NewBaseModel(),
	}
}

// GetDisplayValue retrieves the display value associated with the element.
//
// This method accesses the internal storage and extracts the `display_value`
// attribute if present.
//
// Example:
//
//	displayValue, err := element.GetDisplayValue()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(displayValue)
func (rE *RecordElement) GetDisplayValue() (*ElementValue, error) {
	val, err := internal.DefaultBackedModelAccessorFunc[*RecordElement, ElementValue](rE, displayValueKey)
	return &val, err
}

// SetDisplayValue updates the display value of the element.
//
// Example:
//
//	err := element.SetDisplayValue("Active")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (rE *RecordElement) SetDisplayValue(value any) error {
	if _, ok := value.(*ElementValue); !ok {
		var err error
		if value, err = NewElementValue(value); err != nil {
			return err
		}
	}

	return internal.DefaultBackedModelMutatorFunc(rE, displayValueKey, value)
}

// GetValue retrieves the raw stored value of the element.
//
// This method accesses the internal backing store and fetches the `value` field.
//
// Example:
//
//	value, err := element.GetValue()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(value)
func (rE *RecordElement) GetValue() (*ElementValue, error) {
	val, err := internal.DefaultBackedModelAccessorFunc[*RecordElement, ElementValue](rE, valueKey)
	return &val, err
}

// SetValue updates the stored value of the element.
//
// Example:
//
//	err := element.SetValue("active")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (rE *RecordElement) SetValue(value any) error {
	if _, ok := value.(*ElementValue); !ok {
		var err error
		if value, err = NewElementValue(value); err != nil {
			return err
		}
	}

	return internal.DefaultBackedModelMutatorFunc(rE, valueKey, value)
}

// GetLink retrieves the optional link associated with the element.
//
// If available, this method returns a pointer to the link string.
//
// Example:
//
//	link, err := element.GetLink()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(*link)
func (rE *RecordElement) GetLink() (*string, error) {
	val, err := internal.DefaultBackedModelAccessorFunc[*RecordElement, string](rE, linkKey)
	return &val, err
}

// setLink assigns an optional reference link to the element.
//
// Example:
//
//	link := "https://example.com"
//	if err := element.setLink(&link); err != nil {
//	    log.Fatal(err)
//	}
func (rE *RecordElement) setLink(link *string) error {
	return internal.DefaultBackedModelMutatorFunc(rE, linkKey, link)
}
