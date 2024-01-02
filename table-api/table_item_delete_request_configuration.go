package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v{version}. Use `TableItemDeleteRequestConfiguration2[T]` instead.
//
// TableItemDeleteRequestConfiguration represents request configurations DELETE request.
type TableItemDeleteRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableItemRequestBuilderDeleteQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *TableItemDeleteRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}

// toTableConfig2 converts rC to `TableItemDeleteRequestConfiguration2[TableEntry]`.
func (rC *TableItemDeleteRequestConfiguration) toTableConfig2() *TableItemDeleteRequestConfiguration2[TableEntry] {
	return &TableItemDeleteRequestConfiguration2[TableEntry]{
		header:   rC.Header,
		query:    rC.QueryParameters,
		data:     rC.Data,
		mapping:  rC.ErrorMapping,
		response: rC.response,
	}
}
