//go:build preview.query

package query

import (
	"errors"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

// BaseConditionBuilder represents the base for all conditions, and all operators supported by all types.
type BaseConditionBuilder[T Primitive | time.Time | any] struct {
	// field the field of the query.
	field string
	// query the existing query builder.
	query *QueryBuilder
	// Error error(s) encountered during the building process.
	Error error
}

// NewBaseConditionBuilder instantiates a new base condition builder.
func NewBaseConditionBuilder[T Primitive | time.Time | any](field string, query *QueryBuilder) *BaseConditionBuilder[T] {
	return &BaseConditionBuilder[T]{
		field: field,
		query: query,
		Error: nil,
	}
}

// Is query that field is the provided value.
func (builder *BaseConditionBuilder[T]) Is(value T) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIs, ast.NewLiteralNode(value))
}

// IsNot query that field is not the provided value.
func (builder *BaseConditionBuilder[T]) IsNot(value T) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsNot, ast.NewLiteralNode(value))
}

// IsEmpty query that field is empty.
func (builder *BaseConditionBuilder[T]) IsEmpty() *QueryBuilder {
	return builder.unaryCondition(ast.OperatorIsEmpty)
}

// IsNotsEmpty query that field is not empty.
func (builder *BaseConditionBuilder[T]) IsNotsEmpty() *QueryBuilder {
	return builder.unaryCondition(ast.OperatorIsNotEmpty)
}

// IsAnything query that field is anything.
func (builder *BaseConditionBuilder[T]) IsAnything() *QueryBuilder {
	return builder.unaryCondition(ast.OperatorIsAnything)
}

// IsDynamic query that field is dynamically the sysID provided.
func (builder *BaseConditionBuilder[T]) IsDynamic(sysID string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsNot, ast.NewLiteralNode(sysID))
}

// IsSame query that field is the same the provided value.
func (builder *BaseConditionBuilder[T]) IsSame(sysID string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsSame, ast.NewLiteralNode(sysID))
}

// unaryCondition builds a new unary condition.
func (builder *BaseConditionBuilder[T]) unaryCondition(operator ast.Operator) *QueryBuilder {
	if operator == ast.OperatorUnknown {
		builder.addErrors(UnknownOperatorErr)
	}

	return builder.addCondition(ast.NewUnaryNode(operator, ast.NewLiteralNode(builder.field)))
}

// binaryCondition builds a new binary condition.
func (builder *BaseConditionBuilder[T]) binaryCondition(operator ast.Operator, value ast.Node) *QueryBuilder {
	if operator == ast.OperatorUnknown {
		builder.addErrors(UnknownOperatorErr)
	}

	return builder.addCondition(ast.NewBinaryNode(ast.NewLiteralNode(builder.field), operator, value))
}

// addErrors appends the provided errors to the existing builder errors.
func (builder *BaseConditionBuilder[T]) addErrors(errs ...error) {
	errs = append([]error{builder.Error}, errs...)
	builder.Error = errors.Join(errs...)
}

// addCondition appends the provided condition to query.
func (builder *BaseConditionBuilder[T]) addCondition(condition ast.Node) *QueryBuilder {
	builder.query.addErrors(builder.Error)

	return builder.query.addCondition(condition)
}
