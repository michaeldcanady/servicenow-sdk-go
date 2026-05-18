package accountapi

import (
	"context"
	"maps"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const accountURLTemplate = "{+baseurl}/api/now/v1/account"

// AccountRequestBuilder provides operations to manage accounts.
type AccountRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewAccountRequestBuilderInternal instantiates a new AccountRequestBuilder with the provided request parameters.
func NewAccountRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AccountRequestBuilder {
	return &AccountRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, accountURLTemplate, pathParameters),
	}
}

// ByID returns an AccountItemRequestBuilder for the specified account ID.
func (rB *AccountRequestBuilder) ByID(accountID string) *AccountItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["account_id"] = accountID
	return NewAccountItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// AccountRequestBuilderGetQueryParameters represents the query parameters for a GET request.
type AccountRequestBuilderGetQueryParameters struct{}

// AccountRequestBuilderGetRequestConfiguration represents the configuration for a GET request.
type AccountRequestBuilderGetRequestConfiguration struct {
	Header          *abstractions.RequestHeaders
	Options         []abstractions.RequestOption
	QueryParameters *AccountRequestBuilderGetQueryParameters
}

// Get sends a GET request to retrieve a list of accounts.
func (rB *AccountRequestBuilder) Get(ctx context.Context, config *AccountRequestBuilderGetRequestConfiguration) (AccountCollectionResponse, error) {
	var queryParams *AccountRequestBuilderGetQueryParameters
	if config != nil {
		queryParams = config.QueryParameters
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, queryParams)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAccountCollectionResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(AccountCollectionResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *AccountRequestBuilder) ToGetRequestInformation(ctx context.Context, queryParams *AccountRequestBuilderGetQueryParameters) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformation()
	requestInfo.Method = abstractions.GET
	requestInfo.UrlTemplate = rB.GetURLTemplate()
	requestInfo.PathParameters = rB.GetPathParameters()

	if queryParams != nil {
		requestInfo.AddQueryParameters(queryParams)
	}

	return requestInfo, nil
}
