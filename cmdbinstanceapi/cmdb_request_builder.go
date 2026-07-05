package cmdbinstanceapi

import (
	"maps"

	appserviceapi "github.com/michaeldcanady/servicenow-sdk-go/appserviceapi"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cmdbURLTemplate = "{+baseurl}/api/now/v1/cmdb"
)

// CmdbRequestBuilder provides operations to manage ServiceNow CMDB.
type CmdbRequestBuilder struct {
	core.RequestBuilder
}

// NewCmdbRequestBuilderInternal instantiates a new CmdbRequestBuilder with the provided request parameters.
func NewCmdbRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CmdbRequestBuilder {
	return &CmdbRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, cmdbURLTemplate, pathParameters),
	}
}

// NewCmdbRequestBuilder instantiates a new CmdbRequestBuilder with the provided raw URL.
func NewCmdbRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *CmdbRequestBuilder {
	return NewCmdbRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Instance returns a CmdbInstanceRequestBuilder associated with the CmdbRequestBuilder.
func (rB *CmdbRequestBuilder) Instance() *CmdbInstanceRequestBuilder {
	return NewCmdbInstanceRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// AppService returns an AppServiceRequestBuilder associated with the CmdbRequestBuilder.
func (rB *CmdbRequestBuilder) AppService() *appserviceapi.AppServiceRequestBuilder {
	return appserviceapi.NewAppServiceRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
