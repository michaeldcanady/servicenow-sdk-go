package internal

import "fmt"

func ConvertType[T any](val interface{}) (T, error) {
	v, ok := val.(T)
	if !ok {
		return v, fmt.Errorf("value (%v) cannot be converted to %T", val, v)
	}
	return v, nil
}
