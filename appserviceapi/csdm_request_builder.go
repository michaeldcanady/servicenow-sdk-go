package appserviceapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const csdmAppServiceURLTemplate = "{+baseurl}/api/now/cmdb/csdm/app_service"

// CsdmRequestBuilder provides operations under /api/now/cmdb/csdm/app_service.
type CsdmRequestBuilder struct {
	core.RequestBuilder
}

// NewCsdmRequestBuilderInternal instantiates a new CsdmRequestBuilder.
func NewCsdmRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CsdmRequestBuilder {
	return &CsdmRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, csdmAppServiceURLTemplate, pathParameters),
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
	pathParameters[sysIDKey] = sysID
	return NewCsdmAppServiceItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
