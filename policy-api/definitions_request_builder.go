package policyapi

import (
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	definitionsURLTemplate = "{+baseurl}/api/now/policy/definitions{?sysparm_limit,sysparm_offset,sysparm_query,sysparm_fields}"
)

// DefinitionsRequestBuilder provides operations to manage Service-Now policy definitions.
type DefinitionsRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewDefinitionsRequestBuilderInternal instantiates a new DefinitionsRequestBuilder with the provided path parameters and request adapter.
func NewDefinitionsRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *DefinitionsRequestBuilder {
	return &DefinitionsRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, definitionsURLTemplate, pathParameters),
	}
}
