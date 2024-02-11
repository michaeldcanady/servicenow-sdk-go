package internal

import (
	"net/http"
)

// RequestHeader is a type alias for http request headers
type RequestHeader struct {
	header http.Header
}

func NewRequestHeader() RequestHeader {
	return RequestHeader{
		header: http.Header{},
	}
}

func (rH *RequestHeader) Set(key, value string) {
	rH.header.Set(key, value)
}

func (rH *RequestHeader) Get(key string) string {
	return rH.header.Get(key)
}

func (rH *RequestHeader) SetAll(headers http.Header) {
	rH.header = headers.Clone()
}

func (rH *RequestHeader) Iterate(callback func(string, []string) bool) {
	for key, value := range rH.header {
		keepIterating := callback(key, value)
		if !keepIterating {
			break
		}
	}
}
