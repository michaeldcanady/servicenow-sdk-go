package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimitive_String(t *testing.T) {
	tests := []struct {
		name string
		p    Primitive
		want string
	}{
		{name: "Unknown", p: PrimitiveUnknown, want: "unknown"},
		{name: "Bool", p: PrimitiveBool, want: "bool"},
		{name: "Int8", p: PrimitiveInt8, want: "int8"},
		{name: "Int32", p: PrimitiveInt32, want: "int32"},
		{name: "Int64", p: PrimitiveInt64, want: "int64"},
		{name: "Time", p: PrimitiveTime, want: "time"},
		{name: "Byte", p: PrimitiveByte, want: "byte"},
		{name: "Float32", p: PrimitiveFloat32, want: "float32"},
		{name: "Float64", p: PrimitiveFloat64, want: "float64"},
		{name: "TimeOnly", p: PrimitiveTimeOnly, want: "timeonly"},
		{name: "DateOnly", p: PrimitiveDateOnly, want: "dateonly"},
		{name: "String", p: PrimitiveString, want: "string"},
		{name: "OutOfRange", p: Primitive(9999), want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.p.String())
		})
	}
}
