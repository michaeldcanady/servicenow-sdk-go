package core

import "net/http"

type Parsable[T any] func(*http.Response) (CollectionResponse[T], error)
