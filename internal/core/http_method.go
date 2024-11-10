package core

// Represents the HTTP method used by a request.
type HttpMethod int //nolint:stylecheck

const (
	// The HTTP MethodGet method.
	MethodGet HttpMethod = iota
	// The HTTP MethodPost method.
	MethodPost
	// The HTTP MethodPatch method.
	MethodPatch
	// The HTTP MethodDelete method.
	MethodDelete
	// The HTTP MethodOptions method.
	MethodOptions
	// The HTTP MethodConnect method.
	MethodConnect
	// The HTTP MethodPut method.
	MethodPut
	// The HTTP MethodTrace method.
	MethodTrace
	// The HTTP MethodHead method.
	MethodHead
)

// String returns the string representation of the HTTP method.
func (m HttpMethod) String() string {
	return []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS", "CONNECT", "PUT", "TRACE", "HEAD"}[m]
}
