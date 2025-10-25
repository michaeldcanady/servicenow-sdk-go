# Core Concepts

## RequestBuilder Facade

At the heart of the SDK is the `RequestBuilder`, which acts as a high-level interface for constructing and executing HTTP operations for a specific api path. A `RequestBuilder` is added at each path diversion. (I.E.: <code>https://<var>{host}</var>/api</code>, is shared between **all** apis so it’s superfluous to include).
Example:
The path `…/api/now/{version}/table`  becomes two `requestBuilders`.

- `NowRequestBuilder`
- `TableRequestBuilder`

> Note: Version doesn’t get a `RequestBuilder` as the **whole** SDK represents a single version of the api.

## RequestBuilder Structure

Every RequestBuilder exposes methods aligned with supported HTTP operations, accepting standardized parameters or a method to build the next path segment. 
These methods wrap internal logic for:

1. Building the request: Dynamically constructs the HTTP request, including headers, query parameters, and payload.
2. Sending the request: Routes the request via a pluggable HTTP client, enabling retries, logging, etc.
3. Handling the response: Maps raw HTTP responses into structured SDK types or error classes.

## Batch API Compatibility

All `RequestBuilder` types support batch operations via a shared convention:
`ToXXXRequestInformation(…) (RequestInformation, error)` methods.
These methods generate lightweight RequestInformation object which encapsulate:

* The HTTP method
* Target URL
* Headers and query parameters
* Request body (if applicable)

This makes batch execution possible without sending the request immediately, enabling deferred or grouped interactions.