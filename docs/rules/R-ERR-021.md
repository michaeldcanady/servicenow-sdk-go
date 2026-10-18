---
id: R-ERR-021
applies_to: ["core/**/*.go", "**/*error*.go"]
status: proposed
supersedes: []
conflicts_checked: true
---

<!-- Keywords: error categorization, best effort, ConflictError, ValidationError,
409, raw details, unwrap, DefaultErrorMapping, MainError, generic fallback -->

MUST categorize ServiceNow errors into SDK error types on a best-effort basis via `core.DefaultErrorMapping()`.
MUST attach the raw ServiceNow `detail`, `message`, and `status` to every categorized error so callers can unwrap to the original.
MUST add a `ConflictError` mapped from HTTP 409 and a `ValidationError` for validation failures to the default mapping.
MUST degrade an un-categorizable error to a generic SDK error that still carries the raw payload.
MUST NOT drop the raw ServiceNow details when categorizing — the raw payload is the lossless contract.
MUST NOT treat a category as exhaustive or guaranteed; callers must be able to unwrap for exact handling.
SHOULD keep `UnauthorizedError.Is` matching `snerrors.ErrNotAuthenticated`.
MAY map additional status codes incrementally as they are encountered.

## Verify

```bash
# Passes (zero) once ConflictError (409) is wired into the default mapping.
grep -RInE 'ConflictError' --include='*.go' core/
```

## Rationale

A fixed set of SDK error types can't exhaustively categorize ServiceNow's open-ended error surface, so categorization is best-effort with the raw detail/message/status always attached for unwrap-based exact handling; 409 and validation get first-class types (see ADRs 001, 006, 021).
