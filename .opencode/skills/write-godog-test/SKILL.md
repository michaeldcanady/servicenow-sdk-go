---
name: write-godog-test
description: Writes or extends Godog BDD integration tests for this repo (servicenow-sdk-go) — Gherkin .feature files plus their Go step definitions under tests/integration/ — using the same six-step audit-first discipline as write-unit-tests, adapted to endpoint-level scenarios instead of function-level paths. Use whenever the user asks to "add integration tests for X", "write a feature file for Y", "cover this endpoint with Godog/BDD tests", or after adding/changing a request-builder verb method on a *significant* API surface (a new module, a new verb, a behavior change a consumer would notice) that has no matching .feature coverage yet. This is the repo's only integration-test authoring path — it replaced the old godog-test-writer agent, so don't look for that agent; run this skill directly instead. Not for unit tests (use write-unit-tests) or e2e tests (tests/e2e/, manual, hits a live instance).
---

# Writing Godog integration tests for servicenow-sdk-go

This is `write-unit-tests`' sibling, aimed one level up: instead of auditing a
function for every branch it can take, you're auditing an **endpoint** for
every observable behavior a consumer of the SDK would care about, then
expressing that as a Gherkin scenario backed by `httpmock`. Same six-step
shape, different altitude — don't skip steps or write `.feature`/step code
before step 5.

Work one module or tightly-related group of endpoints at a time (e.g. "the
table API's CRUD verbs", not "every module in the SDK"). If asked to cover
several modules, repeat the loop per module.

## Step 1 — Audit the endpoint(s)

Read the target request builder file(s) directly — don't infer behavior from
the module's name. You need to know:

- Every verb method (`Get`/`Post`/`Patch`/`Put`/`Delete`) it exposes, its
  inputs (path params, query parameters, request body shape), and its
  success response shape (`core.ServiceNowItemResponse[T]` vs
  `core.ServiceNowCollectionResponse[T]`, paginated or not).
- What errors it can actually surface at the HTTP boundary — a 404 on a
  missing record, a 400/422 on bad input, a 5XX mapped through
  `core.DefaultErrorMapping()` — since that's the granularity integration
  tests operate at (unlike unit tests, you're not asserting on
  `snerrors.ErrNilRequestAdapter`-style construction-time guards; those are
  `write-unit-tests`' job, not this skill's).
- Whether it's paginated (uses `core.PageIterator` / `Link` headers) — that's
  its own scenario shape, distinct from plain CRUD.

Then read the closest existing `.feature` + `*_steps_test.go` pair for a
module with a similar shape — `tests/integration/features/table_crud.feature`
+ `tests/integration/table_steps_test.go` is the fullest reference (CRUD +
query + pagination); `attachment_crud.feature` /
`attachment_steps_test.go` is a good non-tableapi comparison. Note their
tag vocabulary, `Background:` shape, and step phrasing style — you'll reuse
it, not invent your own.

## Step 2 — Map the scenarios: happy vs. non-happy

List the distinct observable behaviors a consumer would notice, same
happy/non-happy split as `write-unit-tests` but at endpoint granularity:

- Happy: a successful call with valid input, for each verb this
  endpoint exposes (a single "successful CRUD" scenario chaining
  create→read→update→delete, as `table_crud.feature` does, is normal and
  preferred over one scenario per verb when the verbs share setup).
- Non-happy: the errors from step 1's audit that this endpoint can actually
  produce at the HTTP level — a 404 after delete, a validation error on bad
  input, an empty result set, a boundary pagination case (last page,
  single-item collection).

This repo treats integration tests as coverage for *significant* API
surface, not exhaustive branch coverage — don't add a scenario for a
trivial internal refactor that unit tests already cover, and don't invent a
non-happy scenario the endpoint can't actually produce. If step 1 didn't
turn up a real error path, there isn't a non-happy scenario to write here.

## Step 3 — Check for existing coverage

Search `tests/integration/features/*.feature` for a feature already
covering this endpoint, and its paired `*_steps_test.go`. If one exists:

- Map each existing scenario back to a step-2 behavior. Extend the existing
  `.feature` file and step-definition file with missing scenarios/steps
  rather than starting a second parallel feature file for the same
  endpoint.
- Reuse existing step phrasing verbatim when the behavior is the same (e.g.
  `the response should not be an error`, `the response should be a 404
  error`) — register new Gherkin text with `ctx.Step` only for genuinely new
  assertions, not paraphrases of ones that already exist anywhere in
  `tests/integration/`. Grep across all `*_steps_test.go` files for this,
  not just the one file you're extending — steps are shared repo-wide
  through Godog's step registration.

If nothing exists yet, you're authoring both files fresh, matching the same
target shape.

## Step 4 — Confirm the map is fully covered

Cross-check the step-2 behavior list against what step 3 found plus what
you're about to add. Every behavior lands in exactly one scenario. If a
behavior needs an HTTP response shape `setupGlobalMocks()` in
`tests/integration/setup_test.go` doesn't register yet (a new endpoint path,
a new status code, a new response body), you'll add that responder in step
5 rather than skip the scenario.

## Step 5 — Write the feature file and step definitions

**`.feature` file** (`tests/integration/features/<name>.feature`):

```gherkin
@integration @mock @<module> @<subject>
Feature: <one-line capability statement>
  As a developer using the ServiceNow SDK
  I want to <capability>
  So that <reason>

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @<module> @<subject>
  Scenario: <plain-English description of the behavior>
    When <action>
    Then the response should not be an error
    And <specific assertion>
```

Reuse this repo's existing tag vocabulary (`@integration`, `@mock`,
`@table`, `@query`, `@crud`, `@pagination`, `@batch`, `@attachment`, etc. —
check current tags across `features/*.feature` first) instead of inventing
new tags for concepts this repo already names.

**Step definitions** (`tests/integration/<module>_steps_test.go`):

- `//go:build integration` build tag, `package integration`.
- A `<module>TestContext` struct holding the client, response, error, and
  any scenario-scoped state (`lastSysID`, counters, etc. — see
  `tableTestContext`).
- Step methods as receivers on that struct, named to mirror the Gherkin text.
- `iHaveAValidServiceNowInstanceAndCredentials` / `iHaveInitializedThe...`
  are shared boilerplate — copy the existing implementation, don't
  reinvent it.
- Register everything in an `InitializeXScenario(ctx *godog.ScenarioContext)`
  function with `ctx.Before`/`ctx.After` hooks calling
  `setupGlobalMocks()`/`httpmock.DeactivateAndReset()`, and a
  `TestXFeatures(t *testing.T)` godog suite runner listing the `.feature`
  paths it covers — match `InitializeTableScenario`/`TestTableFeatures`
  exactly.
- If step 4 found a missing HTTP responder, add it to `setupGlobalMocks()`
  in `setup_test.go` (and any fixture JSON it needs to `mock_data_test.go`),
  following the existing per-module URL-building and `httpmock.Register*`
  pattern — don't hand-roll a second mock-setup mechanism.

## Step 6 — Curveballs: now try to break it

Same spirit as `write-unit-tests`' step 6, aimed at the HTTP boundary
instead of function internals. With the expected scenarios passing, try a
few things the endpoint's author probably didn't script for:

- An unexpected status code in the family it maps (e.g. a 403 where only
  404/5XX are mocked) — does `core.DefaultErrorMapping()` actually resolve
  it, or does the SDK mishandle an unmapped code?
- A response body missing a field the deserializer expects (empty
  `result`, a null where a string is expected) — does the SDK return a
  clear error, or does it panic/return a zero-value silently?
- Pagination boundaries: a single-item collection, an empty collection, a
  `Link` header pointing at itself (would the `PageIterator` loop forever?).
- Reusing `c.lastSysID` after a delete (a get-after-delete race a real
  client could trigger) — already modeled in `table_crud.feature`, but
  check the same shape applies to whatever endpoint you're covering.

Add any curveball that reveals a real, reachable behavior as a permanent
scenario in the `.feature` file from step 5. If a curveball actually
breaks something (panic, hang, wrong status mapping, silently swallowed
error), **don't silently patch the production code**. Add the scenario in a
form that documents the break (a comment above it, or `t.Skip` in the
backing step if Godog can't express "expected fail" natively for that
case), then report the concrete finding to the user: what request/response
triggered it, what happened, what you'd expect instead. Filing it as a
tracked issue (via the `write-issue` skill) is often the right next step —
ask the user rather than assuming.

## After writing

```bash
go test -tags=integration ./tests/integration/... -run TestXFeatures -v
```

Confirm the new/extended scenarios pass. If you added or changed shared
mocks in `setup_test.go`, run the full integration suite
(`go test -tags=integration ./tests/integration/...`) to make sure you
didn't break another module's scenarios that share the same base URL
prefix.

## Scope notes

- This is the repo's whole integration-test authoring path — there's no
  separate scanning agent to defer to first; go straight from "what needs
  coverage" to auditing and writing it yourself, module by module.
- Function-level branch/error-path coverage (nil-guards, sentinel errors,
  deserialization edge cases in isolation) is `write-unit-tests` territory,
  not this skill's — don't duplicate that work here just because you're
  already looking at the same request builder.
- `tests/e2e/` was superseded by the E2E runner in `tests/integration/v2/e2e_test.go`
  (`//go:build e2e`, real ServiceNow instance, run manually). E2E scenarios are tagged `@e2e`
  in the same feature files this skill maintains.
