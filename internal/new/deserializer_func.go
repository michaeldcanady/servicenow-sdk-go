package internal

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type DeserializerFunc[T any] func(setter ModelSetter[T]) serialization.NodeParser

func SetValueFromSource[T any](source func() (T, error), setter ModelSetter[T]) error {
	return SetMutatedValueFromSource(source, setter, func(t T) (T, error) { return t, nil })
}

func SetMutatedValueFromSource[T, S any](source func() (T, error), setter ModelSetter[S], mutator func(T) (S, error)) error {
	if source == nil {
		return errors.New("source is nil")
	}

	if mutator == nil {
		return errors.New("mutator is nil")
	}

	val, err := source()
	if err != nil {
		return err
	}

	mutatedT, err := mutator(val)
	if err != nil {
		return err
	}
	return setter(mutatedT)
}

func DeserializeMutatedStringFunc[T any](setter ModelSetter[T], mutator func(*string) (T, error)) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetMutatedValueFromSource(node.GetStringValue, setter, mutator)
	}
}

func DeserializeMutatedValueFunc[T any](setter ModelSetter[T], mutator func(any) (T, error)) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetMutatedValueFromSource(node.GetRawValue, setter, mutator)
	}
}

func DeserializeStringFunc(setter ModelSetter[*string]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetStringValue, setter)
	}
}

func DeserializeBoolFunc(setter ModelSetter[*bool]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetBoolValue, setter)
	}
}

func DeserializeInt8Func(setter ModelSetter[*int8]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetInt8Value, setter)
	}
}

func DeserializeByteFunc(setter ModelSetter[*byte]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetByteValue, setter)
	}
}

func DeserializeFloat32Func(setter ModelSetter[*float32]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetFloat32Value, setter)
	}
}

func DeserializeFloat64Func(setter ModelSetter[*float64]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetFloat64Value, setter)
	}
}

func DeserializeInt32Func(setter ModelSetter[*int32]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetInt32Value, setter)
	}
}

func DeserializeInt64Func(setter ModelSetter[*int64]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetInt64Value, setter)
	}
}

func DeserializeFormattedTime(setter ModelSetter[*time.Time], format string) serialization.NodeParser {
	return DeserializeMutatedStringFunc(setter, StringPtrToTimePtr(format))
}

func DeserializeIsoDurationFunc(setter ModelSetter[*serialization.ISODuration]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetISODurationValue, setter)
	}
}

func DeserializeRawFunc(setter ModelSetter[any]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetRawValue, setter)
	}
}

func DeserializeTimeOnlyFunc(setter ModelSetter[*serialization.TimeOnly]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetTimeOnlyValue, setter)
	}
}

func DeserializeDateOnlyFunc(setter ModelSetter[*serialization.DateOnly]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetDateOnlyValue, setter)
	}
}

func DeserializeUUIDFunc(setter ModelSetter[*uuid.UUID]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetUUIDValue, setter)
	}
}

func DeserializeEnumFunc[T interface{}](setter ModelSetter[T], factory serialization.EnumFactory) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(func() (T, error) {
			var typedValue T
			var ok bool

			value, err := node.GetEnumValue(factory)
			if err != nil {
				return typedValue, err
			}
			typedValue, ok = value.(T)
			if !ok {
				return typedValue, fmt.Errorf("value is %T, expected %T", value, typedValue)
			}

			return typedValue, nil
		}, setter)
	}
}

func DeserializeByteArrayFunc(setter ModelSetter[[]byte]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetByteArrayValue, setter)
	}
}

func DeserializeObjectFunc[T serialization.Parsable](setter ModelSetter[T], factory serialization.ParsableFactory) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(func() (T, error) {
			var typedValue T
			var ok bool

			value, err := node.GetObjectValue(factory)
			if err != nil {
				return typedValue, err
			}

			typedValue, ok = value.(T)
			if !ok {
				return typedValue, fmt.Errorf("value is %T, expected %T", value, typedValue)
			}

			return typedValue, nil
		}, setter)
	}
}

func DeserializeObjectArrayFunc[T serialization.Parsable](setter ModelSetter[[]T], factory serialization.ParsableFactory) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(func() ([]T, error) {
			unknownSlice, err := node.GetCollectionOfObjectValues(factory)
			if err != nil {
				return nil, err
			}

			results := make([]T, len(unknownSlice), 0)

			for index, value := range unknownSlice {
				result, ok := value.(T)
				if !ok {
					return nil, fmt.Errorf("value is not %T", new(T))
				}

				results[index] = result
			}

			return results, nil
		}, setter)
	}
}
