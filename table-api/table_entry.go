package tableapi

import "errors"

const (
	displayValueKey = "display_value"
	valueKey        = "value"
	linkKey         = "link"
)

// TableEntry represents a single Service-Now table entry.
type TableEntry map[string]interface{}

// NewTableEntry creates a new instance of TableEntry.
func NewTableEntry() TableEntry {
	return TableEntry{}
}

// Set assigns the specified key to the given value in the table entry.
// Deprecated: deprecated since v{unreleased}. Please use SetElement or SetValue instead.
//
// Set assigns the specified key to the given value in the table entry.
func (tE TableEntry) Set(key string, value interface{}) {
	tE[key] = value
}

// Value retrieves a TableValue instance for the given key.
func (tE TableEntry) Value(key string) *TableValue {
	value, exists := tE[key]
	if !exists {
		return nil
	}

	var trueVal interface{}

	switch v := value.(type) {
	case map[string]interface{}:
		trueVal = v[valueKey]
	case interface{}:
		trueVal = v
	}
	return &TableValue{value: trueVal}
}

// DisplayValue retrieves a TableValue instance for the given key.
func (tE TableEntry) DisplayValue(key string) *TableValue {
	value, exists := tE[key]
	if !exists {
		return nil
	}

	var trueVal interface{}

	switch v := value.(type) {
	case map[string]interface{}:
		trueVal = v[displayValueKey]
		trueVal = v[valueKey]
	case interface{}:
		trueVal = v
	}
	return &TableValue{value: trueVal}
}

// DisplayValue retrieves a TableValue instance for the given key.
func (tE TableEntry) DisplayValue(key string) *TableValue {
	value, exists := tE[key]
	if !exists {
		return nil
	}

	var trueVal interface{}

	switch v := value.(type) {
	case map[string]interface{}:
		trueVal = v[displayValueKey]
	case interface{}:
		trueVal = v
	}
	return &TableValue{value: trueVal}
}

// Link retrieves a String instance for the given key.
func (tE TableEntry) Link(key string) (*string, error) {
	value, exists := tE[key]
	if !exists {
		return nil, nil
	}

	var trueVal string

	switch v := value.(type) {
	case map[string]interface{}:
		val, ok := v[linkKey]
		if !ok {
			return nil, nil
		}
		trueVal, ok = val.(string)
		if !ok {
			return nil, errors.New("link is not string")
		}
	default:
		return nil, nil
	}
	return &trueVal, nil
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
// Len returns the number of fields in the table entry.
func (tE TableEntry) Len() int {
	return len(tE)
}
