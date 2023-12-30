package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

// TableItemGetRequestConfiguration2 represents the Request Configurations for a GET Table Item request.
type TableItemGetRequestConfiguration2[T Entry] struct {
	header   interface{}
	query    *TableItemRequestBuilderGetQueryParameters
	data     interface{}
	mapping  core.ErrorMapping
	response *TableItemResponse2[T]
}

func (rC *TableItemGetRequestConfiguration2[T]) Header() interface{} {
	return rC.header
}

func (rC *TableItemGetRequestConfiguration2[T]) Query() interface{} {
	return rC.query
}

func (rC *TableItemGetRequestConfiguration2[T]) Data() interface{} {
	return rC.data
}

func (rC *TableItemGetRequestConfiguration2[T]) Mapping() core.ErrorMapping {
	return rC.mapping
}

func (rC *TableItemGetRequestConfiguration2[T]) Response() core.Response {
	return rC.response
}
