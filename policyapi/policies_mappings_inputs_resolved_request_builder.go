package policyapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	policiesMappingsInputsResolvedURLTemplate = "{+baseurl}/api/sn_cdm/v1/policies/mappings/inputs/resolved{?deployable_name,policy_name,sysparm_fields}"
)

// PoliciesMappingsInputsResolvedRequestBuilder provides operations to manage Service-Now policy definitions.
type PoliciesMappingsInputsResolvedRequestBuilder struct {
	core.RequestBuilder
}

// NewPoliciesMappingsInputsResolvedRequestBuilderInternal instantiates a new PoliciesMappingsInputsResolvedRequestBuilder with the provided path parameters and request adapter.
func NewPoliciesMappingsInputsResolvedRequestBuilderRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *PoliciesMappingsInputsResolvedRequestBuilder {
	return &PoliciesMappingsInputsResolvedRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, policiesMappingsInputsResolvedURLTemplate, pathParameters),
	}
}
