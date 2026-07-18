package actsubapi

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
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

// Get sends a GET request to retrieve contexts.
func (rB *ContextsRequestBuilder) Get(ctx context.Context, config *ContextsRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.collectionGetRequestBuilder.Get(ctx, config)
}

// ToGetRequestInformation creates a RequestInformation object for a GET request to retrieve contexts.
func (rB *ContextsRequestBuilder) ToGetRequestInformation(ctx context.Context, config *ContextsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.collectionGetRequestBuilder.ToGetRequestInformation(ctx, config)
}
