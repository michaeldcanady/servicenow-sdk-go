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

const uploadsComponentsVarsURLTemplate = "{+baseurl}/api/sn_cdm/applications/uploads/components/vars"

// UploadsComponentsVarsRequestBuilder provides operations to upload component variables.
type UploadsComponentsVarsRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadsComponentsVarsRequestBuilderInternal instantiates a new [UploadsComponentsVarsRequestBuilder].
func NewUploadsComponentsVarsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadsComponentsVarsRequestBuilder {
	return &UploadsComponentsVarsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadsComponentsVarsURLTemplate, pathParameters),
	}
}

// NewUploadsComponentsVarsRequestBuilder instantiates a new [UploadsComponentsVarsRequestBuilder].
func NewUploadsComponentsVarsRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *UploadsComponentsVarsRequestBuilder {
	return NewUploadsComponentsVarsRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post uploads component variables.
func (rB *UploadsComponentsVarsRequestBuilder) Post(ctx context.Context, body *ComponentVarsUploadRequest, config *UploadsComponentsVarsRequestBuilderPostRequestConfiguration) (UploadStatusResponse, error) {
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
		return nil, nil
	}
	typedRes, ok := res.(UploadStatusResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*UploadStatusResponse)(nil))
	}
	return typedRes, nil
}

// ToPostRequestInformation builds the request information for the Post method.
func (rB *UploadsComponentsVarsRequestBuilder) ToPostRequestInformation(ctx context.Context, body *ComponentVarsUploadRequest, config *UploadsComponentsVarsRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
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
