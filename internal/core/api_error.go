package core

import "fmt"

// ApiError represents an error that occurs during API requests.
type ApiError struct { //nolint:stylecheck
	// Message is the human-readable error message.
	Message string
	// ResponseStatusCode is the HTTP response status code associated with the error.
	ResponseStatusCode int
}

// Error returns the error message as a string. If the message is empty, it returns a default error message.
func (e *ApiError) Error() string {
	if len(e.Message) > 0 {
		return fmt.Sprint(e.Message)
	} else {
		return "error status code received from the API"
	}
}
