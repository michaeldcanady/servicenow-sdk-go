---
id: R-BUILD-003
applies_to: ["**/*.go", "go.mod"]
status: accepted
supersedes: [R-BUILD-008, R-NAME-009]
conflicts_checked: true
---

# The SDK is hand-written on the Kiota runtime

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST: Build on the Kiota runtime libraries (`kiota-abstractions-go`, `kiota-http-go`, and the `kiota-serialization-{json,form,multipart,text}-go` packages).
MUST: Hand-write request builders, models, and response envelopes to match Kiota-generated conventions (builder chaining, `RequestConfiguration` shapes, parsable factories, backed models).
MUST: Use Kiota's `RequestAdapter`, `BaseRequestBuilder`, `Parsable`, and `ParseNode` contracts in hand-written builders.
MUST NOT: Generate the client from OpenAPI with the Kiota CLI.
MUST NOT: Reimplement plumbing the Kiota runtime provides (HTTP transport, retry, auth, URI templates, serialization, error-mapping mechanism).
MUST NOT: Declare a type that aliases or renames Kiota's `RequestBuilder`, `BaseRequestBuilder`, or `RequestInformation`.
SHOULD: Treat msgraph-sdk-go as the tie-breaker when a structural question has no local precedent.
MAY: Import `net/http` in `internal/` transport wiring and in tests.

## Verify

```bash
# grep exit 1 = no match = pass; exit 2 fails closed.
{ grep -RInE '"net/http"' --include='*.go' --exclude='*_test.go' --exclude-dir=internal --exclude-dir=.git .; [ $? -eq 1 ]; } \
  && { grep -RInE 'type\s+(RequestBuilder|BaseRequestBuilder|RequestInformation)\s+(=|struct|interface)' --include='*.go' .; [ $? -eq 1 ]; }
```

## Rationale

ServiceNow publishes no reliable OpenAPI, and its response shapes map poorly to generator output, so the surface is hand-written on Kiota's tested runtime and mirrors msgraph-sdk-go conventions. We accept a dependency on Microsoft packages and the Kiota object model because the alternative is maintaining an equivalent framework. Pagination is SDK-owned (see R-PAGE-005), not a Kiota feature.
