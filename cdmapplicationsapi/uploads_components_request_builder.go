//nolint:dupl // per-verb request-builder methods share the mandatory nil-guard/send boilerplate by convention; each depends on its own outer type, response type, and discriminator factory, so it can't be extracted into a shared helper
package cdmapplicationsapi

import (
	"context"
	"fmt"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const uploadsComponentsURLTemplate = "{+baseurl}/api/sn_cdm/applications/uploads/components"

// UploadsComponentsRequestBuilder provides operations to upload components.
type UploadsComponentsRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadsComponentsRequestBuilderInternal instantiates a new [UploadsComponentsRequestBuilder].
func NewUploadsComponentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadsComponentsRequestBuilder {
	return &UploadsComponentsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadsComponentsURLTemplate, pathParameters),
	}
}

// NewUploadsComponentsRequestBuilder instantiates a new [UploadsComponentsRequestBuilder].
func NewUploadsComponentsRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *UploadsComponentsRequestBuilder {
	return NewUploadsComponentsRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post uploads components.
func (rB *UploadsComponentsRequestBuilder) Post(ctx context.Context, body *ComponentUploadRequest, config *UploadsComponentsRequestBuilderPostRequestConfiguration) (UploadStatusResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
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
func (rB *UploadsComponentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body *ComponentUploadRequest, config *UploadsComponentsRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}
	return requestInfo, nil
}

// Vars returns a [UploadsComponentsVarsRequestBuilder].
func (rB *UploadsComponentsRequestBuilder) Vars() *UploadsComponentsVarsRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewUploadsComponentsVarsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
