package actsubapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	activitiesURLTemplate = "{+baseurl}/api/now/v1/actsub/activities"
)

// ActivitiesRequestBuilder provides operations to manage activities.
type ActivitiesRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewActivitiesRequestBuilderInternal instantiates a new ActivitiesRequestBuilder.
func NewActivitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ActivitiesRequestBuilder {
	return &ActivitiesRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, activitiesURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve activities.
func (rB *ActivitiesRequestBuilder) Get(ctx context.Context, config *ActivitiesRequestBuilderGetRequestConfiguration) (*newInternal.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel], error) {
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowCollectionResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), nil)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*newInternal.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *ActivitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, config *ActivitiesRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if headers := config.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParameters := config.QueryParameters; !internal.IsNil(queryParameters) {
			kiotaRequestInfo.AddQueryParameters(queryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return requestInfo, nil
}
