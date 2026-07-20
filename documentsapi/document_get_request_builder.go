package documentsapi

import (
	"context"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// documentGetRequestConfiguration is the shape shared by every documentsapi GET
// request configuration used by a plain single-document GET (VersionState, Versions)
// - they only differ by exported type name, so each one is convertible to this via
// a pointer cast.
type documentGetRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
}

// documentGetRequestBuilder is the shared implementation behind every documentsapi
// request builder whose only operation is a GET returning a single Document
// (VersionState, Versions) - they differ only in URL template.
type documentGetRequestBuilder struct {
	core.RequestBuilder
}

// newDocumentGetRequestBuilder instantiates a new documentGetRequestBuilder.
func newDocumentGetRequestBuilder(requestAdapter abstractions.RequestAdapter, urlTemplate string, pathParameters map[string]string) *documentGetRequestBuilder {
	return &documentGetRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, urlTemplate, pathParameters),
	}
}

// get retrieves a Document.
func (rB *documentGetRequestBuilder) get(ctx context.Context, requestConfiguration *documentGetRequestConfiguration) (*core.BaseServiceNowItemResponse[Document], error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo, err := rB.toGetRequestInformation(ctx, requestConfiguration)
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

// toGetRequestInformation converts request configurations to Get request information.
func (rB *documentGetRequestBuilder) toGetRequestInformation(_ context.Context, requestConfiguration *documentGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(requestConfiguration) {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
