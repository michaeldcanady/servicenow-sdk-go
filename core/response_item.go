package core

import "net/http"

// Deprecated: deprecated since v{unreleased}.
// ItemResponse[T] represents a single item responses.
type ItemResponse[T any] interface {
	ParseHeaders(http.Header)
}
