package kiota

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DeserializerFunc[T any] func(setter utils.ModelSetter[T]) abstractions.NodeParser

func SetMutatedValueFromSource[T, S any](source func() (T, error), setter utils.ModelSetter[S], mutator utils.Mutator[T, S]) error {
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

func DeserializeMutatedStringFunc[T any](setter utils.ModelSetter[T], mutator utils.Mutator[*string, T]) abstractions.NodeParser {
	return func(node abstractions.ParseNode) error {
		return SetMutatedValueFromSource(node.GetStringValue, setter, mutator)
	}
}

func DeserializeStringFunc(setter utils.ModelSetter[*string]) abstractions.NodeParser {
	return func(node abstractions.ParseNode) error {
		return SetValueFromSource(node.GetStringValue, setter)
	}
}

func DeserializeInt64Func(setter utils.ModelSetter[*int64]) abstractions.NodeParser {
	return func(node abstractions.ParseNode) error {
		return SetValueFromSource(node.GetInt64Value, setter)
	}
}

func SetValueFromSource[T any](source func() (T, error), setter utils.ModelSetter[T]) error {
	return SetMutatedValueFromSource(source, setter, func(t T) (T, error) { return t, nil })
}
