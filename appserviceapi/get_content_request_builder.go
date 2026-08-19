package appserviceapi

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

const getContentURLTemplate = "{+baseurl}/api/now/cmdb/app_service/{sys_id}/getContent{?mode}"

// GetContentRequestBuilder provides operations to retrieve the content of an application service.
type GetContentRequestBuilder struct {
	core.RequestBuilder
}

// NewGetContentRequestBuilderInternal instantiates a new [GetContentRequestBuilder].
func NewGetContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *GetContentRequestBuilder {
	return &GetContentRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, getContentURLTemplate, pathParameters),
	}
}

// NewGetContentRequestBuilder instantiates a new [GetContentRequestBuilder].
func NewGetContentRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *GetContentRequestBuilder {
	return NewGetContentRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Get sends a GET request to retrieve the content of an application service.
func (rB *GetContentRequestBuilder) Get(ctx context.Context, config *GetContentRequestConfiguration) (GetContentResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateGetContentResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := res.(GetContentResponse)
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*GetContentResponse)(nil))
	}

	return typedResp, nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *GetContentRequestBuilder) ToGetRequestInformation(_ context.Context, config *GetContentRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	return requestInfo, nil
}
