package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v1.4.0. removed from public API.
//
// TablePostRequestConfiguration represents request configurations POST request.
type TablePostRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableRequestBuilderPostQueryParameters
	Data            map[string]string
	ErrorMapping    core.ErrorMapping //nolint: staticcheck
	response        *TableItemResponse
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *TablePostRequestConfiguration) toConfiguration() *core.RequestConfiguration { //nolint: staticcheck
	return &core.RequestConfiguration{ //nolint: staticcheck
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
