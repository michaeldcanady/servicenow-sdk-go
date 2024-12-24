package core

import (
	"net/http"
)

// Deprecated: deprecated since v{unreleased}.
// Response represents all possible responses
type Response interface {
	ParseHeaders(headers http.Header)
}
