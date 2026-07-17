package tableapi

import (
	"errors"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ElementValue represents a value of a field in a Service-Now table.
//
// It provides type-safe methods to retrieve the value as various Go types.
type ElementValue struct {
	val any
}

// NewElementValue returns a new ElementValue.
func NewElementValue(val any) (*ElementValue, error) {
	return NewElementVisitor().Visit(val)
}

// CreateElementValueFromDiscriminatorValue is a parsable factory for creating an ElementValue.
func CreateElementValueFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewElementValue(nil)
}

// Serialize writes the objects properties to the current writer.
func (eV *ElementValue) Serialize(writer serialization.SerializationWriter) error {
	return errors.New("Serialize is not supported")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (eV *ElementValue) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{}
}

// IsNil returns whether the element is nil or not.
func (eV *ElementValue) IsNil() bool {
	return conversion.IsNil(eV) || conversion.IsNil(eV.val)
}

func (eV *ElementValue) setValue(val any) error {
	if conversion.IsNil(eV) {
		return nil
	}

	if !internal.IsPointer(val) {
		return errors.New("val is not a pointer")
	}

	eV.val = val

	return nil
}

// getPrimitive converts eV's underlying value to T, returning nil if eV holds no value.
func getPrimitive[T any](eV *ElementValue, strict bool) (*T, error) {
	if eV.IsNil() {
		return nil, nil
	}

	var val T

	if err := conversion.As2(eV.val, &val, strict); err != nil {
		return nil, err
	}

	return &val, nil
}

// GetStringValue returns a String value from the element.
func (eV *ElementValue) GetStringValue() (*string, error) {
	return getPrimitive[string](eV, true)
}

// GetBoolValue returns a Bool value from the element.
func (eV *ElementValue) GetBoolValue() (*bool, error) {
	return getPrimitive[bool](eV, false)
}

// GetInt8Value returns a Int8 value from the element.
func (eV *ElementValue) GetInt8Value() (*int8, error) {
	return getPrimitive[int8](eV, false)
}

// GetByteValue returns a Byte value from the element.
func (eV *ElementValue) GetByteValue() (*byte, error) {
	return getPrimitive[byte](eV, false)
}

// GetFloat32Value returns a Float32 value from the element.
func (eV *ElementValue) GetFloat32Value() (*float32, error) {
	return getPrimitive[float32](eV, false)
}

// GetFloat64Value returns a Float64 value from the element.
func (eV *ElementValue) GetFloat64Value() (*float64, error) {
	return getPrimitive[float64](eV, false)
}

// GetInt32Value returns a Int32 value from the element.
func (eV *ElementValue) GetInt32Value() (*int32, error) {
	return getPrimitive[int32](eV, false)
}

// GetInt64Value returns a Int64 value from the element.
func (eV *ElementValue) GetInt64Value() (*int64, error) {
	return getPrimitive[int64](eV, false)
}

// GetEnumValue returns an enum value from the element.
func (eV *ElementValue) GetEnumValue(parser serialization.EnumFactory) (interface{}, error) {
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
	if s == nil || *s == "" {
		return nil, nil
	}
	return parser(*s)
}

// GetCollectionOfPrimitiveValues returns a Collection of specified primitive values from the element.
func (eV *ElementValue) GetCollectionOfPrimitiveValues(targetType Primitive) ([]interface{}, error) {
	if eV.IsNil() {
		return nil, nil
	}

	if targetType == PrimitiveUnknown {
		return nil, fmt.Errorf("target type can't be %s", PrimitiveUnknown)
	}

	rawCollection, ok := eV.val.([]*ElementValue)
	if !ok {
		return nil, errors.New("val is not a collection")
	}

	collection := make([]interface{}, len(rawCollection))
	for i, v := range rawCollection {
		var (
			val interface{}
			err error
		)
		if v != nil {
			if val, err = v.getPrimitiveValue(targetType); err != nil {
				return nil, err
			}
		}
		collection[i] = val
	}

	return collection, nil
}

// getPrimitiveValue returns the element value as the specified primitive type or error if not of that type
func (eV *ElementValue) getPrimitiveValue(targetType Primitive) (interface{}, error) {
	if eV.IsNil() {
		return nil, nil
	}

	switch targetType {
	case PrimitiveBool:
		return eV.GetBoolValue()
	case PrimitiveByte:
		return eV.GetByteValue()
	//case PrimitiveDateOnly:
	//	return eV.GetDateOnlyValue()
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
	//case PrimitiveTime:
	//	return eV.GetTimeValue()
	//case PrimitiveTimeOnly:
	//	return eV.GetTimeOnlyValue()
	case PrimitiveString:
		return eV.GetStringValue()
	default:
		return nil, fmt.Errorf("unknown primitive %s", targetType)
	}
}

// GetRawValue returns the value of the element as an interface.
func (eV *ElementValue) GetRawValue() (interface{}, error) {
	if eV.IsNil() {
		return nil, nil
	}

	return eV.val, nil
}
