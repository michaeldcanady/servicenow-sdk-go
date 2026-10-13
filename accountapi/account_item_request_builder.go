// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package accountapi

import (
	"context"
	"fmt"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const accountItemURLTemplate = "{+baseurl}/api/now/v1/account/{account_id}"

// AccountItemRequestBuilder provides operations to manage a single account.
type AccountItemRequestBuilder struct {
	core.RequestBuilder
}

// NewAccountItemRequestBuilderInternal instantiates a new [AccountItemRequestBuilder] with the provided request parameters.
func NewAccountItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AccountItemRequestBuilder {
	return &AccountItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, accountItemURLTemplate, pathParameters),
	}
}

// NewAccountItemRequestBuilder instantiates a new [AccountItemRequestBuilder].
func NewAccountItemRequestBuilder(rawUrl string, requestAdapter abstractions.RequestAdapter) *AccountItemRequestBuilder {
	urlParams := map[string]string{internal.RawURLKey: rawUrl}
	return NewAccountItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get sends a GET request to retrieve a single account.
func (rB *AccountItemRequestBuilder) Get(ctx context.Context, config *AccountItemRequestBuilderGetRequestConfiguration) (AccountItemResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAccountItemResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedRes, ok := res.(AccountItemResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T", res)
	}

	return typedRes, nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *AccountItemRequestBuilder) ToGetRequestInformation(_ context.Context, config *AccountItemRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
