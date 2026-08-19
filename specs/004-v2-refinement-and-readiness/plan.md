# Implementation Plan: v2.0 Refinement & Readiness

**Spec**: [specs/004-v2-refinement-and-readiness/spec.md](spec.md) | **Date**: 2026-06-17

## Summary
This plan outlines the final refinements and quality-of-life improvements required for the ServiceNow SDK for Go v2.0 release. It focuses on architectural consistency, closing feature gaps, and ensuring robust test coverage beyond simple code pruning.

## Strategic Objectives
1.  **API Polish**: Standardize naming conventions and remove "experimental" suffixes.
2.  **Feature Parity**: Ensure critical HTTP methods (HEAD) and helpers (Multipart) are available.
3.  **Zero-TODO Policy**: Resolve all remaining technical debt markers in the codebase.
4.  **Contributor Excellence**: Align documentation with modern project standards and Go 1.25+ requirements.

## Phases

### Phase 1: Architectural Standardization
- Rename all `*2` suffixed types and packages to their clean equivalents.
- Consolidate `internal/new` into `internal`.
- Standardize RequestBuilder factory patterns.

### Phase 2: Functional Readiness
- Implement `HEAD` methods for all RequestBuilders.
- Refine error mapping for specific HTTP status codes.
- Implement high-level Multipart body helpers for attachments.

### Phase 3: Testing & Quality Assurance
- Resolve 19+ `TODO: add tests` markers.
- Implement standardized mocking using `mockery` or compatible patterns.
- Validate coverage for critical path serialization/deserialization.

### Phase 4: Documentation & Environment
- Update `CONTRIBUTING.md` for Go 1.25.0 and automated release workflows.
- Add comprehensive GoDoc examples for core modules.
- Audit dependencies and environment configurations.

## Success Criteria
- [ ] No `*2` suffixes remain in the public API.
- [ ] No `TODO` markers remain in the codebase.
- [ ] All RequestBuilders support GET, POST, and HEAD.
- [ ] `CONTRIBUTING.md` accurately reflects the v2 workflow.
- [ ] 100% pass rate on unit and integration tests with improved coverage.
