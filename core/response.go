package core

import (
	"net/http"
)

type Response interface {
	ParseHeaders(headers http.Header)
}
