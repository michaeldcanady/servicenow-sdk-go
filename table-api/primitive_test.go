package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimitive_String(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "PrimitiveUnknown",
			test: func(t *testing.T) {
				assert.Equal(t, PrimitiveUnknown.String(), "unknown")
			},
		},
		{
			title: "PrimitiveBool",
			test: func(t *testing.T) {
				assert.Equal(t, PrimitiveBool.String(), "bool")
			},
		},
		{
			title: "PrimitiveInt8",
			test: func(t *testing.T) {
				assert.Equal(t, PrimitiveInt8.String(), "int8")
			},
		},
		{
			title: "PrimitiveInt32",
			test: func(t *testing.T) {
				assert.Equal(t, PrimitiveInt32.String(), "int32")
			},
		},
		{
			title: "PrimitiveInt64",
			test: func(t *testing.T) {
				assert.Equal(t, PrimitiveInt64.String(), "int64")
			},
		},
		{
			title: "PrimitiveTime",
			test: func(t *testing.T) {
				assert.Equal(t, PrimitiveTime.String(), "time")
			},
		},
		{
			title: "PrimitiveByte",
			test: func(t *testing.T) {
				assert.Equal(t, PrimitiveByte.String(), "byte")
			},
		},
		{
			title: "PrimitiveFloat32",
			test: func(t *testing.T) {
				assert.Equal(t, PrimitiveFloat32.String(), "float32")
			},
		},
		{
			title: "PrimitiveFloat64",
			test: func(t *testing.T) {
				assert.Equal(t, PrimitiveFloat64.String(), "float64")
			},
		},
		{
			title: "PrimitiveTimeOnly",
			test: func(t *testing.T) {
				assert.Equal(t, PrimitiveTimeOnly.String(), "timeonly")
			},
		},
		{
			title: "PrimitiveDateOnly",
			test: func(t *testing.T) {
				assert.Equal(t, PrimitiveDateOnly.String(), "dateonly")
			},
		},
		{
			title: "PrimitiveString",
			test: func(t *testing.T) {
				assert.Equal(t, PrimitiveString.String(), "string")
			},
		},
		{
			title: "unknown primitive",
			test: func(t *testing.T) {
				assert.Equal(t, Primitive(100).String(), "unknown")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

func TestParsePrimitive(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "PrimitiveUnknown",
			test: func(t *testing.T) {
				input := "unknown"

				primitive := ParsePrimitive(input)

				assert.Equal(t, PrimitiveUnknown, primitive)
			},
		},
		{
			title: "PrimitiveBool",
			test: func(t *testing.T) {
				input := "bool"

				primitive := ParsePrimitive(input)

				assert.Equal(t, PrimitiveBool, primitive)
			},
		},
		{
			title: "PrimitiveInt8",
			test: func(t *testing.T) {
				input := "int8"

				primitive := ParsePrimitive(input)

				assert.Equal(t, PrimitiveInt8, primitive)
			},
		},
		{
			title: "PrimitiveInt32",
			test: func(t *testing.T) {
				input := "int32"

				primitive := ParsePrimitive(input)

				assert.Equal(t, PrimitiveInt32, primitive)
			},
		},
		{
			title: "PrimitiveInt64",
			test: func(t *testing.T) {
				input := "int64"

				primitive := ParsePrimitive(input)

				assert.Equal(t, PrimitiveInt64, primitive)
			},
		},
		{
			title: "PrimitiveTime",
			test: func(t *testing.T) {
				input := "time"

				primitive := ParsePrimitive(input)

				assert.Equal(t, PrimitiveTime, primitive)
			},
		},
		{
			title: "PrimitiveByte",
			test: func(t *testing.T) {
				input := "byte"

				primitive := ParsePrimitive(input)

				assert.Equal(t, PrimitiveByte, primitive)
			},
		},
		{
			title: "PrimitiveFloat32",
			test: func(t *testing.T) {
				input := "float32"

				primitive := ParsePrimitive(input)

				assert.Equal(t, PrimitiveFloat32, primitive)
			},
		},
		{
			title: "PrimitiveFloat64",
			test: func(t *testing.T) {
				input := "float64"

				primitive := ParsePrimitive(input)

				assert.Equal(t, PrimitiveFloat64, primitive)
			},
		},
		{
			title: "PrimitiveTimeOnly",
			test: func(t *testing.T) {
				input := "timeonly"

				primitive := ParsePrimitive(input)

				assert.Equal(t, PrimitiveTimeOnly, primitive)
			},
		},
		{
			title: "PrimitiveDateOnly",
			test: func(t *testing.T) {
				input := "dateonly"

				primitive := ParsePrimitive(input)

				assert.Equal(t, PrimitiveDateOnly, primitive)
			},
		},
		{
			title: "PrimitiveString",
			test: func(t *testing.T) {
				input := "string"

				primitive := ParsePrimitive(input)

				assert.Equal(t, PrimitiveString, primitive)
			},
		},
		{
			title: "unknown primitive",
			test: func(t *testing.T) {
				input := "invalid"

				primitive := ParsePrimitive(input)

				assert.Equal(t, PrimitiveUnknown, primitive)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}
