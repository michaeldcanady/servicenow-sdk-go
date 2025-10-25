# Core Concepts

## `RequestBuilder` Facade

The `RequestBuilder` serves as a high-level interface for constructing and executing HTTP operations for a specific API path. A new `RequestBuilder` is introduced at each path diversion. For example, the path `https://{host}/api` is shared across **all** APIs and does not require inclusion.

Example:  
The path `…/api/now/{version}/table` results in two distinct `RequestBuilder` instances:

- `NowRequestBuilder`
- `TableRequestBuilder`

> Note: The version segment does not receive a `RequestBuilder` because the SDK represents a single version of the API.

## `RequestBuilder` Structure

Each `RequestBuilder` provides methods aligned with supported HTTP operations. These methods accept standardized parameters or enable construction of the next path segment. Internally, they handle:

1. **Request construction**: Dynamically builds the HTTP request, including headers, query parameters, and payload.
2. **Request transmission**: Sends the request through an HTTP client, supporting features such as retries and logging.
3. **Response handling**: Converts raw HTTP responses into structured SDK types or error classes.

## Batch API Compatibility

All `RequestBuilder` types support batch operations using a shared convention:  
`ToXXXRequestInformation(…) (RequestInformation, error)` methods.

These methods generate lightweight `RequestInformation` objects that encapsulate:

- The HTTP method  
- Target URL  
- Headers and query parameters  
- Request body (if applicable)

This approach enables batch execution without immediate transmission, allowing deferred or grouped interactions.
