//go:build preview.query

package query

type UnitedConditionBuilder struct {
	*BaseConditionBuilder[any]
	*StringConditionBuilder
	*DateTimeConditionBuilder
	*NumericConditionBuilder
}

func NewUnitedConditionBuilder(field string, query *QueryBuilder) *UnitedConditionBuilder {
	return &UnitedConditionBuilder{
		NewBaseConditionBuilder[any](field, query),
		NewStringConditionBuilder(field, query),
		NewDateTimeConditionBuilder(field, query),
		NewNumericConditionBuilder(field, query),
	}
}
