package tableapi

// Primitive represents a base type in golang
type Primitive int64

const (
	// PrimitiveUnknown represents an unknown base type
	PrimitiveUnknown Primitive = iota - 1
	// PrimitiveBool represents a boolean base type
	PrimitiveBool
	// PrimitiveInt8 represents an int8 base type
	PrimitiveInt8
	// PrimitiveInt32 represents an int32 base type
	PrimitiveInt32
	// PrimitiveInt64 represents an int64 base type
	PrimitiveInt64
	// PrimitiveTime represents the time base type
	PrimitiveTime
	// PrimitiveByte represents a byte base type
	PrimitiveByte
	// PrimitiveFloat32 represents a float32 base type
	PrimitiveFloat32
	// PrimitiveFloat64 represents a float64 base type
	PrimitiveFloat64
	// PrimitiveTimeOnly represents the time only base type
	PrimitiveTimeOnly
	// PrimitiveDateOnly represents the date only base type
	PrimitiveDateOnly
	// PrimitiveString represents a string base type
	PrimitiveString
)

// String return string representation
func (p Primitive) String() string {
	value, ok := map[Primitive]string{
		PrimitiveUnknown:  "unknown",
		PrimitiveBool:     "bool",
		PrimitiveInt8:     "int8",
		PrimitiveInt32:    "int32",
		PrimitiveInt64:    "int64",
		PrimitiveTime:     "time",
		PrimitiveByte:     "byte",
		PrimitiveFloat32:  "float32",
		PrimitiveFloat64:  "float64",
		PrimitiveTimeOnly: "timeonly",
		PrimitiveDateOnly: "dateonly",
		PrimitiveString:   "string",
	}[p]

	if !ok {
		return PrimitiveUnknown.String()
	}

	return value
}
