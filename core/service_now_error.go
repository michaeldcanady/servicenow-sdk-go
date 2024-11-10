package core

import "fmt"

// ServiceNowError represents an error response from the ServiceNow API.
type ServiceNowError struct {
	// Exception is the exception details in the error response.
	Exception Exception `json:"error"`
	// Status is the status of the error response.
	Status string
}

// Error returns a formatted error message that includes both the exception message and detail.
func (e *ServiceNowError) Error() string {
	return fmt.Sprintf("%s: %s", e.Exception.Message, e.Exception.Detail)
}
