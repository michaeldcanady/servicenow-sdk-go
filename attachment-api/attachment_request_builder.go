package attachmentapi

import (
	"context"
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

const (
	// attachmentURLTemplate url template for the Service-Now Attachment endpoint
	attachmentURLTemplate = "{+baseurl}/api/now/v1/attachment{?sysparm_limit,sysparm_offset,sysparm_query}"
)

// AttachmentRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentRequestBuilder struct {
	internal.RequestBuilder
}

// newAttachmentRequestBuilderInternal instantiates a new AttachmentRequestBuilder with the provided requestBuilder
func newAttachmentRequestBuilderInternal(requestBuilder internal.RequestBuilder) *AttachmentRequestBuilder {
	m := &AttachmentRequestBuilder{
		requestBuilder,
	}
	return m
}

// NewAttachmentRequestBuilderInternal instantiates a new AttachmentRequestBuilder with custom parsable for table entries.
func NewAttachmentRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentRequestBuilder {
	return newAttachmentRequestBuilderInternal(
		internal.NewBaseRequestBuilder(requestAdapter, attachmentURLTemplate, pathParameters),
	)
}

// NewAttachmentRequestBuilder instantiates a new AttachmentRequestBuilder with custom parsable for table entries.
func NewAttachmentRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[internal.RawURLKey] = rawURL
	return NewAttachmentRequestBuilderInternal(urlParams, requestAdapter)
}

// ByID provides the way to manage attachment item with provided sys id
func (rB *AttachmentRequestBuilder) ByID(sysID string) *AttachmentItemRequestBuilder {
	if conversion.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters[sysIDKey] = sysID

	return NewAttachmentItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// File provides the way to access Service-Now's attachment file API
func (rB *AttachmentRequestBuilder) File() *AttachmentFileRequestBuilder {
	if conversion.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())

	return NewAttachmentFileRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Upload provides the way to access Service-Now's attachment upload API
func (rB *AttachmentRequestBuilder) Upload() *AttachmentUploadRequestBuilder {
	if conversion.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())

	return NewAttachmentUploadRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Get returns AttachmentCollectionResponse using provided arguments
func (rB *AttachmentRequestBuilder) Get(ctx context.Context, requestConfiguration *AttachmentRequestBuilderGetRequestConfiguration) (*AttachmentCollectionResponse, error) {
	if conversion.IsNil(rB) {
		return nil, nil
	}

	if conversion.IsNil(requestConfiguration) {
		requestConfiguration = &AttachmentRequestBuilderGetRequestConfiguration{}
	}

	headerOpt := nethttplibrary.NewHeadersInspectionOptions()
	headerOpt.InspectResponseHeaders = true

	requestConfiguration.Options = append(requestConfiguration.Options, headerOpt)

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAttachmentCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	snRes, ok := res.(*AttachmentCollectionResponse)
	if !ok {
		return nil, errors.New("res is not *AttachmentCollectionResponse")
	}

	internal.ParseHeaders(snRes, headerOpt.GetResponseHeaders())

	return snRes, nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *AttachmentRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *AttachmentRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
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
