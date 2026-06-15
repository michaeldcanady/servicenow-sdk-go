package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	deleteURLTemplate = "{+baseurl}/api/now/v1/documents/delete{?doc_sys_id,record_sys_id,table_name}"
)

// DeleteRequestBuilder provides operations to manage the delete endpoint.
type DeleteRequestBuilder struct {
	internal.RequestBuilder
}

// NewDeleteRequestBuilderInternal instantiates a new DeleteRequestBuilder.
func NewDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *DeleteRequestBuilder {
	return &DeleteRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, deleteURLTemplate, pathParameters),
	}
}

// Delete removes a document.
func (rB *DeleteRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeleteRequestBuilderDeleteRequestConfiguration) error {
	if conversion.IsNil(rB) {
		return nil
	}

	requestInfo, err := rB.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
}

// ToDeleteRequestInformation ...
func (rB *DeleteRequestBuilder) ToDeleteRequestInformation(_ context.Context, requestConfiguration *DeleteRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParams := requestConfiguration.QueryParameters; !conversion.IsNil(queryParams) {
			kiotaRequestInfo.AddQueryParameters(queryParams)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
