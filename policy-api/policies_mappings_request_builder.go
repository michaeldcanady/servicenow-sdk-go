package policyapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	policiesMappingsURLTemplate = "{+baseurl}/api/sn_cdm/v1/policies/mappings"
)

// PoliciesMappingsRequestBuilder provides operations to manage Service-Now policy definitions.
type PoliciesMappingsRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewPoliciesMappingsRequestBuilderInternal instantiates a new PoliciesMappingsRequestBuilder with the provided path parameters and request adapter.
func NewPoliciesMappingsRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *PoliciesMappingsRequestBuilder {
	return &PoliciesMappingsRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, policiesMappingsURLTemplate, pathParameters),
	}
}

func (rB *PoliciesMappingsRequestBuilder) Inputs() *PoliciesMappingsInputsRequestBuilder {
	if internal.IsNil(rB) {
		return nil
	}

	return NewPoliciesMappingsInputsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
