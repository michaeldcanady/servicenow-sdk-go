---
id: R-BUILD-004
applies_to: ["*api/**/*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

# Query parameters are pointer fields wired through Kiota

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST: Declare every exported `*QueryParameters` field as a pointer, so "unset" differs from zero.
MUST: Declare integer query parameters as `*int32`.
MUST: Wire query parameters through `abstractions.ConfigureRequestInformation` directly.
MUST NOT: Declare a value-typed (non-pointer) query-parameter field.
MUST NOT: Depend on `go-querystring`.
MUST NOT: Wrap `ConfigureRequestInformation` in a bespoke helper.

## Verify

```bash
# Flags `uriparametername` fields that are not pointers/slices, or are an integer pointer other than *int32.
# grep exit 1 = no match = pass; exit 2 fails closed.
{ grep -RInP '^\s*\w+\s+(?:(?![*\[])\S+|\*(?:u?int(?:8|16|64)?|uint32))\s+`uriparametername' --include='*.go' .; [ $? -eq 1 ]; } \
  && { grep -E 'go-querystring' go.mod go.sum; [ $? -eq 1 ]; }
```

## Rationale

Pointer fields wired through `ConfigureRequestInformation` match msgraph-sdk-go and give query parameters the same unset-versus-zero distinction as models (see R-MODEL-002). Kiota supports `*int32` for integer query parameters and not `*int`, and an earlier SDK bug dropped parameters declared as `*int`.
