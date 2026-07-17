package actsubapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	activitiesURLTemplate = "{+baseurl}/api/now/v1/actsub/activities"
)

// ActivitiesRequestBuilder provides operations to manage activities.
type ActivitiesRequestBuilder struct {
	*collectionGetRequestBuilder
}

// NewActivitiesRequestBuilderInternal instantiates a new ActivitiesRequestBuilder.
func NewActivitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ActivitiesRequestBuilder {
	return &ActivitiesRequestBuilder{
		newCollectionGetRequestBuilder(pathParameters, requestAdapter, activitiesURLTemplate),
	}
}

// Get sends a GET request to retrieve activities.
func (rB *ActivitiesRequestBuilder) Get(ctx context.Context, config *ActivitiesRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) {
		return nil, nil
	}
	return rB.collectionGetRequestBuilder.Get(ctx, config)
}

// ToGetRequestInformation creates a RequestInformation object for a GET request to retrieve activities.
func (rB *ActivitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, config *ActivitiesRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, nil
	}
	return rB.collectionGetRequestBuilder.ToGetRequestInformation(ctx, config)
}
