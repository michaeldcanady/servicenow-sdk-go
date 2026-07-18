---
title: Testing guide
description: >-
  The three test rings — unit, integration, e2e — when each earns its keep,
  and the table-driven idiom reviewers expect.
---

# Testing guide

Tests are how a review trusts your change without re-deriving it. The rule
here is simple and non-negotiable: **every exported type and method ships
with tests, in the same PR.** This page shows the three rings of the test
suite, when each one earns its keep, and the idioms the codebase expects.

## The three rings

Think of the suite as three rings around the code, each answering a different
question:

| Ring | Question it answers | Needs an instance? | Runs in CI? |
| ---- | ------------------- | ------------------ | ----------- |
| **Unit** (`*_test.go`, co-located) | Does this component do what it claims, including when things go wrong? | No | Always |
| **Integration** (`tests/integration/`, `-tags integration`) | Do builders, serialization, and error mapping agree end-to-end? | No — HTTP is mocked | Always |
| **E2E** (`tests/e2e/`, `-tags e2e`) | Does the SDK's model match what ServiceNow *actually* sends? | Yes — real credentials | Manual |

```bash
go test ./...                                      # unit — fast, no network
go test -tags integration ./tests/integration/...  # godog BDD over httpmock
go test -tags e2e ./tests/e2e/...                  # live instance from .env
```

Most changes only need the first ring. Add integration coverage when you've
built new request/response plumbing worth exercising end-to-end; run e2e when
you've modeled a new response shape — unit tests only prove your mocks
round-trip through your *own* serializer, and the real instance is the only
authority on what ServiceNow sends.

## Unit tests: the idiom

Unit tests are **table-driven with `testify`**, and HTTP is stubbed with
`httpmock` (plus the `testify/mock`-based doubles in `internal/mocking`).
The canonical shape, matching what you'll find throughout the repo:

```go
func TestTableItemRequestBuilder_Get(t *testing.T) {
    tests := []struct {
        name     string
        status   int
        body     string
        expected string
        wantErr  error
    }{
        {name: "success", status: 200, body: `{"result": {"number": {"value": "INC0010001"}}}`, expected: "INC0010001"},
        {name: "not found", status: 404, body: `{"error": {"message": "No Record found"}}`, wantErr: &core.NotFoundError{}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            httpmock.Activate()
            defer httpmock.DeactivateAndReset()
            httpmock.RegisterResponder("GET",
                "https://instance.service-now.com/api/now/table/incident/abc123",
                httpmock.NewStringResponder(tt.status, tt.body))

            // ... execute the builder and assert with require/assert ...
        })
    }
}
```

What reviewers look for in that table:

- **The failure rows exist.** A mapped API error (401, 404), the nil-guard
  sentinels (`snerrors.ErrNilRequestBuilder` when called on a nil builder),
  and the success path — not just the success path.
- **Data looks like ServiceNow.** `INC0010001`, not `"foo"` — realistic data
  catches realistic bugs (field naming, value nesting).
- **The request is asserted, not just the response.** Verify headers, query
  parameters, and body reached the responder as intended.
- **Bug fixes carry their reproduction.** The test that fails before your fix
  and passes after is the regression guard.

## Integration tests: Gherkin over mocks

The integration ring describes behavior in `.feature` files under
`tests/integration/features/`, with step definitions binding them to the SDK
via `godog`, against `httpmock` responders — readable specs, no live
instance. New steps follow the pattern in the existing
`*_steps_test.go` files, build-tagged `//go:build integration`.

## Coverage

Codecov tracks coverage on every PR. Locally:

```bash
./scripts/test.sh --report     # unit tests + HTML report (coverage.html)
# or the raw tools:
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```

Coverage is a floor, not a target — hold the line or raise it, but reviewers
care more about the failure rows in your tables than the percentage.

## Next steps

- Adding a whole module? The [playbook](add-api-module.md) includes the
  test expectations per file, and its "verify against a live instance" step
  is exactly the e2e discipline described above.
- Not sure which sentinel an error path should return? See
  [error-handling design](design-error-handling.mdx).
