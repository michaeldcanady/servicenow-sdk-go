package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

type TableGetRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableRequestBuilderGetQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableCollectionResponse
}

func (rC *TableGetRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
