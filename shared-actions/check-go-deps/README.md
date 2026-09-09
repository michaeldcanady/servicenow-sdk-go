# check-go-deps

Generic GitHub Action to verify Go module dependencies (`go mod verify`) and check tidiness (`go mod tidy`).

## Usage

```yaml
- uses: NerdIT-Tech/.github/actions/check-go-deps@v1
  with:
    mod-root: "."
    go-version: "stable"
```

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `mod-root` | Root directory containing go.mod/go.sum | No | Repository root |
| `go-version` | Go version to use (e.g., `stable`, `1.22.x`) | No | From go.mod |

## What it does

1. **Verifies dependencies** - Runs `go mod verify` to check checksums
2. **Checks tidiness** - Runs `go mod tidy` and fails if go.mod/go.sum change

## Example: Multi-module repo

```yaml
- uses: NerdIT-Tech/.github/actions/check-go-deps@v1
  with:
    mod-root: "submodule"
    go-version: "stable"
```