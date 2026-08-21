package query

import (
	"fmt"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/ast"
)

// reservedQueryCharacters lists the ServiceNow encoded-query metacharacters that
// cannot appear inside a condition value: "^" separates clauses ("^" / "^OR"),
// "," separates values in a list, and "@" separates the halves of a pair.
// Encoded queries provide no escape sequence for them, so a value containing
// any of these characters would break out of its term and append arbitrary
// clauses to the query.
const reservedQueryCharacters = "^,@" //nolint:gochecknoglobals

// validateQueryValue checks that a consumer-supplied value contains no reserved
// encoded-query characters. It returns nil for safe values and a descriptive
// error otherwise; callers turn that error into an error Condition via
// [NewErrorCondition] so it surfaces through [Condition.Error].
//
// Literals composed internally by this package (for example the
// "label@javascript:expr@javascript:expr" strings built by OnSpecialty) are
// trusted and must bypass this check.
func validateQueryValue(field string, op ast.Operator, val any) error {
	rendered := fmt.Sprintf("%v", val)
	index := strings.IndexAny(rendered, reservedQueryCharacters)
	if index < 0 {
		return nil
	}

	return fmt.Errorf(
		"value %q for field %q (operator %q) contains reserved encoded-query character %q; ^, ,, and @ structure the query itself and cannot be escaped",
		rendered, field, op.String(), rendered[index],
	)
}

// validateQueryFragment checks that a caller-supplied fragment composed into a
// trusted DateTimeValue (JavaScript expressions, OnSpecialty labels and
// expressions) contains no reserved encoded-query characters. Unlike plain
// values, fragments may not contain any of them — including "@": the package
// inserts the only structural separators itself. Callers turn a non-nil error
// into an error Condition via [NewErrorCondition] so it surfaces through
// [Condition.Error].
func validateQueryFragment(name, fragment string) error {
	index := strings.IndexAny(fragment, reservedQueryCharacters)
	if index < 0 {
		return nil
	}

	return fmt.Errorf(
		"%s %q contains reserved encoded-query character %q; ^, ,, and @ structure the query itself and cannot be escaped",
		name, fragment, fragment[index],
	)
}
