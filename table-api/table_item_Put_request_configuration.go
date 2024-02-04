package tableapi

import "github.com/RecoLabs/servicenow-sdk-go/core"

// Deprecated: deprecated since v1.4.0. Removed from public API.
//
// TableItemPutRequestConfiguration represents request configurations GET request.
type TableItemPutRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableItemRequestBuilderPutQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *TableItemPutRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
