# Development Setup

Thank you for your interest in contributing to the ServiceNow SDK for Go! This guide will help you set up your development environment.

## Prerequisites

- **Go**: Version 1.23 or higher is required.
- **Git**: For version control.
- **ServiceNow Instance**: A Personal Developer Instance (PDI) is recommended for testing. You can get one for free at [developer.servicenow.com](https://developer.servicenow.com/).

## Getting Started

1.  **Fork the repository** on GitHub.
2.  **Clone your fork** locally:
    ```bash
    git clone https://github.com/YOUR_USERNAME/servicenow-sdk-go.git
    cd servicenow-sdk-go
    ```
3.  **Install dependencies**:
    ```bash
    go mod download
    ```

## Development Workflow

- **Branching**: Create a new branch for each feature or bug fix.
  ```bash
  git checkout -b feat/my-new-feature
  ```
- **Linting**: We use `golangci-lint`. You can run it locally:
  ```bash
  golangci-lint run
  ```
- **Testing**: Ensure all tests pass before submitting a PR.
  ```bash
  go test ./...
  ```

## Using VS Code

If you use VS Code, we provide a `.devcontainer` configuration to get a consistent environment quickly. We also include recommended extensions and settings in `.vscode/`.
