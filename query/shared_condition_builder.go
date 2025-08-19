//go:build preview.query

package query

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/ast"
)

// SharedConditionBuilder represents the base for all conditions, and all operators supported by all types.
type SharedConditionBuilder[T Primitive | time.Time | any] struct {
	ConditionBuilder[QueryBuilder]
}

// NewSharedConditionBuilder instantiates a new base condition builder.
func NewSharedConditionBuilder[T Primitive | time.Time | any](field string, query *QueryBuilder) *SharedConditionBuilder[T] {
	return &SharedConditionBuilder[T]{
		NewBaseConditionBuilder[T](field, query),
	}
}

// Is query that field is the provided value.
func (builder *SharedConditionBuilder[T]) Is(value T) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIs, ast.NewLiteralNode(value))
}

// IsNot query that field is not the provided value.
func (builder *SharedConditionBuilder[T]) IsNot(value T) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsNot, ast.NewLiteralNode(value))
}

// IsEmpty query that field is empty.
func (builder *SharedConditionBuilder[T]) IsEmpty() *QueryBuilder {
	return builder.unaryCondition(ast.OperatorIsEmpty)
}

// IsNotsEmpty query that field is not empty.
func (builder *SharedConditionBuilder[T]) IsNotsEmpty() *QueryBuilder {
	return builder.unaryCondition(ast.OperatorIsNotEmpty)
}

// IsAnything query that field is anything.
func (builder *SharedConditionBuilder[T]) IsAnything() *QueryBuilder {
	return builder.unaryCondition(ast.OperatorIsAnything)
}

// IsDynamic query that field is dynamically the sysID provided.
func (builder *SharedConditionBuilder[T]) IsDynamic(sysID string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsNot, ast.NewLiteralNode(sysID))
}

// IsSame query that field is the same the provided value.
func (builder *SharedConditionBuilder[T]) IsSame(sysID string) *QueryBuilder {
	return builder.binaryCondition(ast.OperatorIsSame, ast.NewLiteralNode(sysID))
}
