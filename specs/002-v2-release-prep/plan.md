# Implementation Plan: v2 Release Prep

**Branch**: `002-v2-release-prep` | **Date**: 2026-06-15 | **Spec**: [specs/002-v2-release-prep/spec.md](spec.md)

## Summary
Prepare the ServiceNow SDK for its v2.0 release by removing all deprecated code, consolidating the query system on `query2`, and ensuring a clean repository state. This includes updating documentation and validating the migration path for users.

## Technical Context

**Language/Version**: Go 1.25.0

**Primary Dependencies**: Microsoft Kiota Abstractions & HTTP

**Storage**: N/A

**Testing**: go test (Unit), testify, Godog (Integration)

**Target Platform**: Go-supported platforms

**Project Type**: SDK Library

**Performance Goals**: N/A (Maintenance focus)

**Constraints**: Must maintain SemVer 2.0.0; v2.0 is a major breaking release.

**Scale/Scope**: Repository-wide removal of deprecated items and query system consolidation.

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- [x] Principle I: Library-First - The SDK is a collection of libraries.
- [x] Principle V: Versioning - v2.0 release correctly follows SemVer for breaking changes.

## Project Structure

### Documentation (this feature)

```text
specs/002-v2-release-prep/
├── plan.md              # This file
├── research.md          # Deprecated items scan and query selection
├── data-model.md        # Changes to public API surface
├── quickstart.md        # Migration validation guide
├── contracts/           # API Contract changes (removed items)
└── tasks.md             # Implementation tasks
```

### Source Code (repository root)

```text
/
├── attachment-api/      # Cleaned of deprecated methods
├── core/                # Cleaned of deprecated methods and older query types
├── credentials/         # Cleaned of deprecated providers
├── table-api/           # Cleaned of deprecated builders and response types
├── query2/              # Promoted as the standard query system
└── query/               # [TO BE REMOVED]
```

**Structure Decision**: Standard Go package structure remains unchanged, but the contents will be significantly pruned.

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

N/A
