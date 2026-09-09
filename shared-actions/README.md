# Shared GitHub Actions

Reusable composite actions for Go projects, exportable to `NerdIT-Tech/.github/actions/`.

## Actions

| Action | Description | Dependencies |
|--------|-------------|--------------|
| [`detect-changes`](detect-changes/) | Detect changed files matching patterns | `tj-actions/changed-files` |
| [`setup-go-env`](setup-go-env/) | Checkout + install Go toolchain | `actions/checkout`, `actions/setup-go` |
| [`install-tparse`](install-tparse/) | Install tparse (Go test JSON parser) | `actions/setup-go` (prereq) |
| [`go-test`](go-test/) | Run `go test` with tparse, coverage | `install-tparse` (must run first) |
| [`check-go-deps`](check-go-deps/) | Verify & tidy Go module deps | `actions/setup-go` |

## Installation

Copy the `shared-actions/` directory contents to `NerdIT-Tech/.github/actions/`:

```
NerdIT-Tech/.github/
└── actions/
    ├── detect-changes/
    │   ├── action.yml
    │   └── README.md
    ├── setup-go-env/
    │   ├── action.yml
    │   └── README.md
    ├── install-tparse/
    │   ├── action.yml
    │   └── README.md
    ├── go-test/
    │   ├── action.yml
    │   └── README.md
    └── check-go-deps/
        ├── action.yml
        └── README.md
```

## Usage in Workflows

```yaml
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: NerdIT-Tech/.github/actions/setup-go-env@v1
        with:
          go-version: stable
      - uses: NerdIT-Tech/.github/actions/install-tparse@v1
      - uses: NerdIT-Tech/.github/actions/go-test@v1
        with:
          package-paths: ./...
```

## Versioning

Tag releases as `v1`, `v2`, etc. in the target repo. Consumers pin to major versions:
```yaml
uses: NerdIT-Tech/.github/actions/go-test@v1
```