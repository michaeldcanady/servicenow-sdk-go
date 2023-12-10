package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

type TableItemGetRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableItemRequestBuilderGetQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse
}

func (rC *TableItemGetRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
