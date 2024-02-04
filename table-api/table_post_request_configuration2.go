package tableapi

import "github.com/RecoLabs/servicenow-sdk-go/core"

// tablePostRequestConfiguration2[T] represents request configurations POST request.
type tablePostRequestConfiguration2[T Entry] struct {
	header   interface{}
	query    *TableRequestBuilderPostQueryParameters
	data     map[string]string
	mapping  core.ErrorMapping
	response *TableItemResponse
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *tablePostRequestConfiguration2[T]) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.header,
		QueryParameters: rC.query,
		Data:            rC.data,
		ErrorMapping:    rC.mapping,
		Response:        rC.response,
	}
}
