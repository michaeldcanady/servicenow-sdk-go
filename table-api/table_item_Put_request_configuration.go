package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

// Deprecated: deprecated since v{version}. Use `TableItemPutRequestConfiguration2[T]` instead.
//
// TableItemPutRequestConfiguration represents the Request Configurations for a PUT Table Item request.
type TableItemPutRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableItemRequestBuilderPutQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse2[TableEntry]
}

func (rC *TableItemPutRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
