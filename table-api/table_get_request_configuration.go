package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v{version}. Use `TableGetRequestConfiguration2[T]` instead.
//
// TableGetRequestConfiguration represents request configurations GET request.
type TableGetRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableRequestBuilderGetQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableCollectionResponse
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *TableGetRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}

// toTableConfig2 converts rC to `TableGetRequestConfiguration2[TableEntry]`.
func (rC *TableGetRequestConfiguration) toTableConfig2() *TableGetRequestConfiguration2[TableEntry] {
	return &TableGetRequestConfiguration2[TableEntry]{
		header:   rC.Header,
		query:    rC.QueryParameters,
		data:     rC.Data,
		mapping:  rC.ErrorMapping,
		response: rC.response,
	}
}
