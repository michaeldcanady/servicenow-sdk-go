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
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if headers := config.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return requestInfo, nil
}
