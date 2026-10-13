// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package internal

import "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"

// NilPointerError represents a nil pointer when a value pointer is expected.
type NilPointerError struct {
	s string
}

// NewNilPointerError instantiates a new NilPointerError.
func NewNilPointerError(text string) *NilPointerError {
	return &NilPointerError{
		s: text,
	}
}

// Error returns the error string.
func (err *NilPointerError) Error() string {
	if conversion.IsNil(err) {
		return "nil pointer error"
	}
	return err.s
}
