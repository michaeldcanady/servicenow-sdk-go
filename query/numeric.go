//go:build preview.query

package query

// Numeric represents a number value
type Numeric interface {
	Int | Float | Uint
}
