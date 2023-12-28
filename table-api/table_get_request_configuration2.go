package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// TableGetRequestConfiguration2 represents request configurations for GET request.
type TableGetRequestConfiguration2[T Entry] struct {
	Header          interface{}
	QueryParameters *TableRequestBuilderGetQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableCollectionResponse2[T]
}

func (rC *TableGetRequestConfiguration2[T]) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
