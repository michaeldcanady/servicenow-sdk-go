package servicenowsdkgo

import (
	"maps"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	policyapi "github.com/michaeldcanady/servicenow-sdk-go/policy-api"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cdmURLTemplate = "{+baseurl}/api/sn_cdm"
)

// CdmRequestBuilder provides operations to manage Service-Now CDM.
type CdmRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewCdmRequestBuilder instantiates a new CdmRequestBuilder with the provided path parameters and request adapter.
func NewCdmRequestBuilder(url string, requestAdapter abstractions.RequestAdapter) *CdmRequestBuilder {
	pathParameters := map[string]string{"baseurl": url}
	return &CdmRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, cdmURLTemplate, pathParameters),
	}
}

// Policies returns a PolicyRequestBuilder associated with the CdmRequestBuilder.
func (rB *CdmRequestBuilder) Policies() *policyapi.PoliciesRequestBuilder {
	return policyapi.NewPolicyRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
