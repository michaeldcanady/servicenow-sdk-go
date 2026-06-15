# Feature Specification: Prerelease Workflow

**Feature Branch**: `001-prerelease-workflow`

**Created**: 2026-06-15

**Status**: Draft

**Input**: User description: "build a CI/CD workflow for prerelease versions"

## Clarifications

### Session 2026-06-15
- Q: Instead of having a dedicated "develop" branch ideally we'd use the main branch, is this feasible? → A: Yes, we will use the `main` branch as the primary trigger for automated prerelease builds.
- Q: How should we handle commits that are part of a formal release process on the main branch? → A: Exclude commits from the release bot (e.g., "chore: release") from triggering prerelease builds.
- Q: How should this interact with stable releases? → A: Prerelease versions MUST be based on the next intended stable version, supporting manual overrides (e.g., forcing to 1.12.0) via commit metadata or configuration.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Automated Prerelease on Main Push (Priority: P1)

As a developer, I want a prerelease version to be automatically built and published whenever code is merged into the `main` branch, so that integration testing can happen against the latest changes without waiting for a formal production release.

**Why this priority**: High priority as it enables continuous integration and testing across teams using the SDK.

**Independent Test**: Push a commit to `main`. Verify that a new prerelease version (e.g., `v1.12.0-alpha.1`) is created, tagged, and available in the package registry.

**Acceptance Scenarios**:

1. **Given** a new commit on `main`, **When** the workflow completes, **Then** a new git tag with a prerelease suffix is created.
2. **Given** a successful build, **When** the publish step runs, **Then** the Go module is accessible via the registry with the new prerelease version.

---

### User Story 2 - Manual Prerelease Trigger (Priority: P2)

As a release manager, I want to manually trigger a prerelease build for any branch or commit, specifying the prerelease type (alpha, beta, rc), so that I can provide specific versions for targeted testing.

**Why this priority**: Enables flexibility for special testing scenarios or stabilizing a release candidate.

**Independent Test**: Trigger the workflow manually, selecting a branch and entering "beta" as the suffix. Verify the resulting version is `vX.Y.Z-beta.N`.

**Acceptance Scenarios**:

1. **Given** a manual trigger with "rc" selected, **When** the workflow completes, **Then** the version produced includes the "-rc" suffix.

### Edge Cases

- What happens when a build fails during the prerelease process? (Should not publish, should notify developers)
- How does the system handle concurrent pushes to the trigger branch? (Should queue or cancel older builds to ensure only the latest is published)
- What happens if the calculated version already exists in the registry? (Build should fail with a clear error)
- **How to handle release-please commits?** (The workflow MUST NOT trigger for commits authored by the release bot or matching release patterns to avoid circular builds).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The system MUST trigger an automated prerelease build on push to the `main` branch, excluding commits from the automated release bot.
- **FR-002**: The system MUST provide a manual trigger for prerelease builds on any available branch.
- **FR-003**: The system MUST automatically calculate the next unique prerelease version string using a sequential counter (e.g., `alpha.1`, `alpha.2`).
- **FR-004**: The system MUST publish the build artifacts to the GitHub Packages registry.
- **FR-005**: The system MUST create a permanent reference (tag) in the version control system for every successfully published prerelease version.
- **FR-006**: The system MUST ensure all quality gates (tests and linting) pass before publishing a prerelease.

### Key Entities

- **Prerelease Version**: A SemVer-compliant version string with a prerelease suffix (e.g., `1.12.0-alpha.1`).
- **Build Artifact**: The Go module package and associated metadata.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Prerelease versions are available for consumption within 15 minutes of a push to the trigger branch.
- **SC-002**: 100% of published prereleases have a corresponding Git tag and a record in the CI system.
- **SC-003**: No manual editing of version files is required to produce a prerelease version.

## Assumptions

- `main` is the intended branch for automated prerelease builds.
- The project's existing testing infrastructure (Go tests, Godog) is sufficient for validation.
- GitHub Actions is the execution environment.
- The base version (e.g., `1.12.0`) is taken from the `VERSION` file in the repository or calculated from commit history.
