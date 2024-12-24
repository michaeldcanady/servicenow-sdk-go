package batchapi

import (
	"net/url"
)

// Deprecated: deprecated since v{unreleased}.
//
// hasBaseURL ...
type hasBaseURL interface {
	GetBaseURL() string
}

func getBaseURL(adapter hasBaseURL) (*url.URL, error) {
	return url.Parse(adapter.GetBaseURL())
}
