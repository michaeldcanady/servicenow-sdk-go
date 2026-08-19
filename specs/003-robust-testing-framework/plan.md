# Implementation Plan: Robust Testing Framework

**Branch**: `003-robust-testing-framework` | **Date**: 2026-06-16 | **Spec**: [spec.md](spec.md)

**Input**: Feature specification from `/specs/003-robust-testing-framework/spec.md`

## Summary

Build a robust testing framework for the ServiceNow SDK for Go, incorporating standardized unit testing with Testify/httpmock using Test Tables, BDD-style integration testing with Godog, and manual/automated E2E verification against live ServiceNow instances.

## Technical Context

**Language/Version**: Go (Officially supported versions: n and n-1, e.g., 1.25 and 1.26)

**Primary Dependencies**: Microsoft Kiota (abstractions, http), Godog (BDD), Testify (Unit), httpmock (HTTP Mocking)

**Storage**: N/A (SDK Library)

**Testing**: Testify (Unit with Test Tables), Godog (Integration/E2E), httpmock (Mocking)

**Target Platform**: Cross-platform (Go supported environments)

**Project Type**: SDK Library

**Performance Goals**: Fast unit tests (<5s for whole suite), reliable and repeatable integration tests.

**Constraints**: Must maintain consistency with Kiota architecture; surgical changes preferred; no credential logging.

**Scale/Scope**: Repo-wide framework applicable to all current and future API modules.

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- [x] **Gate 1: Surgical Changes**: Does the framework implementation minimize changes to existing core logic while providing new testing capabilities?
- [x] **Gate 2: Idiomatic Go**: Does the testing framework follow standard Go testing conventions and use the established Testify/Godog patterns?
- [x] **Gate 3: Kiota Alignment**: Are the mocking utilities consistent with Kiota's `RequestAdapter` and `ResponseHandler` abstractions?
- [x] **Gate 4: Testing Requirements**: Does the plan ensure unit tests remain in their respective packages and integration tests are centralized in `tests/`?
- [x] **Gate 5: SemVer Consistency**: Are the changes categorized correctly (likely MINOR for a new feature) for `release-please`?
- [x] **Gate 6: Go Version Support**: Does the framework support the two most recent major Go releases?
- [x] **Gate 7: Credential Security**: Does the framework source E2E credentials from environment variables without logging them?
- [x] **Gate 8: Test Table Pattern**: Is the Test Table pattern mandated for all unit tests?

## Project Structure

### Documentation (this feature)

```text
specs/003-robust-testing-framework/
├── plan.md              # This file
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
├── contracts/           # Phase 1 output
└── tasks.md             # Phase 2 output (generated later)
```

### Source Code (repository root)

```text
internal/
├── new/
│   ├── mocking/         # Mock implementations for RequestAdapter/ResponseHandler
│   └── testutils/       # Shared unit testing helpers (loading JSON, etc.)

tests/
├── integration/         # Centralized Godog feature files and step definitions
└── e2e/                 # E2E test suite for live instance verification

scripts/
└── test.sh              # Unified test runner script
```

**Structure Decision**: Centralized `internal/new/mocking` for shared Kiota mocks to avoid duplication, and `tests/` for all non-unit tests.

## Complexity Tracking

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| N/A | | |
