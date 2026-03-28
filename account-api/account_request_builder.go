package accountapi

import (
	"context"
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

const (
	// accountURLTemplate is the URL template for this endpoint.
	accountURLTemplate = "{+baseurl}/api/now/v1/account{?sysparm_limit,sysparm_offset,sysparm_query}"
)

// AccountRequestBuilder provides operations for managing the Account endpoint.
type AccountRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewAccountRequestBuilderInternal instantiates a new AccountRequestBuilder with the provided path parameters and request adapter.
func NewAccountRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,

) *AccountRequestBuilder {
	return &AccountRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, accountURLTemplate, pathParameters),
	}
}

// NewAccountRequestBuilder instantiates a new AccountRequestBuilder with a raw URL.
func NewAccountRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,

) *AccountRequestBuilder {
	urlParams := map[string]string{newInternal.RawURLKey: rawURL}
	return NewAccountRequestBuilderInternal(urlParams, requestAdapter)
}

// ByID provides access to an individual Account record via its sysID.
func (rB *AccountRequestBuilder) ByID(sysID string) *AccountItemRequestBuilder {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["sys_id"] = sysID

	return NewAccountItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Get sends a GET request to the endpoint.
func (rB *AccountRequestBuilder) Get(
	ctx context.Context,
	requestConfiguration *AccountRequestBuilderGetRequestConfiguration,
) (
	newInternal.ServiceNowCollectionResponse[Account],
	error,
) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &AccountRequestBuilderGetRequestConfiguration{}
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

	res, err := rB.GetRequestAdapter().Send(
		ctx,
		requestInfo,
		newInternal.ServiceNowCollectionResponseFromDiscriminatorValue[Account](CreateAccountFromDiscriminatorValue),
		errorMapping,
	)

	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	typedRes, ok := res.(newInternal.ServiceNowCollectionResponse[Account])
	if !ok {
		return nil, errors.New("unexpected response type")
	}

	newInternal.ParseHeaders(typedRes, headerOpt.GetResponseHeaders())

	return typedRes, nil
}

// ToGetRequestInformation creates the RequestInformation for a GET request.
func (rB *AccountRequestBuilder) ToGetRequestInformation(
	ctx context.Context,

	requestConfiguration *AccountRequestBuilderGetRequestConfiguration,
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
