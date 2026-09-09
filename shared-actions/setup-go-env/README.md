# setup-go-env

Generic GitHub Action to checkout (optional) and install a Go toolchain.

## Usage

```yaml
- uses: NerdIT-Tech/.github/actions/setup-go-env@v1
  with:
    go-version: "stable"
    cache: true
    checkout: true
```

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `go-version` | Go version to install (e.g., `stable`, `oldstable`, `1.22.x`) | Yes | - |
| `cache` | Enable Go module/build caching | No | `true` |
| `checkout` | Checkout repository first | No | `true` |