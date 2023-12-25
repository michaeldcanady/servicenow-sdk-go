package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

type TableItemGetRequestConfiguration2[T TableEntry2] struct {
	Header          interface{}
	QueryParameters *TableItemRequestBuilderGetQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse2[T]
}

func (rC *TableItemGetRequestConfiguration2[T]) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
