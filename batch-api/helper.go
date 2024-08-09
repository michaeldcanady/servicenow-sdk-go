package batchapi

import (
	"net/url"
)

type hasBaseURL interface {
	GetBaseURL() string
}

func getBaseURL(adapter hasBaseURL) (*url.URL, error) {
	return url.Parse(adapter.GetBaseURL())
}
