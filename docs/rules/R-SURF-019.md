---
id: R-SURF-019
applies_to: ["**/*_request_builder.go", "**/servicenow_service_client.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

# Request builders mirror ServiceNow's URL hierarchy

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->
<!-- Examples: /api/now/attachment -> client.Now().Attachment(); /api/now/stats/{tableName} -> client.Now().Stats("incident"); /api/now/table/{tableName} -> client.Now().Table2().ByName("incident") (R-SURF-018). -->
<!-- Assumes the client type is `ServiceNowServiceClient`; adjust the Verify if renamed. -->

MUST: Mirror the ServiceNow REST URL hierarchy in the request-builder hierarchy of every `*api` module, with each path segment becoming a method on the parent builder that returns a child builder.
MUST: Represent each path parameter as a string argument to the method that introduces its segment.
MUST: Represent query parameters as typed fields on request-configuration structs (R-BUILD-004).
MUST: Define a leaf builder's complete URL template in that leaf builder's file and nowhere else.
MUST NOT: Put URL template strings in a parent builder or at a call site.
MUST NOT: Expose a flat `Get(path)` or `Request(method, path, params)` method on the client for any ServiceNow endpoint.
SHOULD: Deviate from literal segment names only under R-SURF-018 (collision) or R-SURF-017 (domain affordance), and document the URL mapping.

## Verify

```bash
# URL templates may appear only in *_request_builder.go; the client exposes no flat verbs. Fails closed if no templates exist.
grep -RqE 'URLTemplate|UrlTemplate' --include='*_request_builder.go' . \
  && { grep -RIlE 'URLTemplate|UrlTemplate' --include='*.go' . | grep -vE '(_request_builder|_test)\.go$'; [ $? -eq 1 ]; } \
  && { grep -RInE 'func \([a-z]+ \*?ServiceNowServiceClient\) (Get|Post|Put|Patch|Delete|Request)\(' --include='*.go' .; [ $? -eq 1 ]; }
```

## Rationale

Mirroring the URL makes the SDK self-documenting, lets the compiler catch path typos, and keeps URL knowledge in one place per endpoint, matching Kiota-generated convention (R-BUILD-003). Domain affordances (R-SURF-017) sit above this raw layer.
