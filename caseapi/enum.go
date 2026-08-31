// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package caseapi

import (
	"fmt"
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
)

// unknownEnumValueError reports that s is not a recognized wire value for the named
// enum. The result always matches [snerrors.ErrUnknownEnumValue] under errors.Is, while
// still naming the enum and echoing the offending value.
func unknownEnumValueError(enum, s string) error {
	return fmt.Errorf("%w: %q is not a valid %s", snerrors.ErrUnknownEnumValue, s, enum)
}

// invertEnumStrings builds the lower-cased inverse of an enum's value->string map, so
// Parse<Enum> can resolve a wire value case-insensitively without restating every
// constant in a switch. The unknown sentinel is skipped: it is a parse failure, not a
// value the API ever sends.
func invertEnumStrings[E comparable](strs map[E]string, unknown E) map[string]E {
	values := make(map[string]E, len(strs)-1)
	for value, str := range strs {
		if value == unknown {
			continue
		}
		values[strings.ToLower(str)] = value
	}

	return values
}
