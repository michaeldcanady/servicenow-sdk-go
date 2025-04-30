package tableapi

import (
	"time"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ElementValue represents a generic value.
type ElementValue interface {
	// IsNil returns whether the element is nil or not.
	IsNil() bool
	// GetStringValue returns a String value from the element.
	GetStringValue() (*string, error)
	// GetBoolValue returns a Bool value from the element.
	GetBoolValue() (*bool, error)
	// GetInt8Value returns a Int8 value from the element.
	GetInt8Value() (*int8, error)
	// GetByteValue returns a Byte value from the element.
	GetByteValue() (*byte, error)
	// GetFloat32Value returns a Float32 value from the element.
	GetFloat32Value() (*float32, error)
	// GetFloat64Value returns a Float64 value from the element.
	GetFloat64Value() (*float64, error)
	// GetInt32Value returns a Int32 value from the element.
	GetInt32Value() (*int32, error)
	// GetInt64Value returns a Int64 value from the element.
	GetInt64Value() (*int64, error)
	// GetTimeValue returns a Time value from the element.
	GetTimeValue() (*time.Time, error)
	// GetTimeOnlyValue returns a Time-only value from the element.
	GetTimeOnlyValue() (*serialization.TimeOnly, error)
	// GetDateOnlyValue returns a Date-only value from the element.
	GetDateOnlyValue() (*serialization.DateOnly, error)
	// GetEnumValue returns an enum value from the element.
	GetEnumValue(parser serialization.EnumFactory) (interface{}, error)
	// GetCollectionOfPrimitiveValues returns a Collection of specified primitive values from the element.
	GetCollectionOfPrimitiveValues(targetType Primitive) ([]interface{}, error)
	// GetRawValue returns the value of the element as an interface.
	GetRawValue() (interface{}, error)
	serialization.Parsable
}
