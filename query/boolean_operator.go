package query

type BooleanOperator int64

const (
	BooleanOperatorIs BooleanOperator = iota
	BooleanOperatorIsNot
	BooleanOperatorIsEmpty
	BooleanOperatorIsNotEmpty
	BooleanOperatorIsAnything
	BooleanOperatorIsSame
	BooleanOperatorIsDifferent
)
