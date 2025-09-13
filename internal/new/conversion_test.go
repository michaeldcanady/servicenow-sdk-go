package internal

import (
	"errors"
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDereference(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "String pointer",
			test: func(t *testing.T) {
				input := ToPointer("testing")
				output := Dereference(reflect.ValueOf(input))

				assert.Equal(t, interface{}("testing"), output.Interface())
			},
		},
		{
			name: "String pointer pointer",
			test: func(t *testing.T) {
				input := ToPointer(ToPointer("testing"))
				output := Dereference(reflect.ValueOf(input))

				assert.Equal(t, interface{}("testing"), output.Interface())
			},
		},
		{
			name: "String",
			test: func(t *testing.T) {
				input := "testing"
				output := Dereference(reflect.ValueOf(input))

				assert.Equal(t, interface{}("testing"), output.Interface())
			},
		},
		{
			name: "nil",
			test: func(t *testing.T) {
				input := (*string)(nil)
				output := Dereference(reflect.ValueOf(input))

				assert.Equal(t, "", output.Interface())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestIsNumericKind(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Int8",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(int8(1)))

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Uint8",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(uint8(1)))

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Int16",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(int16(1)))

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Uint16",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(uint16(1)))

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Int32",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(int32(1)))

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Uint32",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(uint32(1)))

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Int64",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(int64(1)))

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Uint64",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(uint64(1)))

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Float32",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(float32(1.00)))

				assert.True(t, isNumeric)
			},
		},
		{
			name: "Float64",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf(float64(1.00)))

				assert.True(t, isNumeric)
			},
		},
		{
			name: "String",
			test: func(t *testing.T) {
				isNumeric := isNumericKind(reflect.ValueOf("string"))

				assert.False(t, isNumeric)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestConvertNumeric(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Int8 to Int8",
			test: func(t *testing.T) {
				input := reflect.ValueOf(int8(0))
				desiredType := reflect.TypeOf(int8(0))
				result, err := convertNumeric(input, desiredType)

				assert.Nil(t, err)
				assert.Equal(t, int8(0), result)
			},
		},
		{
			name: "String to Int8",
			test: func(t *testing.T) {
				input := reflect.ValueOf("string")
				desiredType := reflect.TypeOf(int8(0))
				result, err := convertNumeric(input, desiredType)

				assert.Equal(t, errors.New("string is non-numeric"), err)
				assert.Nil(t, result)
			},
		},
		{
			name: "Int16 to Int8",
			test: func(t *testing.T) {
				input := reflect.ValueOf(math.MaxInt16)
				desiredType := reflect.TypeOf(int8(0))
				result, err := convertNumeric(input, desiredType)

				assert.Equal(t, errors.New("overflow or incompatible decimal converting to int8"), err)
				assert.Nil(t, result)
			},
		},
		{
			name: "Int16 to Int",
			test: func(t *testing.T) {
				input := reflect.ValueOf(math.MaxInt16)
				desiredType := reflect.TypeOf(int(0))
				result, err := convertNumeric(input, desiredType)

				assert.Nil(t, err)
				assert.Equal(t, int(32767), result)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
