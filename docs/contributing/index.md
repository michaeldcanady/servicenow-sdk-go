# Contributor Guide

Thank you for your interest in contributing to the ServiceNow SDK for Go! We
welcome contributions from developers of all skill levels. This guide helps you
get started quickly and ensures your contributions align with the project's
standards.

Whether you're fixing a bug, improving documentation, or adding support for a
new ServiceNow API, your efforts help make this SDK better for everyone.

## Ways to contribute

You can help the project in several ways:

- **Code:** Fix bugs, implement new API modules, or optimize existing logic.
- **Documentation:** Correct typos, clarify instructions, or add new examples.
- **Testing:** Increase test coverage or add complex integration test cases.
- **Community:** Report issues, suggest features, or help others in the issue
  tracker.

## Getting started roadmap

We recommend following this path to start contributing:

1.  **Read the [Architecture](architecture.md) guide:** Understand the SDK's
    structure and how it uses the Microsoft Kiota framework.
2.  **Configure your [Development Setup](setup.md):** Set up your local
    environment and ensure you can run the project.
3.  **Review the [Testing Guide](testing.md):** Learn how to write and run both
    unit and integration tests.
4.  **Find an issue:** Look for "good first issue" or "help wanted" labels on
    the [GitHub issue tracker](https://github.com/michaeldcanady/servicenow-sdk-go/issues).

## Our philosophy

To maintain a high-quality codebase, we follow these core principles:

- **Simplicity:** We prefer clear, maintainable code over complex optimizations.
- **Consistency:** We strictly follow standard Go conventions and established
  Kiota patterns.
- **Reliability:** Every new feature or bug fix must include comprehensive unit
  tests.
- **Documentation:** We keep documentation in sync with code changes to ensure
  users always have accurate information.

## Submitting your changes

Once you've implemented your changes and verified them with tests, follow these
steps to submit your contribution:

1.  **Run the linter:** Ensure your code follows the project's quality
    standards by running `golangci-lint run`.
2.  **Commit your changes:** Use clear, descriptive commit messages.
3.  **Push to your fork:** Push your branch to your GitHub fork.
4.  **Open a Pull Request:** Create a PR against the `main` branch of the
    official repository. Provide a detailed description of your changes and
    link any relevant issues.

We look forward to your contributions!
