package kiota

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type ModelSetter[T any] func(val T) error

type ModelAccessor[T any] func() (T, error)

type Mutator[T, S any] func(input T) (S, error)

type WriterFunc func(serialization.SerializationWriter) error

type SerializerFunc[T any] func(accessor ModelAccessor[T]) WriterFunc

type DeserializerFunc[T any] func(setter ModelSetter[T]) serialization.NodeParser

func SetMutatedValueFromSource[T, S any](source func() (T, error), setter ModelSetter[S], mutator Mutator[T, S]) error {
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

func SetValueFromSource[T any](source func() (T, error), setter ModelSetter[T]) error {
	return SetMutatedValueFromSource(source, setter, func(t T) (T, error) { return t, nil })
}

func WriteMutatedValueToSource[T, S any](writer func(S) error, accessor ModelAccessor[T], mutator Mutator[T, S]) error {
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

func WriteValueToSource[T any](writer func(T) error, accessor ModelAccessor[T]) error {
	return WriteMutatedValueToSource(writer, accessor, func(t T) (T, error) { return t, nil })
}

func ConfigureRequestInformation[T any](request *KiotaRequestInformation, config *abstractions.RequestConfiguration[T]) {
	if request == nil {
		return
	}
	if config == nil {
		return
	}
	if headers := config.Headers; !utils.IsNil(headers) {
		request.Headers.AddAll(headers)
	}
	if options := config.Options; !utils.IsNil(options) {
		request.AddRequestOptions(options)
	}
	if queryParams := config.QueryParameters; !utils.IsNil(queryParams) {
		request.AddQueryParameters(queryParams)
	}
}
