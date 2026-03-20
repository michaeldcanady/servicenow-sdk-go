package policyapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	policyURLTemplate = "{+baseurl}/api/sn_cdm/v1/policies"
)

// PoliciesRequestBuilder provides operations to manage Service-Now policies.
type PoliciesRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewPolicyRequestBuilderInternal instantiates a new PolicyRequestBuilder with the provided path parameters and request adapter.
func NewPolicyRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *PoliciesRequestBuilder {
	return &PoliciesRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, policyURLTemplate, pathParameters),
	}
}

// Mappings provides the way to access Service-Now's policy definitions API
func (rB *PoliciesRequestBuilder) Mappings() *PoliciesMappingsRequestBuilder {
	if internal.IsNil(rB) {
		return nil
	}

	return NewPoliciesMappingsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
