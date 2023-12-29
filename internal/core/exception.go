package core

// Exception represents an exception in the ServiceNow error response.
type Exception struct {
	// Detail is a detailed description of the exception.
	Detail string
	// Message is a brief description of the exception.
	Message string
}
