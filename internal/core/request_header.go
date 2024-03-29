package core

import "net/http"

type RequestHeader interface {
	Set(key, value string)
	Get(key string) string
	SetAll(headers http.Header)
	Iterate(callback func(string, []string) bool)
}
