package accountapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// accountItemURLTemplate is the URL template for this endpoint.
	accountItemURLTemplate = "{+baseurl}/api/now/v1/account/{sys_id}"
)

// AccountItemRequestBuilder provides operations for managing the AccountItem endpoint.
type AccountItemRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewAccountItemRequestBuilderInternal instantiates a new AccountItemRequestBuilder with the provided path parameters and request adapter.
func NewAccountItemRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *AccountItemRequestBuilder {
	return &AccountItemRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, accountItemURLTemplate, pathParameters),
	}
}

// NewAccountItemRequestBuilder instantiates a new AccountItemRequestBuilder with a raw URL.
func NewAccountItemRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AccountItemRequestBuilder {
	urlParams := map[string]string{newInternal.RawURLKey: rawURL}
	return NewAccountItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get sends a GET request to the endpoint.
func (rB *AccountItemRequestBuilder) Get(
	ctx context.Context,
	requestConfiguration *AccountItemRequestBuilderGetRequestConfiguration,
) (
	newInternal.ServiceNowItemResponse[Account],
	error,
) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &AccountItemRequestBuilderGetRequestConfiguration{}
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(
		ctx,
		requestInfo,
		newInternal.ServiceNowItemResponseFromDiscriminatorValue[Account](CreateAccountFromDiscriminatorValue),
		errorMapping,
	)

	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	typedRes, ok := res.(newInternal.ServiceNowItemResponse[Account])
	if !ok {
		return nil, errors.New("unexpected response type")
	}

	return typedRes, nil
}

// ToGetRequestInformation creates the RequestInformation for a GET request.
func (rB *AccountItemRequestBuilder) ToGetRequestInformation(
	ctx context.Context,
	requestConfiguration *AccountItemRequestBuilderGetRequestConfiguration,
) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(
		abstractions.GET,
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

	kiotaRequestInfo.Headers.TryAdd(
		internalHttp.RequestHeaderAccept.String(),
		newInternal.ContentTypeApplicationJSON,
	)

	return kiotaRequestInfo.RequestInformation, nil
}
