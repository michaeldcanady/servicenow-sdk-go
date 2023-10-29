package tableapi

type TableEntry map[string]interface{}

func (tE TableEntry) Value(key string) *TableValue {
	value, exists := tE[key]
	if !exists {
		return nil
	}

	var _v interface{}

	switch v := value.(type) {
	case map[string]interface{}:
		_v = v["value"]
	case interface{}:
		_v = v
	}
	return &TableValue{value: _v}
}
