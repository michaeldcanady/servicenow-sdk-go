package cmdbinstanceapi

import (
	"maps"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cmdbURLTemplate = "{+baseurl}/api/now/v1/cmdb"
)

// CmdbRequestBuilder provides operations to manage ServiceNow CMDB.
type CmdbRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewCmdbRequestBuilderInternal instantiates a new CmdbRequestBuilder with the provided request parameters.
func NewCmdbRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CmdbRequestBuilder {
	return &CmdbRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, cmdbURLTemplate, pathParameters),
	}
}

// NewCmdbRequestBuilder instantiates a new CmdbRequestBuilder with the provided raw URL.
func NewCmdbRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *CmdbRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewCmdbRequestBuilderInternal(urlParams, requestAdapter)
}

// Instance returns a CmdbInstanceRequestBuilder associated with the CmdbRequestBuilder.
func (rB *CmdbRequestBuilder) Instance() *CmdbInstanceRequestBuilder {
	return NewCmdbInstanceRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
