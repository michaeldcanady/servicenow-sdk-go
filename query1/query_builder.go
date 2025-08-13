package query

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

type QueryBuilder struct {
	query ast.Node
	// and/or
	logicalOperator ast.Operator
	Error           error
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		query:           nil,
		logicalOperator: ast.OperatorAnd,
		Error:           nil,
	}
}

func NewQuery() *FieldBuilder {
	qb := NewQueryBuilder()

	return NewFieldBuilder(qb)
}

func (builder *QueryBuilder) And() *FieldBuilder {
	return builder.setLogicalOperator(ast.OperatorAnd)
}

func (builder *QueryBuilder) Or() *FieldBuilder {
	return builder.setLogicalOperator(ast.OperatorOr)
}

func (builder *QueryBuilder) setLogicalOperator(operator ast.Operator) *FieldBuilder {
	if builder.logicalOperator != ast.OperatorUnknown {
		builder.Error = errors.New("logicalOperator is set")
	}
	builder.logicalOperator = operator
	return NewFieldBuilder(builder)
}

func (builder *QueryBuilder) Field(name string) *ConditionBuilder {
	return NewConditionBuilder(name, builder)
}

func (builder *QueryBuilder) addCondition(condition ast.Node) *QueryBuilder {
	if builder.logicalOperator == ast.OperatorUnknown {
		builder.Error = errors.New("logicalOperator is unset")
		return builder
	}

	if builder.query == nil {
		builder.query = condition
		return builder
	}
	builder.query = &ast.BinaryNode{
		LeftExpression:  builder.query,
		Operator:        builder.logicalOperator,
		RightExpression: condition,
	}
	return builder
}

func (builder *QueryBuilder) Build() string {
	visitor := NewStringerVisitor()

	visitor.Visit(builder.query)

	return visitor.String()
}
