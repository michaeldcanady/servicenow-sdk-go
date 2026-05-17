package cmdbinstanceapi

import (
	"maps"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cmdbInstanceURLTemplate = "{+baseurl}/api/now/v1/cmdb/instance"
)

// CmdbInstanceRequestBuilder2 provides operations to manage ServiceNow CMDB instances.
type CmdbInstanceRequestBuilder2 struct {
	newInternal.RequestBuilder
}

// NewCmdbInstanceRequestBuilder2Internal instantiates a new CmdbInstanceRequestBuilder2 with the provided request parameters.
func NewCmdbInstanceRequestBuilder2Internal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CmdbInstanceRequestBuilder2 {
	return &CmdbInstanceRequestBuilder2{
		newInternal.NewBaseRequestBuilder(requestAdapter, cmdbInstanceURLTemplate, pathParameters),
	}
}

// NewCmdbInstanceRequestBuilder2 instantiates a new CmdbInstanceRequestBuilder2 with the provided raw URL.
func NewCmdbInstanceRequestBuilder2(rawURL string, requestAdapter abstractions.RequestAdapter) *CmdbInstanceRequestBuilder2 {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewCmdbInstanceRequestBuilder2Internal(urlParams, requestAdapter)
}

// ByClass provides operations to manage a specific CMDB class.
func (rB *CmdbInstanceRequestBuilder2) ByClass(className string) *CmdbClassRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["className"] = className
	return NewCmdbClassRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
