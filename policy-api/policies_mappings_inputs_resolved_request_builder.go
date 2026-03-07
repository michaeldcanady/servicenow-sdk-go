package policyapi

import (
	"context"
	"errors"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

const (
	policiesMappingsInputsResolvedURLTemplate = "{+baseurl}/api/sn_cdm/v1/policies/mappings/inputs/resolved{?deployable_name,policy_name,sysparm_fields}"
)

// PoliciesMappingsInputsResolvedRequestBuilder provides operations to manage Service-Now policy definitions.
type PoliciesMappingsInputsResolvedRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewPoliciesMappingsInputsResolvedRequestBuilderInternal instantiates a new PoliciesMappingsInputsResolvedRequestBuilder with the provided path parameters and request adapter.
func NewPoliciesMappingsInputsResolvedRequestBuilderRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *PoliciesMappingsInputsResolvedRequestBuilder {
	return &PoliciesMappingsInputsResolvedRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, policiesMappingsInputsResolvedURLTemplate, pathParameters),
	}
}

// Get sends an HTTP GET request and returns a collection of resolved policy mappings inputs.
func (rB *PoliciesMappingsInputsResolvedRequestBuilder) Get(ctx context.Context, requestConfiguration *PoliciesMappingsInputsResolvedRequestBuilderGetRequestConfiguration) (newInternal.ServiceNowCollectionResponse[*PoliciesMappingsInput], error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &PoliciesMappingsInputsResolvedRequestBuilderGetRequestConfiguration{}
	}

	headerOpt := nethttplibrary.NewHeadersInspectionOptions()
	headerOpt.InspectResponseHeaders = true

	requestConfiguration.Options = append(requestConfiguration.Options, headerOpt)

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowCollectionResponseFromDiscriminatorValue[*PoliciesMappingsInput](CreatePoliciesMappingsInputFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("response is nil")
	}

	typedResp, ok := resp.(newInternal.ServiceNowCollectionResponse[*PoliciesMappingsInput])
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*newInternal.ServiceNowCollectionResponse[*PoliciesMappingsInput])(nil))
	}

	newInternal.ParseHeaders(typedResp, headerOpt.GetResponseHeaders())

	return typedResp, nil
}

// ToGetRequestInformation converts provided parameters into request information.
func (rB *PoliciesMappingsInputsResolvedRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *PoliciesMappingsInputsResolvedRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		newInternal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
