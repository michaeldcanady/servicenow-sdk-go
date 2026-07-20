package cdmapplicationsapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	applicationsURLTemplate                          = "{+baseurl}/api/sn_cdm/applications"
	deployablesURLTemplate                           = "{+baseurl}/api/sn_cdm/applications/deployables{?appName,name}"
	sharedComponentsURLTemplate                      = "{+baseurl}/api/sn_cdm/applications/shared_components{?appName,sharedComponentName,name}"
	uploadStatusURLTemplate                          = "{+baseurl}/api/sn_cdm/applications/upload-status/{upload_id}"
	exportsURLTemplate                               = "{+baseurl}/api/sn_cdm/applications/deployables/exports{?appName,deployableName}"
	exportItemURLTemplate                            = "{+baseurl}/api/sn_cdm/applications/deployables/exports/{export_id}"
	exportItemStatusURLTemplate                      = "{+baseurl}/api/sn_cdm/applications/deployables/exports/{export_id}/status"
	exportItemContentURLTemplate                     = "{+baseurl}/api/sn_cdm/applications/deployables/exports/{export_id}/content"
	sharedLibrariesComponentsApplicationsURLTemplate = "{+baseurl}/api/sn_cdm/applications/shared_libraries/components/applications{?appName,sharedComponentName,name}"
	uploadsComponentsURLTemplate                     = "{+baseurl}/api/sn_cdm/applications/uploads/components"
	uploadsComponentsVarsURLTemplate                 = "{+baseurl}/api/sn_cdm/applications/uploads/components/vars"
	uploadsCollectionsURLTemplate                    = "{+baseurl}/api/sn_cdm/applications/uploads/collections"
	uploadsCollectionsFileURLTemplate                = "{+baseurl}/api/sn_cdm/applications/uploads/collections/file{?appName,collectionName}"
	uploadsDeployablesFileURLTemplate                = "{+baseurl}/api/sn_cdm/applications/uploads/deployables/file{?appName,deployableName}"
)

// ApplicationsRequestBuilder provides operations to manage applications.
type ApplicationsRequestBuilder struct {
	core.RequestBuilder
}

// NewApplicationsRequestBuilderInternal instantiates a new ApplicationsRequestBuilder.
func NewApplicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ApplicationsRequestBuilder {
	return &ApplicationsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, applicationsURLTemplate, pathParameters),
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

// SharedLibraries returns a SharedLibrariesRequestBuilder.
func (rB *ApplicationsRequestBuilder) SharedLibraries() *SharedLibrariesRequestBuilder {
	return NewSharedLibrariesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Uploads returns a UploadsRequestBuilder.
func (rB *ApplicationsRequestBuilder) Uploads() *UploadsRequestBuilder {
	return NewUploadsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// DeployablesRequestBuilder provides operations to manage deployables.
type DeployablesRequestBuilder struct {
	core.RequestBuilder
}

// NewDeployablesRequestBuilderInternal instantiates a new DeployablesRequestBuilder.
func NewDeployablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *DeployablesRequestBuilder {
	return &DeployablesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, deployablesURLTemplate, pathParameters),
	}
}

// Delete deletes a deployable.
func (rB *DeployablesRequestBuilder) Delete(ctx context.Context, config *DeployablesRequestBuilderDeleteRequestConfiguration) error {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
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
	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, nil)
}

// Put updates deployables.
func (rB *DeployablesRequestBuilder) Put(ctx context.Context, body *DeployableUpdateRequest, config *DeployablesRequestBuilderPutRequestConfiguration) (DeployableUpdateResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateDeployableUpdateResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(DeployableUpdateResponse), nil
}

// Exports returns an ExportsRequestBuilder.
func (rB *DeployablesRequestBuilder) Exports() *ExportsRequestBuilder {
	return NewExportsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// SharedComponentsRequestBuilder provides operations to manage shared components.
type SharedComponentsRequestBuilder struct {
	core.RequestBuilder
}

// NewSharedComponentsRequestBuilderInternal instantiates a new SharedComponentsRequestBuilder.
func NewSharedComponentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SharedComponentsRequestBuilder {
	return &SharedComponentsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, sharedComponentsURLTemplate, pathParameters),
	}
}

// Delete deletes shared component references.
func (rB *SharedComponentsRequestBuilder) Delete(ctx context.Context, config *SharedComponentsRequestBuilderDeleteRequestConfiguration) error {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
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
	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, nil)
}

// Put updates shared components.
func (rB *SharedComponentsRequestBuilder) Put(ctx context.Context, body *SharedComponentUpdateRequest, config *SharedComponentsRequestBuilderPutRequestConfiguration) (SharedComponentUpdateResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateSharedComponentUpdateResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(SharedComponentUpdateResponse), nil
}

// UploadStatusRequestBuilder provides operations to access upload status.
type UploadStatusRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadStatusRequestBuilderInternal instantiates a new UploadStatusRequestBuilder.
func NewUploadStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadStatusRequestBuilder {
	return &UploadStatusRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadStatusURLTemplate, pathParameters),
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
	core.RequestBuilder
}

// NewUploadStatusItemRequestBuilderInternal instantiates a new UploadStatusItemRequestBuilder.
func NewUploadStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadStatusItemRequestBuilder {
	return &UploadStatusItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadStatusURLTemplate, pathParameters),
	}
}

// Get gets the status of a specific upload.
func (rB *UploadStatusItemRequestBuilder) Get(ctx context.Context, config *UploadStatusItemRequestBuilderGetRequestConfiguration) (UploadStatusResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateUploadStatusResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(UploadStatusResponse), nil
}

// ExportsRequestBuilder provides operations to manage deployable exports.
type ExportsRequestBuilder struct {
	core.RequestBuilder
}

// NewExportsRequestBuilderInternal instantiates a new ExportsRequestBuilder.
func NewExportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExportsRequestBuilder {
	return &ExportsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, exportsURLTemplate, pathParameters),
	}
}

// Get gets the collection of deployable exports.
func (rB *ExportsRequestBuilder) Get(ctx context.Context, config *ExportsRequestBuilderGetRequestConfiguration) (ExportsResponse, error) {
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
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateExportsResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(ExportsResponse), nil
}

// ByID returns a ExportItemRequestBuilder.
func (rB *ExportsRequestBuilder) ByID(exportId string) *ExportItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["export_id"] = exportId
	return NewExportItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// ExportItemRequestBuilder provides operations to manage a specific deployable export.
type ExportItemRequestBuilder struct {
	core.RequestBuilder
}

// NewExportItemRequestBuilderInternal instantiates a new ExportItemRequestBuilder.
func NewExportItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExportItemRequestBuilder {
	return &ExportItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, exportItemURLTemplate, pathParameters),
	}
}

// Content returns a ExportItemContentRequestBuilder.
func (rB *ExportItemRequestBuilder) Content() *ExportItemContentRequestBuilder {
	return NewExportItemContentRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Status returns a ExportItemStatusRequestBuilder.
func (rB *ExportItemRequestBuilder) Status() *ExportItemStatusRequestBuilder {
	return NewExportItemStatusRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ExportItemStatusRequestBuilder provides operations to check export status.
type ExportItemStatusRequestBuilder struct {
	core.RequestBuilder
}

// NewExportItemStatusRequestBuilderInternal instantiates a new ExportItemStatusRequestBuilder.
func NewExportItemStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExportItemStatusRequestBuilder {
	return &ExportItemStatusRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, exportItemStatusURLTemplate, pathParameters),
	}
}

// Get gets the status of a specific export.
func (rB *ExportItemStatusRequestBuilder) Get(ctx context.Context, config *ExportItemStatusRequestBuilderGetRequestConfiguration) (ExportStatusResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateExportStatusResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(ExportStatusResponse), nil
}

// ExportItemContentRequestBuilder provides operations to download export content.
type ExportItemContentRequestBuilder struct {
	core.RequestBuilder
}

// NewExportItemContentRequestBuilderInternal instantiates a new ExportItemContentRequestBuilder.
func NewExportItemContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExportItemContentRequestBuilder {
	return &ExportItemContentRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, exportItemContentURLTemplate, pathParameters),
	}
}

// Get fetches the export content.
func (rB *ExportItemContentRequestBuilder) Get(ctx context.Context, config *ExportItemContentRequestBuilderGetRequestConfiguration) ([]byte, error) {
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
	res, err := rB.GetRequestAdapter().SendPrimitive(ctx, requestInfo, "[]byte", nil)
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, nil
	}
	return res.([]byte), nil
}

// SharedLibrariesRequestBuilder provides operations to manage shared libraries.
type SharedLibrariesRequestBuilder struct {
	core.RequestBuilder
}

// NewSharedLibrariesRequestBuilderInternal instantiates a new SharedLibrariesRequestBuilder.
func NewSharedLibrariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SharedLibrariesRequestBuilder {
	return &SharedLibrariesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_cdm/applications/shared_libraries", pathParameters),
	}
}

// Components returns a SharedLibrariesComponentsRequestBuilder.
func (rB *SharedLibrariesRequestBuilder) Components() *SharedLibrariesComponentsRequestBuilder {
	return NewSharedLibrariesComponentsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// SharedLibrariesComponentsRequestBuilder provides operations to manage shared library components.
type SharedLibrariesComponentsRequestBuilder struct {
	core.RequestBuilder
}

// NewSharedLibrariesComponentsRequestBuilderInternal instantiates a new SharedLibrariesComponentsRequestBuilder.
func NewSharedLibrariesComponentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SharedLibrariesComponentsRequestBuilder {
	return &SharedLibrariesComponentsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_cdm/applications/shared_libraries/components", pathParameters),
	}
}

// Applications returns a SharedLibrariesComponentsApplicationsRequestBuilder.
func (rB *SharedLibrariesComponentsRequestBuilder) Applications() *SharedLibrariesComponentsApplicationsRequestBuilder {
	return NewSharedLibrariesComponentsApplicationsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// SharedLibrariesComponentsApplicationsRequestBuilder provides operations to access shared library component applications.
type SharedLibrariesComponentsApplicationsRequestBuilder struct {
	core.RequestBuilder
}

// NewSharedLibrariesComponentsApplicationsRequestBuilderInternal instantiates a new SharedLibrariesComponentsApplicationsRequestBuilder.
func NewSharedLibrariesComponentsApplicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SharedLibrariesComponentsApplicationsRequestBuilder {
	return &SharedLibrariesComponentsApplicationsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, sharedLibrariesComponentsApplicationsURLTemplate, pathParameters),
	}
}

// Get gets the collection of shared library component applications.
func (rB *SharedLibrariesComponentsApplicationsRequestBuilder) Get(ctx context.Context, config *SharedLibrariesComponentsApplicationsRequestBuilderGetRequestConfiguration) (SharedLibrariesComponentsApplicationsResponse, error) {
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
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateSharedLibrariesComponentsApplicationsResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(SharedLibrariesComponentsApplicationsResponse), nil
}

// UploadsRequestBuilder provides operations to manage uploads.
type UploadsRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadsRequestBuilderInternal instantiates a new UploadsRequestBuilder.
func NewUploadsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadsRequestBuilder {
	return &UploadsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_cdm/applications/uploads", pathParameters),
	}
}

// Components returns a UploadsComponentsRequestBuilder.
func (rB *UploadsRequestBuilder) Components() *UploadsComponentsRequestBuilder {
	return NewUploadsComponentsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Collections returns a UploadsCollectionsRequestBuilder.
func (rB *UploadsRequestBuilder) Collections() *UploadsCollectionsRequestBuilder {
	return NewUploadsCollectionsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Deployables returns a UploadsDeployablesRequestBuilder.
func (rB *UploadsRequestBuilder) Deployables() *UploadsDeployablesRequestBuilder {
	return NewUploadsDeployablesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// UploadsComponentsRequestBuilder provides operations to upload components.
type UploadsComponentsRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadsComponentsRequestBuilderInternal instantiates a new UploadsComponentsRequestBuilder.
func NewUploadsComponentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadsComponentsRequestBuilder {
	return &UploadsComponentsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadsComponentsURLTemplate, pathParameters),
	}
}

// Post uploads components.
func (rB *UploadsComponentsRequestBuilder) Post(ctx context.Context, body *ComponentUploadRequest, config *UploadsComponentsRequestBuilderPostRequestConfiguration) (UploadStatusResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateUploadStatusResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(UploadStatusResponse), nil
}

// Vars returns a UploadsComponentsVarsRequestBuilder.
func (rB *UploadsComponentsRequestBuilder) Vars() *UploadsComponentsVarsRequestBuilder {
	return NewUploadsComponentsVarsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// UploadsComponentsVarsRequestBuilder provides operations to upload component variables.
type UploadsComponentsVarsRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadsComponentsVarsRequestBuilderInternal instantiates a new UploadsComponentsVarsRequestBuilder.
func NewUploadsComponentsVarsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadsComponentsVarsRequestBuilder {
	return &UploadsComponentsVarsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadsComponentsVarsURLTemplate, pathParameters),
	}
}

// Post uploads component variables.
func (rB *UploadsComponentsVarsRequestBuilder) Post(ctx context.Context, body *ComponentVarsUploadRequest, config *UploadsComponentsVarsRequestBuilderPostRequestConfiguration) (UploadStatusResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateUploadStatusResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(UploadStatusResponse), nil
}

// UploadsCollectionsRequestBuilder provides operations to upload collections.
type UploadsCollectionsRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadsCollectionsRequestBuilderInternal instantiates a new UploadsCollectionsRequestBuilder.
func NewUploadsCollectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadsCollectionsRequestBuilder {
	return &UploadsCollectionsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadsCollectionsURLTemplate, pathParameters),
	}
}

// Post uploads collections.
func (rB *UploadsCollectionsRequestBuilder) Post(ctx context.Context, body *CollectionUploadRequest, config *UploadsCollectionsRequestBuilderPostRequestConfiguration) (UploadStatusResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateUploadStatusResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(UploadStatusResponse), nil
}

// File returns a UploadsCollectionsFileRequestBuilder.
func (rB *UploadsCollectionsRequestBuilder) File() *UploadsCollectionsFileRequestBuilder {
	return NewUploadsCollectionsFileRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// UploadsCollectionsFileRequestBuilder provides operations to upload collection files.
type UploadsCollectionsFileRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadsCollectionsFileRequestBuilderInternal instantiates a new UploadsCollectionsFileRequestBuilder.
func NewUploadsCollectionsFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadsCollectionsFileRequestBuilder {
	return &UploadsCollectionsFileRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadsCollectionsFileURLTemplate, pathParameters),
	}
}

// Post uploads collection files using a stream payload (like attachment-api)
func (rB *UploadsCollectionsFileRequestBuilder) Post(ctx context.Context, media *Media, config *UploadsCollectionsFileRequestBuilderPostRequestConfiguration) (UploadStatusResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
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
	requestInfo.SetStreamContentAndContentType(media.GetData(), media.GetContentType())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateUploadStatusResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(UploadStatusResponse), nil
}

// UploadsDeployablesRequestBuilder provides operations to manage deployables uploads.
type UploadsDeployablesRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadsDeployablesRequestBuilderInternal instantiates a new UploadsDeployablesRequestBuilder.
func NewUploadsDeployablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadsDeployablesRequestBuilder {
	return &UploadsDeployablesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_cdm/applications/uploads/deployables", pathParameters),
	}
}

// File returns a UploadsDeployablesFileRequestBuilder.
func (rB *UploadsDeployablesRequestBuilder) File() *UploadsDeployablesFileRequestBuilder {
	return NewUploadsDeployablesFileRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// UploadsDeployablesFileRequestBuilder provides operations to upload deployable files.
type UploadsDeployablesFileRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadsDeployablesFileRequestBuilderInternal instantiates a new UploadsDeployablesFileRequestBuilder.
func NewUploadsDeployablesFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadsDeployablesFileRequestBuilder {
	return &UploadsDeployablesFileRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadsDeployablesFileURLTemplate, pathParameters),
	}
}

// Post uploads deployable files.
func (rB *UploadsDeployablesFileRequestBuilder) Post(ctx context.Context, media *Media, config *UploadsDeployablesFileRequestBuilderPostRequestConfiguration) (UploadStatusResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
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
	requestInfo.SetStreamContentAndContentType(media.GetData(), media.GetContentType())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateUploadStatusResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(UploadStatusResponse), nil
}
