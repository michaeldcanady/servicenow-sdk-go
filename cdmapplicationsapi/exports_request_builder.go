package cdmapplicationsapi

import (
	"context"
	"fmt"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const exportsURLTemplate = "{+baseurl}/api/sn_cdm/applications/deployables/exports{?appName,deployableName}"

// ExportsRequestBuilder provides operations to manage deployable exports.
type ExportsRequestBuilder struct {
	core.RequestBuilder
}

// NewExportsRequestBuilderInternal instantiates a new [ExportsRequestBuilder].
func NewExportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExportsRequestBuilder {
	return &ExportsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, exportsURLTemplate, pathParameters),
	}
}

// NewExportsRequestBuilder instantiates a new [ExportsRequestBuilder].
func NewExportsRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *ExportsRequestBuilder {
	return NewExportsRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Get gets the collection of deployable exports.
func (rB *ExportsRequestBuilder) Get(ctx context.Context, config *ExportsRequestBuilderGetRequestConfiguration) (ExportsResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateExportsResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(ExportsResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*ExportsResponse)(nil))
	}
	return typedRes, nil
}

// ToGetRequestInformation builds the request information for the Get method.
func (rB *ExportsRequestBuilder) ToGetRequestInformation(_ context.Context, config *ExportsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
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
	return requestInfo, nil
}

// ByID returns a [ExportItemRequestBuilder].
func (rB *ExportsRequestBuilder) ByID(exportID string) *ExportItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["export_id"] = exportID
	return NewExportItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
