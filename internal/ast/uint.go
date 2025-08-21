//go:build preview.query

package ast

// Uint represents any unsigned integer type.
type Uint interface {
	~uint | ~uint16 | ~uint32 | ~uint64 | ~uint8
}
