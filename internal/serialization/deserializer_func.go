package serialization

import (
	"errors"
	"time"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type DeserializerFunc[T any] func(setter ModelSetter[T]) serialization.NodeParser

type Mutator[T, S any] func(input T) (S, error)

func SetMutatedValueFromSource[T, S any](source func() (T, error), setter ModelSetter[S], mutator Mutator[T, S]) error {
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

func DeserializeMutatedStringFunc[T any](setter ModelSetter[T], mutator Mutator[*string, T]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetMutatedValueFromSource(node.GetStringValue, setter, mutator)
	}
}

func DeserializeStringFunc(setter ModelSetter[*string]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetStringValue, setter)
	}
}

func DeserializeInt64Func(setter ModelSetter[*int64]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetInt64Value, setter)
	}
}

func DeserializeBoolFunc(setter ModelSetter[*bool]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetBoolValue, setter)
	}
}

func DeserializeFloat64Func(setter ModelSetter[*float64]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetFloat64Value, setter)
	}
}

func DeserializeTimeFunc(setter ModelSetter[*time.Time]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetTimeValue, setter)
	}
}

func DeserializeByteArrayFunc(setter ModelSetter[[]byte]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetValueFromSource(node.GetByteArrayValue, setter)
	}
}

func DeserializeMutatedByteArrayFunc[T any](setter ModelSetter[T], mutator Mutator[[]byte, T]) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		return SetMutatedValueFromSource(node.GetByteArrayValue, setter, mutator)
	}
}

func DeserializeCollectionOfObjectValuesFunc[T serialization.Parsable](setter ModelSetter[[]T], factory serialization.ParsableFactory) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		val, err := node.GetCollectionOfObjectValues(factory)
		if err != nil {
			return err
		}
		res := make([]T, len(val))
		for i, v := range val {
			res[i] = v.(T)
		}
		return setter(res)
	}
}

func DeserializeObjectValueFunc[T serialization.Parsable](setter ModelSetter[T], factory serialization.ParsableFactory) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		val, err := node.GetObjectValue(factory)
		if err != nil {
			return err
		}
		return setter(val.(T))
	}
}

func DeserializeEnumFunc[T any](setter ModelSetter[*T], factory serialization.EnumFactory) serialization.NodeParser {
	return func(node serialization.ParseNode) error {
		val, err := node.GetEnumValue(factory)
		if err != nil {
			return err
		}
		if val == nil {
			return setter(nil)
		}

		if v, ok := val.(*T); ok {
			return setter(v)
		}

		if v, ok := val.(T); ok {
			return setter(&v)
		}

		return errors.New("unexpected type from enum factory")
	}
}
