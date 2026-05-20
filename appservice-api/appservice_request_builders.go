package appserviceapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	appServiceURLTemplate            = "{+baseurl}/api/now/v1/cmdb/app_service"
	createOrUpdateServiceURLTemplate = "{+baseurl}/api/now/v1/cmdb/app_service/createOrUpdateService"
	appServiceItemURLTemplate        = "{+baseurl}/api/now/v1/cmdb/app_service/{sys_id}"
	getContentURLTemplate            = "{+baseurl}/api/now/v1/cmdb/app_service/{sys_id}/getContent{?Mode}"
)

// AppServiceRequestBuilder provides operations to manage ServiceNow Application Services.
type AppServiceRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewAppServiceRequestBuilderInternal instantiates a new AppServiceRequestBuilder with path parameters.
func NewAppServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AppServiceRequestBuilder {
	return &AppServiceRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, appServiceURLTemplate, pathParameters),
	}
}

// NewAppServiceRequestBuilder instantiates a new AppServiceRequestBuilder with a raw URL.
func NewAppServiceRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *AppServiceRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewAppServiceRequestBuilderInternal(urlParams, requestAdapter)
}

// CreateOrUpdateService returns a CreateOrUpdateServiceRequestBuilder.
func (rB *AppServiceRequestBuilder) CreateOrUpdateService() *CreateOrUpdateServiceRequestBuilder {
	return NewCreateOrUpdateServiceRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ByID returns an AppServiceItemRequestBuilder.
func (rB *AppServiceRequestBuilder) ByID(sysID string) *AppServiceItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["sys_id"] = sysID
	return NewAppServiceItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// CreateOrUpdateServiceRequestBuilder provides operations to create or update an application service.
type CreateOrUpdateServiceRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewCreateOrUpdateServiceRequestBuilderInternal instantiates a new CreateOrUpdateServiceRequestBuilder.
func NewCreateOrUpdateServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CreateOrUpdateServiceRequestBuilder {
	return &CreateOrUpdateServiceRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, createOrUpdateServiceURLTemplate, pathParameters),
	}
}

// Post sends a POST request to create or update an application service.
func (rB *CreateOrUpdateServiceRequestBuilder) Post(ctx context.Context, body *CreateOrUpdateServiceRequest, config *AppServiceRequestBuilderCreateOrUpdateServiceRequestConfiguration) (CreateOrUpdateServiceResponse, error) {
	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}
	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateCreateOrUpdateServiceResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CreateOrUpdateServiceResponse), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *CreateOrUpdateServiceRequestBuilder) ToPostRequestInformation(ctx context.Context, body *CreateOrUpdateServiceRequest, config *AppServiceRequestBuilderCreateOrUpdateServiceRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if headers := config.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)
	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), newInternal.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}
	return requestInfo, nil
}

// AppServiceItemRequestBuilder provides operations to manage a single application service.
type AppServiceItemRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewAppServiceItemRequestBuilderInternal instantiates a new AppServiceItemRequestBuilder.
func NewAppServiceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AppServiceItemRequestBuilder {
	return &AppServiceItemRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, appServiceItemURLTemplate, pathParameters),
	}
}

// GetContent returns a GetContentRequestBuilder.
func (rB *AppServiceItemRequestBuilder) GetContent() *GetContentRequestBuilder {
	return NewGetContentRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// GetContentRequestBuilder retrieves the content of an application service.
type GetContentRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewGetContentRequestBuilderInternal instantiates a new GetContentRequestBuilder.
func NewGetContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *GetContentRequestBuilder {
	return &GetContentRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, getContentURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve the content of an application service.
func (rB *GetContentRequestBuilder) Get(ctx context.Context, config *GetContentRequestBuilderGetRequestConfiguration) (GetContentResponse, error) {
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}
	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateGetContentResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(GetContentResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *GetContentRequestBuilder) ToGetRequestInformation(_ context.Context, config *GetContentRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
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
