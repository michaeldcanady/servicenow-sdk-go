//go:build preview.query

package query

// Primitive represents any base type.
type Primitive interface {
	Numeric | ~string
}
