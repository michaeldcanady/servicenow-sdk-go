package tableapi

import "encoding/json"

// NewTableEntry creates a new table entry instance.
func NewTableEntry2() *TableEntry2 {
	return &TableEntry2{}
}

// TableEntry2 represents a single Service-Now Table Entry.
type TableEntry2 struct {
	value map[string]TableValue2
}

func (tE *TableEntry2) UnmarshalJSON(data []byte) error {

	var temp map[string]TableValue2

	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}

	tE.value = temp

	return nil
}

func (tE *TableEntry2) MarshalJSON() ([]byte, error) {
	return json.Marshal(tE.value)
}

// Deprecated: deprecated since v{version}. It was added to match the requirements but has no use.
// Use `Value(key).[Link|Value|DisplayValue] instead.`
//
// Set sets the specified key to the provided value.
func (tE *TableEntry2) Set(key string, value interface{}) {
	// It was added to match the requirements but has no use.
}

// Value returns a tE if a valid key is provided.
func (tE *TableEntry2) Value(key string) *TableValue2 {
	value, exists := tE.value[key]
	if !exists {
		return nil
	}
	return &value
}

// Keys returns a slice of the tE's keys.
func (tE *TableEntry2) Keys() []string {
	keys := make([]string, 0, tE.Len())
	for k := range tE.value {
		keys = append(keys, k)
	}
	return keys
}

// Len returns the length of the tE.
func (tE *TableEntry2) Len() int {

	if tE == nil {
		return 0
	}
	return len((*tE).value)
}
