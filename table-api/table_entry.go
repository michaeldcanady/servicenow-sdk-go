package tableapi

// NewTableEntry creates a new table entry instance.
func NewTableEntry() TableEntry {
	return TableEntry{}
}

// TableEntry represents a single Service-Now Table Entry.
type TableEntry map[string]interface{}

// Set sets the specified key to the provided value.
func (tE TableEntry) Set(key string, value interface{}) {
	tE[key] = value
}

// Value returns a tE if a valid key is provided.
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

// Keys returns a slice of the tE's keys.
func (tE TableEntry) Keys() []string {
	keys := make([]string, 0, tE.Len())
	for k := range tE {
		keys = append(keys, k)
	}
	return keys
}

// Len returns the length of the tE.
func (tE TableEntry) Len() int {
	return len(tE)
}
