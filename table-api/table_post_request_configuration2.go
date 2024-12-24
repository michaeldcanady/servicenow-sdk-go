package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v{unreleased}.
//
// tablePostRequestConfiguration2[T] represents request configurations POST request.
type tablePostRequestConfiguration2[T Entry] struct {
	header   interface{}
	query    *TableRequestBuilderPostQueryParameters
	data     map[string]string
	mapping  core.ErrorMapping //nolint: staticcheck
	response *TableItemResponse
}

// toConfiguration converts rC to `core.RequestConfiguration`.
func (rC *tablePostRequestConfiguration2[T]) toConfiguration() *core.RequestConfiguration { //nolint: staticcheck
	return &core.RequestConfiguration{ //nolint: staticcheck
		Header:          rC.header,
		QueryParameters: rC.query,
		Data:            rC.data,
		ErrorMapping:    rC.mapping,
		Response:        rC.response,
	}
}
