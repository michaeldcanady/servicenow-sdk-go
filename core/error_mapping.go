package core

import "strconv"

// ErrorMapping is a map that maps error codes to human-readable error messages.
type ErrorMapping map[string]string

// NewErrorMapping creates a new instance of the ErrorMapping.
// It initializes an empty map to store error code to error message mappings.
func NewErrorMapping() ErrorMapping {
	return ErrorMapping(make(map[string]string))
}

// Set sets the error message for the specified error code.
// It associates the given error code with the provided error message in the mapping.
func (eM ErrorMapping) Set(code, err string) {
	eM[code] = err
}

// Len returns the number of error code to error message mappings in the ErrorMapping.
// It provides the total count of error mappings.
func (eM ErrorMapping) Len() int {
	return len(eM)
}

// Get retrieves the error message associated with the provided error code.
// It returns the error message along with a boolean indicating whether the error code was found.
// If the error code is not found, an empty string is returned along with `false`.
func (eM ErrorMapping) Get(code int) (string, bool) {
	// Check if the provided code is within the error code range (400 or greater).
	if code <= 399 {
		return "", false
	}

	statusAsString := strconv.Itoa(code)

	// Try to find an exact match for the provided error code.
	if msg, found := eM[statusAsString]; found {
		return msg, true
	}

	// Try to find a relative match (e.g., 4XX for client errors or 5XX for server errors).
	if code >= 500 && code < 600 && eM["5XX"] != "" {
		return eM["5XX"], true
	} else if code >= 400 && code < 500 && eM["4XX"] != "" {
		return eM["4XX"], true
	}

	// If no match is found, return an empty string and `false`.
	return "", false
}
