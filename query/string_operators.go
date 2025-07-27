package query

type StringOperator int64

const (
	StringOperatorStartWith StringOperator = iota
	StringOperatorEndsWith
	StringOperatorContains
	StringOperatorDoesNotContain
	StringOperatorIs
	StringOperatorIsNot
	StringOperatorIsEmpty
	StringOperatorIsNotEmpty
	StringOperatorIsAnything
	StringOperatorIsEmptyString
	StringOperatorLessThanOrIs
	StringOperatorGreaterThanOrIs
	StringOperatorBetween
	StringOperatorIsSame
	StringOperatorIsDifferent
)
