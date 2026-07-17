package actsubapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	subObjectsURLTemplate = "{+baseurl}/api/now/v1/actsub/subobjects"
)

// SubObjectsRequestBuilder provides operations to manage subscribable objects.
type SubObjectsRequestBuilder struct {
	*collectionGetRequestBuilder
}

// NewSubObjectsRequestBuilderInternal instantiates a new SubObjectsRequestBuilder.
func NewSubObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubObjectsRequestBuilder {
	return &SubObjectsRequestBuilder{
		newCollectionGetRequestBuilder(pathParameters, requestAdapter, subObjectsURLTemplate),
	}
}
