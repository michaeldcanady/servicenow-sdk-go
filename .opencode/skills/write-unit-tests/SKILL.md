---
name: write-unit-tests
description: >-
  Writes or extends Go unit tests for this repo (servicenow-sdk-go) using a
  six-step audit-first process — audit the target code, map every path (happy
  and non-happy), check for existing tests and reconcile with them, write
  table-driven testify tests, then deliberately try to break the code with
  curveball inputs and report anything that actually breaks. Use whenever the
  user asks to "write tests for X", "add unit tests", "cover this
  function/file", "improve test coverage for Y", or after writing/editing a
  non-test .go file that has no matching coverage yet — proactively suggest
  or run this even if the user only asked for the code change, not the
  tests. Distinct from the principal-qa-engineer agent's repo-wide
  scan-and-report loop: this is the hands-on, one-file-at-a-time process for
  actually writing the test table once you're already looking at the code.
---

# Writing unit tests for servicenow-sdk-go

This skill is a six-step discipline for writing Go tests that actually earn
their keep, rather than tests that just exercise the one path already in
your head. Do not skip steps or write test code before step 5 — steps 1-4
are what keep the test table honest instead of a copy of whatever the
happy-path call site already does.

Work one source file (or one tightly-related group of functions/methods) at
a time. If asked to cover a whole package or the whole diff, repeat this
loop per file rather than trying to hold every file's path map in your head
at once.

## Step 1 — Audit the code under test

Read the actual source file(s), not just the function signature. You need
to know, concretely:

- Every parameter and its zero/nil value — does the function guard against
  it, or trust the caller?
- Every branch, including ones buried in a called helper that lives in the
  same package (e.g. `conversion.IsNil`, `internal/store` accessors,
  `internal/serialization` helpers) — a one-line call can hide several
  return paths.
- Every possible return value, especially every distinct `error` it can
  produce and where each one comes from (a sentinel from `errors/errors.go`,
  a package-local sentinel, a wrapped adapter/HTTP error, a
  `fmt.Errorf`-constructed one).
- What state the function depends on or mutates (backing-store fields via
  `internal/store` accessor/mutator funcs, request builder path parameters,
  etc.) — untested state interactions are a common silent gap here since
  models are backing-store-backed, not plain structs.

If the function is a request-builder verb method (`Get`/`Post`/`Patch`/`Put`/
`Delete`), you already know its opening shape from this repo's convention:
the nil-guard block (`conversion.IsNil(rB)` / `conversion.IsNil(rB.RequestBuilder)`
→ `snerrors.ErrNilRequestBuilder`, `conversion.IsNil(rB.GetRequestAdapter())`
→ `snerrors.ErrNilRequestAdapter`), then `ToXRequestInformation`, then
`Send`/`SendPrimitive` with `core.DefaultErrorMapping()`. Confirm it's
actually there rather than assuming — some older or hand-written methods
predate the convention.

## Step 2 — Map the paths: happy vs. non-happy

From the audit, write out (mentally or as a scratch list, not necessarily in
the final code) every distinct path through the function. A path is
"happy" if it's the intended success case with valid input; everything else
— nil/zero receivers or arguments, a nil adapter, malformed input, an
adapter/HTTP error, a deserialization failure, an out-of-range enum value —
is non-happy.

Don't invent a non-happy path a function can't actually produce. A pure
getter with no error return and no guard has exactly one path — a
happy-path-only test fully covers it, and forcing a fake failure case onto
it is padding, not coverage. The point of this step is precision: name the
real paths, not a fixed quota of "at least one failure case."

## Step 3 — Check for existing tests

Find the co-located `_test.go` file (this repo's convention: one test file
per source file — `tableapi/` is the fullest reference,
`policyapi/state_test.go` is a clean minimal example of the table-driven
style this repo expects). If tests already exist:

- Map each existing test case back to a path from step 2. Cases that don't
  match this repo's standard (table-driven `testify`, not free-form
  `if got != want { t.Fatalf(...) }`, or asserting on error *string text*
  instead of `assert.ErrorIs(t, err, snerrors.SomeSentinel)`) should be
  rewritten into the standard shape as you touch that test function, not
  left inconsistent next to new table-driven cases.
- Extend the existing test table with rows for any path from step 2 that
  isn't already a row, rather than writing a second, parallel test function
  for the same symbol.

If nothing exists yet, you're writing the file fresh — follow the same
target shape.

## Step 4 — Confirm the map is fully covered

Before writing code, cross-check your step-2 path list against what step 3
found (existing rows) plus what you're about to add. Every path should land
in exactly one test row. If a path genuinely can't be exercised through this
repo's test seams (e.g. `internal/mocking`'s shared mocks would need a type
assertion that panics on a bare nil), don't silently drop it — see the
workaround note in Step 5, and only report it as untestable if that doesn't
work either.

## Step 5 — Write the tests as test tables

Match this repo's established shape exactly:

```go
func TestX(t *testing.T) {
    tests := []struct {
        name    string
        // ...inputs...
        want    Y
        wantErr error // or bool, if the package doesn't use a sentinel here
    }{
        {"describes the case", /* ... */},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := X(tt.input)
            if tt.wantErr != nil {
                assert.ErrorIs(t, err, tt.wantErr)
                return
            }
            require.NoError(t, err)
            assert.Equal(t, tt.want, got)
        })
    }
}
```

Conventions to hold to, per this repo's `CLAUDE.md` and existing packages:

- `testify` `assert`/`require`, table-driven, co-located `_test.go`.
- HTTP interactions via `httpmock`; internal collaborators via the
  `testify/mock`-based mocks in `internal/mocking`.
- Assert on sentinel identity (`assert.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)`),
  never on error message text — this repo has been burned by duplicated
  sentinel-shaped `errors.New(...)` calls that break `errors.Is`, so a
  string-text assertion silently tolerates that regression instead of
  catching it.
- Reuse an existing sentinel from `errors/errors.go` or the package-local
  `errors.go` by exact identity; don't invent a new one just for the test.
- If a shared `internal/mocking` mock would need to return a `nil` of a
  concrete interface/struct type and that panics a type assertion inside
  the code under test, try returning a valid non-nil dummy value of the
  right concrete type alongside the injected error first — most call sites
  check the error before touching the value, so this usually sidesteps the
  panic without changing the shared mock.
- Every new or touched exported symbol needs a test; don't leave a function
  you opened only partially covered.

## Step 6 — Curveballs: now try to break it

This is the step that's easy to skip once the "expected" table passes, and
it's the one that actually earns the audit you did in step 1. With the
expected paths covered, spend a few minutes deliberately trying to break the
function with inputs its author probably didn't think about:

- Boundary values: empty string vs. whitespace-only string, zero-length
  slice/map vs. nil slice/map, `0`/`-1`/max-int enum values, an enum cast
  from an out-of-range int (`State(999)`-style).
- Concurrent/repeated calls where the code touches a backing store, if
  that's plausible for the function.
- Malformed JSON/field-deserializer input for anything touching
  `GetFieldDeserializers`/`Serialize`.
- Passing a builder/receiver that's the zero value vs. explicitly `nil` —
  these are not always guarded identically, and `conversion.IsNil` is a
  reflect-based check specifically because a typed-nil interface and a
  literal `nil` don't always behave the same in a naive `== nil` check.
- A context that's already cancelled or has a zero deadline, for anything
  taking a `context.Context`.

Add any curveball that reveals a *new* path to the test table from step 5 —
it belongs there permanently, not just in a scratch exploration.

If a curveball actually breaks the code (panic, wrong value, silently
swallowed error, hang), **do not silently patch the production code as a side
effect of writing tests.** Add a test that captures the broken behavior
(marked or named so it's clearly documented as a known-bad case, e.g.
`t.Skip("BUG: ...")` or a `_test.go` comment above the case), then report
the concrete failure to the user in your response: what input, what
happened, what you'd expect instead. Let them decide whether it's a real bug
worth a separate fix versus expected/acceptable behavior — a test-writing
pass silently changing production logic hides exactly the kind of gap this
process exists to surface.

## After writing

Run the package's tests and lint before considering the work done:

```bash
go test ./<package>/... -run TestXxx -v
golangci-lint run ./<package>/...
```

## Scope notes

- Unit tests only — co-located `_test.go`, no build tags. Integration
  (`tests/integration/`, `//go:build integration`) and e2e
  (`tests/e2e/`, `//go:build e2e`) suites are out of scope for this skill.
- This is the file-at-a-time authoring loop. If you need a repo-wide
  coverage audit across many files before deciding what to write, that's
  the `principal-qa-engineer` agent's job (classify-then-report across a
  diff or package) — use it first, then come back to this skill per file it
  flags.
