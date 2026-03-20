package policyapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	policyURLTemplate = "{+baseurl}/api/sn_cdm/v1/policies"
)

// PoliciesRequestBuilder provides operations to manage Service-Now policies.
type PoliciesRequestBuilder struct {
	kiota.RequestBuilder
}

// NewPolicyRequestBuilderInternal instantiates a new PolicyRequestBuilder with the provided path parameters and request adapter.
func NewPolicyRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *PoliciesRequestBuilder {
	return &PoliciesRequestBuilder{
		RequestBuilder: kiota.NewBaseRequestBuilder(requestAdapter, policyURLTemplate, pathParameters),
	}
}

// Mappings provides the way to access Service-Now's policy definitions API
func (rB *PoliciesRequestBuilder) Mappings() *PoliciesMappingsRequestBuilder {
	if utils.IsNil(rB) {
		return nil
	}

	return NewPoliciesMappingsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
