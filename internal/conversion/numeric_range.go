package conversion

import (
	"math"
	"reflect"
	"sync"
)

var (
	int8Range  = sync.OnceValue(func() *numericRange { return newNumericRange(math.MaxInt8, math.MinInt8, false) })()
	int16Range = sync.OnceValue(func() *numericRange { return newNumericRange(math.MaxInt16, math.MinInt16, false) })()
	int32Range = sync.OnceValue(func() *numericRange { return newNumericRange(math.MaxInt32, math.MinInt32, false) })()
	int64Range = sync.OnceValue(func() *numericRange { return newNumericRange(math.MaxInt64, math.MinInt64, false) })()
	intRange   = sync.OnceValue(func() *numericRange { return newNumericRange(math.MaxInt64, math.MinInt64, false) })()

	float32Range = sync.OnceValue(func() *numericRange { return newNumericRange(math.MaxFloat32, -math.MaxFloat32, true) })()
	float64Range = sync.OnceValue(func() *numericRange { return newNumericRange(math.MaxFloat64, -math.MaxFloat64, true) })()

	uint8Range  = sync.OnceValue(func() *numericRange { return newNumericRange(math.MaxUint8, 0, false) })()
	uint16Range = sync.OnceValue(func() *numericRange { return newNumericRange(math.MaxUint16, 0, false) })()
	uint32Range = sync.OnceValue(func() *numericRange { return newNumericRange(math.MaxUint32, 0, false) })()
	uint64Range = sync.OnceValue(func() *numericRange { return newNumericRange(math.MaxUint64, 0, false) })()
	uintRange   = sync.OnceValue(func() *numericRange { return newNumericRange(math.MaxUint64, 0, false) })()

	ranges = map[reflect.Kind]*numericRange{
		reflect.Int8:  int8Range,
		reflect.Int16: int16Range,
		reflect.Int32: int32Range,
		reflect.Int64: int64Range,
		reflect.Int:   intRange,

		reflect.Uint8:  uint8Range,
		reflect.Uint16: uint16Range,
		reflect.Uint32: uint32Range,
		reflect.Uint64: uint64Range,
		reflect.Uint:   uintRange,

		reflect.Float32: float32Range,
		reflect.Float64: float64Range,
	}
)

func newNumericRange(max, min float64, allowDecimal bool) *numericRange {
	return &numericRange{
		min:          min,
		max:          max,
		allowDecimal: allowDecimal,
	}
}

func (nR *numericRange) Within(value float64) bool {
	return value >= nR.min && value <= nR.max
}

func (nR *numericRange) Compatible(value float64) bool {
	if !nR.Within(value) {
		return false
	}
	if hasDecimalPlace(value) {
		return nR.allowDecimal
	}
	return true
}
