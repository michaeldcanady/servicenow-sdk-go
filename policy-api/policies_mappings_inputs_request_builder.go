package policyapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	policiesMappingsInputsURLTemplate = "{+baseurl}/api/sn_cdm/v1/policies/mappings/inputs"
)

// PoliciesMappingsRequestBuilder provides operations to manage Service-Now policy definitions.
type PoliciesMappingsInputsRequestBuilder struct {
	kiota.RequestBuilder
}

// NewPoliciesMappingsInputsRequestBuilderInternal instantiates a new PoliciesMappingsInputsRequestBuilder with the provided path parameters and request adapter.
func NewPoliciesMappingsInputsRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *PoliciesMappingsInputsRequestBuilder {
	return &PoliciesMappingsInputsRequestBuilder{
		RequestBuilder: kiota.NewBaseRequestBuilder(requestAdapter, policiesMappingsInputsURLTemplate, pathParameters),
	}
}

func (rB *PoliciesMappingsInputsRequestBuilder) Resolved() *PoliciesMappingsInputsResolvedRequestBuilder {
	if utils.IsNil(rB) {
		return nil
	}

	return NewPoliciesMappingsInputsResolvedRequestBuilderRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
