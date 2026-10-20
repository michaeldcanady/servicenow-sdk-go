---
id: R-MODEL-002
applies_to: ["model/**/*.go", "*api/**/*model*.go"]
status: accepted
supersedes: []
conflicts_checked: true
---

# Models keep property data in a Kiota backing store

<!-- Keywords MUST, MUST NOT, SHOULD, SHOULD NOT, MAY are interpreted per RFC 2119 / RFC 8174 (https://www.rfc-editor.org/rfc/rfc2119). -->

MUST: Embed `core.BaseModel` in every model.
MUST: Keep property data in a Kiota `BackingStore`.
MUST: Expose each property as a `GetX() (*T, error)` / `SetX(*T) error` accessor pair.
MUST: Return a pointer from each getter, so `nil` ("not sent") stays distinct from a zero value.
MUST: Record each `SetX` call in the backing store so the property enters the dirty set.
MUST: Serialize only the backing store's dirty set, so request bodies carry only properties the caller set.
MUST: Constrain generic model code with `model.ServiceNowItem`.
MUST NOT: Declare exported model data as plain struct fields.
SHOULD: Build accessors with the store/serialization helper idiom instead of hand-rolling each one.

## Verify

```bash
# A backed model needs no encoding/json tags. grep exit 1 = no match = pass; exit 2 fails closed.
grep -RInE 'json:"' --include='*.go' model/; [ $? -eq 1 ]
```

## Rationale

A plain struct field cannot distinguish "the instance sent an empty value" from "the instance sent nothing", which caused absent/zero confusion and accidental whole-object overwrites. A backing store with change tracking fixes both.
