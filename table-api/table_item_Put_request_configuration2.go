package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

// TableItemPutRequestConfiguration2[T] represents the Request Configurations for a PUT Table Item request.
type TableItemPutRequestConfiguration2[T Entry] struct {
	header   interface{}
	query    *TableItemRequestBuilderPutQueryParameters
	data     interface{}
	mapping  core.ErrorMapping
	response *TableItemResponse2[T]
}

func (rC *TableItemPutRequestConfiguration2[T]) Header() interface{} {
	return rC.header
}

func (rC *TableItemPutRequestConfiguration2[T]) Query() interface{} {
	return rC.query
}

func (rC *TableItemPutRequestConfiguration2[T]) Data() interface{} {
	return rC.data
}

func (rC *TableItemPutRequestConfiguration2[T]) Mapping() core.ErrorMapping {
	return rC.mapping
}

func (rC *TableItemPutRequestConfiguration2[T]) Response() core.Response {
	return rC.response
}
