# Tasks: v2 Release Prep

**Input**: Design documents from `/specs/002-v2-release-prep/`

**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/removal.md

**Tests**: Tests are encouraged for validating the migration and ensuring no regressions in the pruned codebase.

**Organization**: Tasks are grouped into Setup, Foundational (Code Removal), and User Stories (Documentation & Metadata).

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (US1, US2, US3)

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Initial prep and verification

- [x] T001 Verify current v1.x baseline tests pass (`go test ./...`)
- [x] T002 [P] Ensure Go 1.25.0 toolchain is active in `.devcontainer/devcontainer.json` or environment

---

## Phase 2: Foundational (Blocking Pruning & Consolidation)

**Purpose**: Remove deprecated code and consolidate query systems. This MUST be completed before documenting changes.

- [ ] T003 [P] Remove deprecated entities and methods in `attachment-api/` (refer to `data-model.md`)
- [ ] T004 [P] Remove deprecated methods in `core/helper.go`, `core/page_iterator.go`, `core/request_builder_deprecated.go`, and `core/url_information.go`
- [ ] T005 [P] Remove deprecated credential providers in `credentials/` (refer to `data-model.md`)
- [ ] T006 [P] Remove deprecated types and builders in `table-api/` (refer to `data-model.md`)
- [ ] T007 [P] Delete `servicenow_client_deprecated.go` and `now_request_builder_deprecated.go`
- [ ] T008 [P] Remove `query/` directory entirely (FR-008)
- [ ] T009 Update internal references to `query/` in `docs/` and `internal/` to use `query2/`
- [ ] T010 [P] Update `go.mod` module path to include `/v2` suffix (Standard Go v2 practice)
- [ ] T011 Update all internal imports in the repository to use the new `.../v2/...` module path (depends on T010)
- [ ] T012 Run `go test ./...` to verify structural integrity and no regressions in surviving code

**Checkpoint**: Foundation ready - Codebase is clean and follows v2 standards.

---

## Phase 3: User Story 1 - Document Breaking Changes (Priority: P1)

**Goal**: Clearly identify and document all breaking changes.

**Independent Test**: Verify `docs/v2-breaking-changes.md` exists and lists all items removed in Phase 2.

### Implementation for User Story 1

- [ ] T013 [P] [US1] Create `docs/v2-breaking-changes.md` with a detailed list of all removed APIs and their replacements
- [ ] T014 [US1] Update `CHANGELOG.md` with a summary of the v2.0 breaking changes
- [ ] T015 [US1] Finalize breaking changes list based on actual removals in T003-T007

**Checkpoint**: Breaking changes are fully documented for external consumers.

---

## Phase 4: User Story 2 - Provide Migration Guide (Priority: P2)

**Goal**: Provide a clear path for users to upgrade from v1.x to v2.0.

**Independent Test**: Follow the migration guide with a sample v1 project and verify successful build against v2.

### Implementation for User Story 2

- [ ] T016 [P] [US2] Create `docs/v2-migration-guide.md` based on `specs/002-v2-release-prep/quickstart.md`
- [ ] T017 [US2] Add "Before and After" code examples for Client Initialization, Query Building, and Request Execution
- [ ] T018 [US2] Link the migration guide from the root `Readme.md`

**Checkpoint**: Migration guide is ready for public consumption.

---

## Phase 5: User Story 3 - Finalize v2.0 Release Metadata (Priority: P3)

**Goal**: Update repository metadata and CI pipelines for the major release.

**Independent Test**: Dry-run the release pipeline and verify v2.0.0 versioning.

### Implementation for User Story 3

- [ ] T019 [US3] Update `VERSION` file to `2.0.0`
- [ ] T020 [US3] Update `release-please-config.json` to handle the v2.0 release
- [ ] T021 [US3] Validate GitHub Actions workflows (`.github/workflows/`) for major version tagging

**Checkpoint**: Repository metadata is finalized for the v2.0 release.

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Final cleanup and verification

- [ ] T022 [P] Run `go fmt ./...` and `go vet ./...` across the entire project
- [ ] T023 Run full suite of unit and integration tests (`go test ./...` and `godog`)
- [ ] T024 Perform a final review of the generated documentation for clarity and accuracy
- [ ] T025 Run `specs/002-v2-release-prep/quickstart.md` validation scenarios

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: Prerequisite for all tasks.
- **Foundational (Phase 2)**: MUST be completed before documentation (US1, US2) as it defines what was actually changed.
- **User Stories (Phase 3-5)**: Depend on Codebase pruning. US1 and US2 are closely related. US3 can run in parallel with documentation.
- **Polish (Phase 6)**: Final verification step.

### Parallel Opportunities

- All removals in Phase 2 (T003-T007) can run in parallel.
- Documentation creation (T013, T016) can run in parallel once removals are identified.
- Metadata updates (T019-T021) can run in parallel with documentation.

---

## Implementation Strategy

### MVP First (Pruning & Documentation)

1. Complete Setup (Phase 1).
2. Complete Foundational removal of deprecated code (Phase 2).
3. Generate the Breaking Changes list (US1).
4. **VALIDATE**: Ensure the repo compiles and tests pass for remaining code.

### Incremental Delivery

1. Pruning → Clean Codebase.
2. Breaking Changes Doc → Commmunicated changes.
3. Migration Guide → Enabled upgrades.
4. Metadata & Release → Completed v2 Prep.
