//go:build preview.query

package query

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

// QueryBuilder represents the assembling of a new Service-Now query.
type QueryBuilder struct {
	// query the query being built.
	query ast.Node
	// logicalOperator the logical operator to join the next condition to the current query.
	logicalOperator ast.Operator
	// Error errors that arose during the building process.
	Error error
}

// NewQueryBuilder instantiates a new query builder.
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		query:           nil,
		logicalOperator: ast.OperatorUnknown,
		Error:           nil,
	}
}

// NewQuery The entry point for a new query.
func NewQuery() *FieldBuilder {
	qb := NewQueryBuilder()

	return NewFieldBuilder(qb)
}

// And adds an and logical operator to join the new condition.
func (builder *QueryBuilder) And() *FieldBuilder {
	return builder.setLogicalOperator(ast.OperatorAnd)
}

// Or adds an or logical operator to join the new condition.
func (builder *QueryBuilder) Or() *FieldBuilder {
	return builder.setLogicalOperator(ast.OperatorOr)
}

// setLogicalOperator sets the current logical operator to join the new condition.
func (builder *QueryBuilder) setLogicalOperator(operator ast.Operator) *FieldBuilder {
	if builder.logicalOperator != ast.OperatorUnknown {
		builder.Error = errors.Join(builder.Error, errors.New("logicalOperator already is set"))
	}
	builder.logicalOperator = operator
	return NewFieldBuilder(builder)
}

// addCondition appends the condition to the end of the current query.
func (builder *QueryBuilder) addCondition(condition ast.Node) *QueryBuilder {
	if builder.query == nil {
		builder.query = condition
		return builder
	}

	if builder.logicalOperator == ast.OperatorUnknown {
		builder.Error = errors.Join(builder.Error, errors.New("logicalOperator is unset"))
		return builder
	}

	builder.query = ast.NewBinaryNode(builder.query, builder.logicalOperator, condition)
	return builder
}

// String returns the assembled query as a string.
func (builder *QueryBuilder) String() string {
	visitor := ast.NewStringerVisitor()

	visitor.Visit(builder.query)

	return visitor.String()
}

func (builder *QueryBuilder) addErrors(errs ...error) {
	errs = append([]error{builder.Error}, errs...)
	builder.Error = errors.Join(errs...)
}

// Build returns the assembled query.
func (builder *QueryBuilder) Build() (ast.Node, error) {
	if builder.Error != nil {
		return nil, builder.Error
	}
	return builder.query, nil
}
