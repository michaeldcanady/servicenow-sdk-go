# ADR: Standardizing Error Handling and Messaging

## Status
Proposed

## Context
The ServiceNow SDK for Go has inconsistent error handling patterns, including varying message phrasing ("can't be nil" vs "is nil") and hard-coded error messages scattered throughout the codebase. This hinders maintainability and compromises the developer experience for v2.0.

## Decision
We will standardize error handling with the following principles:
1. **Centralized Error Package**: Create a new `/errors` package to house sentinel errors, custom error types, and error-related utilities.
2. **Standardized Messaging**: Adopt strict phrasing for error messages to ensure consistency:
   - Use `"[parameter] cannot be nil"` for nil checks.
   - Use `"[parameter] is required"` for missing inputs.
   - Avoid contractions (use `cannot`, not `can't`).
3. **Behavioral Consistency**: Centralize validation logic (for example, parameter nil checks) within the `errors` package where feasible, and apply consistent validation across all constructors and API methods.

## Consequences
- **Pros**: Improved code consistency, easier maintenance, and a cleaner, more professional developer experience.
- **Cons**: Requires a refactoring effort across all API packages.
