package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	deleteURLTemplate = "{+baseurl}/api/now/documents/delete{?doc_sys_id,record_sys_id,table_name}"
)

// DeleteRequestBuilder provides operations to manage the delete endpoint.
type DeleteRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewDeleteRequestBuilderInternal instantiates a new DeleteRequestBuilder.
func NewDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *DeleteRequestBuilder {
	return &DeleteRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, deleteURLTemplate, pathParameters),
	}
}

// Delete removes a document.
func (rB *DeleteRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeleteRequestBuilderDeleteRequestConfiguration) error {
	if internal.IsNil(rB) {
		return nil
	}

	requestInfo, err := rB.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
}

// ToDeleteRequestInformation ...
func (rB *DeleteRequestBuilder) ToDeleteRequestInformation(_ context.Context, requestConfiguration *DeleteRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParams := requestConfiguration.QueryParameters; !internal.IsNil(queryParams) {
			kiotaRequestInfo.AddQueryParameters(queryParams)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
