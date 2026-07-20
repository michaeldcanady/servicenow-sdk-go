package documentsapi

import (
	"context"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	deleteURLTemplate = "{+baseurl}/api/now/v1/documents/delete{?doc_sys_id,record_sys_id,table_name}"
)

// DeleteRequestBuilder provides operations to manage the delete endpoint.
type DeleteRequestBuilder struct {
	core.RequestBuilder
}

// NewDeleteRequestBuilderInternal instantiates a new DeleteRequestBuilder.
func NewDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *DeleteRequestBuilder {
	return &DeleteRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, deleteURLTemplate, pathParameters),
	}
}

// Delete removes a document.
func (rB *DeleteRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeleteRequestBuilderDeleteRequestConfiguration) error {
	if conversion.IsNil(rB) {
		return snerrors.ErrNilRequestBuilder
	}

	requestInfo, err := rB.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	errorMapping := core.DefaultErrorMapping()
	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
}

// ToDeleteRequestInformation ...
func (rB *DeleteRequestBuilder) ToDeleteRequestInformation(_ context.Context, requestConfiguration *DeleteRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(requestConfiguration) {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*requestConfiguration.QueryParameters)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
