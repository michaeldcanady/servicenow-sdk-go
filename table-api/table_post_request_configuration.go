package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v{version}. removed from public API.
//
// TablePostRequestConfiguration represents request configurations POST request.
type TablePostRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableRequestBuilderPostQueryParameters
	Data            map[string]string
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *TablePostRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}

// toTableConfig2 converts rC to `TableItemDeleteRequestConfiguration2[TableEntry]`.
func (rC *TablePostRequestConfiguration) toTableConfig2() *tablePostRequestConfiguration2[TableEntry] {
	return &tablePostRequestConfiguration2[TableEntry]{
		header:   rC.Header,
		query:    rC.QueryParameters,
		data:     rC.Data,
		mapping:  rC.ErrorMapping,
		response: rC.response,
	}
}
