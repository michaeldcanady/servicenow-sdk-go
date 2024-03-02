package internal

import (
	"net/http"
)

// Response represents all possible responses
type Response interface {
	ParseHeaders(headers http.Header)
}
