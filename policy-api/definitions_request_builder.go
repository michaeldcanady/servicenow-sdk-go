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

// Get sends an HTTP GET request and returns a collection of policy definitions.
func (rB *DefinitionsRequestBuilder) Get(ctx context.Context, requestConfiguration *DefinitionsRequestBuilderGetRequestConfiguration) (newInternal.ServiceNowCollectionResponse[*PolicyDefinition], error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &DefinitionsRequestBuilderGetRequestConfiguration{}
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

	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowCollectionResponseFromDiscriminatorValue[*PolicyDefinition](CreatePolicyDefinitionFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("response is nil")
	}

	typedResp, ok := resp.(newInternal.ServiceNowCollectionResponse[*PolicyDefinition])
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*newInternal.ServiceNowCollectionResponse[*PolicyDefinition])(nil))
	}

	newInternal.ParseHeaders(typedResp, headerOpt.GetResponseHeaders())

	return typedResp, nil
}

// ToGetRequestInformation converts provided parameters into request information.
func (rB *DefinitionsRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *DefinitionsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
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
