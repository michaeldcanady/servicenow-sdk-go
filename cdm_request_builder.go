package servicenowsdkgo

import (
	"maps"

	cdmapplicationsapi "github.com/michaeldcanady/servicenow-sdk-go/cdmapplicationsapi"
	cdmchangesetapi "github.com/michaeldcanady/servicenow-sdk-go/cdmchangesetapi"
	cdmeditorapi "github.com/michaeldcanady/servicenow-sdk-go/cdmeditorapi"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal"
	policyapi "github.com/michaeldcanady/servicenow-sdk-go/policyapi"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cdmURLTemplate = "{+baseurl}/api/sn_cdm"
)

// CdmRequestBuilder provides operations to manage Service-Now CDM.
type CdmRequestBuilder struct {
	core.RequestBuilder
}

// NewCdmRequestBuilderInternal instantiates a new CdmRequestBuilder from raw path parameters.
// It is exported so sibling packages can construct a CdmRequestBuilder while chaining through
// this SDK's fluent builder tree; consumers should generally use [NewCdmRequestBuilder] instead.
func NewCdmRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CdmRequestBuilder {
	return &CdmRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, cdmURLTemplate, pathParameters),
	}
}

// NewCdmRequestBuilder instantiates a new [CdmRequestBuilder] with the provided path parameters and request adapter.
func NewCdmRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *CdmRequestBuilder {
	return NewCdmRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Policies returns a [policyapi.PolicyRequestBuilder] associated with the [CdmRequestBuilder].
func (rB *CdmRequestBuilder) Policies() *policyapi.PoliciesRequestBuilder {
	return policyapi.NewPolicyRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Editor returns a [cdmeditorapi.CdmEditorRequestBuilder] associated with the [CdmRequestBuilder].
func (rB *CdmRequestBuilder) Editor() *cdmeditorapi.CdmEditorRequestBuilder {
	return cdmeditorapi.NewCdmEditorRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Changesets returns a [cdmchangesetapi.ChangesetsRequestBuilder] associated with the [CdmRequestBuilder].
func (rB *CdmRequestBuilder) Changesets() *cdmchangesetapi.ChangesetsRequestBuilder {
	return cdmchangesetapi.NewChangesetsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Applications returns a [cdmapplicationsapi.ApplicationsRequestBuilder] associated with the [CdmRequestBuilder].
func (rB *CdmRequestBuilder) Applications() *cdmapplicationsapi.ApplicationsRequestBuilder {
	return cdmapplicationsapi.NewApplicationsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
