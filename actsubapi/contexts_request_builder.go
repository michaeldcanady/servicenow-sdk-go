package actsubapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	contextsURLTemplate = "{+baseurl}/api/now/v1/actsub/contexts"
)

// ContextsRequestBuilder provides operations to manage contexts.
type ContextsRequestBuilder struct {
	*collectionGetRequestBuilder
}

// NewContextsRequestBuilderInternal instantiates a new ContextsRequestBuilder.
func NewContextsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ContextsRequestBuilder {
	return &ContextsRequestBuilder{
		newCollectionGetRequestBuilder(pathParameters, requestAdapter, contextsURLTemplate),
	}
}
