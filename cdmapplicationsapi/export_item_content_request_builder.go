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

const exportItemContentURLTemplate = "{+baseurl}/api/sn_cdm/applications/deployables/exports/{export_id}/content"

// ExportItemContentRequestBuilder provides operations to download export content.
type ExportItemContentRequestBuilder struct {
	core.RequestBuilder
}

// NewExportItemContentRequestBuilderInternal instantiates a new [ExportItemContentRequestBuilder].
func NewExportItemContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExportItemContentRequestBuilder {
	return &ExportItemContentRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, exportItemContentURLTemplate, pathParameters),
	}
}

// NewExportItemContentRequestBuilder instantiates a new [ExportItemContentRequestBuilder].
func NewExportItemContentRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *ExportItemContentRequestBuilder {
	return NewExportItemContentRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Get fetches the export content.
func (rB *ExportItemContentRequestBuilder) Get(ctx context.Context, config *ExportItemContentRequestBuilderGetRequestConfiguration) ([]byte, error) {
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

	res, err := rB.GetRequestAdapter().SendPrimitive(ctx, requestInfo, "[]byte", core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.([]byte)
	if !ok {
		return nil, fmt.Errorf("res is not %T", []byte(nil))
	}
	return typedRes, nil
}

// ToGetRequestInformation builds the request information for the Get method.
func (rB *ExportItemContentRequestBuilder) ToGetRequestInformation(_ context.Context, config *ExportItemContentRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationOctetStream)
	return requestInfo, nil
}
