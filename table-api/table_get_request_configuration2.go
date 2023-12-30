package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

// TableGetRequestConfiguration2 represents request configurations for GET request.
type TableGetRequestConfiguration2[T Entry] struct {
	header   interface{}
	query    *TableRequestBuilderGetQueryParameters
	data     interface{}
	mapping  core.ErrorMapping
	response *TableCollectionResponse2[T]
}

func (rC *TableGetRequestConfiguration2[T]) Header() interface{} {
	return rC.header
}

func (rC *TableGetRequestConfiguration2[T]) Query() interface{} {
	return rC.query
}

func (rC *TableGetRequestConfiguration2[T]) Data() interface{} {
	return rC.data
}

func (rC *TableGetRequestConfiguration2[T]) Mapping() core.ErrorMapping {
	return rC.mapping
}

func (rC *TableGetRequestConfiguration2[T]) Response() core.Response {
	return rC.response
}
