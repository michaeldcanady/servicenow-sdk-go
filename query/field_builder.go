//go:build preview.query

package query

// FieldBuilder represents the selection of a specific type of field.
type FieldBuilder struct {
	// query the current query.
	query *QueryBuilder
}

// NewFieldBuilder instantiates a new field builder.
func NewFieldBuilder(query *QueryBuilder) *FieldBuilder {
	return &FieldBuilder{
		query: query,
	}
}

// Field a generic field.
func (builder *FieldBuilder) Field(name string) *ConditionBuilder {
	return NewConditionBuilder(name, builder.query)
}

// StringField a string field.
func (builder *FieldBuilder) StringField(name string) *StringConditionBuilder {
	return NewStringConditionBuilder(name, builder.query)
}

// DateTimeField a date-time field.
func (builder *FieldBuilder) DateTimeField(name string) *DateTimeConditionBuilder {
	return NewDateTimeConditionBuilder(name, builder.query)
}

// NumericField a numeric field.
func (builder *FieldBuilder) NumericField(name string) *NumericConditionBuilder {
	return NewNumericConditionBuilder(name, builder.query)
}
