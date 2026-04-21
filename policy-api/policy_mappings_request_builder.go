package policyapi

import (
	"context"
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// policyMappingsURLTemplate is the URL template for this endpoint.
	policyMappingsURLTemplate = "{+baseurl}/api/sn_cdm/v1/policies/mappings{?appName,deployableName,policyName,returnFields}"
)

// PolicyMappingsRequestBuilder provides operations for managing the PolicyMappings endpoint.
type PolicyMappingsRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewPolicyMappingsRequestBuilderInternal instantiates a new PolicyMappingsRequestBuilder with the provided path parameters and request adapter.
func NewPolicyMappingsRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *PolicyMappingsRequestBuilder {
	return &PolicyMappingsRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, policyMappingsURLTemplate, pathParameters),
	}
}

// NewPolicyMappingsRequestBuilder instantiates a new PolicyMappingsRequestBuilder with a raw URL.
func NewPolicyMappingsRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *PolicyMappingsRequestBuilder {
	urlParams := map[string]string{newInternal.RawURLKey: rawURL}
	return NewPolicyMappingsRequestBuilderInternal(urlParams, requestAdapter)
}

// Inputs provides access to the PolicyMappingInputs endpoint.
func (rB *PolicyMappingsRequestBuilder) Inputs() *PolicyMappingInputsRequestBuilder {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	return NewPolicyMappingInputsRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Post sends a POST request to the endpoint.
func (rB *PolicyMappingsRequestBuilder) Post(
	ctx context.Context,
	requestConfiguration *PolicyMappingsRequestBuilderPostRequestConfiguration,
) (
	newInternal.ServiceNowItemResponse[PolicyMapping],
	error,
) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &PolicyMappingsRequestBuilderPostRequestConfiguration{}
	}

	// Validate required query parameters
	queryParameters := requestConfiguration.QueryParameters

	if queryParameters.AppName == nil || *queryParameters.AppName == "" {
		return nil, errors.New("AppName is required")
	}

	if queryParameters.DeployableName == nil || *queryParameters.DeployableName == "" {
		return nil, errors.New("DeployableName is required")
	}

	if queryParameters.PolicyName == nil || *queryParameters.PolicyName == "" {
		return nil, errors.New("PolicyName is required")
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(
		ctx,
		requestInfo,
		newInternal.ServiceNowItemResponseFromDiscriminatorValue[PolicyMapping](CreatePolicyMappingFromDiscriminatorValue),
		errorMapping,
	)

	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	typedRes, ok := res.(newInternal.ServiceNowItemResponse[PolicyMapping])
	if !ok {
		return nil, errors.New("unexpected response type")
	}

	return typedRes, nil
}

// ToPostRequestInformation creates the RequestInformation for a POST request.
func (rB *PolicyMappingsRequestBuilder) ToPostRequestInformation(
	ctx context.Context,
	requestConfiguration *PolicyMappingsRequestBuilderPostRequestConfiguration,
) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(
		abstractions.POST,
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

// Delete sends a DELETE request to the endpoint.
func (rB *PolicyMappingsRequestBuilder) Delete(
	ctx context.Context,
	requestConfiguration *PolicyMappingsRequestBuilderDeleteRequestConfiguration,
) error {
	if internal.IsNil(rB) {
		return nil
	}

	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &PolicyMappingsRequestBuilderDeleteRequestConfiguration{}
	}

	// Validate required query parameters
	queryParameters := requestConfiguration.QueryParameters

	if queryParameters.AppName == nil || *queryParameters.AppName == "" {
		return errors.New("AppName is required")
	}

	if queryParameters.DeployableName == nil || *queryParameters.DeployableName == "" {
		return errors.New("DeployableName is required")
	}

	if queryParameters.PolicyName == nil || *queryParameters.PolicyName == "" {
		return errors.New("PolicyName is required")
	}

	requestInfo, err := rB.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
}

// ToDeleteRequestInformation creates the RequestInformation for a DELETE request.
func (rB *PolicyMappingsRequestBuilder) ToDeleteRequestInformation(
	ctx context.Context,
	requestConfiguration *PolicyMappingsRequestBuilderDeleteRequestConfiguration,
) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(
		abstractions.DELETE,
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
