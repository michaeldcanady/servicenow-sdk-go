//go:build preview.query

package query

type UnitedConditionBuilder struct {
	*StringConditionBuilder
	*DateTimeConditionBuilder
	*NumericConditionBuilder
}

func NewUnitedConditionBuilder(field string, query *QueryBuilder) *UnitedConditionBuilder {
	return &UnitedConditionBuilder{
		NewStringConditionBuilder(field, query),
		NewDateTimeConditionBuilder(field, query),
		NewNumericConditionBuilder(field, query),
	}
}
