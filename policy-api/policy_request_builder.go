package policyapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	policyURLTemplate = "{+baseurl}/api/now/policy"
)

// PolicyRequestBuilder provides operations to manage Service-Now policies.
type PolicyRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewPolicyRequestBuilderInternal instantiates a new PolicyRequestBuilder with the provided path parameters and request adapter.
func NewPolicyRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *PolicyRequestBuilder {
	return &PolicyRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, policyURLTemplate, pathParameters),
	}
}

// Definitions provides the way to access Service-Now's policy definitions API
func (rB *PolicyRequestBuilder) Definitions() *DefinitionsRequestBuilder {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())

	return NewDefinitionsRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
