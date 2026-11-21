---
title: 'ADR 026: mise.toml is the source of tool versions, and CI keeps native actions with a drift check'
description: >-
  CI keeps its native GitHub Actions. Actions read a shared version file where
  they can, and a drift check fails the PR when any other CI pin differs from
  mise.toml.
---

# ADR 026: mise.toml is the source of tool versions, and CI keeps native actions with a drift check

## Status

Proposed (2026-11-21). The plan is in
`specs/010-holistic-delivery-pipeline/spec.md`.

## Context

`mise.toml` pins the tools that `mise run check` uses locally. CI pins the
same tools again in workflow inputs, and a comment next to each `mise.toml`
pin names the CI location to keep in step. On `main`, these pins are
duplicated:

- golangci-lint 2.14.0 (`reusable-lint-go.yaml` and `release-verify.yml`)
- Vale 3.18.0 (`reusable-vale-lint.yaml`, `reusable-docs-gate.yaml`, and
  `docs-preview.yml`)
- zizmor 1.29.0 (`reusable-zizmor.yaml`)
- Node 20 (`.github/actions/docs-build` and `docs-archive.yml`)

Nothing checks the comments, so a pin can drift and the local check then
disagrees with CI.

Go is different. CI runs a matrix of the `stable` and `oldstable` aliases on
purpose, and the `go` pin in `mise.toml` tracks `stable`. `govulncheck` and
`tparse` install `latest` in CI but are pinned in `mise.toml`.

Alternatives considered:

1. **CI runs `mise` tasks** through `jdx/mise-action`. This gives one set of
   pins, and a local run matches CI by construction. Rejected:
   - Every job installs `mise` first.
   - CI loses the PR annotations that native actions add, such as
     golangci-lint's.
   - Dependabot keeps action pins current today, and moving tools into
     `mise.toml` would need a dependency bot that updates that file.
2. **Keep today's comments.** Rejected: nothing enforces them.
3. **Native actions, shared version files where possible, and a drift check
   for the rest** (chosen).

## Decision

1. **`mise.toml` is the one source of tool versions.** When the version of a
   tool changes, change `mise.toml` first.
2. **CI keeps native actions.** Dependabot keeps their commit SHAs current.
   The drift check doesn't cover action SHAs.
3. **Read a shared version file where an action supports it.** For example,
   `setup-go` reads `go-version-file: go.mod`, and `setup-node` reads
   `node-version-file`. Use this only when `mise` reads the same file, so
   the version still has one source.
4. **A drift check covers the remaining duplicated pins.** A script, run as
   a `mise` task and as a `PR Gate` job when `mise.toml` or `.github/**`
   changes, fails when a workflow's version input differs from `mise.toml`.
   It covers golangci-lint, Vale, zizmor, and Node until Node moves to a
   version file.
5. **The CI Go matrix stays on the `stable` and `oldstable` aliases.** It is
   outside the drift check.
6. **`govulncheck` and `tparse` either get pinned in CI and join the drift
   check, or carry a comment in `mise.toml` saying that CI tracks `latest` on
   purpose.** The spec's open question 3 settles which.

## Consequences

- **Pros:**
  - Local and CI tool versions can't drift without a failing PR.
  - CI keeps action annotations and Dependabot updates.
- **Cons:**
  - Dependabot doesn't update version inputs such as `with: version:`, so a
    tool bump stays a manual edit to `mise.toml` plus the workflows that the
    check flags.
  - The drift check is one more script to maintain.
- **Rule for future CI work:** a new tool that CI and `mise.toml` both pin
  either reads a shared version file or joins the drift check in the same PR.
