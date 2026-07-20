package documentsapi

import (
	"context"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// documentPostRequestConfiguration is the shape shared by every documentsapi POST
// request configuration (Attach, CreateDocument, Create, SyncDown) - they only
// differ by exported type name, so each one is convertible to this via a pointer cast.
type documentPostRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
	Data    Document
}

// documentPostRequestBuilder is the shared implementation behind every documentsapi
// request builder whose only operation is a POST returning a single Document
// (Attach, CreateDocument, Create, SyncDown) - they differ only in URL template.
type documentPostRequestBuilder struct {
	core.RequestBuilder
}

// newDocumentPostRequestBuilder instantiates a new documentPostRequestBuilder.
func newDocumentPostRequestBuilder(requestAdapter abstractions.RequestAdapter, urlTemplate string, pathParameters map[string]string) *documentPostRequestBuilder {
	return &documentPostRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, urlTemplate, pathParameters),
	}
}

// post sends a POST request to create or link a document.
func (rB *documentPostRequestBuilder) post(ctx context.Context, requestConfiguration *documentPostRequestConfiguration) (*core.BaseServiceNowItemResponse[Document], error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo, err := rB.toPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[Document](CreateDocumentFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*core.BaseServiceNowItemResponse[Document]), nil
}

// toPostRequestInformation converts request configurations to Post request information.
func (rB *documentPostRequestBuilder) toPostRequestInformation(ctx context.Context, requestConfiguration *documentPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(requestConfiguration) {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
		if data := requestConfiguration.Data; !conversion.IsNil(data) {
			if err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), data); err != nil {
				return nil, err
			}
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
