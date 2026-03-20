package kiota

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type ModelSetter[T any] = utils.ModelSetter[T]

type ModelAccessor[T any] = utils.ModelAccessor[T]

type Mutator[T, S any] = utils.Mutator[T, S]

type WriterFunc = utils.WriterFunc

type SerializerFunc[T any] = utils.SerializerFunc[T]

type DeserializerFunc[T any] = utils.DeserializerFunc[T]

func SetMutatedValueFromSource[T, S any](source utils.ModelGetter[T], setter ModelSetter[S], mutator Mutator[T, S]) error {
	return utils.SetMutatedValueFromSource(source, setter, mutator)
}

func SetValueFromSource[T any](source func() (T, error), setter ModelSetter[T]) error {
	return SetMutatedValueFromSource(source, setter, func(t T) (T, error) { return t, nil })
}

func WriteMutatedValueToSource[T, S any](writer utils.ModelSetter[S], accessor ModelAccessor[T], mutator Mutator[T, S]) error {
	return utils.WriteMutatedValueToSource(writer, accessor, mutator)
}

func WriteValueToSource[T any](writer utils.ModelSetter[T], accessor ModelAccessor[T]) error {
	return utils.WriteValueToSource(writer, accessor)
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
