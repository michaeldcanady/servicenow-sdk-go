//go:build preview.query

package ast

// Primitive represents any base type.
type Primitive interface {
	Numeric | ~string | ~bool
}

// Numeric represents a number value
type Numeric interface {
	Int | Float | Uint
}

// Int represents an integer type.
type Int interface {
	~int | ~int16 | ~int32 | ~int64 | ~int8
}

// Float represents a float type.
type Float interface {
	~float32 | ~float64
}

// Uint represents an unsigned integer type.
type Uint interface {
	~uint | ~uint16 | ~uint32 | ~uint64 | ~uint8
}
