# Core concepts

## `RequestBuilder` façade

The `RequestBuilder` serves as a high-level interface for constructing and executing HTTP operations for a specific API path. At each path diversion (after the base URI) introduce a new `RequestBuilder`.

**Example:**  
Absolute URI: `https://{instance}.service-now.com/api/now/{version}/table`

To generate the appropriate `RequestBuilder` types:

1. **Remove the shared base URI**: Strip the common prefix `https://{instance}.service-now.com/api` to isolate the relative path `now/{version}/table`.
2. **Segment the relative path**: Split the path by `/`, omitting the version segment. This yields two distinct `RequestBuilder` types:
   - `NowRequestBuilder`
   - `TableRequestBuilder`

> The base URI is `https://{instance}.service-now.com/api`  
> Note: The SDK targets a single API version, so the version segment can be excluded from `RequestBuilder` generation.

## `RequestBuilder` structure

Each `RequestBuilder` provides methods aligned with supported HTTP operations. These methods accept standardized parameters or enable construction of the next path segment. Internally, they handle:

1. **Request construction**: Dynamically builds the HTTP request, including headers, query parameters, and payload.
2. **Request transmission**: Sends the request through an HTTP client, supporting features such as retries and logging.
3. **Response handling**: Converts raw HTTP responses into structured SDK types or error classes.

## Batch API compatibility

All `RequestBuilder` types support batch operations using a shared convention:  
`ToXXXRequestInformation(…) (RequestInformation, error)` methods.

These methods generate lightweight `RequestInformation` objects that encapsulate:

- The HTTP method  
- Target address  
- Headers and query parameters  
- Request body (if applicable)

This approach enables batch execution without immediate transmission, allowing deferred or grouped interactions.
