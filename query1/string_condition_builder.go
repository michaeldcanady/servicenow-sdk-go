package query

import (
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

type StringConditionBuilder struct {
	*BaseConditionBuilder[string]
}

func NewStringConditionBuilder(field string, query *QueryBuilder) *StringConditionBuilder {
	return &StringConditionBuilder{
		NewBaseConditionBuilder[string](field, query),
	}
}

func (builder *StringConditionBuilder) StartsWith(value string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorStartsWith, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *StringConditionBuilder) EndsWith(value string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorEndsWith, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *StringConditionBuilder) Contains(value string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorContains, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *StringConditionBuilder) DoesNotContain(value string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorDoesNotContain, &ast.LiteralNode{
		Value: fmt.Sprintf("%v", value),
	})
}

func (builder *StringConditionBuilder) IsEmptyString() *QueryBuilder {
	return builder.unaryCondition(ast.OperatorIsEmptyString)
}
