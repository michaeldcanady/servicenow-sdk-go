package tableapi

import (
	"errors"
	"fmt"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
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

// elementValue is an implementation of ElementValue.
type elementValue struct {
	val interface{}
}

// newElementValue returns a new Element Value
func newElementValue(val interface{}) ElementValue {
	return &elementValue{
		val: val,
	}
}

func CreateElementValuetFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newElementValue(nil), nil
}

// Serialize writes the objects properties to the current writer.
func (eV *elementValue) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(eV) {
		return nil
	}

	return errors.New("Serialize not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (eV *elementValue) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{}
}

// IsNil returns whether the element is nil or not.
func (eV *elementValue) IsNil() bool {
	return internal.IsNil(eV) || internal.IsNil(eV.val)
}

func (eV *elementValue) setValue(val interface{}) error {
	if eV.IsNil() {
		return nil
	}
	eV.val = val
	return nil
}

// GetStringValue returns a String value from the element.
func (eV *elementValue) GetStringValue() (*string, error) {
	if eV.IsNil() {
		return nil, nil
	}

	val, ok := eV.val.(*string)
	if !ok {
		return nil, fmt.Errorf("type '%T' is not compatible with type string", eV.val)
	}
	return val, nil
}

// GetBoolValue returns a Bool value from the element.
func (eV *elementValue) GetBoolValue() (*bool, error) {
	if eV.IsNil() {
		return nil, nil
	}

	val, ok := eV.val.(*bool)
	if !ok {
		return nil, fmt.Errorf("type '%T' is not compatible with type bool", eV.val)
	}
	return val, nil
}

// GetInt8Value returns a Int8 value from the element.
func (eV *elementValue) GetInt8Value() (*int8, error) {
	if eV.IsNil() {
		return nil, nil
	}

	var val int8

	if err := internal.As(eV.val, &val); err != nil {
		return nil, err
	}

	return &val, nil
}

// GetByteValue returns a Byte value from the element.
func (eV *elementValue) GetByteValue() (*byte, error) {
	if eV.IsNil() {
		return nil, nil
	}

	var val byte

	if err := internal.As(eV.val, &val); err != nil {
		return nil, err
	}

	return &val, nil
}

// GetFloat32Value returns a Float32 value from the element.
func (eV *elementValue) GetFloat32Value() (*float32, error) {
	if eV.IsNil() {
		return nil, nil
	}

	var val float32

	if err := internal.As(eV.val, &val); err != nil {
		return nil, err
	}

	return &val, nil
}

// GetFloat64Value returns a Float64 value from the element.
func (eV *elementValue) GetFloat64Value() (*float64, error) {
	if eV.IsNil() {
		return nil, nil
	}

	var val float64

	if err := internal.As(eV.val, &val); err != nil {
		return nil, err
	}

	return &val, nil
}

// GetInt32Value returns a Int32 value from the element.
func (eV *elementValue) GetInt32Value() (*int32, error) {
	if eV.IsNil() {
		return nil, nil
	}

	var val int32

	if err := internal.As(eV.val, &val); err != nil {
		return nil, err
	}

	return &val, nil
}

// GetInt64Value returns a Int64 value from the element.
func (eV *elementValue) GetInt64Value() (*int64, error) {
	if eV.IsNil() {
		return nil, nil
	}

	var val int64

	if err := internal.As(eV.val, &val); err != nil {
		return nil, err
	}

	return &val, nil
}

// GetTimeValue returns a Time value from the element.
func (eV *elementValue) GetTimeValue() (*time.Time, error) {
	if eV.IsNil() {
		return nil, nil
	}
	v, err := eV.GetStringValue()
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}

	parsed, err := time.Parse(time.RFC3339, *v)
	return &parsed, err
}

// GetTimeOnlyValue returns a Time-only value from the element.
func (eV *elementValue) GetTimeOnlyValue() (*serialization.TimeOnly, error) {
	if eV.IsNil() {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}

// GetDateOnlyValue returns a Date-only value from the element.
func (eV *elementValue) GetDateOnlyValue() (*serialization.DateOnly, error) {
	if eV.IsNil() {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}

// GetEnumValue returns an enum value from the element.
func (eV *elementValue) GetEnumValue(parser serialization.EnumFactory) (interface{}, error) {
	if eV.IsNil() {
		return nil, nil
	}
	if parser == nil {
		return nil, errors.New("parser is nil")
	}
	s, err := eV.GetStringValue()
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, nil
	}
	return parser(*s)
}

// GetCollectionOfPrimitiveValues returns a Collection of specified primitive values from the element.
func (eV *elementValue) GetCollectionOfPrimitiveValues(targetType Primitive) ([]interface{}, error) {
	if eV.IsNil() {
		return nil, nil
	}

	if targetType == PrimitiveUnknown {
		return nil, fmt.Errorf("target type can't be %s", PrimitiveUnknown)
	}

	rawCollection, ok := eV.val.([]interface{})
	if !ok {
		return nil, errors.New("val is not a collection")
	}

	collection := make([]interface{}, 0, len(rawCollection))
	for i, v := range rawCollection {
		var (
			val interface{}
			err error
		)
		if v != nil {
			val, err = eV.getPrimitiveValue(targetType)
			if err != nil {
				return nil, err
			}
		}
		collection[i] = val
	}

	return collection, nil
}

// getPrimitiveValue returns the element value as the specified primitive type or error if not of that type
func (eV *elementValue) getPrimitiveValue(targetType Primitive) (interface{}, error) {
	if eV.IsNil() {
		return nil, nil
	}

	switch targetType {
	case PrimitiveBool:
		return eV.GetBoolValue()
	case PrimitiveByte:
		return eV.GetByteValue()
	case PrimitiveDateOnly:
		return eV.GetDateOnlyValue()
	case PrimitiveFloat32:
		return eV.GetFloat32Value()
	case PrimitiveFloat64:
		return eV.GetFloat64Value()
	case PrimitiveInt32:
		return eV.GetInt32Value()
	case PrimitiveInt64:
		return eV.GetInt64Value()
	case PrimitiveInt8:
		return eV.GetInt8Value()
	case PrimitiveTime:
		return eV.GetTimeValue()
	case PrimitiveTimeOnly:
		return eV.GetTimeOnlyValue()
	case PrimitiveString:
		return eV.GetStringValue()
	default:
		return nil, fmt.Errorf("unknown primitive %s", targetType)
	}
}

// GetRawValue returns the value of the element as an interface.
func (eV *elementValue) GetRawValue() (interface{}, error) {
	if eV.IsNil() {
		return nil, nil
	}

	return eV.val, nil
}
