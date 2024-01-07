package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v{version}. Use `tableGetRequestConfiguration3[T]` instead.
//
// tableGetRequestConfiguration2[T] represents T entry request configurations GET request.
type tableGetRequestConfiguration2[T Entry] struct {
	header   interface{}
	query    *TableRequestBuilderGetQueryParameters
	data     interface{}
	mapping  core.ErrorMapping
	response *TableCollectionResponse2[T]
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *tableGetRequestConfiguration2[T]) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.header,
		QueryParameters: rC.query,
		Data:            rC.data,
		ErrorMapping:    rC.mapping,
		Response:        rC.response,
	}
}
