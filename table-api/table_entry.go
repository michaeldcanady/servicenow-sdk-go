package tableapi

type tableEntry map[string]interface{}

func NewTableEntry() tableEntry {
	return tableEntry{}
}

type TableEntry2 interface {
	Value(string) *TableValue
	Set(string, interface{})
	Keys() []string
	Len() int
}

// Deprecated: deprecated since v{version}.
// TableEntry represents a single Service-Now Table Entry.
type TableEntry = tableEntry

// Value returns a TableValue if a valid key is provided.
func (tE tableEntry) Value(key string) *TableValue {
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

// Set sets the provided key, value pair.
func (tE tableEntry) Set(key string, value interface{}) {

	if tE == nil {
		tE = make(tableEntry)
	}

	tE[key] = value
}

// Keys returns a slice of the TableEntry's keys
func (tE tableEntry) Keys() []string {
	keys := make([]string, 0, len(tE))
	for k := range tE {
		keys = append(keys, k)
	}
	return keys
}

// Len returns the length of tE
func (tE tableEntry) Len() int {
	return len(tE)
}
