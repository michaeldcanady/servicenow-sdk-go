// Copyright (c) 2026 Michael Canady
// SPDX-License-Identifier: MIT

package query

import (
	"fmt"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/ast"
)

// clauseSeparator is the encoded-query metacharacter that terminates a term
// and starts a new clause ("^" / "^OR"). Encoded queries provide no escape
// sequence for it, so a value containing it breaks out of its term and appends
// arbitrary clauses to the query. It is rejected in every caller-supplied
// value and fragment.
const clauseSeparator = "^" //nolint:gochecknoglobals

// The remaining metacharacters are structural only in the operator context
// that renders them, and ordinary literals everywhere else (for example, the
// comma in "Smith, John" or the @ in "user@example.com").
const (
	// listSeparator joins the values of an IN / NOT IN list.
	listSeparator = "," //nolint:gochecknoglobals

	// pairSeparator joins the halves of a BETWEEN pair and the segments of an
	// ON date-time composite ("label@javascript:expr@javascript:expr").
	pairSeparator = "@" //nolint:gochecknoglobals
)

// reservedValueCharacters returns the metacharacters that take on structural
// meaning inside a value rendered with op: the clause separator everywhere,
// plus the list separator for IN / NOT IN values and the pair separator for
// BETWEEN values.
func reservedValueCharacters(op ast.Operator) string {
	switch op {
	case ast.OperatorIsOneOf, ast.OperatorIsNotOneOf:
		return clauseSeparator + listSeparator
	case ast.OperatorBetween:
		return clauseSeparator + pairSeparator
	default:
		return clauseSeparator
	}
}

// validateQueryValue checks that a consumer-supplied value contains none of
// the metacharacters that would be structural for op (see
// [reservedValueCharacters]). It returns nil for safe values and a descriptive
// error otherwise; callers turn that error into an error Condition via
// [NewErrorCondition] so it surfaces through [Condition.Error].
//
// Literals composed internally by this package (for example the
// "label@javascript:expr@javascript:expr" strings built by OnSpecialty) are
// trusted and must bypass this check.
func validateQueryValue(field string, op ast.Operator, val any) error {
	rendered := fmt.Sprintf("%v", val)
	reserved := reservedValueCharacters(op)
	index := strings.IndexAny(rendered, reserved)
	if index < 0 {
		return nil
	}

	return fmt.Errorf(
		"value %q for field %q (operator %q) contains reserved encoded-query character %q; it structures the query itself and cannot be escaped",
		rendered, field, op.String(), rendered[index],
	)
}

// fragmentReservedCharacters lists the metacharacters rejected in fragments
// composed into trusted DateTimeValues: the package inserts the only
// structural separators itself ("@" between composite segments), so both it
// and the clause separator must come from the caller verbatim. Commas are
// deliberately absent — they are ordinary literals inside JavaScript
// expressions (for example, multi-argument gs.* calls) and labels.
const fragmentReservedCharacters = clauseSeparator + pairSeparator //nolint:gochecknoglobals

// validateQueryFragment checks that a caller-supplied fragment composed into a
// trusted DateTimeValue (JavaScript expressions, OnSpecialty labels and
// expressions) contains no reserved fragment characters. Callers turn a
// non-nil error into an error Condition via [NewErrorCondition] so it
// surfaces through [Condition.Error].
func validateQueryFragment(name, fragment string) error {
	index := strings.IndexAny(fragment, fragmentReservedCharacters)
	if index < 0 {
		return nil
	}

	return fmt.Errorf(
		"%s %q contains reserved encoded-query character %q; it structures the query itself and cannot be escaped",
		name, fragment, fragment[index],
	)
}
