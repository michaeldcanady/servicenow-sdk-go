package cmdbinstanceapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cmdbInstanceURLTemplate = "{+baseurl}/api/now/v1/cmdb/instance"
)

// CmdbInstanceRequestBuilder provides operations to manage ServiceNow CMDB instances.
type CmdbInstanceRequestBuilder struct {
	internal.RequestBuilder
}

// NewCmdbInstanceRequestBuilderInternal instantiates a new CmdbInstanceRequestBuilder with the provided request parameters.
func NewCmdbInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CmdbInstanceRequestBuilder {
	return &CmdbInstanceRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, cmdbInstanceURLTemplate, pathParameters),
	}
}

// NewCmdbInstanceRequestBuilder instantiates a new CmdbInstanceRequestBuilder with the provided raw URL.
func NewCmdbInstanceRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *CmdbInstanceRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[internal.RawURLKey] = rawURL
	return NewCmdbInstanceRequestBuilderInternal(urlParams, requestAdapter)
}

// ByClass provides operations to manage a specific CMDB class.
func (rB *CmdbInstanceRequestBuilder) ByClass(className string) *CmdbClassRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["className"] = className
	return NewCmdbClassRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
