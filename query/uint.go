//go:build preview.query

package query

// Uint represents any unsigned integer type.
type Uint interface {
	~uint | ~uint32 | ~uint64
}
