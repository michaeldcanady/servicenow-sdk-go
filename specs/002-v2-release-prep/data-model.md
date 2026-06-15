# Data Model Changes: v2 Release Prep

## Public API Surface Pruning

The following entities and types will be removed or replaced as part of the v2.0 cleanup:

### attachment-api
- Removed `Attachment`, `AttachmentCollectionResponse`, `AttachmentItemResponse`.
- Removed `AttachmentRequestBuilder` (deprecated methods only).
- Unified on `Attachment`, `AttachmentCollectionResponse2`, etc.

### core
- Removed `RequestBuilder` (older versions like `NewRequestBuilder`).
- Removed `SendGet`, `SendPost`, `SendPut`, `SendDelete` (replaced by `SendGet2`, `SendPost2`, etc.).
- Removed `PageIterator` (replaced by `PageIterator2[T]`).
- Removed `URLInformation` (older versions).

### credentials
- Removed `Credential` (deprecated since v1.11.0).
- Removed `UsernamePasswordCredential` (replaced by `BasicAuthenticationProvider`).
- Removed `TokenCredential` (replaced by `ROPCCredential`).

### table-api
- Removed `DisplayValue`, `Entry`, `Fragment`, `LogicalOperator`, `OrderBy`, `OrderDirection`, `Query`, `RelationalOperator` (older versions).
- Removed `TableCollectionResponse` (replaced by `TableCollectionResponse2[T]`).
- Removed `TableEntry` (replaced by `TableRecord`).
- Removed `TableItemRequestBuilder` (deprecated versions).

## Query System Consolidation
- **Winner**: `query2`
- **Loser**: `query` (to be deleted)
- The public `query2` package becomes the recommended way to build encoded queries.
