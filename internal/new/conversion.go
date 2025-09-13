package internal

import (
	"reflect"
	"slices"
)

var (
	numericKinds = []reflect.Kind{
		reflect.Int8,
		reflect.Uint8,
		reflect.Int16,
		reflect.Uint16,
		reflect.Int32,
		reflect.Uint32,
		reflect.Int64,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
	}
)

// Dereference recursively unwraps nested pointers.
func Dereference(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return reflect.Zero(v.Type().Elem())
		}
		v = v.Elem()
	}
	return v
}

// isNumericKind checks if value is a numeric value.
func isNumericKind(v reflect.Value) bool {
	return slices.Contains(numericKinds, v.Kind())
}
