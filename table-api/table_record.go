//go:build preview.tableApiV2

package tableapi

import (
	"errors"
	"fmt"
	"slices"

	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
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
	*internal.BaseModel
}

func CreateTableRecordFromDiscriminatorValue(node serialization.ParseNode) (serialization.Parsable, error) {
	record := NewTableRecord()

	value, err := node.GetRawValue()
	if err != nil {
		return nil, err
	}

	typedValue, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("unsupported type %T", value)
	}

	for key, _ := range typedValue {
		record.keys = append(record.keys, key)
	}

	return record, nil
}

// GetFieldDeserializers implements serialization.Parsable.
func (tR *TableRecord) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	fieldDeserializers := map[string]func(serialization.ParseNode) error{}

	for _, key := range tR.keys {
		fieldDeserializers[key] = internal.DeserializeMutatedValueFunc(func(value *RecordElement) error {
			return tR.SetElement(key, value)
		}, func(a any) (*RecordElement, error) {
			// TODO: improve
			elem := NewRecordElement()

			value, ok := a.(map[string]any)
			if !ok {
				if err := elem.SetValue(value); err != nil {
					return nil, err
				}
				return elem, nil
			}
			if link, ok := value[linkKey]; ok {
				typedLink, ok := link.(*string)
				if !ok {
					return nil, errors.New("link is not string")
				}
				if err := elem.setLink(typedLink); err != nil {
					return nil, err
				}
			}
			if displayValue, ok := value[displayValueKey]; ok {
				if err := elem.SetDisplayValue(displayValue); err != nil {
					return nil, err
				}
			}
			if realValue, ok := value[valueKey]; ok {
				if err := elem.SetValue(realValue); err != nil {
					return nil, err
				}
			}
			return elem, nil
		})
	}

	return fieldDeserializers
}

// Serialize implements serialization.Parsable.
func (tR *TableRecord) Serialize(writer serialization.SerializationWriter) error {
	return errors.New("unimplemented")
}

func NewTableRecord() *TableRecord {
	return &TableRecord{
		keys:      make([]string, 0),
		BaseModel: internal.NewBaseModel(),
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
	return internal.DefaultBackedModelAccessorFunc[*TableRecord, *RecordElement](tR, key)
}

// SetElement assigns a RecordElement to the specified key.
//
// Example:
//
//	element := NewRecordElement()
//	err := record.SetElement("status", element)
func (tR *TableRecord) SetElement(key string, element *RecordElement) error {
	return internal.DefaultBackedModelMutatorFunc(tR, key, element)
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
