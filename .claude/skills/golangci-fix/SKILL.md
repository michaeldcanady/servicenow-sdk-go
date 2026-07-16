---
name: golangci-fix
description: Run golangci-lint across the repo (or a given package), group findings by linter, and fix them in the project's established idiom. Use for cleanup passes after a refactor or before opening a PR.
---

# golangci-fix

Runs this repo's configured lint suite and fixes findings without introducing
unrelated refactors.

## When to use

- After a refactor touching multiple packages (e.g. an accessor/mutator pass
  across `tableapi`, `policyapi`, etc.)
- Before opening a PR, to catch what CI's lint job (`.github/workflows/ci.yml`)
  would flag
- When the user asks to "fix lint errors" or "clean up lint warnings"

## Steps

1. Determine scope: default to `./...`; if the user names a package or the
   conversation has been focused on specific directories, scope to those
   (`golangci-lint run ./tableapi/...`).

2. Run the linter and capture output:
   ```
   golangci-lint run <scope> --output.text.path stdout
   ```

3. **Group findings by linter** (this repo enables: `dogsled`, `dupl`, `errcheck`,
   `gocognit`, `gocyclo`, `gosec`, `misspell`, `nakedret`, `staticcheck`,
   `unconvert`, `unparam`, `whitespace`, plus `gofmt` as a formatter). Note that
   `_test.go` files are already exempted from `dupl`, `errcheck`, `gocyclo`,
   `gosec`, `typecheck` per `.golangci.yml` — don't "fix" findings there that
   the config already excludes; if they show up, treat it as a config drift
   signal, not a real fix target.

4. Fix findings **without changing behavior**:
   - `errcheck` — handle or explicitly discard (`_ = ...`) with a reason if
     truly ignorable; never blanket-suppress.
   - `gocognit`/`gocyclo` — extract helper functions instead of restructuring
     control flow into something less idiomatic for the codebase.
   - `gosec` — fix the actual issue (e.g. unchecked file perms, weak rand);
     if it's a false positive, use a `//nolint:gosec // reason` comment
     sparingly, matching any existing `//nolint` usage patterns in the repo.
   - `dupl` — note that `(actsub-api|documents-api|table-api)/.*` paths are
     already excluded from `dupl` in `.golangci.yml`; don't force
     deduplication there just because two files look similar — that's
     accepted duplication in this codebase.
   - `unparam`/`unconvert`/`whitespace`/`misspell`/`nakedret` — straightforward
     mechanical fixes.

5. Re-run `gofmt -s -w <scope>` and `golangci-lint run <scope>` to confirm
   the findings are resolved and nothing new was introduced.

6. Run `go build ./...` and `go test <scope>` to confirm no regressions.

## Notes

- Do not disable linters or widen `.golangci.yml` exclusions to make findings
  disappear — fix the code, not the config, unless the user explicitly asks
  for a config change.
- Keep fixes surgical per this repo's engineering standard ("Surgical Changes"
  in `GEMINI.md`) — don't bundle drive-by refactors into a lint-fix pass.
