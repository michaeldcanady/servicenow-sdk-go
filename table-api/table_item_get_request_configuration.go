package tableapi

import "github.com/RecoLabs/servicenow-sdk-go/core"

// Deprecated: deprecated since v1.4.0. Removed from public API.
//
// TableItemGetRequestConfiguration represents request configurations GET request.
type TableItemGetRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableItemRequestBuilderGetQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *TableItemGetRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
