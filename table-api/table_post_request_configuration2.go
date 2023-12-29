package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// TablePostRequestConfiguration2[T] represents Request Configuration for POST Table Collection Request.
type TablePostRequestConfiguration2[T Entry] struct {
	header   interface{}
	query    *TableRequestBuilderPostQueryParameters
	data     T
	mapping  core.ErrorMapping
	response *TableItemResponse2[T]
}

func (rC *TablePostRequestConfiguration2[T]) Header() interface{} {
	return rC.header
}

func (rC *TablePostRequestConfiguration2[T]) Query() interface{} {
	return rC.query
}

func (rC *TablePostRequestConfiguration2[T]) Data() interface{} {
	return rC.data
}

func (rC *TablePostRequestConfiguration2[T]) Mapping() core.ErrorMapping {
	return rC.mapping
}

func (rC *TablePostRequestConfiguration2[T]) Response() core.Response {
	return rC.response
}
