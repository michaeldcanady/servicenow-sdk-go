package accountapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const accountItemURLTemplate = "{+baseurl}/api/now/v1/account/{account_id}"

// AccountItemRequestBuilder provides operations to manage a single account.
type AccountItemRequestBuilder struct {
	internal.RequestBuilder
}

// NewAccountItemRequestBuilderInternal instantiates a new AccountItemRequestBuilder with the provided request parameters.
func NewAccountItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AccountItemRequestBuilder {
	return &AccountItemRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, accountItemURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve a single account.
func (rB *AccountItemRequestBuilder) Get(ctx context.Context, config *AccountItemRequestBuilderGetRequestConfiguration) (AccountItemResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	errorMapping := internal.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAccountItemResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(AccountItemResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *AccountItemRequestBuilder) ToGetRequestInformation(_ context.Context, config *AccountItemRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
