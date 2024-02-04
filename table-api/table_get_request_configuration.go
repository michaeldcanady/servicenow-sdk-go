package tableapi

import "github.com/RecoLabs/servicenow-sdk-go/core"

// Deprecated: deprecated since v1.4.0. Removed from public API.
//
// TableGetRequestConfiguration represents request configurations GET request.
type TableGetRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableRequestBuilderGetQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableCollectionResponse
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *TableGetRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
