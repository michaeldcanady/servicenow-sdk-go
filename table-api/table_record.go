package tableapi

import (
	"errors"
	"fmt"
	"slices"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

var _ serialization.Parsable = (*TableRecord)(nil)

// TableRecord represents a structured record in a Service-Now table.
//
// This model provides a flexible way to store and retrieve attributes from a table entry
// using a backing store. It supports both raw values and structured RecordElement instances.
type TableRecord struct {
	keys []string
	internal.Model
}

// CreateTableRecordFromDiscriminatorValue creates a new TableRecord from a ParseNode.
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

func recordElementParser(node serialization.ParseNode) (*RecordElement, error) {
	rawValue, err := node.GetRawValue()
	if err != nil {
		return nil, err
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
				return nil, errors.New("link is not *string")
			}
			link = strLink
		}
	case any:
		value = typedVal
	default:
		return nil, errors.New("value is one of expected types")
	}

	elem := NewRecordElement()

	if err := elem.SetDisplayValue(displayValue); err != nil {
		return nil, err
	}

	if err := elem.SetValue(value); err != nil {
		return nil, err
	}

	if err := elem.SetLink(link); err != nil {
		return nil, err
	}

	return elem, nil
}

// GetSysID returns the sys_id of the record if it exists.
func (tR *TableRecord) GetSysID() (*string, error) {
	element, err := tR.Get("sys_id")
	if err != nil {
		return nil, err
	}

	value, err := element.GetValue()
	if err != nil {
		return nil, err
	}

	return value.GetStringValue()
}

// GetFieldDeserializers implements serialization.Parsable.
func (tR *TableRecord) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	fieldDeserializers := map[string]func(serialization.ParseNode) error{}

	for _, key := range tR.keys {
		fieldDeserializers[key] = func(node serialization.ParseNode) error {
			element, err := recordElementParser(node)
			if err != nil {
				return err
			}
			return tR.SetElement(key, element)
		}
	}

	return fieldDeserializers
}

// Serialize implements serialization.Parsable.
func (tR *TableRecord) Serialize(writer serialization.SerializationWriter) error {
	return errors.New("unimplemented")
}

// NewTableRecord creates a new instance of TableRecord.
func NewTableRecord() *TableRecord {
	return &TableRecord{
		keys:  make([]string, 0),
		Model: model.NewBaseModel(),
	}
}

// Get retrieves a RecordElement associated with the specified key.
func (tR *TableRecord) Get(key string) (*RecordElement, error) {
	elem, err := store.DefaultBackedModelAccessorFunc[*TableRecord, RecordElement](tR, key)

	return &elem, err
}

// SetElement assigns a RecordElement to the specified key.
func (tR *TableRecord) SetElement(key string, element *RecordElement) error {
	return store.DefaultBackedModelMutatorFunc(tR, key, element)
}

// SetValue assigns a value to the specified key using a RecordElement wrapper.
func (tR *TableRecord) SetValue(key string, value any) error {
	elem := NewRecordElement()
	if err := elem.SetValue(value); err != nil {
		return err
	}
	return tR.SetElement(key, elem)
}

// HasAttribute checks whether the specified key exists in the record.
func (tR *TableRecord) HasAttribute(key string) bool {
	return slices.Contains(tR.keys, key)
}
