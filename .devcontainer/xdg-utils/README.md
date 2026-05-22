# XDG Utils Feature

Installs `xdg-utils` and configures a mount for the host browser pipe. This allows tools in the container to open the browser on the host.

## Configuration

- **Host Browser Pipe Path**: Configurable via the `hostBrowserPipePath` option. Default: `/home/michael/hostbrowserpipe`
- **Container Pipe Path**: `/tmp/hostbrowserpipe`
