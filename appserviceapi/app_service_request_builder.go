package appserviceapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const appServiceURLTemplate = "{+baseurl}/api/now/v1/cmdb/app_service"

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
