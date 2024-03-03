package batchapi

import (
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

func getBaseURL(adapter core.Client2) (*url.URL, error) {
	return url.Parse(adapter.GetBaseURL())
}
