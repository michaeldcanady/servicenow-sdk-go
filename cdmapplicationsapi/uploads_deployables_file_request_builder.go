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

const uploadsDeployablesFileURLTemplate = "{+baseurl}/api/sn_cdm/applications/uploads/deployables/file{?appName,deployableName}"

// UploadsDeployablesFileRequestBuilder provides operations to upload deployable files.
type UploadsDeployablesFileRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadsDeployablesFileRequestBuilderInternal instantiates a new [UploadsDeployablesFileRequestBuilder].
func NewUploadsDeployablesFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadsDeployablesFileRequestBuilder {
	return &UploadsDeployablesFileRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadsDeployablesFileURLTemplate, pathParameters),
	}
}

// NewUploadsDeployablesFileRequestBuilder instantiates a new [UploadsDeployablesFileRequestBuilder].
func NewUploadsDeployablesFileRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *UploadsDeployablesFileRequestBuilder {
	return NewUploadsDeployablesFileRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post uploads deployable files.
func (rB *UploadsDeployablesFileRequestBuilder) Post(ctx context.Context, media *Media, config *UploadsDeployablesFileRequestBuilderPostRequestConfiguration) (UploadStatusResponse, error) {
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
func (rB *UploadsDeployablesFileRequestBuilder) ToPostRequestInformation(_ context.Context, media *Media, config *UploadsDeployablesFileRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
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
