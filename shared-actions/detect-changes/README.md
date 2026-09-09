# detect-changes

Generic GitHub Action to detect changed files matching given patterns.

## Usage

```yaml
- uses: NerdIT-Tech/.github/actions/detect-changes@v1
  id: changes
  with:
    files: |
      **.go
      **/*.md
    base_sha: ${{ github.event.pull_request.base.sha }}
    fetch-depth: 1
    checkout: true

- if: steps.changes.outputs.any_changed == 'true'
  run: echo "Files changed!"
```

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `files` | Newline-separated glob patterns to watch | No | All files |
| `base_sha` | Base SHA/ref to diff against | No | Auto-detect |
| `fetch-depth` | Git fetch depth for checkout | No | `1` |
| `checkout` | Whether to checkout first | No | `true` |

## Outputs

| Output | Description |
|--------|-------------|
| `any_changed` | `'true'` if any matching file changed |