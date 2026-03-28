package policyapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// policyMappingInputsURLTemplate is the URL template for this endpoint.
	policyMappingInputsURLTemplate = "{+baseurl}/api/sn_cdm/v1/policies/mappings/inputs{?appName,deployableName,inputName,inputValue,policyName,returnFields}"
)

// PolicyMappingInputsRequestBuilder provides operations for managing the PolicyMappingInputs endpoint.
type PolicyMappingInputsRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewPolicyMappingInputsRequestBuilderInternal instantiates a new PolicyMappingInputsRequestBuilder with the provided path parameters and request adapter.
func NewPolicyMappingInputsRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,

) *PolicyMappingInputsRequestBuilder {
	return &PolicyMappingInputsRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, policyMappingInputsURLTemplate, pathParameters),
	}
}

// NewPolicyMappingInputsRequestBuilder instantiates a new PolicyMappingInputsRequestBuilder with a raw URL.
func NewPolicyMappingInputsRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,

) *PolicyMappingInputsRequestBuilder {
	urlParams := map[string]string{newInternal.RawURLKey: rawURL}
	return NewPolicyMappingInputsRequestBuilderInternal(urlParams, requestAdapter)
}

// Put sends a PUT request to the endpoint.
func (rB *PolicyMappingInputsRequestBuilder) Put(
	ctx context.Context,
	requestConfiguration *PolicyMappingInputsRequestBuilderPutRequestConfiguration,
) (
	newInternal.ServiceNowItemResponse[MappingInputVariable],
	error,
) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &PolicyMappingInputsRequestBuilderPutRequestConfiguration{}
	}

	// Validate required query parameters
	queryParameters := requestConfiguration.QueryParameters

	if queryParameters.AppName == nil || *queryParameters.AppName == "" {
		return nil, errors.New("AppName is required")
	}

	if queryParameters.DeployableName == nil || *queryParameters.DeployableName == "" {
		return nil, errors.New("DeployableName is required")
	}

	if queryParameters.InputName == nil || *queryParameters.InputName == "" {
		return nil, errors.New("InputName is required")
	}

	if queryParameters.InputValue == nil || *queryParameters.InputValue == "" {
		return nil, errors.New("InputValue is required")
	}

	if queryParameters.PolicyName == nil || *queryParameters.PolicyName == "" {
		return nil, errors.New("PolicyName is required")
	}

	requestInfo, err := rB.ToPutRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(
		ctx,
		requestInfo,
		newInternal.ServiceNowItemResponseFromDiscriminatorValue[MappingInputVariable](CreateMappingInputVariableFromDiscriminatorValue),
		errorMapping,
	)

	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	typedRes, ok := res.(newInternal.ServiceNowItemResponse[MappingInputVariable])
	if !ok {
		return nil, errors.New("unexpected response type")
	}

	return typedRes, nil
}

// ToPutRequestInformation creates the RequestInformation for a PUT request.
func (rB *PolicyMappingInputsRequestBuilder) ToPutRequestInformation(
	ctx context.Context,
	requestConfiguration *PolicyMappingInputsRequestBuilderPutRequestConfiguration,
) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(
		abstractions.PUT,
		rB.GetURLTemplate(),
		rB.GetPathParameters(),
	)

	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}

	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParams := requestConfiguration.QueryParameters; !internal.IsNil(queryParams) {
			kiotaRequestInfo.AddQueryParameters(queryParams)
		}
	}

	return kiotaRequestInfo.RequestInformation, nil
}
