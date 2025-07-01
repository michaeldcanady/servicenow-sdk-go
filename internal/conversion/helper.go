package conversion

import (
	"fmt"
	"math"
	"reflect"

	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
)

type numericRange struct {
	min          float64
	max          float64
	allowDecimal bool
}

var (
	numericTypeRanges = map[reflect.Kind]numericRange{
		reflect.Int8:    {math.MinInt8, math.MaxInt8, false},
		reflect.Uint8:   {0, math.MaxUint8, false},
		reflect.Int16:   {math.MinInt16, math.MaxInt16, false},
		reflect.Uint16:  {0, math.MaxUint16, false},
		reflect.Int32:   {math.MinInt32, math.MaxInt32, false},
		reflect.Uint32:  {0, math.MaxUint32, false},
		reflect.Int64:   {math.MinInt64, math.MaxInt64, false},
		reflect.Uint64:  {0, math.MaxUint64, false},
		reflect.Float32: {-math.MaxFloat32, math.MaxFloat32, true},
		reflect.Float64: {-math.MaxFloat64, math.MaxFloat64, true},
	}
)

// isCompatible checks if the value is compatible with the type tp.
func isCompatible(value interface{}, tp reflect.Type, strict bool) bool {
	if isNumericType(value) && isNumericType(tp) {
		return isCompatibleInt(value, tp)
	}

	if strict {
		return reflect.TypeOf(value) == tp
	}
	return reflect.TypeOf(value).ConvertibleTo(tp)
}

// As converts the value to the type T.
func As[T any](in interface{}, out T) error {
	if internal.IsNil(in) {
		return nil
	}

	valValue := reflect.ValueOf(in)
	for valValue.Kind() == reflect.Ptr {
		valValue = valValue.Elem()
		in = valValue.Interface()
	}

	outVal := reflect.ValueOf(out)
	if outVal.Kind() != reflect.Pointer || internal.IsNil(out) {
		return fmt.Errorf("out is not pointer or is nil")
	}

	nestedOutVal := outVal.Elem()
	if nestedOutVal.Kind() == reflect.Interface && !nestedOutVal.IsNil() {
		nestedOutVal = nestedOutVal.Elem()
	}

	outType := nestedOutVal.Type()

	if !isCompatible(in, outType, true) {
		return fmt.Errorf("value '%v' is not compatible with type %T", in, nestedOutVal.Interface())
	}

	outVal.Elem().Set(valValue.Convert(outType))
	return nil
}

// As2 converts the input value to the specified type T and assigns it to out if compatible.
// It supports strict and non-strict mode.
func As2[T any](in interface{}, out T, strict bool) error {
	// Early return if input is nil
	if internal.IsNil(in) {
		return nil
	}

	// Validate output is a non-nil pointer
	outVal := reflect.ValueOf(out)
	if outVal.Kind() != reflect.Pointer || internal.IsNil(out) {
		return fmt.Errorf("out must be a non-nil pointer")
	}

	// If types match, set directly
	if typedIn, ok := in.(T); ok {
		reflect.ValueOf(out).Elem().Set(reflect.ValueOf(typedIn))
		return nil
	}

	// Unwrap pointer layers of input
	valValue := reflect.ValueOf(in)
	for valValue.Kind() == reflect.Ptr && !valValue.IsNil() {
		valValue = valValue.Elem()
	}

	// Get the concrete type behind the output pointer
	nestedOutVal := outVal.Elem()
	if nestedOutVal.Kind() == reflect.Interface && !nestedOutVal.IsNil() {
		nestedOutVal = nestedOutVal.Elem()
	}

	outType := nestedOutVal.Type()
	inVal := valValue.Interface()

	// Compatibility check
	if !isCompatible(inVal, outType, strict) {
		return fmt.Errorf("cannot convert '%v' to type %s", inVal, outType)
	}

	converted := reflect.ValueOf(inVal).Convert(outType)
	outVal.Elem().Set(converted)

	return nil
}

// isNumericType checks if the given type is a numeric type.
func isNumericType(in interface{}) bool {
	if in == nil {
		return false
	}

	tp, ok := in.(reflect.Type)
	if !ok {
		tp = reflect.TypeOf(in)
	}

	switch tp.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// isCompatibleInt checks if the given value is compatible with the specified integer type.
func isCompatibleInt(in interface{}, tp reflect.Type) bool {
	if !isNumericType(in) || !isNumericType(tp) {
		return false
	}

	inFloat := reflect.ValueOf(in).Convert(reflect.TypeOf(float64(0))).Float()
	hasDecimal := hasDecimalPlace(inFloat)

	if rangeInfo, ok := numericTypeRanges[tp.Kind()]; ok {
		if inFloat >= rangeInfo.min && inFloat <= rangeInfo.max {
			return rangeInfo.allowDecimal || !hasDecimal
		}
	}
	return false
}

// hasDecimalPlace checks if the given float64 value has a decimal place.
func hasDecimalPlace(value float64) bool {
	return value != float64(int64(value))
}
