package accountapi

import (
	"context"

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

// AccountItemRequestBuilderGetRequestConfiguration represents the configuration for a GET request.
type AccountItemRequestBuilderGetRequestConfiguration struct {
	Header  *abstractions.RequestHeaders
	Options []abstractions.RequestOption
}

// Get sends a GET request to retrieve a single account.
func (rB *AccountItemRequestBuilder) Get(ctx context.Context, config *AccountItemRequestBuilderGetRequestConfiguration) (AccountItemResponse, error) {
	requestInfo, err := rB.ToGetRequestInformation(ctx)
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
func (rB *AccountItemRequestBuilder) ToGetRequestInformation(ctx context.Context) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformation()
	requestInfo.Method = abstractions.GET
	requestInfo.UrlTemplate = rB.GetURLTemplate()
	requestInfo.PathParameters = rB.GetPathParameters()

	return requestInfo, nil
}
