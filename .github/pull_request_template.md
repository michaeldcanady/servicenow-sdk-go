<!--
  PR title MUST be a Conventional Commit line — CI lints it and it becomes the squash commit.
  Examples:
    feat(tableapi): add display-value support
    fix(credentials): handle expired refresh token
    chore(ci): update golangci-lint version
    docs: fix broken link in tableapi README
-->

## What & why

<!-- Short description of the change and the problem it solves. Link related issues with "Closes #123". -->

## Type of change

<!-- Check one. This determines which checklist items apply. -->

- [ ] `feat` — new or changed exported API surface
- [ ] `fix` — bug fix
- [ ] `refactor` — internal improvement, no behavior change
- [ ] `docs` — documentation only
- [ ] `chore` — CI, tooling, dependencies, or other non-release work
- [ ] `test` — test coverage or infrastructure

## Checklist

<!-- Only check items relevant to your change type. -->

- [ ] `gofmt -s -w .` and `golangci-lint run ./...` pass
- [ ] `go test ./...` passes
- [ ] New or changed exported surface has unit tests
- [ ] `website/` is updated (or explain why not below)
- [ ] Code samples use `website/snippets/*.go` region markers, not inline in pages
- [ ] `VERSION` and `CHANGELOG.md` are untouched (release-please manages them)

<!-- If docs or tests are intentionally untouched, say why: -->
