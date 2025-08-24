//go:build preview.query

package query

// ErrorAdder represents a type that stores an extendable error.
type ErrorAdder interface {
	// addErrors adds the provided errors to the existing error.
	addErrors(...error)
}
