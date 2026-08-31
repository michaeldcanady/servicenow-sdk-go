// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package query

import (
	"time"
)

// DateTimeValue represents a value for a date-time field.
type DateTimeValue struct {
	literal string
}

// NewDateTimeValue creates a new DateTimeValue from a time.Time object.
func NewDateTimeValue(val time.Time) DateTimeValue {
	return DateTimeValue{
		literal: val.Format("2006-01-02 15:04:05"),
	}
}

// String returns the value's ServiceNow encoded-query literal.
func (v DateTimeValue) String() string {
	return v.literal
}

// Time is a shorter alias for NewDateTimeValue, useful in the fluent API.
func Time(t time.Time) DateTimeValue {
	return NewDateTimeValue(t)
}

// JS wraps a JavaScript expression with the required "javascript:" prefix as a DateTimeValue.
func JS(expr string) DateTimeValue {
	return DateTimeValue{literal: "javascript:" + expr}
}
