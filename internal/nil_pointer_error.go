package internal

// NilPointerError represents a nil pointer when a value pointer is expected.
type NilPointerError struct {
	s string
}

// NewNilPointerError instantiates a new NilPointerError.
func NewNilPointerError(text string) *NilPointerError {
	return &NilPointerError{
		s: text,
	}
}

// Error returns the error string.
func (err *NilPointerError) Error() string {
	return err.s
}
