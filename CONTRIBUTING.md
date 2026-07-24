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
3. Create a new branch for your changes, named `<type>/<kebab-description>` (see
   [Branching convention](#branching-convention) below): `git checkout -b fix/nil-pointer-in-tableapi`
4. Make your changes and commit them: `git commit -am 'Add some feature'`
   1. Include tests that cover your changes.
   2. Update the documentation to reflect your changes, where appropriate.
   3. Add and entry to the `changelog.md` file describing your changes if appropriate.
5. Push your changes to your fork: `git push origin fix/nil-pointer-in-tableapi`
6. Create a pull request from your fork to the main repository: `gh pr create` (With the GitHub CLI).
   The PR description must reference the issue it addresses (e.g. `Closes #123` or `Part of #123`) —
   see [Linking issues](#linking-issues) below.

## Branching convention

This repo practices trunk-based development with typed feature branches. Branch names must match:

```
<type>/<kebab-description>
```

where `<type>` is one of `fix`, `feat`, `chore`, `docs`, `style`, `refactor`, `perf`, `test` — the
same vocabulary enforced on PR titles — and `<kebab-description>` is a short, lowercase,
hyphen-separated summary (e.g. `feat/add-cdm-changeset-api`, `docs/publish-support-policy`).
`.github/workflows/branch-policy.yml` enforces this on every PR targeting `main` (automated
`release-please--*` and `dependabot/*` branches are exempt).

## Linking issues

Every PR must reference an issue in its description (e.g. `Closes #123`, `Fixes #123`, or
`Part of #123` for PRs that only address part of a larger issue).
`.github/workflows/branch-policy.yml` enforces this on every PR targeting `main`. If a PR is a
genuinely trivial chore with no associated issue, apply the `no-issue-required` label instead of
opening a throwaway issue.

## Reporting Bugs

If you find a bug in Servicenow-SDK-Go, please report it by opening a new issue in the issue tracker. Please include as much detail as possible, including steps to reproduce the bug and any relevant error messages.
