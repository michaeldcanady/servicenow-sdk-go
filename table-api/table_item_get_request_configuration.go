package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v{version}. Use `TableItemGetRequestConfiguration2[T]` instead.
//
// TableItemGetRequestConfiguration represents request configurations GET request.
type TableItemGetRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableItemRequestBuilderGetQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *TableItemGetRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}

// toTableConfig2 converts rC to `TableItemDeleteRequestConfiguration2[TableEntry]`.
func (rC *TableItemGetRequestConfiguration) toTableConfig2() *TableItemGetRequestConfiguration2[TableEntry] {
	return &TableItemGetRequestConfiguration2[TableEntry]{
		header:   rC.Header,
		query:    rC.QueryParameters,
		data:     rC.Data,
		mapping:  rC.ErrorMapping,
		response: rC.response,
	}
}
