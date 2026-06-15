package policyapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	policiesMappingsInputsURLTemplate = "{+baseurl}/api/sn_cdm/v1/policies/mappings/inputs"
)

// PoliciesMappingsRequestBuilder provides operations to manage Service-Now policy definitions.
type PoliciesMappingsInputsRequestBuilder struct {
	internal.RequestBuilder
}

// NewPoliciesMappingsInputsRequestBuilderInternal instantiates a new PoliciesMappingsInputsRequestBuilder with the provided path parameters and request adapter.
func NewPoliciesMappingsInputsRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *PoliciesMappingsInputsRequestBuilder {
	return &PoliciesMappingsInputsRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, policiesMappingsInputsURLTemplate, pathParameters),
	}
}

func (rB *PoliciesMappingsInputsRequestBuilder) Resolved() *PoliciesMappingsInputsResolvedRequestBuilder {
	if conversion.IsNil(rB) {
		return nil
	}

	return NewPoliciesMappingsInputsResolvedRequestBuilderRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
