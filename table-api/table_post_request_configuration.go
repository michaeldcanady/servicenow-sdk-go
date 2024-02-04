package tableapi

import "github.com/RecoLabs/servicenow-sdk-go/core"

// Deprecated: deprecated since v1.4.0. removed from public API.
//
// TablePostRequestConfiguration represents request configurations POST request.
type TablePostRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableRequestBuilderPostQueryParameters
	Data            map[string]string
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *TablePostRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
