package internal

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"sync"
)

var (
	numericKinds = sync.OnceValue(func() []reflect.Kind {
		keys := make([]reflect.Kind, len(ranges))
		i := 0
		for key := range ranges {
			keys[i] = key
			i++
		}
		return keys
	})()
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
func isNumericKind(v reflect.Kind) bool {
	return slices.Contains(numericKinds, v)
}

// convertNumeric converts value to desired type
func convertNumeric(srcVal reflect.Value, targetType reflect.Type) (any, error) {
	if !isNumericKind(srcVal.Kind()) {
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

// convertValue converts the provide value to the target type.
func convertValue(srcVal reflect.Value, targetType reflect.Type) (reflect.Value, error) {
	srcKind := srcVal.Kind()
	dstKind := targetType.Kind()

	// Direct assignable
	if srcVal.Type().AssignableTo(targetType) {
		return srcVal, nil
	}

	// Numeric conversions
	if isNumericKind(srcKind) && isNumericKind(dstKind) {
		converted, err := convertNumeric(srcVal, targetType)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(converted).Convert(targetType), nil
	}

	// String to numeric
	if srcKind == reflect.String && isNumericKind(dstKind) {
		num, err := strconv.ParseFloat(srcVal.String(), 64)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("unable to convert %s to float: %s", srcVal, err)
		}
		converted, err := convertNumeric(reflect.ValueOf(num), targetType)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(converted).Convert(targetType), nil
	}

	// Numeric to string
	if isNumericKind(srcKind) && dstKind == reflect.String {
		return reflect.ValueOf(fmt.Sprintf("%v", srcVal.Interface())), nil
	}

	return reflect.Value{}, fmt.Errorf("unsupported conversion: %s to %s", srcKind, dstKind)
}
