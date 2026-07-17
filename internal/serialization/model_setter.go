package serialization

import (
	"errors"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type ModelSetter[T any] func(val T) error

type ModelAccessor[T any] func() (T, error)

type WriterFunc func(serialization.SerializationWriter) error

// pipeMutatedValue reads a T from get, converts it to S via mutator, and passes it to put.
// It's the shared shape behind both SetMutatedValueFromSource (node -> model) and
// WriteMutatedValueToSource (model -> writer): a get/mutate/put pipeline that only differs
// in which direction the value is flowing and what to call the missing-function errors.
func pipeMutatedValue[T, S any](get func() (T, error), getErr error, put func(S) error, putErr error, mutator conversion.Mutator[T, S]) error {
	if get == nil {
		return getErr
	}

	if mutator == nil {
		return snerrors.ErrNilMutator
	}

	if put == nil {
		return putErr
	}

	val, err := get()
	if err != nil {
		return err
	}

	mutatedT, err := mutator(val)
	if err != nil {
		return err
	}
	return put(mutatedT)
}

func identityMutator[T any](t T) (T, error) { return t, nil }

func SetMutatedValueFromSource[T, S any](source func() (T, error), setter ModelSetter[S], mutator conversion.Mutator[T, S]) error {
	return pipeMutatedValue(source, errors.New("source is nil"), setter, errors.New("setter is nil"), mutator)
}

func SetValueFromSource[T any](source func() (T, error), setter ModelSetter[T]) error {
	return SetMutatedValueFromSource(source, setter, identityMutator[T])
}

func WriteMutatedValueToSource[T, S any](writer func(S) error, accessor ModelAccessor[T], mutator conversion.Mutator[T, S]) error {
	return pipeMutatedValue(accessor, errors.New("accessor is nil"), writer, errors.New("writer is nil"), mutator)
}

func WriteValueToSource[T any](writer func(T) error, accessor ModelAccessor[T]) error {
	return WriteMutatedValueToSource(writer, accessor, identityMutator[T])
}
