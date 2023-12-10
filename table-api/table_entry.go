package tableapi

type TableEntry map[string]interface{}

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
