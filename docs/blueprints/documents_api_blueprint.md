# Technical Blueprint: Documents API Support (`documentsapi`)

This blueprint outlines the architectural plan for implementing the ServiceNow Documents API in the Go SDK, derived from `spec/now_documents_api_latest_spec.json`.

## 1. Package Overview
- **Package Name**: `documentsapi`
- **Base Path**: `/api/now/documents`
- **Internal Alias**: `newInternal` (`github.com/michaeldcanady/servicenow-sdk-go/internal`)

## 2. RequestBuilder Hierarchy

| Path | RequestBuilder | Method | Description |
| :--- | :--- | :--- | :--- |
| `/api/now/documents` | `DocumentsRequestBuilder` | `Post` | Base documents operations |
| `/api/now/documents/create` | `CreateRequestBuilder` | `Post` | Create document |
| `/api/now/documents/createDocument` | `CreateDocumentRequestBuilder` | `Post` | Create or link document |
| `/api/now/documents/delete` | `DeleteRequestBuilder` | `Delete` | Delete document (Query based) |
| `/api/now/documents/explore` | `ExploreRequestBuilder` | `Get` | Search/Explore docs and folders |
| `/api/now/documents/versions/{document_sys_id}` | `VersionsRequestBuilder` | `ByID(docID).Get()` | List document versions |
| `/api/now/documents/{document_sys_id}/content` | `ContentRequestBuilder` | `ByID(docID).Content().Get()` | Fetch/Stream content |
| `/api/now/documents/action/{action}/document/{documentSysId}/version/{versionSysId}` | `ActionRequestBuilder` | `Action(a).Document(d).Version(v).Patch()` | Complex action execution |

## 3. Operations & Models

### 3.1. Explore Operation
- **Path**: `/api/now/documents/explore`
- **Method**: `Get`
- **Query Parameters**:
    - `page` (int)
    - `limit` (int)
    - `query` (string)
    - `table_name` (string)
    - `folder_sys_id` (string)
    - `record_sys_id` (string)
- **Response**: `ServiceNowCollectionResponse[Document]` (Generic `internal/new` model)

### 3.2. Content Streaming
- **Path**: `/api/now/documents/{document_sys_id}/content`
- **Method**: `Get`
- **Response**: Binary Stream (`[]byte`)

### 3.3. Complex Path: Action
- **Path**: `/api/now/documents/action/{action}/document/{docId}/version/{vId}`
- **Pattern**: `Documents.Action(action).Document(docId).Version(versionId).Patch(ctx, config)`

## 4. Model Definitions (Proposed)
*Note: Spec schemas are empty; using standard ServiceNow model patterns.*

### 4.1. `Document` Model
- `SysID` (string)
- `Name` (string)
- `Type` (string)
- `ParentFolder` (string)

## 5. Implementation Strategy
1. **Scaffold Directory**: Create `documents-api/`.
2. **Base Builder**: Implement `DocumentsRequestBuilder2` using `internal/new.BaseRequestBuilder`.
3. **Sub-Builders**: Implement nested builders for `/explore`, `/versions`, `/content`.
4. **Configurations**: Generate `RequestConfiguration` and `QueryParameters` for all operations.
5. **Validation**: Add unit tests in `documents_request_builder_test.go`.
