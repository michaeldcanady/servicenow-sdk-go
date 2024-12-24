package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v1.4.0. Removed from public API.
//
// TableItemDeleteRequestConfiguration represents request configurations DELETE request.
type TableItemDeleteRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableItemRequestBuilderDeleteQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping //nolint: staticcheck
	response        *TableItemResponse
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *TableItemDeleteRequestConfiguration) toConfiguration() *core.RequestConfiguration { //nolint: staticcheck
	return &core.RequestConfiguration{ //nolint: staticcheck
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
