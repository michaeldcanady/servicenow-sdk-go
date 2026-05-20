package cdmapplicationsapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	applicationsURLTemplate     = "{+baseurl}/api/sn_cdm/applications"
	deployablesURLTemplate      = "{+baseurl}/api/sn_cdm/applications/deployables{?appName,name}"
	sharedComponentsURLTemplate = "{+baseurl}/api/sn_cdm/applications/shared_components{?appName,sharedComponentName,name}"
	uploadStatusURLTemplate     = "{+baseurl}/api/sn_cdm/applications/upload-status/{upload_id}"
)

// ApplicationsRequestBuilder provides operations to manage applications.
type ApplicationsRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewApplicationsRequestBuilderInternal instantiates a new ApplicationsRequestBuilder.
func NewApplicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ApplicationsRequestBuilder {
	return &ApplicationsRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, applicationsURLTemplate, pathParameters),
	}
}

// Deployables returns a DeployablesRequestBuilder.
func (rB *ApplicationsRequestBuilder) Deployables() *DeployablesRequestBuilder {
	return NewDeployablesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// SharedComponents returns a SharedComponentsRequestBuilder.
func (rB *ApplicationsRequestBuilder) SharedComponents() *SharedComponentsRequestBuilder {
	return NewSharedComponentsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// UploadStatus returns a UploadStatusRequestBuilder.
func (rB *ApplicationsRequestBuilder) UploadStatus() *UploadStatusRequestBuilder {
	return NewUploadStatusRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// DeployablesRequestBuilder provides operations to manage deployables.
type DeployablesRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewDeployablesRequestBuilderInternal instantiates a new DeployablesRequestBuilder.
func NewDeployablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *DeployablesRequestBuilder {
	return &DeployablesRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, deployablesURLTemplate, pathParameters),
	}
}

// Delete deletes a deployable.
func (rB *DeployablesRequestBuilder) Delete(ctx context.Context, config *DeployablesRequestBuilderDeleteRequestConfiguration) error {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
		if config.QueryParameters != nil {
			kiotaRequestInfo.AddQueryParameters(config.QueryParameters)
		}
	}
	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, nil)
}

// SharedComponentsRequestBuilder provides operations to manage shared components.
type SharedComponentsRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewSharedComponentsRequestBuilderInternal instantiates a new SharedComponentsRequestBuilder.
func NewSharedComponentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SharedComponentsRequestBuilder {
	return &SharedComponentsRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, sharedComponentsURLTemplate, pathParameters),
	}
}

// Delete deletes shared component references.
func (rB *SharedComponentsRequestBuilder) Delete(ctx context.Context, config *SharedComponentsRequestBuilderDeleteRequestConfiguration) error {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
		if config.QueryParameters != nil {
			kiotaRequestInfo.AddQueryParameters(config.QueryParameters)
		}
	}
	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, nil)
}

// UploadStatusRequestBuilder provides operations to access upload status.
type UploadStatusRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewUploadStatusRequestBuilderInternal instantiates a new UploadStatusRequestBuilder.
func NewUploadStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadStatusRequestBuilder {
	return &UploadStatusRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, uploadStatusURLTemplate, pathParameters),
	}
}

// ByID returns a UploadStatusItemRequestBuilder.
func (rB *UploadStatusRequestBuilder) ByID(uploadId string) *UploadStatusItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["upload_id"] = uploadId
	return NewUploadStatusItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// UploadStatusItemRequestBuilder provides operations to access a specific upload status.
type UploadStatusItemRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewUploadStatusItemRequestBuilderInternal instantiates a new UploadStatusItemRequestBuilder.
func NewUploadStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadStatusItemRequestBuilder {
	return &UploadStatusItemRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, uploadStatusURLTemplate, pathParameters),
	}
}

// Get gets the status of a specific upload.
func (rB *UploadStatusItemRequestBuilder) Get(ctx context.Context, config *UploadStatusItemRequestBuilderGetRequestConfiguration) (UploadStatusResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateUploadStatusResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(UploadStatusResponse), nil
}
