# Go Cache Persistence Feature

This feature configures persistent volumes for Go package and build caches to speed up development and rebuilds.

## Mounted Volumes

- `go-pkg`: Mounted to `/go/pkg`
- `go-cache`: Mounted to `/home/vscode/.cache/go-build`

## Permissions

Note: On some systems, these volumes might be mounted with root permissions. The main `devcontainer.json` includes a `postCreateCommand` to ensure these directories are owned by the `vscode` user.
