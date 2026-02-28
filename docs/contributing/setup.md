# Development Setup

This guide walks you through setting up your environment to contribute to the
ServiceNow SDK for Go. You can choose between a fully automated setup using
VS Code Dev Containers or a manual local setup.

## Prerequisites

Regardless of your setup choice, you need the following:

- **Git:** For version control and cloning the repository.
- **ServiceNow Instance:** A Personal Developer Instance (PDI) is essential for
  running integration tests. You can obtain one at no cost at
  [developer.servicenow.com](https://developer.servicenow.com/).

## Option 1: VS Code Dev Containers (Recommended)

The easiest way to get started is using VS Code Dev Containers. This method
automatically configures a consistent environment with all necessary tools
(Go, linters, and extensions).

1.  **Install [Docker](https://www.docker.com/get-started/)** or
    **[Podman](https://podman.io/)**.
2.  **Install the [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)**
    in VS Code.
3.  **Fork and clone** the repository:
    ```bash
    git clone https://github.com/YOUR_USERNAME/servicenow-sdk-go.git
    ```
4.  **Open the folder** in VS Code.
5.  **Click "Reopen in Container"** when the notification appears in the bottom
    right corner.

## Option 2: Local development

If you prefer to manage your own environment, ensure you meet these additional
requirements:

- **Go:** Version 1.23 or higher.
- **golangci-lint:** For running code quality checks.
- **just:** (Optional) A command runner used for some development tasks.

### Configure your environment

1.  **Fork and clone** the repository:
    ```bash
    git clone https://github.com/YOUR_USERNAME/servicenow-sdk-go.git
    cd servicenow-sdk-go
    ```
2.  **Download dependencies:**
    ```bash
    go mod download
    ```
3.  **Install the linter:**
    Follow the installation instructions at
    [golangci-lint.run](https://golangci-lint.run/usage/install/).

## Configure environment variables

To run integration tests, you must configure your ServiceNow instance
credentials. Create a `.env` file in the project root (this file is ignored by
git):

```env
SN_INSTANCE=your_instance_name
SN_USERNAME=your_username
SN_PASSWORD=your_password
```

## Verify your setup

Run the unit tests to ensure your environment is working correctly:

```bash
go test ./...
```

## Next steps

Now that your environment is ready, learn about the
[project architecture](architecture.md) to understand how the SDK is structured.
