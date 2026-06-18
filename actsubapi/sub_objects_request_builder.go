package actsubapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	subObjectsURLTemplate = "{+baseurl}/api/now/v1/actsub/subobjects"
)

// SubObjectsRequestBuilder provides operations to manage subscribable objects.
type SubObjectsRequestBuilder struct {
	internal.RequestBuilder
}

// NewSubObjectsRequestBuilderInternal instantiates a new SubObjectsRequestBuilder.
func NewSubObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubObjectsRequestBuilder {
	return &SubObjectsRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, subObjectsURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve subscribable objects.
func (rB *SubObjectsRequestBuilder) Get(ctx context.Context, config *SubObjectsRequestBuilderGetRequestConfiguration) (*internal.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowCollectionResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), internal.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*internal.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *SubObjectsRequestBuilder) ToGetRequestInformation(ctx context.Context, config *SubObjectsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return requestInfo, nil
}
