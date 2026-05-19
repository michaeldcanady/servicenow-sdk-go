package documentsapi

import (
	"maps"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// documentsURLTemplate is the url template for the Documents API endpoint
	documentsURLTemplate = "{+baseurl}/api/now/v1/documents"
)

// DocumentsRequestBuilder provides operations to manage ServiceNow documents.
type DocumentsRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewDocumentsRequestBuilderInternal instantiates a new DocumentsRequestBuilder with the provided request parameters.
func NewDocumentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *DocumentsRequestBuilder {
	return &DocumentsRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, documentsURLTemplate, pathParameters),
	}
}

// NewDocumentsRequestBuilder instantiates a new DocumentsRequestBuilder with the provided raw URL.
func NewDocumentsRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *DocumentsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewDocumentsRequestBuilderInternal(urlParams, requestAdapter)
}

// Explore provides operations to manage the explore endpoint.
func (rB *DocumentsRequestBuilder) Explore() *ExploreRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	return NewExploreRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Create provides operations to manage the create endpoint.
func (rB *DocumentsRequestBuilder) Create() *CreateRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	return NewCreateRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// CreateDocument provides operations to manage the createDocument endpoint.
func (rB *DocumentsRequestBuilder) CreateDocument() *CreateDocumentRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	return NewCreateDocumentRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// VersionState provides operations to manage document version states.
func (rB *DocumentsRequestBuilder) VersionState(versionSysID string) *VersionStateRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["version_sys_id"] = versionSysID
	return NewVersionStateRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Attach provides operations to manage document attachments via a provider.
func (rB *DocumentsRequestBuilder) Attach(providerID string) *AttachRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["provider_id"] = providerID
	return NewAttachRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Delete provides operations to manage the delete endpoint.
func (rB *DocumentsRequestBuilder) Delete() *DeleteRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	return NewDeleteRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Versions provides operations to manage document versions.
func (rB *DocumentsRequestBuilder) Versions(documentSysID string) *VersionsRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["document_sys_id"] = documentSysID
	return NewVersionsRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Content provides operations to manage document content.
func (rB *DocumentsRequestBuilder) Content(documentSysID string) *ContentRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["document_sys_id"] = documentSysID
	return NewContentRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// SyncDown provides operations to manage document synchronization.
func (rB *DocumentsRequestBuilder) SyncDown(documentSysID string) *SyncDownRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["documentSysId"] = documentSysID
	return NewSyncDownRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Action provides operations to manage document actions.
func (rB *DocumentsRequestBuilder) Action(action string) *ActionRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["action"] = action
	return NewActionRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
