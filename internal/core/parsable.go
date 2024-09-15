package core

import "net/http"

type Parsable[T Response] func(*http.Response) (T, error)
