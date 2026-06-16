# Tasks: Robust Testing Framework

**Input**: Design documents from `/specs/003-robust-testing-framework/`

**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: Tests are MANDATORY as this feature is a testing framework itself.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3, US4)
- Include exact file paths in descriptions

## Path Conventions

- Paths shown below are based on the implementation plan structure.

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 Create project structure for testing utilities in `internal/new/mocking/`, `internal/new/testutils/`, `tests/integration/`, `tests/e2e/`, and `scripts/`
- [x] T002 [P] Initialize `.env` support and `.env.example` in root directory for E2E credentials
- [x] T003 [P] Add BDD dependencies (Godog) to `go.mod`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T004 Implement JSON loading helper in `internal/new/testutils/json_helper.go`
- [x] T005 [P] Implement base `MockRequestAdapter` in `internal/new/mocking/request_adapter.go`
- [x] T006 [P] Implement base `MockResponseHandler` in `internal/new/mocking/response_handler.go`
- [x] T007 Implement structured logging (JSON) configuration for test failures in `internal/new/testutils/logger.go`

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Standardized Unit Testing (Priority: P1) 🎯 MVP

**Goal**: Consistent way to write unit tests for API modules without a real ServiceNow instance.

**Independent Test**: Create a dummy request builder and verify it can be tested using `MockRequestAdapter` and `testutils.LoadJSON`.

### Tests for User Story 1

- [x] T008 [P] [US1] Create unit test for `testutils.LoadJSON` in `internal/new/testutils/json_helper_test.go`
- [x] T009 [P] [US1] Create unit test for `MockRequestAdapter` in `internal/new/mocking/request_adapter_test.go`

### Implementation for User Story 1

- [x] T010 [US1] Finalize `MockRequestAdapter` with support for capturing `RequestInformation` in `internal/new/mocking/request_adapter.go`
- [x] T011 [US1] Ensure `MockResponseHandler` correctly maps to models in `internal/new/mocking/response_handler.go`
- [x] T012 [US1] Update `internal/new/testutils/json_helper.go` to support relative `testdata/` paths
- [x] T013 [US1] Refactor an existing unit test (e.g., `account-api/account_request_builder_test.go`) to use the new Test Table pattern and `mocking` utilities

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently.

---

## Phase 4: User Story 2 - BDD Integration Testing (Priority: P2)

**Goal**: Use Gherkin feature files to verify SDK interactions.

**Independent Test**: Running a Godog feature file in `tests/integration/` returns successful Gherkin steps.

### Tests for User Story 2

- [x] T014 [P] [US2] Create initial Gherkin feature for Table API in `tests/integration/features/table_api.feature`
- [x] T015 [P] [US2] Create Godog test runner in `tests/integration/integration_test.go`

### Implementation for User Story 2

- [x] T016 [US2] Implement Table API step definitions in `tests/integration/steps/table_api_steps.go`
- [x] T017 [US2] Configure Godog to use `MockRequestAdapter` for local integration runs in `tests/integration/integration_test.go`

**Checkpoint**: User Story 2 (BDD) should work independently using mocks.

---

## Phase 5: User Story 4 - End-to-End (E2E) Verification (Priority: P2)

**Goal**: Run tests against a live ServiceNow instance.

**Independent Test**: Running tests in `tests/e2e/` successfully authenticates and performs a basic GET against a live sandbox.

### Tests for User Story 4

- [x] T018 [P] [US4] Create E2E test suite initialization in `tests/e2e/e2e_test.go`
- [x] T019 [P] [US4] Create basic CRUD E2E test for incidents in `tests/e2e/incident_e2e_test.go`

### Implementation for User Story 4

- [x] T020 [US4] Implement credential loading from environment variables in `tests/e2e/e2e_test.go`
- [x] T021 [US4] Add manual trigger logic (e.g., `go test -tags=e2e`) to prevent accidental runs in `tests/e2e/e2e_test.go`
- [x] T022 [US4] Add authentication failure reporting in `tests/e2e/e2e_test.go`

**Checkpoint**: E2E tests should be manually executable against live instances.

---

## Phase 6: User Story 3 - Automated Coverage & CI/CD Integration (Priority: P3)

**Goal**: Automated test execution and combined coverage reporting.

**Independent Test**: Running `./scripts/test.sh --report` generates a combined `coverage.html` file.

### Implementation for User Story 3

- [x] T023 [US3] Create unified test runner script in `scripts/test.sh`
- [x] T024 [US3] Implement combined coverage profile merging in `scripts/test.sh`
- [x] T025 [US3] Add HTML report generation to `scripts/test.sh`
- [x] T026 [US3] [P] Configure GitHub Actions workflow in `.github/workflows/test.yml` to run unit and integration tests

---

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [x] T027 [P] Update `Readme.md` with instructions on writing tests using the new framework
- [x] T028 Finalize `quickstart.md` validation by running all example commands
- [x] T029 [P] Ensure all unit tests across the repo adhere to the Test Table pattern
- [x] T030 Ensure JSON logging for failures is active across all test types

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies.
- **Foundational (Phase 2)**: Depends on Phase 1 completion. BLOCKS all user stories.
- **User Stories (Phase 3+)**: All depend on Phase 2. US1 is the MVP. US2 and US4 are independent of each other. US3 integrates them all.

### Parallel Opportunities

- T002 and T003 in Setup.
- T005 and T006 in Foundational.
- Tests (T008, T009) and implementation can be parallelized within US1 if multiple devs are present.
- US2 and US4 can proceed in parallel once Phase 2 is done.

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Setup + Foundational.
2. Complete Phase 3: User Story 1.
3. Validate by refactoring one existing test package to use the new pattern.

### Incremental Delivery

1. Deliver Unit Testing framework (US1).
2. Deliver BDD Integration framework (US2).
3. Deliver E2E Live Verification (US4).
4. Deliver Automated Coverage & CI/CD (US3).
