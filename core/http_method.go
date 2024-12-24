package core

// Deprecated: deprecated since v{unreleased}.
// Represents the HTTP method used by a request.
type HttpMethod int //nolint:stylecheck

const (
	// Deprecated: deprecated since v{unreleased}.
	// The HTTP GET method.
	GET HttpMethod = iota
	// Deprecated: deprecated since v{unreleased}.
	// The HTTP POST method.
	POST
	// Deprecated: deprecated since v{unreleased}.
	// The HTTP PATCH method.
	PATCH
	// Deprecated: deprecated since v{unreleased}.
	// The HTTP DELETE method.
	DELETE
	// Deprecated: deprecated since v{unreleased}.
	// The HTTP OPTIONS method.
	OPTIONS
	// Deprecated: deprecated since v{unreleased}.
	// The HTTP CONNECT method.
	CONNECT
	// Deprecated: deprecated since v{unreleased}.
	// The HTTP PUT method.
	PUT
	// Deprecated: deprecated since v{unreleased}.
	// The HTTP TRACE method.
	TRACE
	// Deprecated: deprecated since v{unreleased}.
	// The HTTP HEAD method.
	HEAD
)

// String returns the string representation of the HTTP method.
func (m HttpMethod) String() string {
	return []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS", "CONNECT", "PUT", "TRACE", "HEAD"}[m]
}
