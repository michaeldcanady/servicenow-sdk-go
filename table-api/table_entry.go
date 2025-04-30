package tableapi

import "fmt"

// TableEntry represents a single Service-Now table entry.
type TableEntry map[string]interface{}

// NewTableEntry creates a new instance of TableEntry.
func NewTableEntry() TableEntry {
	return TableEntry{}
}

// Deprecated: deprecated since v{unreleased}. Please use SetElement or SetValue instead.
//
// Set assigns the specified key to the given value in the table entry.
func (tE TableEntry) Set(key string, value interface{}) {
	tE[key] = value
}

// Deprecated: deprecated since v{unreleased}. Please use Get instead.
//
// Value retrieves a TableValue instance for the given key.
func (tE TableEntry) Value(key string) *TableValue {
	value, exists := tE[key]
	if !exists {
		return nil
	}

	var trueVal interface{}

	switch v := value.(type) {
	case map[string]interface{}:
		trueVal = v["value"]
	case interface{}:
		trueVal = v
	}
	return &TableValue{value: trueVal}
}

// Keys returns a slice of all field names stored in the table entry.
func (tE TableEntry) Keys() []string {
	keys := make([]string, 0, tE.Len())
	for k := range tE {
		keys = append(keys, k)
	}
	return keys
}

// Len returns the number of fields in the table entry.
func (tE TableEntry) Len() int {
	return len(tE)
}

// Get retrieves a RecordElementModel from the table entry based on the provided key.
//
// If the key is found, it processes the stored value and extracts necessary
// metadata such as display value, value, and links.
func (tE TableEntry) Get(key string) (*RecordElementModel, error) {
	value, ok := tE[key]
	if !ok {
		return nil, nil
	}

	model := NewRecordElement()

	switch v := value.(type) {
	case map[string]interface{}:
		if err := model.SetDisplayValue(v[displayValueKey]); err != nil {
			return nil, err
		}
		if err := model.SetValue(v[valueKey]); err != nil {
			return nil, err
		}
		if val, ok := v[linkKey]; ok {
			if link, ok := val.(*string); ok {
				if err := model.setLink(link); err != nil {
					return nil, err
				}
			}
		}
	case interface{}:
		if err := model.SetValue(v); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported type: %T", v)
	}
	return model, nil
}

// SetElement assigns a RecordElement to the specified key.
func (tE TableEntry) SetElement(key string, element RecordElement) error {
	tE[key] = element

	return nil
}

// SetValue assigns a simple value to the specified key using a RecordElement wrapper.
func (tE TableEntry) SetValue(key string, value any) error {
	model := NewRecordElement()
	if err := model.SetValue(value); err != nil {
		return err
	}
	return tE.SetElement(key, model)
}

// HasAttribute checks if the specified key exists in the table entry.
func (tE TableEntry) HasAttribute(key string) bool {
	_, ok := tE[key]
	return ok
}
