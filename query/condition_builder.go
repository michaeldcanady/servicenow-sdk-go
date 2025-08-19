//go:build preview.query

package query

type ConditionBuilder[T QueryBuilder] interface {
	UnaryConditionBuilder[T]
	BinaryConditionBuilder[T]
	ErrorAdder
}
