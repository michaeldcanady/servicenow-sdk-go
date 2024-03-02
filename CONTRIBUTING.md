# Contributing to ServiceNow SDK Go

Thanks for your interest in contributing to ServiceNow SDK Go! We welcome contributions from everyone, regardless of skill level or experience. Here are some guidelines to help you get started:

## Getting Started

To get started, you'll need to have the following tools installed:

- [Golang v1.21+](https://go.dev/doc/install)

## Recommended tools

- [Visual Studio Code (VS Code)](https://code.visualstudio.com/)

## Running the tests

    ```bash
    go test ./...
    ```

## Contributing Code

1. Fork the repository
2. Clone it to your local machine: `git clone {url}`
3. Create a new branch for your changes: `git checkout -b my-new-feature`
4. Make your changes and commit them: `git commit -am 'Add some feature'`
   1. Include tests that cover your changes.
   2. Update the documentation to reflect your changes, where appropriate.
   3. Add and entry to the `changelog.md` file describing your changes if appropriate.
5. Push your changes to your fork: `git push origin my-new-feature`
6. Create a pull request from your fork to the main repository: `gh pr create` (With the GitHub CLI)

## Reporting Bugs

If you find a bug in Servicenow-SDK-Go, please report it by opening a new issue in the issue tracker. Please include as much detail as possible, including steps to reproduce the bug and any relevant error messages.
