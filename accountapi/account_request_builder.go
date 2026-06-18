package accountapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const accountURLTemplate = "{+baseurl}/api/now/v1/account"

// AccountRequestBuilder provides operations to manage accounts.
type AccountRequestBuilder struct {
	internal.RequestBuilder
}

// NewAccountRequestBuilderInternal instantiates a new AccountRequestBuilder with the provided request parameters.
func NewAccountRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AccountRequestBuilder {
	return &AccountRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, accountURLTemplate, pathParameters),
	}
}

// ByID returns an AccountItemRequestBuilder for the specified account ID.
func (rB *AccountRequestBuilder) ByID(accountID string) *AccountItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["account_id"] = accountID
	return NewAccountItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Get sends a GET request to retrieve a list of accounts.
func (rB *AccountRequestBuilder) Get(ctx context.Context, config *AccountRequestBuilderGetRequestConfiguration) (AccountCollectionResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	errorMapping := internal.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAccountCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(AccountCollectionResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *AccountRequestBuilder) ToGetRequestInformation(_ context.Context, config *AccountRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
