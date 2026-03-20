package policyapi

import (
	"context"
	"errors"
	"fmt"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	policiesMappingsURLTemplate = "{+baseurl}/api/sn_cdm/v1/policies/mappings"
)

// PoliciesMappingsRequestBuilder provides operations to manage Service-Now policy definitions.
type PoliciesMappingsRequestBuilder struct {
	kiota.RequestBuilder
}

// NewPoliciesMappingsRequestBuilderInternal instantiates a new PoliciesMappingsRequestBuilder with the provided path parameters and request adapter.
func NewPoliciesMappingsRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *PoliciesMappingsRequestBuilder {
	return &PoliciesMappingsRequestBuilder{
		RequestBuilder: kiota.NewBaseRequestBuilder(requestAdapter, policiesMappingsURLTemplate, pathParameters),
	}
}

func (rB *PoliciesMappingsRequestBuilder) Inputs() *PoliciesMappingsInputsRequestBuilder {
	if utils.IsNil(rB) {
		return nil
	}

	return NewPoliciesMappingsInputsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

func (rB *PoliciesMappingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *PoliciesMappingsRequestBuilderDeleteRequestConfiguration) error {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil
	}

	if utils.IsNil(requestConfiguration) {
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

	errorMapping := abstractions.ErrorMappings{
		"XXX": model.CreateServiceNowErrorFromDiscriminatorValue,
	}

	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
}

func (rB *PoliciesMappingsRequestBuilder) Post(ctx context.Context, requestConfiguration *PoliciesMappingsRequestBuilderPostRequestConfiguration) (model.ServiceNowItemResponse[*PoliciesMapping], error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if utils.IsNil(requestConfiguration) {
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

	errorMapping := abstractions.ErrorMappings{
		"XXX": model.CreateServiceNowErrorFromDiscriminatorValue,
	}

	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, model.ServiceNowItemResponseFromDiscriminatorValue[*PoliciesMapping](CreatePoliciesMappingsInputFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("response is nil")
	}

	typedResp, ok := resp.(model.ServiceNowItemResponse[*PoliciesMapping])
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*model.ServiceNowItemResponse[*PoliciesMapping])(nil))
	}

	return typedResp, nil
}

func (rB *PoliciesMappingsRequestBuilder) ToPostRequestInformation(_ context.Context, requestConfiguration *PoliciesMappingsRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &kiota.KiotaRequestInformation{RequestInformation: requestInfo}
	if !utils.IsNil(requestConfiguration) {
		kiota.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)
	}

	return kiotaRequestInfo.RequestInformation, nil
}

func (rB *PoliciesMappingsRequestBuilder) ToDeleteRequestInformation(_ context.Context, requestConfiguration *PoliciesMappingsRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &kiota.KiotaRequestInformation{RequestInformation: requestInfo}
	if !utils.IsNil(requestConfiguration) {
		kiota.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)
	}

	return kiotaRequestInfo.RequestInformation, nil
}
