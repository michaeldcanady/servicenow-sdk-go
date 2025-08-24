//go:build preview.query

package query

// Float represents any float type.
type Float interface {
	~float32 | ~float64
}
