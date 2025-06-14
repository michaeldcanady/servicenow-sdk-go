package attachmentapi

import (
	"context"
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

const (
	// attachmentURLTemplate url template for the Service-Now Attachment endpoint
	attachmentURLTemplate = "{+baseurl}/api/now/v1/attachment{?sysparm_limit,sysparm_offset,sysparm_query}"
)

// AttachmentRequestBuilder2 provides operations to manage Service-Now attachments.
type AttachmentRequestBuilder2 struct {
	newInternal.RequestBuilder
}

// newAttachmentRequestBuilder2Internal instantiates a new AttachmentRequestBuilder2 with the provided requestBuilder
func newAttachmentRequestBuilder2Internal(requestBuilder newInternal.RequestBuilder) *AttachmentRequestBuilder2 {
	m := &AttachmentRequestBuilder2{
		requestBuilder,
	}
	return m
}

// NewV1CompatibleAttachmentRequestBuilder2 instantiates a new AttachmentRequestBuilder2 with custom parsable for table entries.
func NewV1CompatibleAttachmentRequestBuilder2(
	pathParameters map[string]string,
	client core.Client,
) *AttachmentRequestBuilder2 {

	authProvider := core.NewAPIV1ClientAdapter(client)
	adapter, _ := nethttplibrary.NewNetHttpRequestAdapter(authProvider)

	return newAttachmentRequestBuilder2Internal(

		newInternal.NewBaseRequestBuilder(adapter, attachmentURLTemplate, pathParameters),
	)
}

// NewAttachmentRequestBuilder2Internal instantiates a new AttachmentRequestBuilder2 with custom parsable for table entries.
func NewAttachmentRequestBuilder2Internal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentRequestBuilder2 {
	return newAttachmentRequestBuilder2Internal(
		newInternal.NewBaseRequestBuilder(requestAdapter, attachmentURLTemplate, pathParameters),
	)
}

// NewAttachmentRequestBuilder2 instantiates a new AttachmentRequestBuilder2 with custom parsable for table entries.
func NewAttachmentRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentRequestBuilder2 {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewAttachmentRequestBuilder2Internal(urlParams, requestAdapter)
}

// ByID provides the way to manage attachment item with provided sys id
func (rB *AttachmentRequestBuilder2) ByID(sysID string) *AttachmentItemRequestBuilder {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters[sysIDKey] = sysID

	return NewAttachmentItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// File provides the way to access Service-Now's attachment file API
func (rB *AttachmentRequestBuilder2) File() *AttachmentFileRequestBuilder {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())

	return NewAttachmentFileRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Upload provides the way to access Service-Now's attachment upload API
func (rB *AttachmentRequestBuilder2) Upload() *AttachmentUploadRequestBuilder {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())

	return NewAttachmentUploadRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Get returns AttachmentCollectionResponse using provided arguments
func (rB *AttachmentRequestBuilder2) Get(ctx context.Context, requestConfiguration *AttachmentRequestBuilder2GetRequestConfiguration) (*AttachmentCollectionResponse2Model, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &AttachmentRequestBuilder2GetRequestConfiguration{}
	}

	opts := nethttplibrary.NewHeadersInspectionOptions()
	opts.InspectResponseHeaders = true

	requestConfiguration.Options = append(requestConfiguration.Options, opts)

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAttachmentCollectionResponse2FromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	snRes, ok := res.(*AttachmentCollectionResponse2Model)
	if !ok {
		return nil, errors.New("res is not *AttachmentCollectionResponse2Model")
	}

	return snRes, nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *AttachmentRequestBuilder2) ToGetRequestInformation(_ context.Context, requestConfiguration *AttachmentRequestBuilder2GetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if parameter := requestConfiguration.QueryParameters; !internal.IsNil(parameter) {
			kiotaRequestInfo.AddQueryParameters(parameter)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return kiotaRequestInfo.RequestInformation, nil
}
