package tableapi

// TableEntry represents a single Service-Now Table Entry.
type TableEntry map[string]interface{}

// Value returns a TableValue if a valid key is provided.
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

// Keys returns a slice of the TableEntry's keys
func (tE TableEntry) Keys() []string {
	keys := make([]string, 0, len(tE))
	for k := range tE {
		keys = append(keys, k)
	}
	return keys
}
