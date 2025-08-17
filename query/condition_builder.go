//go:build preview.query

package query

type ConditionBuilder struct {
	*BaseConditionBuilder[any]
	*StringConditionBuilder
	*DateTimeConditionBuilder
	*NumericConditionBuilder
}

func NewConditionBuilder(field string, query *QueryBuilder) *ConditionBuilder {
	return &ConditionBuilder{
		NewBaseConditionBuilder[any](field, query),
		NewStringConditionBuilder(field, query),
		NewDateTimeConditionBuilder(field, query),
		NewNumericConditionBuilder(field, query),
	}
}
