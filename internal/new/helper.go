package internal

import (
	"reflect"
)

// IsNil checks if a value is nil or a nil interface
func IsNil(a interface{}) bool {
	defer func() { _ = recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

// ToPointer
func ToPointer[T any](value T) *T {
	return &value
}
