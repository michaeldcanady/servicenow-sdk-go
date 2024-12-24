package core

import "net/http"

// Deprecated: deprecated since v{unreleased}.
//
// RequestHeader
type RequestHeader interface {
	Set(key, value string)
	Get(key string) string
	SetAll(headers http.Header)
	Iterate(callback func(string, []string) bool)
}

// RequestHeader is a type alias for http request headers
type requestHeader struct {
	header http.Header
}

func NewRequestHeader() RequestHeader {
	return &requestHeader{
		header: http.Header{},
	}
}

func (rH *requestHeader) Set(key, value string) {
	rH.header.Set(key, value)
}

func (rH *requestHeader) Get(key string) string {
	return rH.header.Get(key)
}

func (rH *requestHeader) SetAll(headers http.Header) {
	rH.header = headers.Clone()
}

func (rH *requestHeader) Iterate(callback func(string, []string) bool) {
	for key, value := range rH.header {
		keepIterating := callback(key, value)
		if !keepIterating {
			break
		}
	}
}
