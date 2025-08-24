//go:build preview.query

package query

// Int represents an integer type.
type Int interface {
	~int | ~int32 | ~int64
}
