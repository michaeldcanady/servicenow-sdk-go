package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// TableItemDeleteRequestConfiguration2 represents request configurations for a DELETE request.
type TableItemDeleteRequestConfiguration2[T Entry] struct {
	header   interface{}
	query    *TableItemRequestBuilderDeleteQueryParameters
	data     interface{}
	mapping  core.ErrorMapping
	response *TableItemResponse2[TableEntry]
}

func (rC *TableItemDeleteRequestConfiguration2[T]) Header() interface{} {
	return rC.header
}

func (rC *TableItemDeleteRequestConfiguration2[T]) Query() interface{} {
	return rC.query
}

func (rC *TableItemDeleteRequestConfiguration2[T]) Data() interface{} {
	return rC.data
}

func (rC *TableItemDeleteRequestConfiguration2[T]) Mapping() core.ErrorMapping {
	return rC.mapping
}

func (rC *TableItemDeleteRequestConfiguration2[T]) Response() core.Response {
	return rC.response
}
