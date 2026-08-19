package documentsapi

import (
	"context"
	"fmt"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	contentURLTemplate = "{+baseurl}/api/now/v1/documents/{document_sys_id}/content"
)

// ContentRequestBuilder provides operations to manage document content.
type ContentRequestBuilder struct {
	core.RequestBuilder
}

// NewContentRequestBuilderInternal instantiates a new ContentRequestBuilder.
func NewContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ContentRequestBuilder {
	return &ContentRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, contentURLTemplate, pathParameters),
	}
}

// Get fetches and streams the default version attachment for the document.
func (rB *ContentRequestBuilder) Get(ctx context.Context, requestConfiguration *ContentRequestBuilderGetRequestConfiguration) ([]byte, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := res.([]byte)
	if !ok {
		return nil, fmt.Errorf("resp is not %T", ([]byte)(nil))
	}

	return typedResp, nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *ContentRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *ContentRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(requestConfiguration) {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationOctetStream)

	return requestInfo, nil
}
