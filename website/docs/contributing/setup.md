---
title: Development setup
description: >-
  Clone to green build in under ten minutes — one click with the dev
  container, or a short local-toolchain path.
---

# Development setup

Goal: from clone to a green build in under ten minutes, with the same
toolchain CI uses. There are two paths — the dev container gives you
everything preconfigured in one click; the local path is for people who
prefer their own environment.

## Prerequisites

- **Git**, and a GitHub account for your fork.
- **A ServiceNow instance** — only needed for integration/e2e work, not for
  building or unit tests. A free
  [Personal Developer Instance](https://developer.servicenow.com/) is enough.

## Path 1: Dev container (recommended)

The repository ships a dev-container definition with Go, `golangci-lint`,
`just`, and the editor extensions already installed — the environment CI runs
is the environment you develop in.

1.  Install [Docker](https://www.docker.com/get-started/) or
    [Podman](https://podman.io/), plus the
    [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
    for VS Code.
2.  **Fork and clone**:
    ```bash
    git clone https://github.com/YOUR_USERNAME/servicenow-sdk-go.git
    ```
3.  Open the folder in VS Code and click **"Reopen in Container"** when
    prompted. First launch builds the image; later launches are instant.

## Path 2: Local toolchain

- **Go** 1.25 or higher.
- **[golangci-lint](https://golangci-lint.run/usage/install/)** — the lint
  gate CI enforces (config: `.golangci.yml`).
- **[just](https://github.com/casey/just)** *(optional)* — a command runner
  wrapping the common tasks (`just build`, `just lint`, `just fmt`).

```bash
git clone https://github.com/YOUR_USERNAME/servicenow-sdk-go.git
cd servicenow-sdk-go
go mod download
```

## The moment of success

Three commands, three green results — after this you're ready to make
changes:

```bash
go build ./...            # compiles every package
go test ./...             # unit tests (fast, no network, no instance)
golangci-lint run ./...   # the same lint gate as CI
```

If all three pass on a fresh clone, your environment is correct. If you'll be
editing docs too, `just setup-docs` then `just serve-docs` runs this site
locally (Node 20+).

## Credentials for integration and e2e tests

Unit tests never need credentials. The integration suite mocks HTTP too — but
its configuration, and the e2e suite's real connections, read a `.env` file
in the project root (git ignores it):

```env
SN_INSTANCE=your_instance_name
SN_USERNAME=your_username
SN_PASSWORD=your_password
```

The [testing guide](testing.md) covers when each suite runs and how to invoke
them.

## Next steps

Environment ready — where next depends on what you came to do:

- **Shipping your first contribution?** Go straight to
  [Your first PR](first-pr.mdx) — it picks up exactly here, with the timer
  already running.
- **Building features?** Read the [architecture](architecture.md) page to
  follow one request through the codebase before you start changing it.
