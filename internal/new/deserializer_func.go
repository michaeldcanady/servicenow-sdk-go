package internal

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type ModelSetter[T any] func(val T) error

type DeserializerFunc[T any] func(setter ModelSetter[T]) serialization.NodeParser

type TypeConverter[T, S any] func(input T) (S, error)

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

func StringPtrToInt64Ptr(input *string) (*int64, error) {
	if input == nil {
		return nil, nil
	}
	intVal, err := strconv.Atoi(*input)
	if err != nil {
		return nil, err
	}
	int64Val := int64(intVal)
	return &int64Val, nil
}

func StringPtrToFloat64Ptr(input *string) (*float64, error) {
	if input == nil {
		return nil, nil
	}
	floatVal, err := strconv.ParseFloat(*input, 64)
	if err != nil {
		return nil, err
	}
	return &floatVal, nil
}

func StringPtrToBoolPtr(input *string) (*bool, error) {
	if input == nil {
		return nil, nil
	}
	boolVal, err := strconv.ParseBool(*input)
	if err != nil {
		return nil, err
	}
	return &boolVal, nil
}

func StringPtrToTimePtr(format string) TypeConverter[*string, *time.Time] {
	return func(input *string) (*time.Time, error) {
		if input == nil {
			return nil, nil
		}
		dateTime, err := StringToTime(format)(*input)
		if err != nil {
			return nil, err
		}
		return &dateTime, nil
	}
}

func StringToTime(format string) TypeConverter[string, time.Time] {
	return func(input string) (time.Time, error) {
		dateTime, err := time.Parse(format, input)
		if err != nil {
			return time.Time{}, err
		}

		return dateTime, nil
	}

}

func StringPtrToPrimitiveSlice[T any](delimiter string, mutator func(string) (T, error)) TypeConverter[*string, []T] {
	return func(input *string) ([]T, error) {
		sliceString := *input

		stringSlice := strings.Split(sliceString, delimiter)

		primitiveSlice := make([]T, len(stringSlice))

		for index, stringVal := range stringSlice {
			primitiveVal, err := mutator(stringVal)
			if err != nil {
				return nil, err
			}
			primitiveSlice[index] = primitiveVal
		}
		return primitiveSlice, nil
	}
}
