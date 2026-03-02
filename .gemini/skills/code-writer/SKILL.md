---
name: code-writer
description: Always use this skill when the task involves writing, editing, refactoring, or reviewing Go source code in the odc repository, including files in `cmd/`, `internal2/`, `pkg/`, or any `.go`, `.json`, `.yaml`, or configuration files that affect runtime behavior.
---

# `code-writer` skill instructions

As an expert Go engineer for the **odc (go-onedrive)** project, you produce
idiomatic, maintainable, and testable Go code that aligns with the project’s
architecture, domain model, and contribution workflow. When asked to write,
edit, or review code, you must ensure your output is correct, consistent, and
fully aligned with odc’s engineering conventions.

Follow the contribution process in `CONTRIBUTING.md` and the standards below.

---

# **Phase 1: Engineering standards**

Adhere to these principles when writing or modifying Go code.

## Go style and idioms

- **Clarity:** Prefer explicit, readable code over cleverness.
- **Naming:** Use intention‑revealing names. Follow Go naming conventions
  (exported identifiers use PascalCase; unexported use camelCase).
- **Zero values:** Design APIs that work naturally with Go’s zero values.
- **Errors:** Return `error` as the last return value. Use `fmt.Errorf` with
  `%w` for wrapping. Avoid vague messages like `"failed"`.
- **Context:** Accept `context.Context` as the first parameter for all
  long‑running or I/O‑bound operations.
- **Concurrency:** Use goroutines and channels only when necessary. Avoid
  unnecessary shared state.
- **Interfaces:** Define small, behavior‑focused interfaces in the consumer
  package, not the provider.
- **Type organization:** Each type should be within it's own file of the same name in camel-case.
- **Function, variable, and constant organization:** Non-types (functions, variables, and constants) should be organized by function (i.e., like with like).
- **Enum implementation:** All enums should be an int enum, if it needs a string value give it a `String()` method.

## Architecture and design

- **Separation of concerns:** Keep Graph API calls, domain logic, and CLI
  orchestration separate.
- **Dependency boundaries:** Use constructor functions to inject dependencies.
- **Extensibility:** Prefer patterns that support additional Graph endpoints,
  auth flows, and filesystem operations without breaking existing behavior.
- **Consistency:** Match existing odc patterns for services, repositories,
  adapters, and command wiring.
- **CLI command structure:** All new CLI commands must follow the template pattern
  found in the `templates/` directory of this skill. This includes separating the
  command factory, execution logic, and options into distinct files:
  - `command.go`: Cobra command initialization and flag definition.
  - `command_cmd.go`: Command execution logic and dependency resolution.
  - `options.go`: Option structure and validation logic.
- **Subcommand nesting:** Each command and subcommand must be contained within
  its own directory named after the command. Subcommands must be nested
  recursively within their parent command's directory (e.g.,
  `interface/cli/edit/sessions/list/`).

## Testing and reliability

- **Table‑driven tests:** Use table‑driven patterns for all logic with multiple
  cases.
- **Mocks:** Use mock repositories, mock Graph clients, and mock identity
  providers for deterministic tests.
- **No network:** Tests must never hit the real Graph API.
- **Fixtures:** Use realistic OneDrive metadata and content fixtures.
- **Coverage:** Add or update tests when modifying or adding behavior.

## Formatting and structure

- **Formatting:** All code must pass `gofmt` and `go vet`.
- **Imports:** Group imports into stdlib, external, and internal sections.
- **Comments:** Use comments to explain *why*, not *what*. Keep GoDoc concise
  and accurate. All functions, methods, types, constants, and variables should have GoDocs.

---


# **Phase 2: Preparation**

Before writing or modifying code, thoroughly investigate the request and its
context.

1. **Clarify:** Understand the exact change requested. If ambiguous (e.g.,
   “fix this”), ask for clarification.
2. **Investigate:** Examine relevant services, repositories, adapters, and
   command implementations.
3. **Search:** Identify all files affected by the change, including tests,
   mocks, and configuration.
4. **Assess:** Determine whether the change is additive, a refactor, or a
   behavioral modification.
5. **Plan:** Produce a step‑by‑step plan before making any modifications.

---

# **Phase 3: Execution**

Implement your plan using the appropriate file system tools. Use `replace` for
targeted edits and `write_file` for new files or large rewrites.

## Writing new Go code

- **Structure:** Follow odc’s patterns for services, repositories, and command
  wiring.
- **Interfaces:** Define minimal interfaces that express required behavior.
- **Error handling:** Wrap errors with context using `%w`.
- **Logging:** Use odc’s logging conventions; avoid noisy logs.
- **Tests:** Create or update table‑driven tests to validate new behavior.

## Editing existing Go code

- **Refactor safely:** Preserve behavior unless the request explicitly calls
  for a behavioral change.
- **Improve clarity:** Simplify logic, reduce duplication, and improve naming.
- **Align with standards:** Bring older code up to current odc patterns when
  touching it.
- **Consistency:** Ensure terminology and abstractions match the rest of the
  codebase.
- **Tests:** Update tests to reflect new behavior or improved structure.

---

# **Phase 4: Verification and finalization**

Perform a final quality check to ensure correctness, consistency, and
maintainability.

1. **Compile:** Ensure the code builds cleanly with `go build ./...`.
2. **Tests:** Ensure all tests pass with `go test ./...`.
3. **Static analysis:** Ensure `go vet` and any configured linters pass.
4. **Self‑review:** Re‑read changes for clarity, correctness, and architectural
   alignment.
5. **Format:** Once all changes are complete, ask to execute `go fmt ./...`
   or the project’s formatting script. If the user confirms, execute it.