package cmdbinstanceapi

import (
	"maps"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cmdbInstanceURLTemplate = "{+baseurl}/api/now/v1/cmdb/instance"
)

// CmdbInstanceRequestBuilder provides operations to manage ServiceNow CMDB instances.
type CmdbInstanceRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewCmdbInstanceRequestBuilderInternal instantiates a new CmdbInstanceRequestBuilder with the provided request parameters.
func NewCmdbInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CmdbInstanceRequestBuilder {
	return &CmdbInstanceRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, cmdbInstanceURLTemplate, pathParameters),
	}
}

// NewCmdbInstanceRequestBuilder instantiates a new CmdbInstanceRequestBuilder with the provided raw URL.
func NewCmdbInstanceRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *CmdbInstanceRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewCmdbInstanceRequestBuilderInternal(urlParams, requestAdapter)
}

// ByClass provides operations to manage a specific CMDB class.
func (rB *CmdbInstanceRequestBuilder) ByClass(className string) *CmdbClassRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["className"] = className
	return NewCmdbClassRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
