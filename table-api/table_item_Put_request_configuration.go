package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v{version}. Use `TableItemPutRequestConfiguration2[T]` instead.
//
// TableItemPutRequestConfiguration represents request configurations GET request.
type TableItemPutRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableItemRequestBuilderPutQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *TableItemPutRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}

// toTableConfig2 converts rC to `TableItemDeleteRequestConfiguration2[TableEntry]`.
func (rC *TableItemPutRequestConfiguration) toTableConfig2() *TableItemPutRequestConfiguration2[TableEntry] {
	return &TableItemPutRequestConfiguration2[TableEntry]{
		header:   rC.Header,
		query:    rC.QueryParameters,
		data:     rC.Data,
		mapping:  rC.ErrorMapping,
		response: rC.response,
	}
}
