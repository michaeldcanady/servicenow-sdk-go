---
title: Development setup
description: >-
  Clone to green build in under ten minutes — one click with the dev
  container, or a short local-toolchain path.
---

# Development setup

Goal: from clone to a green build in under ten minutes, with the same
toolchain CI uses. You have two paths — the dev container gives you
everything preconfigured in one click; the local path is for people who
prefer their own environment.

## Prerequisites

- **Git**, and a GitHub account for your fork.
- **A ServiceNow instance** — only needed for integration/e2e work, not for
  building or unit tests. A free
  [Personal Developer Instance](https://developer.servicenow.com/) is enough.

## Path 1: Dev container (recommended)

The repository ships a dev-container definition with
[mise](https://mise.jdx.dev/) and the editor extensions already installed.
When the container is created, it runs `mise install`, which installs Go,
`golangci-lint`, Vale, and every other tool at the versions that CI uses, and
it activates mise in bash.

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

- **[mise](https://mise.jdx.dev/getting-started.html)**: installs the tools
  that `mise.toml` pins, at the versions that CI uses, and runs the repository
  tasks. Follow the mise guide to activate it in your shell.

```bash
git clone https://github.com/YOUR_USERNAME/servicenow-sdk-go.git
cd servicenow-sdk-go
mise trust
mise install
go mod download
```

If you don't want to use mise, install Go 1.25 or later and
[golangci-lint](https://golangci-lint.run/usage/install/). The plain `go` and
`golangci-lint` commands on this page work without mise.

## The moment of success

Three commands, three green results — after this you're ready to make
changes:

```bash
go build ./...            # compiles every package
go test ./...             # unit tests (fast, no network, no instance)
golangci-lint run ./...   # the same lint gate as CI
```

If all three pass on a fresh clone, your environment is correct. If you're
editing docs too, run `mise run docs:setup` and then `mise run docs:serve` to
serve this site locally (Node 20 or later).

## Run the pull request checks before you push

The repository tasks live in `mise.toml` and `.mise/tasks/`. To list them, run
`mise tasks`. These two tasks run the checks as CI does, so you can fix a
failure before you push:

- `mise run check` runs every pull request check that works locally: build,
  lint, unit tests on both Go releases that CI tests, mocked integration tests,
  the Go dependency and vulnerability checks, the docs snippet check, Vale on
  your changed Markdown files, the license-header check, the CI configuration
  linters, and the branch-name and pull request title rules.
- `mise run check:all` also runs the checks that run after merge: Vale on the
  whole docs site, and the docs build.

Some checks run only on GitHub, because they need GitHub or a service account:
CodeQL, OpenSSF Scorecard, the CodeRabbit review, GitGuardian, and the
linked-issue check. The e2e tests need a live ServiceNow instance. To run them
against your own instance, set up the credentials in the next section, and
then run `mise run go:test:e2e`.

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
