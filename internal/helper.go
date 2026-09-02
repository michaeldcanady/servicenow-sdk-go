// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package internal

import (
	"reflect"
)

// ToPointer Converts provided value to pointer.
func ToPointer[T any](value T) *T {
	return &value
}

// Logger interface for custom logging.
type Logger interface {
	Log(message string, args ...interface{})
}

// NoOpLogger is a default logger that does nothing.
type NoOpLogger struct{}

func (n *NoOpLogger) Log(_ string, _ ...interface{}) {}

// IsPointer
func IsPointer(value any) bool {
	return reflect.ValueOf(value).Kind() == reflect.Pointer
}
