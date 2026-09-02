// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package ast

import "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"

//https://www.servicenow.com/docs/bundle/vancouver-platform-user-interface/page/use/common-ui-elements/reference/r_OpAvailableFiltersQueries.html

// Operators for conditions
type Operator int64

const (
	OperatorUnknown Operator = iota - 1
	// all
	OperatorIs
	OperatorIsNot
	OperatorIsEmpty
	OperatorIsNotEmpty
	OperatorIsAnything
	// string
	OperatorStartsWith
	OperatorEndsWith
	OperatorContains
	OperatorIsEmptyString
	OperatorDoesNotContain
	OperatorMatchesPattern
	OperatorIsInHierarchy
	// field
	OperatorIsDynamic
	OperatorIsSame
	// date/time
	OperatorOn
	OperatorNotOn
	OperatorBefore
	OperatorAtOrBefore
	OperatorAfter
	OperatorAtOrAfter
	OperatorTrendOnOrAfter
	OperatorTrendOnOrBefore
	OperatorTrendAfter
	OperatorTrendBefore
	OperatorTrendOn
	OperatorRelativeAfter
	OperatorRelativeBefore
	// numeric
	OperatorLessThan
	OperatorGreaterThan
	OperatorLessThanOrIs
	OperatorGreaterThanOrIs
	OperatorBetween
	OperatorIsDifferent
	OperatorGreaterThanField
	OperatorLessThanField
	OperatorGreaterThanOrIsField
	OperatorLessThanOrIsField
	OperatorIsMoreThan
	OperatorIsLessThan
	OperatorIsOneOf
	OperatorIsNotOneOf
	// logical
	OperatorAnd
	OperatorOr
)

var operatorStrings = map[Operator]string{
	OperatorUnknown:              "unknown",
	OperatorIs:                   "=",
	OperatorIsNot:                "!=",
	OperatorIsEmpty:              "ISEMPTY",
	OperatorIsNotEmpty:           "ISNOTEMPTY",
	OperatorLessThan:             "<",
	OperatorGreaterThan:          ">",
	OperatorLessThanOrIs:         "<=",
	OperatorGreaterThanOrIs:      ">=",
	OperatorBetween:              "BETWEEN",
	OperatorIsAnything:           "ANYTHING",
	OperatorIsSame:               "SAMEAS",
	OperatorIsDifferent:          "NSAMEAS",
	OperatorGreaterThanField:     "GT_FIELD",
	OperatorLessThanField:        "LT_FIELD",
	OperatorGreaterThanOrIsField: "GT_OR_EQUALS_FIELD",
	OperatorLessThanOrIsField:    "LT_OR_EQUALS_FIELD",
	OperatorOn:                   "ON",
	OperatorNotOn:                "NOTON",
	OperatorBefore:               "<",
	OperatorAtOrBefore:           "<=",
	OperatorAfter:                ">",
	OperatorAtOrAfter:            ">=",
	OperatorTrendOnOrAfter:       "DATEPART",
	OperatorTrendOnOrBefore:      "DATEPART",
	OperatorTrendAfter:           "DATEPART",
	OperatorTrendBefore:          "DATEPART",
	OperatorTrendOn:              "DATEPART",
	OperatorRelativeAfter:        "DATEPART",
	OperatorRelativeBefore:       "DATEPART",
	OperatorIsMoreThan:           "MORETHAN",
	OperatorIsLessThan:           "LESSTHAN",
	OperatorIsOneOf:              "IN",
	OperatorIsNotOneOf:           "NOT IN",
	OperatorStartsWith:           "STARTSWITH",
	OperatorEndsWith:             "ENDSWITH",
	OperatorContains:             "LIKE",
	OperatorIsEmptyString:        "EMPTYSTRING",
	OperatorIsDynamic:            "DYNAMIC",
	OperatorDoesNotContain:       "NOT LIKE",
	OperatorMatchesPattern:       "MATCHES PATTERN",
	OperatorIsInHierarchy:        "IN HIERARCHY",
	OperatorAnd:                  "^",
	OperatorOr:                   "^OR",
}

func (o Operator) String() string {
	return conversion.EnumString(operatorStrings, o, operatorStrings[OperatorUnknown])
}

// known reports whether o has a defined encoded-query rendering. Operators
// outside the declared set (including OperatorUnknown) render as "unknown"
// via String, which would corrupt a query silently; traversal treats them as
// structural faults instead.
func (o Operator) known() bool {
	if o == OperatorUnknown {
		return false
	}
	_, ok := operatorStrings[o]
	return ok
}
