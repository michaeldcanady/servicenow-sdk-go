package utils

import (
	"reflect"
)

// IsNil checks if a value is nil or a nil interface.
func IsNil(a interface{}) bool {
	defer func() { _ = recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

// ToPointer Converts provided value to pointer.
func ToPointer[T any](value T) *T {
	return &value
}

// IsPointer
func IsPointer(value any) bool {
	return reflect.ValueOf(value).Kind() == reflect.Pointer
}

// ModelSetter represents a function that sets a value.
type ModelSetter[T any] func(val T) error

// Mutator represents a function that transforms a value.
type Mutator[T, S any] func(input T) (S, error)
