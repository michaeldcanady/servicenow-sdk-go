package accountapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const accountItemURLTemplate = "{+baseurl}/api/now/v1/account/{account_id}"

// AccountItemRequestBuilder provides operations to manage a single account.
type AccountItemRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewAccountItemRequestBuilderInternal instantiates a new AccountItemRequestBuilder with the provided request parameters.
func NewAccountItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AccountItemRequestBuilder {
	return &AccountItemRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, accountItemURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve a single account.
func (rB *AccountItemRequestBuilder) Get(ctx context.Context, config *AccountItemRequestBuilderGetRequestConfiguration) (AccountItemResponse, error) {
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAccountItemResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(AccountItemResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *AccountItemRequestBuilder) ToGetRequestInformation(ctx context.Context, config *AccountItemRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if headers := config.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return requestInfo, nil
}
