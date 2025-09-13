package internal

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// IsNil checks if a value is nil or a nil interface.
func IsNil(a interface{}) bool {
	defer func() { _ = recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

// ToPointer Converts provided value to pointer.
func ToPointer[T any](value T) *T {
	return &value
}

// IsPointer
func IsPointer(value any) bool {
	return reflect.ValueOf(value).Kind() == reflect.Pointer
}

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

// Dereference recursively unwraps nested pointers
func Dereference(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return reflect.Zero(v.Type().Elem())
		}
		v = v.Elem()
	}
	return v
}

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
	if IsNil(in) {
		return nil
	}

	valValue := reflect.ValueOf(in)
	for valValue.Kind() == reflect.Ptr {
		valValue = valValue.Elem()
		in = valValue.Interface()
	}

	outVal := reflect.ValueOf(out)
	if outVal.Kind() != reflect.Pointer || IsNil(out) {
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
func As2[T any](in any, out T, strict bool) error {
	// Early return if input is nil
	if IsNil(in) {
		return nil
	}

	// Validate output is a non-nil pointer
	outVal := reflect.ValueOf(out)
	if outVal.Kind() != reflect.Pointer {
		return fmt.Errorf("out must be a pointer")
	}
	outVal = outVal.Elem()

	// Unwrap pointer layers of input
	valValue := reflect.ValueOf(in)
	for valValue.Kind() == reflect.Ptr && !valValue.IsNil() {
		valValue = valValue.Elem()
	}

	// If types match, set directly
	if typedIn, ok := valValue.Interface().(T); ok {
		fmt.Println("here")
		reflect.ValueOf(out).Set(reflect.ValueOf(typedIn))
		return nil
	}

	outType := outVal.Type()
	if outType.Kind() == reflect.Pointer {
		outType = outType.Elem()
	}
	inVal := valValue.Interface()

	// Compatibility check
	if !isCompatible(inVal, outType, strict) {
		return fmt.Errorf("cannot convert '%v' to type %s", inVal, outType)
	}

	converted := reflect.ValueOf(inVal).Convert(outType)
	outVal.Elem().Set(converted)

	return nil
}

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
			return reflect.Value{}, err
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

	return reflect.Value{}, fmt.Errorf("unsupported conversion: %s → %s", srcKind, dstKind)
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

func isNumericKind(k reflect.Kind) bool {
	return (k >= reflect.Int && k <= reflect.Int64) ||
		(k >= reflect.Uint && k <= reflect.Uint64) ||
		(k == reflect.Float32 || k == reflect.Float64)
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

func convertNumeric(srcVal reflect.Value, targetType reflect.Type) (any, error) {
	srcFloat := srcVal.Convert(reflect.TypeOf(float64(0))).Float()

	targetKind := targetType.Kind()

	rng, ok := ranges[targetKind]
	if !ok {
		return nil, fmt.Errorf("unsupported numeric target type: %s", targetType.Kind())
	}

	if !rng.Compatible(srcFloat) {
		return nil, fmt.Errorf("overflow or incompatible decimal converting to %s", targetType.Kind())
	}

	// Safe to convert
	return reflect.ValueOf(srcFloat).Convert(targetType).Interface(), nil
}
