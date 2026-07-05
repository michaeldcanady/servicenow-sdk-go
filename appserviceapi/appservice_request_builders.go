package appserviceapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	appServiceURLTemplate         = "{+baseurl}/api/now/v1/cmdb/app_service"
	createURLTemplate             = "{+baseurl}/api/now/v1/cmdb/app_service/create"
	csdmAppServiceURLTemplate     = "{+baseurl}/api/now/v1/cmdb/csdm/app_service"
	findServiceURLTemplate        = "{+baseurl}/api/now/v1/cmdb/csdm/app_service/find_service{?name*,number*}"
	registerServiceURLTemplate    = "{+baseurl}/api/now/v1/cmdb/csdm/app_service/register_service"
	csdmAppServiceItemURLTemplate = "{+baseurl}/api/now/v1/cmdb/csdm/app_service/{sys_id}"
	populateServiceURLTemplate    = "{+baseurl}/api/now/v1/cmdb/csdm/app_service/{sys_id}/populate_service"
	serviceDetailsURLTemplate     = "{+baseurl}/api/now/v1/cmdb/csdm/app_service/{sys_id}/service_details"
)

// AppServiceRequestBuilder provides operations to manage ServiceNow Application Services.
type AppServiceRequestBuilder struct {
	core.RequestBuilder
}

// NewAppServiceRequestBuilderInternal instantiates a new AppServiceRequestBuilder with path parameters.
func NewAppServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AppServiceRequestBuilder {
	return &AppServiceRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, appServiceURLTemplate, pathParameters),
	}
}

// NewAppServiceRequestBuilder instantiates a new AppServiceRequestBuilder with a raw URL.
func NewAppServiceRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *AppServiceRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[internal.RawURLKey] = rawURL
	return NewAppServiceRequestBuilderInternal(urlParams, requestAdapter)
}

// Create returns a CreateRequestBuilder for POST /api/now/v1/cmdb/app_service/create.
func (rB *AppServiceRequestBuilder) Create() *CreateRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewCreateRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Csdm returns a CsdmRequestBuilder for /api/now/v1/cmdb/csdm/app_service.
func (rB *AppServiceRequestBuilder) Csdm() *CsdmRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewCsdmRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// CsdmAppServiceItemRequestBuilder provides operations for a specific CSDM application service.
type CsdmAppServiceItemRequestBuilder struct {
	core.RequestBuilder
}

// NewCsdmAppServiceItemRequestBuilderInternal instantiates a new CsdmAppServiceItemRequestBuilder.
func NewCsdmAppServiceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CsdmAppServiceItemRequestBuilder {
	return &CsdmAppServiceItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, csdmAppServiceItemURLTemplate, pathParameters),
	}
}

// PopulateService returns a PopulateServiceRequestBuilder.
func (rB *CsdmAppServiceItemRequestBuilder) PopulateService() *PopulateServiceRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewPopulateServiceRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ServiceDetails returns a ServiceDetailsRequestBuilder.
func (rB *CsdmAppServiceItemRequestBuilder) ServiceDetails() *ServiceDetailsRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewServiceDetailsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// PopulateServiceRequestBuilder provides operations to populate a CSDM service.
type PopulateServiceRequestBuilder struct {
	core.RequestBuilder
}

// NewPopulateServiceRequestBuilderInternal instantiates a new PopulateServiceRequestBuilder.
func NewPopulateServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *PopulateServiceRequestBuilder {
	return &PopulateServiceRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, populateServiceURLTemplate, pathParameters),
	}
}

// Put sends a PUT request to populate a service.
func (rB *PopulateServiceRequestBuilder) Put(ctx context.Context, body *PopulateServiceRequest, config *PopulateServiceRequestConfiguration) (PopulateServiceResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToPutRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}
	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreatePopulateServiceResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(PopulateServiceResponse), nil
}

// ToPutRequestInformation creates a RequestInformation object for a PUT request.
func (rB *PopulateServiceRequestBuilder) ToPutRequestInformation(ctx context.Context, body *PopulateServiceRequest, config *PopulateServiceRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}
	return kiotaRequestInfo.RequestInformation, nil
}

// ServiceDetailsRequestBuilder provides operations to update details of a CSDM service.
type ServiceDetailsRequestBuilder struct {
	core.RequestBuilder
}

// NewServiceDetailsRequestBuilderInternal instantiates a new ServiceDetailsRequestBuilder.
func NewServiceDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ServiceDetailsRequestBuilder {
	return &ServiceDetailsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, serviceDetailsURLTemplate, pathParameters),
	}
}

// Put sends a PUT request to update service details.
func (rB *ServiceDetailsRequestBuilder) Put(ctx context.Context, body *ServiceDetailsRequest, config *ServiceDetailsRequestConfiguration) (ServiceDetailsResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToPutRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}
	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateServiceDetailsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ServiceDetailsResponse), nil
}

// ToPutRequestInformation creates a RequestInformation object for a PUT request.
func (rB *ServiceDetailsRequestBuilder) ToPutRequestInformation(ctx context.Context, body *ServiceDetailsRequest, config *ServiceDetailsRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}
	return kiotaRequestInfo.RequestInformation, nil
}
