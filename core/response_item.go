package core

import "net/http"

// ItemResponse[T] represents a single item responses.
type ItemResponse[T any] interface {
	ParseHeaders(http.Header)
}
