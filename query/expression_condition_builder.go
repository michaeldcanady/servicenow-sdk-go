//go:build preview.query

package query

type ExpressionConditionBuilder[T QueryBuilder] interface {
	UnaryConditionBuilder[T]
	BinaryConditionBuilder[T]
	ErrorBuilder
}
