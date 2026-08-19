//nolint:dupl // per-verb request-builder methods share the mandatory nil-guard/send boilerplate by convention; each depends on its own outer type, response type, and discriminator factory, so it can't be extracted into a shared helper
package cdmapplicationsapi

import (
	"context"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const uploadsCollectionsFileURLTemplate = "{+baseurl}/api/sn_cdm/applications/uploads/collections/file{?appName,collectionName}"

// UploadsCollectionsFileRequestBuilder provides operations to upload collection files.
type UploadsCollectionsFileRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadsCollectionsFileRequestBuilderInternal instantiates a new [UploadsCollectionsFileRequestBuilder].
func NewUploadsCollectionsFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadsCollectionsFileRequestBuilder {
	return &UploadsCollectionsFileRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadsCollectionsFileURLTemplate, pathParameters),
	}
}

// NewUploadsCollectionsFileRequestBuilder instantiates a new [UploadsCollectionsFileRequestBuilder].
func NewUploadsCollectionsFileRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *UploadsCollectionsFileRequestBuilder {
	return NewUploadsCollectionsFileRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post uploads collection files using a stream payload (like attachment-api)
func (rB *UploadsCollectionsFileRequestBuilder) Post(ctx context.Context, media *Media, config *UploadsCollectionsFileRequestBuilderPostRequestConfiguration) (UploadStatusResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, media, config)
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

// ToPostRequestInformation builds the request information for the Post method.
func (rB *UploadsCollectionsFileRequestBuilder) ToPostRequestInformation(_ context.Context, media *Media, config *UploadsCollectionsFileRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
		if config.QueryParameters != nil {
			requestInfo.AddQueryParameters(*config.QueryParameters)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	requestInfo.SetStreamContentAndContentType(media.GetData(), media.GetContentType())
	return requestInfo, nil
}
