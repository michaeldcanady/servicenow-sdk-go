//go:build preview.query

package query2

import (
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast2"
)

// NumberField represents a numeric field in ServiceNow.
type NumberField struct {
	BaseField
}

// Is query that field is the provided value.
func (f NumberField) Is(val float64) Condition {
	return f.binary(ast2.OperatorIs, val)
}

// IsNot query that field is not the provided value.
func (f NumberField) IsNot(val float64) Condition {
	return f.binary(ast2.OperatorIsNot, val)
}

// LessThan query that field is less than the provided value.
func (f NumberField) LessThan(val float64) Condition {
	return f.binary(ast2.OperatorLessThan, val)
}

// GreaterThan query that field is greater than the provided value.
func (f NumberField) GreaterThan(val float64) Condition {
	return f.binary(ast2.OperatorGreaterThan, val)
}

// LessThanOrIs query that field is less than or is the provided value.
func (f NumberField) LessThanOrIs(val float64) Condition {
	return f.binary(ast2.OperatorLessThanOrIs, val)
}

// GreaterThanOrIs query that field is greater than or is the provided value.
func (f NumberField) GreaterThanOrIs(val float64) Condition {
	return f.binary(ast2.OperatorGreaterThanOrIs, val)
}

// Between query that field is between the provided lower and upper values.
func (f NumberField) Between(lower, upper float64) Condition {
	if lower >= upper {
		return NewErrorCondition(fmt.Errorf("%v is greater or equal to %v", lower, upper))
	}
	return f.pair(ast2.OperatorBetween, lower, upper)
}

// IsOneOf query that field is one of the provided values.
func (f NumberField) IsOneOf(values ...float64) Condition {
	return f.multi(ast2.OperatorIsOneOf, convertSliceToArrayNode(values...))
}

// IsNotOneOf query that field is not one of the provided values.
func (f NumberField) IsNotOneOf(values ...float64) Condition {
	return f.multi(ast2.OperatorIsNotOneOf, convertSliceToArrayNode(values...))
}

// GreaterThanField query that field is greater than the provided field.
func (f NumberField) GreaterThanField(otherField string) Condition {
	return f.binary(ast2.OperatorGreaterThanField, otherField)
}

// LessThanField query that field is less than the provided field.
func (f NumberField) LessThanField(otherField string) Condition {
	return f.binary(ast2.OperatorLessThanField, otherField)
}

// GreaterThanOrIsField query that field is greater than or is the provided field.
func (f NumberField) GreaterThanOrIsField(otherField string) Condition {
	return f.binary(ast2.OperatorGreaterThanOrIsField, otherField)
}

// LessThanOrIsField query that field is less than or is the provided field.
func (f NumberField) LessThanOrIsField(otherField string) Condition {
	return f.binary(ast2.OperatorLessThanOrIsField, otherField)
}

// IsMoreThan query that field is more than the provided value.
func (f NumberField) IsMoreThan(val float64) Condition {
	return f.binary(ast2.OperatorIsMoreThan, val)
}

// IsLessThan query that field is less than the provided value.
func (f NumberField) IsLessThan(val float64) Condition {
	return f.binary(ast2.OperatorIsLessThan, val)
}
