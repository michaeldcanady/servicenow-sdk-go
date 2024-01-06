package core

import "net/http"

// CollectionResponse[T] represents collection responses.
type CollectionResponse[T any] interface {
	ToPage() PageResult[T]
	ParseHeaders(http.Header)
}
