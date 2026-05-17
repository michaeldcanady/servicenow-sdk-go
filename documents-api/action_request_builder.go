package documentsapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	actionURLTemplate = "{+baseurl}/api/now/documents/action/{action}/document/{documentSysId}/version/{versionSysId}"
)

// ActionRequestBuilder provides operations to manage document actions.
type ActionRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewActionRequestBuilderInternal instantiates a new ActionRequestBuilder.
func NewActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ActionRequestBuilder {
	return &ActionRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, actionURLTemplate, pathParameters),
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
	newInternal.RequestBuilder
}

// NewDocumentActionRequestBuilderInternal ...
func NewDocumentActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *DocumentActionRequestBuilder {
	return &DocumentActionRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, actionURLTemplate, pathParameters),
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
	newInternal.RequestBuilder
}

// NewVersionActionRequestBuilderInternal ...
func NewVersionActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *VersionActionRequestBuilder {
	return &VersionActionRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, actionURLTemplate, pathParameters),
	}
}

// Patch executes the specified action on the document version.
func (rB *VersionActionRequestBuilder) Patch(ctx context.Context, requestConfiguration *VersionActionRequestBuilderPatchRequestConfiguration) error {
	if internal.IsNil(rB) {
		return nil
	}

	requestInfo, err := rB.ToPatchRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
}

// ToPatchRequestInformation converts request configurations to Patch request information.
func (rB *VersionActionRequestBuilder) ToPatchRequestInformation(ctx context.Context, requestConfiguration *VersionActionRequestBuilderPatchRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PATCH, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if data := requestConfiguration.Data; !internal.IsNil(data) {
			err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), newInternal.ContentTypeApplicationJSON, data)
			if err != nil {
				return nil, err
			}
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
