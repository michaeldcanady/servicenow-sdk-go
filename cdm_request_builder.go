package servicenowsdkgo

import (
	"maps"

	cdmapplicationsapi "github.com/michaeldcanady/servicenow-sdk-go/v2/cdm-applications-api"
	cdmchangesetapi "github.com/michaeldcanady/servicenow-sdk-go/v2/cdm-changeset-api"
	cdmeditorapi "github.com/michaeldcanady/servicenow-sdk-go/v2/cdm-editor-api"
	internal "github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	policyapi "github.com/michaeldcanady/servicenow-sdk-go/v2/policy-api"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cdmURLTemplate = "{+baseurl}/api/sn_cdm"
)

// CdmRequestBuilder provides operations to manage Service-Now CDM.
type CdmRequestBuilder struct {
	internal.RequestBuilder
}

func NewCdmRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CdmRequestBuilder {
	return &CdmRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, cdmURLTemplate, pathParameters),
	}
}

// NewCdmRequestBuilder instantiates a new CdmRequestBuilder with the provided path parameters and request adapter.
func NewCdmRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *CdmRequestBuilder {
	return NewCdmRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Policies returns a PolicyRequestBuilder associated with the CdmRequestBuilder.
func (rB *CdmRequestBuilder) Policies() *policyapi.PoliciesRequestBuilder {
	return policyapi.NewPolicyRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Editor returns a CdmEditorRequestBuilder associated with the CdmRequestBuilder.
func (rB *CdmRequestBuilder) Editor() *cdmeditorapi.CdmEditorRequestBuilder {
	return cdmeditorapi.NewCdmEditorRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Changesets returns a ChangesetsRequestBuilder associated with the CdmRequestBuilder.
func (rB *CdmRequestBuilder) Changesets() *cdmchangesetapi.ChangesetsRequestBuilder {
	return cdmchangesetapi.NewChangesetsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Applications returns a ApplicationsRequestBuilder associated with the CdmRequestBuilder.
func (rB *CdmRequestBuilder) Applications() *cdmapplicationsapi.ApplicationsRequestBuilder {
	return cdmapplicationsapi.NewApplicationsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
