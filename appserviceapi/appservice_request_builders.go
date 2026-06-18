package appserviceapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
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
	internal.RequestBuilder
}

// NewAppServiceRequestBuilderInternal instantiates a new AppServiceRequestBuilder with path parameters.
func NewAppServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AppServiceRequestBuilder {
	return &AppServiceRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, appServiceURLTemplate, pathParameters),
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

// CreateRequestBuilder provides operations to create an application service.
type CreateRequestBuilder struct {
	internal.RequestBuilder
}

// NewCreateRequestBuilderInternal instantiates a new CreateRequestBuilder.
func NewCreateRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, createURLTemplate, pathParameters),
	}
}

// Post sends a POST request to create an application service.
func (rB *CreateRequestBuilder) Post(ctx context.Context, body *CreateServiceRequest, config *CreateRequestConfiguration) (CreateServiceResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}
	errorMapping := internal.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateCreateServiceResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CreateServiceResponse), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *CreateRequestBuilder) ToPostRequestInformation(ctx context.Context, body *CreateServiceRequest, config *CreateRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)
	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}
	return kiotaRequestInfo.RequestInformation, nil
}

// CsdmRequestBuilder provides operations under /api/now/v1/cmdb/csdm/app_service.
type CsdmRequestBuilder struct {
	internal.RequestBuilder
}

// NewCsdmRequestBuilderInternal instantiates a new CsdmRequestBuilder.
func NewCsdmRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CsdmRequestBuilder {
	return &CsdmRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, csdmAppServiceURLTemplate, pathParameters),
	}
}

// FindService returns a FindServiceRequestBuilder.
func (rB *CsdmRequestBuilder) FindService() *FindServiceRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewFindServiceRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// RegisterService returns a RegisterServiceRequestBuilder.
func (rB *CsdmRequestBuilder) RegisterService() *RegisterServiceRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewRegisterServiceRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ByID returns a CsdmAppServiceItemRequestBuilder.
func (rB *CsdmRequestBuilder) ByID(sysID string) *CsdmAppServiceItemRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["sys_id"] = sysID
	return NewCsdmAppServiceItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// FindServiceRequestBuilder provides operations to find an application service.
type FindServiceRequestBuilder struct {
	internal.RequestBuilder
}

// NewFindServiceRequestBuilderInternal instantiates a new FindServiceRequestBuilder.
func NewFindServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FindServiceRequestBuilder {
	return &FindServiceRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, findServiceURLTemplate, pathParameters),
	}
}

// Get sends a GET request to find an application service.
func (rB *FindServiceRequestBuilder) Get(ctx context.Context, config *FindServiceRequestConfiguration) (FindServiceResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}
	errorMapping := internal.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateFindServiceResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(FindServiceResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *FindServiceRequestBuilder) ToGetRequestInformation(ctx context.Context, config *FindServiceRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)
	return kiotaRequestInfo.RequestInformation, nil
}

// RegisterServiceRequestBuilder provides operations to register a CSDM service.
type RegisterServiceRequestBuilder struct {
	internal.RequestBuilder
}

// NewRegisterServiceRequestBuilderInternal instantiates a new RegisterServiceRequestBuilder.
func NewRegisterServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *RegisterServiceRequestBuilder {
	return &RegisterServiceRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, registerServiceURLTemplate, pathParameters),
	}
}

// Post sends a POST request to register a service.
func (rB *RegisterServiceRequestBuilder) Post(ctx context.Context, body *RegisterServiceRequest, config *RegisterServiceRequestConfiguration) (RegisterServiceResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}
	errorMapping := internal.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateRegisterServiceResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(RegisterServiceResponse), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *RegisterServiceRequestBuilder) ToPostRequestInformation(ctx context.Context, body *RegisterServiceRequest, config *RegisterServiceRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)
	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}
	return kiotaRequestInfo.RequestInformation, nil
}

// CsdmAppServiceItemRequestBuilder provides operations for a specific CSDM application service.
type CsdmAppServiceItemRequestBuilder struct {
	internal.RequestBuilder
}

// NewCsdmAppServiceItemRequestBuilderInternal instantiates a new CsdmAppServiceItemRequestBuilder.
func NewCsdmAppServiceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CsdmAppServiceItemRequestBuilder {
	return &CsdmAppServiceItemRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, csdmAppServiceItemURLTemplate, pathParameters),
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
	internal.RequestBuilder
}

// NewPopulateServiceRequestBuilderInternal instantiates a new PopulateServiceRequestBuilder.
func NewPopulateServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *PopulateServiceRequestBuilder {
	return &PopulateServiceRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, populateServiceURLTemplate, pathParameters),
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
	errorMapping := internal.DefaultErrorMapping()
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

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)
	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}
	return kiotaRequestInfo.RequestInformation, nil
}

// ServiceDetailsRequestBuilder provides operations to update details of a CSDM service.
type ServiceDetailsRequestBuilder struct {
	internal.RequestBuilder
}

// NewServiceDetailsRequestBuilderInternal instantiates a new ServiceDetailsRequestBuilder.
func NewServiceDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ServiceDetailsRequestBuilder {
	return &ServiceDetailsRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, serviceDetailsURLTemplate, pathParameters),
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
	errorMapping := internal.DefaultErrorMapping()
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

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)
	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}
	return kiotaRequestInfo.RequestInformation, nil
}
