package query

type FieldBuilder struct {
	query *QueryBuilder
}

func NewFieldBuilder(query *QueryBuilder) *FieldBuilder {
	return &FieldBuilder{
		query: query,
	}
}

func (builder *FieldBuilder) Field(name string) *ConditionBuilder {
	return NewConditionBuilder(name, builder.query)
}

func (builder *FieldBuilder) StringField(name string) *StringConditionBuilder {
	return NewStringConditionBuilder(name, builder.query)
}

func (builder *FieldBuilder) DateTimeField(name string) *DateTimeConditionBuilder {
	return NewDateTimeConditionBuilder(name, builder.query)
}

func (builder *FieldBuilder) NumericField(name string) *NumericConditionBuilder {
	return NewNumericConditionBuilder(name, builder.query)
}
