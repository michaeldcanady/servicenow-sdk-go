# go-test

Generic GitHub Action to run `go test` with JSON output via tparse, coverage, and configurable options.

## Usage

```yaml
- uses: NerdIT-Tech/.github/actions/install-tparse@v1
  with:
    tparse-version: "v0.18.0"

- uses: NerdIT-Tech/.github/actions/go-test@v1
  with:
    format-json: "true"
    verbose: "true"
    coverage-profile: "coverage.out"
    build-flags: "-tags=integration"
    package-paths: "./..."
    working-directory: "."
```

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `format-json` | Run with `-json` and pipe through tparse | No | `true` |
| `verbose` | Run with `-v` | No | `true` |
| `coverage-profile` | Coverage output file (empty to skip) | No | `""` |
| `build-flags` | Additional flags for `go build` | No | `""` |
| `package-paths` | Packages to test (space/newline separated) | Yes | - |
| `working-directory` | Directory to run tests in | No | `.` |

## Outputs

| Output | Description |
|--------|-------------|
| `coverage-profile` | Path to generated coverage profile |

## Prerequisites

Requires `tparse` on PATH. Use the `install-tparse` action first.