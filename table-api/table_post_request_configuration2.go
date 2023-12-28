package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// TablePostRequestConfiguration2[T] represents Request Configuration for POST Table Collection Request.
type TablePostRequestConfiguration2[T Entry] struct {
	Header          interface{}
	QueryParameters *TableRequestBuilderPostQueryParameters
	Data            T
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse2[T]
}

func (rC *TablePostRequestConfiguration2[T]) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
