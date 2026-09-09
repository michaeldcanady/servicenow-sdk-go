# install-tparse

Generic GitHub Action to install tparse (Go test output parser) and add it to PATH.

## Usage

```yaml
- uses: NerdIT-Tech/.github/actions/install-tparse@v1
  with:
    tparse-version: "v0.18.0"
    install-dir: "$HOME/go/bin"
    use-sudo: "false"
```

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `tparse-version` | tparse version (e.g., `v0.18.0`, `latest`) | No | `latest` |
| `install-dir` | Installation directory (must be on PATH) | No | `$HOME/go/bin` |
| `use-sudo` | Use sudo for installation | No | `false` |

## Notes

- Defaults to `$HOME/go/bin` which is on PATH in GitHub-hosted runners
- No longer writes to `GITHUB_PATH` (zizmor compliant)