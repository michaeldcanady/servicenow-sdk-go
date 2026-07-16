package tableapi

const (
	primitiveUnknown  = "unknown"
	primitiveBool     = "bool"
	primitiveInt8     = "int8"
	primitiveInt32    = "int32"
	primitiveInt64    = "int64"
	primitiveTime     = "time"
	primitiveByte     = "byte"
	primitiveFloat32  = "float32"
	primitiveFloat64  = "float64"
	primitiveTimeOnly = "timeonly"
	primitiveDateOnly = "dateonly"
	primitiveString   = "string"
)

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

var primitiveStrings = map[Primitive]string{
	PrimitiveUnknown:  primitiveUnknown,
	PrimitiveBool:     primitiveBool,
	PrimitiveInt8:     primitiveInt8,
	PrimitiveInt32:    primitiveInt32,
	PrimitiveInt64:    primitiveInt64,
	PrimitiveTime:     primitiveTime,
	PrimitiveByte:     primitiveByte,
	PrimitiveFloat32:  primitiveFloat32,
	PrimitiveFloat64:  primitiveFloat64,
	PrimitiveTimeOnly: primitiveTimeOnly,
	PrimitiveDateOnly: primitiveDateOnly,
	PrimitiveString:   primitiveString,
}

// String return string representation
func (p Primitive) String() string {
	value, ok := primitiveStrings[p]

	if !ok {
		return PrimitiveUnknown.String()
	}

	return value
}
