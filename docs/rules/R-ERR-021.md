---
id: R-ERR-021
applies_to: ["core/**/*.go", "errors/**/*.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

# ServiceNow error categorization is best-effort and lossless

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->
<!-- TODO(owner): list the HTTP statuses and response bodies that count as a "validation failure", and cite where ServiceNow returns 409. -->

MUST: Categorize ServiceNow errors into SDK error types through `core.DefaultErrorMapping()`.
MUST: Attach the raw ServiceNow `message` and `detail`, the response-body `status`, and the HTTP status code to every categorized error, so callers can unwrap to the original.
MUST: Map HTTP 409 to a `ConflictError` in the default mapping.
MUST: Map validation failures to a `ValidationError` in the default mapping.
MUST: Degrade an un-categorizable error to a generic SDK error that still carries the raw payload.
MUST NOT: Drop raw ServiceNow details when categorizing.
MUST NOT: Document or treat a category as exhaustive or guaranteed.
SHOULD: Keep `UnauthorizedError` matching `snerrors.ErrNotAuthenticated` under `errors.Is`.
MAY: Map additional status codes as they are encountered.

## Verify

```bash
# Passes only once 409 is mapped to a ConflictError factory. Expected to fail while this rule is proposed.
grep -RqE '"409"\s*:\s*\w*ConflictError' --include='*.go' core/
```

## Rationale

A fixed set of SDK error types cannot exhaustively categorize ServiceNow's open-ended error surface, so categorization is best-effort with the raw payload always attached for exact handling by unwrapping. 409 and validation failures get first-class types (see R-ERR-001).
