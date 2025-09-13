package internal

import (
	"fmt"
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
		reflect.Int,
		reflect.Float32,
		reflect.Float64,
	}
	floatType = reflect.TypeOf(float64(0))
)

//TODO: add converter type?

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

// convertNumeric converts value to desired type
func convertNumeric(srcVal reflect.Value, targetType reflect.Type) (any, error) {
	if !isNumericKind(srcVal) {
		return nil, fmt.Errorf("%s is non-numeric", srcVal)
	}

	srcFloat := srcVal.Convert(floatType).Float()

	targetKind := targetType.Kind()

	rng, ok := ranges[targetKind]
	if !ok {
		return nil, fmt.Errorf("unsupported numeric target type: %s", targetType.Kind())
	}

	if !rng.Compatible(srcFloat) {
		return nil, fmt.Errorf("overflow or incompatible decimal converting to %s", targetType.Kind())
	}

	return reflect.ValueOf(srcFloat).Convert(targetType).Interface(), nil
}
