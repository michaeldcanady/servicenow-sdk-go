package core

import "net/http"

// Deprecated: deprecated since v{unreleased}.
// CollectionResponse[T] represents collection responses.
type CollectionResponse[T any] interface {
	ToPage() PageResult[T]
	ParseHeaders(http.Header)
}
