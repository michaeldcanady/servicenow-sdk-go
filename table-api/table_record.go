package tableapi

import (
	"fmt"

	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

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
type TableRecord interface {
	// Get retrieves a RecordElement associated with the specified key.
	//
	// This method returns the stored element for the given field name.
	//
	// Example:
	//
	//      element := record.Get("status")
	//      fmt.Println(element)
	Get(string) RecordElement
	// SetElement assigns a RecordElement to the specified key.
	//
	// Example:
	//
	//      element := NewRecordElement()
	//      err := record.SetElement("status", element)
	SetElement(string, RecordElement) error
	// SetValue assigns a value to the specified key using a RecordElement wrapper.
	//
	// This method ensures that the stored data conforms to the RecordElement interface.
	//
	// Example:
	//
	//      err := record.SetValue("status", "active")
	SetValue(string, any) error
	// HasAttribute checks whether the specified key exists in the record.
	//
	// Example:
	//
	//      exists := record.HasAttribute("status")
	//      fmt.Println(exists) // Output: true or false
	HasAttribute(string) bool
}

type TableRecordModel struct {
	internal.Model
}

// GetFieldDeserializers implements serialization.Parsable.
func (tR *TableRecordModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {

	fieldDeserializers := map[string]func(serialization.ParseNode) error{}

	for key := range tR.GetBackingStore().Enumerate() {
		fieldDeserializers[key] = func(pn serialization.ParseNode) error {
			panic("not implemented")
		}
	}

	return fieldDeserializers
}

// Serialize implements serialization.Parsable.
func (tR *TableRecordModel) Serialize(writer serialization.SerializationWriter) error {
	panic("unimplemented")
}

func NewTableRecord() *TableRecordModel {
	return &TableRecordModel{internal.NewBaseModel()}
}

func (tR *TableRecordModel) Get(key string) (*RecordElementModel, error) {
	panic("unimplemented")
}

func (tR *TableRecordModel) Set(key string, element RecordElement) error {
	panic("unimplemented")
}

func (tR *TableRecordModel) SetValue(key string, value any) error {
	panic("unimplemented")
}

func (tR *TableRecordModel) HasAttribute(key string) bool {
	panic("unimplemented")
}

func NewTableRecordFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	record := NewTableRecord()

	rawValue, err := parseNode.GetRawValue()
	if err != nil {
		return nil, err
	}

	value, ok := rawValue.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("value is not of type %T", value)
	}

	for k := range value {
		if err := record.SetValue(k, nil); err != nil {
			return nil, err
		}
	}

	return record, nil
}
