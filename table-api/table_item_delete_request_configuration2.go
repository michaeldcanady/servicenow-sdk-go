package tableapi

import "github.com/RecoLabs/servicenow-sdk-go/core"

// tableItemDeleteRequestConfiguration2[T] represents request configurations DELETE request.
type tableItemDeleteRequestConfiguration2[T Entry] struct {
	header   interface{}
	query    *TableItemRequestBuilderDeleteQueryParameters
	data     interface{}
	mapping  core.ErrorMapping
	response *TableItemResponse2[T]
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *tableItemDeleteRequestConfiguration2[T]) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.header,
		QueryParameters: rC.query,
		Data:            rC.data,
		ErrorMapping:    rC.mapping,
		Response:        rC.response,
	}
}
