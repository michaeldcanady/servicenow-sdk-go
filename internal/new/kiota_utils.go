package internal

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

func SetStringValueFromSource(setter func(*string) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetStringValue()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetStringCollectionValueFromSource(setter func([]string) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}

		collection, err := CastCollection[interface{}, string](value)
		if err != nil {
			return err
		}

		return setter(collection)
	}
}

func SetBoolValueFromSource(setter func(*bool) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetBoolValue()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetBoolCollectionValueFromSource(setter func([]bool) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetCollectionOfPrimitiveValues("bool")
		if err != nil {
			return err
		}

		collection, err := CastCollection[interface{}, bool](value)
		if err != nil {
			return err
		}

		return setter(collection)
	}
}

func SetInt8ValueFromSource(setter func(*int8) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetInt8Value()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetInt8CollectionValueFromSource(setter func([]int8) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetCollectionOfPrimitiveValues("int8")
		if err != nil {
			return err
		}

		collection, err := CastCollection[interface{}, int8](value)
		if err != nil {
			return err
		}

		return setter(collection)
	}
}

func SetByteValueFromSource(setter func(*byte) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetByteValue()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetByteArrayValueFromSource(setter func([]byte) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetByteArrayValue()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetFloat32ValueFromSource(setter func(*float32) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetFloat32Value()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetFloat32CollectionValueFromSource(setter func([]float32) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetCollectionOfPrimitiveValues("float32")
		if err != nil {
			return err
		}

		collection, err := CastCollection[interface{}, float32](value)
		if err != nil {
			return err
		}

		return setter(collection)
	}
}

func SetFloat64ValueFromSource(setter func(*float64) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetFloat64Value()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetFloat64CollectionValueFromSource(setter func([]float64) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetCollectionOfPrimitiveValues("float64")
		if err != nil {
			return err
		}

		collection, err := CastCollection[interface{}, float64](value)
		if err != nil {
			return err
		}

		return setter(collection)
	}
}

func SetInt32ValueFromSource(setter func(*int32) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetInt32Value()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetInt32CollectionValueFromSource(setter func([]int32) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetCollectionOfPrimitiveValues("int32")
		if err != nil {
			return err
		}

		collection, err := CastCollection[interface{}, int32](value)
		if err != nil {
			return err
		}

		return setter(collection)
	}
}

func SetInt64ValueFromSource(setter func(*int64) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetInt64Value()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetInt64CollectionValueFromSource(setter func([]int64) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetCollectionOfPrimitiveValues("int64")
		if err != nil {
			return err
		}

		collection, err := CastCollection[interface{}, int64](value)
		if err != nil {
			return err
		}

		return setter(collection)
	}
}

func SetTimeValueFromSource(setter func(*time.Time) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetTimeValue()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetTimeCollectionValueFromSource(setter func([]time.Time) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetCollectionOfPrimitiveValues("time")
		if err != nil {
			return err
		}

		collection, err := CastCollection[interface{}, time.Time](value)
		if err != nil {
			return err
		}

		return setter(collection)
	}
}

func SetISODurationValueFromSource(setter func(*serialization.ISODuration) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetISODurationValue()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetISODurationCollectionValueFromSource(setter func([]serialization.ISODuration) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetCollectionOfPrimitiveValues("isoduration")
		if err != nil {
			return err
		}

		collection, err := CastCollection[interface{}, serialization.ISODuration](value)
		if err != nil {
			return err
		}

		return setter(collection)
	}
}

func SetTimeOnlyValueFromSource(setter func(*serialization.TimeOnly) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetTimeOnlyValue()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetTimeOnlyCollectionValueFromSource(setter func([]serialization.TimeOnly) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetCollectionOfPrimitiveValues("timeonly")
		if err != nil {
			return err
		}

		collection, err := CastCollection[interface{}, serialization.TimeOnly](value)
		if err != nil {
			return err
		}

		return setter(collection)
	}
}

func SetDateOnlyValueFromSource(setter func(*serialization.DateOnly) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetDateOnlyValue()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetDateOnlyCollectionValueFromSource(setter func([]serialization.DateOnly) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetCollectionOfPrimitiveValues("dateonly")
		if err != nil {
			return err
		}

		collection, err := CastCollection[interface{}, serialization.DateOnly](value)
		if err != nil {
			return err
		}

		return setter(collection)
	}
}

func SetUUIDValueFromSource(setter func(*uuid.UUID) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetUUIDValue()
		if err != nil {
			return err
		}
		return setter(value)
	}
}

func SetUUIDCollectionValueFromSource(setter func([]uuid.UUID) error) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetCollectionOfPrimitiveValues("uuid")
		if err != nil {
			return err
		}

		collection, err := CastCollection[interface{}, uuid.UUID](value)
		if err != nil {
			return err
		}

		return setter(collection)
	}
}

func SetEnumValueFromSource[T any](setter func(*T) error, factory serialization.EnumFactory) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetEnumValue(factory)
		if err != nil {
			return err
		}

		typedValue, ok := value.(*T)
		if !ok {
			return fmt.Errorf("value is not %T", typedValue)
		}

		return setter(typedValue)
	}
}

// TODO: add SetEnumCollectionValueFromSource

func SetObjectValueFromSource[T serialization.Parsable](setter func(T) error, factory serialization.ParsableFactory) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetObjectValue(factory)
		if err != nil {
			return err
		}

		typedValue, ok := value.(T)
		if !ok {
			return fmt.Errorf("value is not %T", typedValue)
		}

		return setter(typedValue)
	}
}

// TODO: add SetObjectCollectionValueFromSource
func SetObjectCollectionValueFromSource[T serialization.Parsable](setter func([]T) error, factory serialization.ParsableFactory) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		value, err := node.GetCollectionOfObjectValues(factory)
		if err != nil {
			return err
		}

		collection, err := CastCollection[serialization.Parsable, T](value)
		if err != nil {
			return err
		}

		return setter(collection)
	}
}

// TODO: move to conversion
func CastCollection[T, R any](collection []T) ([]R, error) {
	var err error
	newCollection := CollectionApply(collection, func(in T) (R, bool) {
		out, ok := any(in).(R)
		if !ok {
			var emptyR R
			err = fmt.Errorf("item not %T", emptyR)
			return emptyR, false
		}
		return out, true
	})
	if err != nil {
		return nil, err
	}

	return newCollection, nil
}

// TODO: move to conversion
func CollectionApply[T, R any](collection []T, mutator func(in T) (R, bool)) []R {
	outputs := make([]R, len(collection))
	for i, item := range collection {
		output, ok := mutator(item)
		if !ok {
			break
		}

		outputs[i] = output
	}
	return outputs
}
