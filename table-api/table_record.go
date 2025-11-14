package tableapi

import (
	"errors"
	"fmt"
	"slices"

	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

var _ serialization.Parsable = (*TableRecord)(nil)

// TableRecord defines an interface for managing structured records in a table.
//
// This interface provides methods for retrieving, updating, and checking the existence
// of attributes stored within a table entry.
//
// Implementing types should ensure proper storage and retrieval of RecordElement instances.
//
// Example usage:
//
//	var record TableRecord
//	element := record.Get("status")
//	fmt.Println(element)
type TableRecord struct {
	keys []string
	internal.Model
}

func CreateTableRecordFromDiscriminatorValue(node serialization.ParseNode) (serialization.Parsable, error) {
	value, err := node.GetRawValue()
	if err != nil {
		return nil, err
	}

	if typedValue, ok := value.(map[string]any); ok {
		record := NewTableRecord()
		for key := range typedValue {
			record.keys = append(record.keys, key)
		}
		return record, nil
	}
	return nil, fmt.Errorf("unsupported type %T", value)
}

func recordElementParser(node serialization.ParseNode) error {
	rawValue, err := node.GetRawValue()
	if err != nil {
		return err
	}
	var (
		displayValue any
		link         *string
		value        any
	)

	switch typedVal := rawValue.(type) {
	case map[string]any:
		if dv, ok := typedVal[displayValueKey]; ok {
			displayValue = dv
		}

		if v, ok := typedVal[valueKey]; ok {
			value = v
		}

		if rawLink, ok := typedVal[linkKey]; ok {
			strLink, ok := rawLink.(*string)
			if !ok {
				return errors.New("link is not *string")
			}
			link = strLink
		}
	case any:
		value = typedVal
	default:
		return errors.New("value is one of expected types")
	}

	elem := NewRecordElement()

	if err := elem.SetDisplayValue(displayValue); err != nil {
		return err
	}

	if err := elem.SetValue(value); err != nil {
		return err
	}

	if err := elem.SetLink(link); err != nil {
		return err
	}

	return nil
}

// GetFieldDeserializers implements serialization.Parsable.
func (tR *TableRecord) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	fieldDeserializers := map[string]func(serialization.ParseNode) error{}

	for _, key := range tR.keys {
		fieldDeserializers[key] = recordElementParser
	}

	return fieldDeserializers
}

// Serialize implements serialization.Parsable.
func (tR *TableRecord) Serialize(writer serialization.SerializationWriter) error {
	return errors.New("unimplemented")
}

func NewTableRecord() *TableRecord {
	return &TableRecord{
		keys:  make([]string, 0),
		Model: internal.NewBaseModel(),
	}
}

// Get retrieves a RecordElement associated with the specified key.
//
// This method returns the stored element for the given field name.
//
// Example:
//
//	element := record.Get("status")
//	fmt.Println(element)
func (tR *TableRecord) Get(key string) (*RecordElement, error) {
	elem, err := store.DefaultBackedModelAccessorFunc[*TableRecord, RecordElement](tR, key)

	return &elem, err
}

// SetElement assigns a RecordElement to the specified key.
//
// Example:
//
//	element := NewRecordElement()
//	err := record.SetElement("status", element)
func (tR *TableRecord) SetElement(key string, element *RecordElement) error {
	return store.DefaultBackedModelMutatorFunc(tR, key, element)
}

// SetValue assigns a value to the specified key using a RecordElement wrapper.
//
// This method ensures that the stored data conforms to the RecordElement interface.
//
// Example:
//
//	err := record.SetValue("status", "active")
func (tR *TableRecord) SetValue(key string, value any) error {
	elem := NewRecordElement()
	if err := elem.SetValue(value); err != nil {
		return err
	}
	return tR.SetElement(key, elem)
}

// HasAttribute checks whether the specified key exists in the record.
//
// Example:
//
//	exists := record.HasAttribute("status")
//	fmt.Println(exists) // Output: true or false
func (tR *TableRecord) HasAttribute(key string) bool {
	return slices.Contains(tR.keys, key)
}
