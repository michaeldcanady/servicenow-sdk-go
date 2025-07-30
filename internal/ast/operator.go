package ast

//https://www.servicenow.com/docs/bundle/vancouver-platform-user-interface/page/use/common-ui-elements/reference/r_OpAvailableFiltersQueries.html

// Operators for conditions
type Operator int64

const (
	OperatorUnknown Operator = iota - 1
	OperatorIs
	OperatorIsNot
	OperatorIsEmpty
	OperatorIsNotEmpty
	OperatorLessThan
	OperatorGreaterThan
	OperatorLessThanOrIs
	OperatorGreaterThanOrIs
	OperatorBetween
	OperatorIsAnything
	OperatorIsSame
	OperatorIsDifferent
	OperatorGreaterThanField
	OperatorLessThanField
	OperatorGreaterThanOrIsField
	OperatorLessThanOrIsField
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
	OperatorIsMoreThan
	OperatorIsLessThan
	OperatorIsOneOf
	OperatorIsNotOneOf
	OperatorStartsWith
	OperatorEndsWith
	OperatorContains
	OperatorIsEmptyString
	OperatorIsDynamic
	OperatorDoesNotContain
	OperatorStartWith
	OperatorAnd
	OperatorOr
)

func (o Operator) String() string {
	str, ok := map[Operator]string{
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
		OperatorStartWith:            "",
		OperatorAnd:                  "^",
		OperatorOr:                   "^OR",
	}[o]
	if !ok {
		return OperatorUnknown.String()
	}
	return str
}
