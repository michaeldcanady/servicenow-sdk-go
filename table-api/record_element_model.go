package tableapi

import (
	"errors"

	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecordElementModel implements the RecordElement interface.
//
// This model encapsulates structured data storage for table entries and provides
// methods for retrieving and updating values related to a specific record.
type RecordElementModel struct {
	internal.Model
	//singleValue possible to have only display value, value, or link
	singleValue bool
}

// GetFieldDeserializers implements serialization.Parsable.
func (rE *RecordElementModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if rE.singleValue {
		return make(map[string]func(serialization.ParseNode) error)
	}

	return map[string]func(serialization.ParseNode) error{
		displayValueKey: func(pn serialization.ParseNode) error {
			rawValue, err := pn.GetRawValue()
			if err != nil {
				return err
			}

			return rE.SetDisplayValue(rawValue)
		},
		valueKey: func(pn serialization.ParseNode) error {
			rawValue, err := pn.GetRawValue()
			if err != nil {
				return err
			}

			return rE.SetValue(rawValue)
		},
		linkKey: func(pn serialization.ParseNode) error {
			rawValue, err := pn.GetStringValue()
			if err != nil {
				return err
			}

			return rE.setLink(rawValue)
		},
	}
}

// Serialize implements serialization.Parsable.
func (rE *RecordElementModel) Serialize(writer serialization.SerializationWriter) error {
	panic("unimplemented")
}

// NewRecordElement creates a new instance of RecordElementModel.
//
// This function initializes a RecordElementModel with a new backing store.
func NewRecordElement() *RecordElementModel {
	return &RecordElementModel{
		internal.NewBaseModel(),
		false,
	}
}

func NewRecordElementFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	element := NewRecordElement()

	rawValue, err := parseNode.GetRawValue()
	if err != nil {
		return nil, err
	}

	switch rawValue.(type) {
	case map[string]any:
		element.singleValue = false
	default:
		element.singleValue = true
		element.GetBackingStore().Set(valueKey, rawValue)
	}

	return element, nil
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
func (rE *RecordElementModel) GetDisplayValue() (ElementValue, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	store := rE.GetBackingStore()
	if internal.IsNil(store) {
		return nil, nil
	}

	value, err := store.Get(displayValueKey)
	if err != nil {
		return nil, err
	}

	return NewElementValue(value), nil
}

// SetDisplayValue updates the display value of the element.
//
// Example:
//
//	err := element.SetDisplayValue("Active")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (rE *RecordElementModel) SetDisplayValue(value any) error {
	if internal.IsNil(rE) {
		return nil
	}

	store := rE.GetBackingStore()
	if internal.IsNil(store) {
		return nil
	}

	return store.Set(displayValueKey, value)
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
func (rE *RecordElementModel) GetValue() (ElementValue, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	store := rE.GetBackingStore()
	if internal.IsNil(store) {
		return nil, nil
	}

	value, err := store.Get(valueKey)
	if err != nil {
		return nil, err
	}

	return NewElementValue(value), nil
}

// SetValue updates the stored value of the element.
//
// Example:
//
//	err := element.SetValue("active")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (rE *RecordElementModel) SetValue(value any) error {
	store := rE.GetBackingStore()
	if internal.IsNil(store) {
		return nil
	}

	return store.Set(valueKey, value)
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
func (rE *RecordElementModel) GetLink() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	store := rE.GetBackingStore()
	if internal.IsNil(store) {
		return nil, nil
	}

	value, err := store.Get(linkKey)
	if err != nil {
		return nil, err
	}

	link, ok := value.(*string)
	if !ok {
		return nil, errors.New("value is not *string")
	}

	return link, nil
}

// setLink assigns an optional reference link to the element.
//
// Example:
//
//	link := "https://example.com"
//	if err := element.setLink(&link); err != nil {
//	    log.Fatal(err)
//	}
func (rE *RecordElementModel) setLink(link *string) error {
	store := rE.GetBackingStore()
	if internal.IsNil(store) {
		return nil
	}

	return store.Set(linkKey, link)
}
