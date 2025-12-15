package conversion

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"sync"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
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

// Convert converts provided input to and sets output
func Convert(input any, output any) error {
	if output == nil {
		return errors.New("output cannot be nil")
	}

	outVal := reflect.ValueOf(output)
	if outVal.Kind() != reflect.Ptr || outVal.IsNil() {
		return errors.New("output must be a non-nil pointer")
	}

	targetVal := outVal.Elem()
	targetType := targetVal.Type()

	if input == nil {
		targetVal.Set(reflect.Zero(targetType))
		return nil
	}

	srcVal := reflect.ValueOf(input)

	// If input is a pointer, deref for conversion but keep original for pointer assignment
	derefSrc := Dereference(srcVal)

	// Case 1: Direct assignable (including pointer-to-pointer)
	if srcVal.Type().AssignableTo(targetType) {
		targetVal.Set(srcVal)
		return nil
	}

	// Case 2: Target is a pointer type
	if targetType.Kind() == reflect.Ptr {
		elemType := targetType.Elem()

		// If input is already pointer to correct type
		if srcVal.Type().AssignableTo(targetType) {
			targetVal.Set(srcVal)
			return nil
		}

		// Convert underlying value to element type
		convertedVal, err := convertValue(derefSrc, elemType)
		if err != nil {
			return err
		}

		ptr := reflect.New(elemType)
		ptr.Elem().Set(convertedVal)
		targetVal.Set(ptr)
		return nil
	}

	// Case 3: Non-pointer target
	convertedVal, err := convertValue(derefSrc, targetType)
	if err != nil {
		return err
	}
	targetVal.Set(convertedVal)
	return nil
}

func StringToTime(format string) serialization.Mutator[string, time.Time] {
	return func(input string) (time.Time, error) {
		dateTime, err := time.Parse(format, input)
		if err != nil {
			return time.Time{}, err
		}

		return dateTime, nil
	}
}
