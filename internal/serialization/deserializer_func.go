package serialization

import (
	"errors"

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
