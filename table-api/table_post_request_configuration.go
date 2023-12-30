package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

// Deprecated: deprecated since v{version}. Use `TablePostRequestConfiguration2[T]` instead.
//
// TablePostRequestConfiguration represents Request Configuration for POST Table Collection Request.
type TablePostRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableRequestBuilderPostQueryParameters
	Data            map[string]string
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse
}

func (rC *TablePostRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
