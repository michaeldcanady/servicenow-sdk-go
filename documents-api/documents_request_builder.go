package documentsapi

import (
	"maps"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// documentsURLTemplate is the url template for the Documents API endpoint
	documentsURLTemplate = "{+baseurl}/api/now/documents"
)

// DocumentsRequestBuilder2 provides operations to manage ServiceNow documents.
type DocumentsRequestBuilder2 struct {
	newInternal.RequestBuilder
}

// NewDocumentsRequestBuilder2Internal instantiates a new DocumentsRequestBuilder2 with the provided request parameters.
func NewDocumentsRequestBuilder2Internal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *DocumentsRequestBuilder2 {
	return &DocumentsRequestBuilder2{
		newInternal.NewBaseRequestBuilder(requestAdapter, documentsURLTemplate, pathParameters),
	}
}

// NewDocumentsRequestBuilder2 instantiates a new DocumentsRequestBuilder2 with the provided raw URL.
func NewDocumentsRequestBuilder2(rawURL string, requestAdapter abstractions.RequestAdapter) *DocumentsRequestBuilder2 {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewDocumentsRequestBuilder2Internal(urlParams, requestAdapter)
}

// Explore provides operations to manage the explore endpoint.
func (rB *DocumentsRequestBuilder2) Explore() *ExploreRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	return NewExploreRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Create provides operations to manage the create endpoint.
func (rB *DocumentsRequestBuilder2) Create() *CreateRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	return NewCreateRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// CreateDocument provides operations to manage the createDocument endpoint.
func (rB *DocumentsRequestBuilder2) CreateDocument() *CreateDocumentRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	return NewCreateDocumentRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// VersionState provides operations to manage document version states.
func (rB *DocumentsRequestBuilder2) VersionState(versionSysID string) *VersionStateRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["version_sys_id"] = versionSysID
	return NewVersionStateRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Attach provides operations to manage document attachments via a provider.
func (rB *DocumentsRequestBuilder2) Attach(providerID string) *AttachRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["provider_id"] = providerID
	return NewAttachRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Delete provides operations to manage the delete endpoint.
func (rB *DocumentsRequestBuilder2) Delete() *DeleteRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	return NewDeleteRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Versions provides operations to manage document versions.
func (rB *DocumentsRequestBuilder2) Versions(documentSysID string) *VersionsRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["document_sys_id"] = documentSysID
	return NewVersionsRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Content provides operations to manage document content.
func (rB *DocumentsRequestBuilder2) Content(documentSysID string) *ContentRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["document_sys_id"] = documentSysID
	return NewContentRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// SyncDown provides operations to manage document synchronization.
func (rB *DocumentsRequestBuilder2) SyncDown(documentSysID string) *SyncDownRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["documentSysId"] = documentSysID
	return NewSyncDownRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Action provides operations to manage document actions.
func (rB *DocumentsRequestBuilder2) Action(action string) *ActionRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["action"] = action
	return NewActionRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
