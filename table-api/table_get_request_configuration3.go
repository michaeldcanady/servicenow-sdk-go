package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// tableGetRequestConfiguration3[T] represents T entry request configurations GET request.
type tableGetRequestConfiguration3[T Entry] struct {
	header   interface{}
	query    *TableRequestBuilderGetQueryParameters
	data     interface{}
	mapping  core.ErrorMapping
	response *TableCollectionResponse3[T]
}

func (rC *tableGetRequestConfiguration3[T]) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.header,
		QueryParameters: rC.query,
		Data:            rC.data,
		ErrorMapping:    rC.mapping,
		Response:        rC.response,
	}
}
