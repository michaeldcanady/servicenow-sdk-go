package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v{version}. Use `TableItemPutRequestConfiguration2[T]` instead.
//
// TableItemPutRequestConfiguration2[T] represents the Request Configurations for a PUT Table Item request.
type TableItemPutRequestConfiguration2[T TableEntry2] struct {
	Header          interface{}
	QueryParameters *TableItemRequestBuilderPutQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse2[T]
}

func (rC *TableItemPutRequestConfiguration2[T]) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
