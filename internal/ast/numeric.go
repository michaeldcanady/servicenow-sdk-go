//go:build preview.query

package ast

// Numeric represents a number value
type Numeric interface {
	Int | Float | Uint
}
