//go:build preview.query

package ast

// Primitive represents any base type.
type Primitive interface {
	Numeric | ~string | ~bool
}
