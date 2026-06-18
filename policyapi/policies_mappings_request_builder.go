package policyapi

import (
	"context"
	"errors"
	"fmt"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	policiesMappingsURLTemplate = "{+baseurl}/api/sn_cdm/v1/policies/mappings"
)

// PoliciesMappingsRequestBuilder provides operations to manage Service-Now policy definitions.
type PoliciesMappingsRequestBuilder struct {
	core.RequestBuilder
}

// NewPoliciesMappingsRequestBuilderInternal instantiates a new PoliciesMappingsRequestBuilder with the provided path parameters and request adapter.
func NewPoliciesMappingsRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *PoliciesMappingsRequestBuilder {
	return &PoliciesMappingsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, policiesMappingsURLTemplate, pathParameters),
	}
}

func (rB *PoliciesMappingsRequestBuilder) Inputs() *PoliciesMappingsInputsRequestBuilder {
	if conversion.IsNil(rB) {
		return nil
	}

	return NewPoliciesMappingsInputsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

func (rB *PoliciesMappingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *PoliciesMappingsRequestBuilderDeleteRequestConfiguration) error {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	if conversion.IsNil(requestConfiguration) {
		return errors.New("requestConfiguration is nil")
	}

	queryParameters := requestConfiguration.QueryParameters
	if queryParameters.AppName == "" {
		return errors.New("AppName is required")
	}

	if queryParameters.DeployableName == "" {
		return errors.New("DeployableName is required")
	}

	if queryParameters.PolicyName == "" {
		return errors.New("PolicyName is required")
	}

	requestInfo, err := rB.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	errorMapping := core.DefaultErrorMapping()
	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
}

func (rB *PoliciesMappingsRequestBuilder) Post(ctx context.Context, requestConfiguration *PoliciesMappingsRequestBuilderPostRequestConfiguration) (core.ServiceNowItemResponse[*PoliciesMapping], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if conversion.IsNil(requestConfiguration) {
		return nil, errors.New("requestConfiguration is nil")
	}

	queryParameters := requestConfiguration.QueryParameters
	if queryParameters.AppName == "" {
		return nil, errors.New("AppName is required")
	}

	if queryParameters.DeployableName == "" {
		return nil, errors.New("DeployableName is required")
	}

	if queryParameters.PolicyName == "" {
		return nil, errors.New("PolicyName is required")
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[*PoliciesMapping](CreatePoliciesMappingsInputFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("response is nil")
	}

	typedResp, ok := resp.(core.ServiceNowItemResponse[*PoliciesMapping])
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*core.ServiceNowItemResponse[*PoliciesMapping])(nil))
	}

	return typedResp, nil
}

func (rB *PoliciesMappingsRequestBuilder) ToPostRequestInformation(_ context.Context, requestConfiguration *PoliciesMappingsRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)
	}

	return kiotaRequestInfo.RequestInformation, nil
}

func (rB *PoliciesMappingsRequestBuilder) ToDeleteRequestInformation(_ context.Context, requestConfiguration *PoliciesMappingsRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)
	}

	return kiotaRequestInfo.RequestInformation, nil
}
