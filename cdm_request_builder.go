package servicenowsdkgo

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	policyapi "github.com/michaeldcanady/servicenow-sdk-go/policy-api"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cdmURLTemplate = "{+baseurl}/api/sn_cdm"
)

// CdmRequestBuilder provides operations to manage Service-Now CDM.
type CdmRequestBuilder struct {
	kiota.RequestBuilder
}

func NewCdmRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CdmRequestBuilder {
	return &CdmRequestBuilder{
		RequestBuilder: kiota.NewBaseRequestBuilder(requestAdapter, cdmURLTemplate, pathParameters),
	}
}

// NewCdmRequestBuilder instantiates a new CdmRequestBuilder with the provided path parameters and request adapter.
func NewCdmRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *CdmRequestBuilder {
	return NewCdmRequestBuilderInternal(map[string]string{utils.RawURLKey: rawURL}, requestAdapter)
}

// Policies returns a PolicyRequestBuilder associated with the CdmRequestBuilder.
func (rB *CdmRequestBuilder) Policies() *policyapi.PoliciesRequestBuilder {
	return policyapi.NewPolicyRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
