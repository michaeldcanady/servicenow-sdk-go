package ast

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
	OperatorDosNotContain
	OperatorIsEmptyString
	OperatorIsDynamic
	OperatorDoesNotContain
	OperatorStartWith
	OperatorAnd
	OperatorOr
)

func (o Operator) String() string {
	str, ok := map[Operator]string{
		OperatorUnknown: "unknown",
		OperatorIs:      "IS",
	}[o]
	if !ok {
		return OperatorUnknown.String()
	}
	return str
}
