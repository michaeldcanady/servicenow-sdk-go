package conversion

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// TypeConverter[T, S] converts input, type T, to output, type S, and returns an error if there is one.
type TypeConverter[T, S any] func(input T) (S, error)

// StringPtrToInt64Ptr Converts string pointer to int64 pointer.
func StringPtrToInt64Ptr(input *string) (*int64, error) {
	if input == nil {
		return nil, errors.New("input is nil")
	}
	intVal, err := strconv.Atoi(*input)
	if err != nil {
		return nil, err
	}
	int64Val := int64(intVal)
	return &int64Val, nil
}

// StringPtrToFloat64Ptr Converts string pointer to float64 pointer.
func StringPtrToFloat64Ptr(input *string) (*float64, error) {
	if input == nil {
		return nil, errors.New("input is nil")
	}
	floatVal, err := strconv.ParseFloat(*input, 64)
	if err != nil {
		return nil, err
	}
	return &floatVal, nil
}

// StringPtrToBoolPtr Converts string pointer to bool pointer.
func StringPtrToBoolPtr(input *string) (*bool, error) {
	if input == nil {
		return nil, errors.New("input is nil")
	}
	boolVal, err := strconv.ParseBool(*input)
	if err != nil {
		return nil, err
	}
	return &boolVal, nil
}

// StringPtrToTimePtr Converts string pointer to formatted time pointer.
func StringPtrToTimePtr(format string) TypeConverter[*string, *time.Time] {
	return func(input *string) (*time.Time, error) {
		if input == nil {
			return nil, errors.New("input is nil")
		}
		dateTime, err := time.Parse(format, *input)
		if err != nil {
			return &time.Time{}, err
		}

		return &dateTime, nil
	}
}

// StringPtrToPrimitiveSlice Converts string pointer to slice of primitive type T
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
