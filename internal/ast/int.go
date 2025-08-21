//go:build preview.query

package ast

// Int represents an integer type.
type Int interface {
	~int | ~int16 | ~int32 | ~int64 | ~int8
}
