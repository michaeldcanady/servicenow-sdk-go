package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// TableGetRequestConfiguration2[T] represents T entry request configurations GET request.
type TableGetRequestConfiguration2[T Entry] struct {
	header   interface{}
	query    *TableRequestBuilderGetQueryParameters
	data     interface{}
	mapping  core.ErrorMapping
	response *TableCollectionResponse2[T]
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *TableGetRequestConfiguration2[T]) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.header,
		QueryParameters: rC.query,
		Data:            rC.data,
		ErrorMapping:    rC.mapping,
		Response:        rC.response,
	}
}
