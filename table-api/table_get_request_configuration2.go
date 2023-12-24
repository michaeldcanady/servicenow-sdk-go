package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

type TableGetRequestConfiguration2[T TableEntry2] struct {
	Header          interface{}
	QueryParameters *TableRequestBuilderGetQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableCollectionResponse[T]
}

func (rC *TableGetRequestConfiguration[T]) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}