package core

import "net/http"

// CollectionResponse[T] represents collection responses.
type CollectionResponse[T any] interface {
	ParseHeaders(http.Header)
	Results() []T
}
