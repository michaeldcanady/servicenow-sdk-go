package tableapi

import "github.com/RecoLabs/servicenow-sdk-go/core"

// tableItemGetRequestConfiguration2[T] represents request configurations GET request.
type tableItemGetRequestConfiguration2[T Entry] struct {
	header   interface{}
	query    *TableItemRequestBuilderGetQueryParameters
	data     interface{}
	mapping  core.ErrorMapping
	response *TableItemResponse2[T]
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *tableItemGetRequestConfiguration2[T]) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.header,
		QueryParameters: rC.query,
		Data:            rC.data,
		ErrorMapping:    rC.mapping,
		Response:        rC.response,
	}
}
