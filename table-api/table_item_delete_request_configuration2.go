package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v{unreleased}.
//
// tableItemDeleteRequestConfiguration2[T] represents request configurations DELETE request.
type tableItemDeleteRequestConfiguration2[T Entry] struct {
	header   interface{}
	query    *TableItemRequestBuilderDeleteQueryParameters
	data     interface{}
	mapping  core.ErrorMapping //nolint: staticcheck
	response *TableItemResponse2[T]
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *tableItemDeleteRequestConfiguration2[T]) toConfiguration() *core.RequestConfiguration { //nolint: staticcheck
	return &core.RequestConfiguration{ //nolint: staticcheck
		Header:          rC.header,
		QueryParameters: rC.query,
		Data:            rC.data,
		ErrorMapping:    rC.mapping,
		Response:        rC.response,
	}
}
