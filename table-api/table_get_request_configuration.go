package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

// Deprecated: deprecated since v{version}. Use `TableGetRequestConfiguration2[T]` instead.
//
// TableGetRequestConfiguration represents request configurations for a GET request.
type TableGetRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableRequestBuilderGetQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableCollectionResponse2[TableEntry]
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
