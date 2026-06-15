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
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
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
func (rB *AccountRequestBuilder) ToGetRequestInformation(ctx context.Context, config *AccountRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if headers := config.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParameters := config.QueryParameters; !conversion.IsNil(queryParameters) {
			kiotaRequestInfo.AddQueryParameters(queryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return requestInfo, nil
}
