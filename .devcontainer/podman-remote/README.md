# Podman Remote Client Feature

Installs the Podman CLI and configures it to talk to the host Podman socket. This allows running containers from within the devcontainer ("DooD" - Docker outside of Docker pattern).

## Configuration

- **Host Socket Path**: Configurable via the `hostSocketPath` option. Default: `/var/run/user/1000/podman/podman.sock`
- **Container Socket Path**: `/tmp/podman.sock`
- **Environment Variable**: `CONTAINER_HOST` is set to `unix:///tmp/podman.sock`
