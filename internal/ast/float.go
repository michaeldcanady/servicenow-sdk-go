//go:build preview.query

package ast

// Float represents any float type.
type Float interface {
	~float32 | ~float64
}
