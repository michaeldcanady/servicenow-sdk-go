package cdmapplicationsapi

import (
	"context"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// UploadStatusItemRequestBuilder provides operations to access a specific upload status.
type UploadStatusItemRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadStatusItemRequestBuilderInternal instantiates a new [UploadStatusItemRequestBuilder].
func NewUploadStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadStatusItemRequestBuilder {
	return &UploadStatusItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadStatusURLTemplate, pathParameters),
	}
}

// NewUploadStatusItemRequestBuilder instantiates a new [UploadStatusItemRequestBuilder].
func NewUploadStatusItemRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *UploadStatusItemRequestBuilder {
	return NewUploadStatusItemRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Get gets the status of a specific upload.
func (rB *UploadStatusItemRequestBuilder) Get(ctx context.Context, config *UploadStatusItemRequestBuilderGetRequestConfiguration) (UploadStatusResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateUploadStatusResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(UploadStatusResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*UploadStatusResponse)(nil))
	}
	return typedRes, nil
}

// ToGetRequestInformation builds the request information for the Get method.
func (rB *UploadStatusItemRequestBuilder) ToGetRequestInformation(_ context.Context, config *UploadStatusItemRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	return requestInfo, nil
}
