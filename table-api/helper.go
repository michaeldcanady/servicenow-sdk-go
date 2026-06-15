package tableapi

import (
	"fmt"
	"reflect"
)

func convertType[T any](val interface{}) (T, error) {
	v, ok := val.(T)
	if !ok {
		return v, fmt.Errorf("value (%v) cannot be converted to %T", val, v)
	}
	return v, nil
}

// isNil checks if a value is nil or a nil interface
func isNil(a interface{}) bool {
	defer func() { _ = recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}
