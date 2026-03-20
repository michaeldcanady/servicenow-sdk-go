package utils

import (
	"errors"
)

// ModelSetter represents a function that sets a value.
type ModelSetter[T any] func(val T) error

func SetMutatedValueFromSource[T, S any](source ModelGetter[T], setter ModelSetter[S], mutator Mutator[T, S]) error {
	if source == nil {
		return errors.New("source is nil")
	}

	if mutator == nil {
		return errors.New("mutator is nil")
	}

	if setter == nil {
		return errors.New("setter is nil")
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

func SetValueFromSource[T any](source ModelGetter[T], setter ModelSetter[T]) error {
	return SetMutatedValueFromSource(source, setter, NoOpMutator)
}

func WriteMutatedValueToSource[T, S any](writer ModelSetter[S], accessor ModelAccessor[T], mutator Mutator[T, S]) error {
	if writer == nil {
		return errors.New("writer is nil")
	}

	if mutator == nil {
		return errors.New("mutator is nil")
	}

	if accessor == nil {
		return errors.New("accessor is nil")
	}

	val, err := accessor()
	if err != nil {
		return err
	}

	mutatedT, err := mutator(val)
	if err != nil {
		return err
	}

	return writer(mutatedT)
}

func WriteValueToSource[T any](writer ModelSetter[T], accessor ModelAccessor[T]) error {
	return WriteMutatedValueToSource(writer, accessor, func(t T) (T, error) { return t, nil })
}
