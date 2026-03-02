# Contributor guide

Thank you for your interest in contributing to the ServiceNow SDK for Go! The
project welcomes contributions from developers of all skill levels. This guide
helps you get started and ensures your contributions align with the project's
standards.

Whether you're fixing a bug, improving documentation, or adding support for a
new ServiceNow API, your efforts help make this SDK better for everyone.

## Ways to contribute

You can help the project in many ways:

- **Code:** Fix bugs, implement new API modules, or optimize existing logic.
- **Documentation:** Correct typos, clarify instructions, or add new examples.
- **Testing:** Increase test coverage or add complex integration test cases.
- **Community:** Report issues, suggest features, or help others in the issue
  tracker.

## Getting started roadmap

Follow this path to start contributing:

1.  **Read the [Architecture](architecture.md) guide:** Understand the SDK's
    structure and how it uses the Microsoft Kiota framework.
2.  **Configure your [Development Setup](setup.md):** Set up your local
    environment and make sure you can run the project.
3.  **Review the [Testing Guide](testing.md):** Learn how to write and run both
    unit and integration tests.
4.  **Find an issue:** Look for "good first issue" or "help wanted" labels on
    the [GitHub issue tracker](https://github.com/michaeldcanady/servicenow-sdk-go/issues).

## Philosophy

To maintain a high-quality codebase, this project follows these core principles:

- **Simplicity:** The project prefers clear, maintainable code over complex
  optimizations.
- **Consistency:** The project follows standard Go conventions and established
  Kiota patterns.
- **Reliability:** Every new feature or bug fix must include comprehensive unit
  tests.
- **Documentation:** Documentation stays in sync with code changes to make sure
  users always have accurate information.

## Submitting your changes

Once you've implemented your changes and verified them with tests, follow these
steps to submit your contribution:

1.  **Run the linter:** Make sure your code follows the project's quality
    standards by running `golangci-lint run`.
2.  **Commit your changes:** Use clear, descriptive commit messages.
3.  **Push to your fork:** Push your branch to your GitHub fork.
4.  **Open a Pull Request:** Create a PR for the `main` branch of the
    official repository. Provide a detailed description of your changes and
    link any relevant issues.

The project looks forward to your contributions!
