package internal

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
	return &TableValue{Value: trueVal}
}

// Set sets the provided key, value pair.
func (tE TableEntry) Set(key string, value interface{}) {
	if tE == nil {
		tE = make(TableEntry)
	}

	tE[key] = value
}

// Keys returns a slice of the TableEntry's keys
func (tE TableEntry) Keys() []string {
	keys := make([]string, 0, tE.Len())
	for k := range tE {
		keys = append(keys, k)
	}
	return keys
}

// Len returns the length of tE
func (tE TableEntry) Len() int {
	return len(tE)
}
