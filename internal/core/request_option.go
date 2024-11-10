package core

// Represents a request option.
type RequestOption interface {
	// GetKey returns the key to store the current option under.
	GetKey() RequestOptionKey
}
