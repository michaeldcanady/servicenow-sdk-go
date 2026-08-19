<!-- PR title must be a Conventional Commit line, e.g. `feat(tableapi): add display-value support` — CI lints it and it becomes the squash commit. -->

## What & why

<!-- Short description of the change and the problem it solves. Link related issues. -->

## Checklist

- [ ] `gofmt -s -w .`, `golangci-lint run ./...`, and `go test ./...` pass locally
- [ ] New or changed exported surface has unit tests
- [ ] **Docs:** the docs site (`website/`) is updated — or no exported surface changed / explain below why no docs change is needed
- [ ] Code samples live in `website/snippets/*.go` region markers (compiled by CI), not inline in pages
- [ ] `VERSION` and `CHANGELOG.md` are untouched (release-please manages them)

<!-- If docs are intentionally untouched, say why: -->
