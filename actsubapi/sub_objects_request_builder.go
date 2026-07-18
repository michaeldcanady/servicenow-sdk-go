package actsubapi

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
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

// Get sends a GET request to retrieve subscribable objects.
func (rB *SubObjectsRequestBuilder) Get(ctx context.Context, config *SubObjectsRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.collectionGetRequestBuilder.Get(ctx, config)
}

// ToGetRequestInformation creates a RequestInformation object for a GET request to retrieve subscribable objects.
func (rB *SubObjectsRequestBuilder) ToGetRequestInformation(ctx context.Context, config *SubObjectsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.collectionGetRequestBuilder.ToGetRequestInformation(ctx, config)
}
