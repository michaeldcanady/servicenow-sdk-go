package documentsapi

import (
	"context"
	"maps"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	actionURLTemplate = "{+baseurl}/api/now/v1/documents/action/{action}/document/{documentSysId}/version/{versionSysId}"
)

// ActionRequestBuilder provides operations to manage document actions.
type ActionRequestBuilder struct {
	core.RequestBuilder
}

// NewActionRequestBuilderInternal instantiates a new ActionRequestBuilder.
func NewActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ActionRequestBuilder {
	return &ActionRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, actionURLTemplate, pathParameters),
	}
}

// Document returns a DocumentActionRequestBuilder for the specified document.
func (rB *ActionRequestBuilder) Document(documentSysID string) *DocumentActionRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["documentSysId"] = documentSysID
	return NewDocumentActionRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// DocumentActionRequestBuilder ...
type DocumentActionRequestBuilder struct {
	core.RequestBuilder
}

// NewDocumentActionRequestBuilderInternal ...
func NewDocumentActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *DocumentActionRequestBuilder {
	return &DocumentActionRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, actionURLTemplate, pathParameters),
	}
}

// Version returns a VersionActionRequestBuilder for the specified version.
func (rB *DocumentActionRequestBuilder) Version(versionSysID string) *VersionActionRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["versionSysId"] = versionSysID
	return NewVersionActionRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// VersionActionRequestBuilder ...
type VersionActionRequestBuilder struct {
	core.RequestBuilder
}

// NewVersionActionRequestBuilderInternal ...
func NewVersionActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *VersionActionRequestBuilder {
	return &VersionActionRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, actionURLTemplate, pathParameters),
	}
}

// Patch executes the specified action on the document version.
func (rB *VersionActionRequestBuilder) Patch(ctx context.Context, requestConfiguration *VersionActionRequestBuilderPatchRequestConfiguration) error {
	if conversion.IsNil(rB) {
		return snerrors.ErrNilRequestBuilder
	}

	requestInfo, err := rB.ToPatchRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	errorMapping := core.DefaultErrorMapping()
	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
}

// ToPatchRequestInformation converts request configurations to Patch request information.
func (rB *VersionActionRequestBuilder) ToPatchRequestInformation(ctx context.Context, requestConfiguration *VersionActionRequestBuilderPatchRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PATCH, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(requestConfiguration) {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
		if data := requestConfiguration.Data; !conversion.IsNil(data) {
			if err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), data); err != nil {
				return nil, err
			}
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
