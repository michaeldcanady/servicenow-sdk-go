package core

import "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

// Deprecated: deprecated since v{version}. Will be removed from public API
type RequestBuilder = core.RequestBuilder

// Deprecated: deprecated since v{version}. Will be removed from public API
func NewRequestBuilder(client core.Client, urlTemplate string, pathParameters map[string]string) *RequestBuilder {
	return core.NewRequestBuilder(client, urlTemplate, pathParameters)
}
