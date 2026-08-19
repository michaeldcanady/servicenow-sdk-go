package appserviceapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const appServiceURLTemplate = "{+baseurl}/api/now/cmdb/app_service"

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

// Create returns a [CreateRequestBuilder].
func (rB *AppServiceRequestBuilder) Create() *CreateRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewCreateRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Csdm returns a [CsdmRequestBuilder].
func (rB *AppServiceRequestBuilder) Csdm() *CsdmRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewCsdmRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ByID returns an [AppServiceItemRequestBuilder] for the specified application service.
func (rB *AppServiceRequestBuilder) ByID(sysID string) *AppServiceItemRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters[sysIDKey] = sysID
	return NewAppServiceItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// ConvertToDynamicService returns a [ConvertToDynamicServiceRequestBuilder].
func (rB *AppServiceRequestBuilder) ConvertToDynamicService() *ConvertToDynamicServiceRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewConvertToDynamicServiceRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ConvertToManualService returns a [ConvertToManualServiceRequestBuilder].
func (rB *AppServiceRequestBuilder) ConvertToManualService() *ConvertToManualServiceRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewConvertToManualServiceRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// CreateDynamicService returns a [CreateDynamicServiceRequestBuilder].
func (rB *AppServiceRequestBuilder) CreateDynamicService() *CreateDynamicServiceRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewCreateDynamicServiceRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// UpdateDynamicNumberOfLevels returns an [UpdateDynamicNumberOfLevelsRequestBuilder].
func (rB *AppServiceRequestBuilder) UpdateDynamicNumberOfLevels() *UpdateDynamicNumberOfLevelsRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewUpdateDynamicNumberOfLevelsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
