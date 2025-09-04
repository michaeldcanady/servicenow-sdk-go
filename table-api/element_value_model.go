//go:build preview.tableApiV2

package tableapi

import (
	"errors"
	"fmt"
	"reflect"

	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ElementValue is an implementation of ElementValue.
type ElementValue struct {
	val any
}

// NewElementValue returns a new Element Value
func NewElementValue(val any) (*ElementValue, error) {
	return loadTree(val)
}

func loadTree(val any) (*ElementValue, error) {
	var err error

	rv := reflect.ValueOf(val)

	for {
		switch rv.Kind() {
		case reflect.Array, reflect.Slice:
			array := make([]*ElementValue, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				if array[i], err = loadTree(rv.Index(i).Interface()); err != nil {
					return nil, err
				}
			}
			return &ElementValue{val: array}, nil
		case reflect.Map:
			mapping := make(map[string]*ElementValue, rv.Len())
			for _, valKey := range rv.MapKeys() {
				key := valKey.Interface().(string)
				if mapping[key], err = loadTree(rv.MapIndex(valKey).Interface()); err != nil {
					return nil, err
				}
			}
			return &ElementValue{val: mapping}, nil
		case reflect.Pointer:
			if rv.IsNil() {
				return nil, nil
			}
			return loadTree(rv.Elem().Interface())
		default:
			return &ElementValue{val: val}, nil
		}
	}
}

// CreateElementValueFromDiscriminatorValue is a parsable factory for creating a ElementValueModel
func CreateElementValueFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewElementValue(nil)
}

// Serialize writes the objects properties to the current writer.
func (eV *ElementValue) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(eV) {
		return nil
	}

	return errors.New("Serialize is not supported")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (eV *ElementValue) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{}
}

// IsNil returns whether the element is nil or not.
func (eV *ElementValue) IsNil() bool {
	return internal.IsNil(eV) || internal.IsNil(eV.val)
}

func (eV *ElementValue) setValue(val interface{}) error { //nolint: unused
	if internal.IsNil(eV) {
		return nil
	}

	if !internal.IsPointer(val) {
		return errors.New("val is not a pointer")
	}

	eV.val = val

	return nil
}

// GetStringValue returns a String value from the element.
func (eV *ElementValue) GetStringValue() (*string, error) {
	if eV.IsNil() {
		return nil, nil
	}

	var val string

	if err := internal.As(eV.val, &val); err != nil {
		return nil, err
	}

	return &val, nil
}

// GetBoolValue returns a Bool value from the element.
func (eV *ElementValue) GetBoolValue() (*bool, error) {
	if eV.IsNil() {
		return nil, nil
	}

	var val bool

	if err := internal.As(eV.val, &val); err != nil {
		return nil, err
	}

	return &val, nil
}

// GetInt8Value returns a Int8 value from the element.
func (eV *ElementValue) GetInt8Value() (*int8, error) {
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
func (eV *ElementValue) GetByteValue() (*byte, error) {
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
func (eV *ElementValue) GetFloat32Value() (*float32, error) {
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
func (eV *ElementValue) GetFloat64Value() (*float64, error) {
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
func (eV *ElementValue) GetInt32Value() (*int32, error) {
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
func (eV *ElementValue) GetInt64Value() (*int64, error) {
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
//func (eV *ElementValueModel) GetTimeValue() (*time.Time, error) {
//	if eV.IsNil() {
//		return nil, nil
//	}
//	v, err := eV.GetStringValue()
//	if err != nil {
//		return nil, err
//	}
//	if v == nil {
//		return nil, nil
//	}
//
//	parsed, err := time.Parse(time.RFC3339, *v)
//	return &parsed, err
//}

// GetTimeOnlyValue returns a Time-only value from the element.
//func (eV *ElementValueModel) GetTimeOnlyValue() (*serialization.TimeOnly, error) {
//	if eV.IsNil() {
//		return nil, nil
//	}
//
//	return nil, errors.New("not implemented")
//}

// GetDateOnlyValue returns a Date-only value from the element.
//func (eV *ElementValueModel) GetDateOnlyValue() (*serialization.DateOnly, error) {
//	if eV.IsNil() {
//		return nil, nil
//	}
//
//	return nil, errors.New("not implemented")
//}

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
	if s == nil {
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
			val, err = v.getPrimitiveValue(targetType)
			if err != nil {
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
