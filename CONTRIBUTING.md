# Contributing to the ServiceNow SDK for Go

Thanks for your interest in contributing! To keep guidance in one place, the
full contributor documentation lives on the docs site:

- **[Your first PR](https://michaeldcanady.github.io/servicenow-sdk-go/contributing/first-pr)** —
  new here? A time-boxed walkthrough from fork to open PR in about an hour.
- **[Contributor guide](https://michaeldcanady.github.io/servicenow-sdk-go/contributing/)** —
  workflow, branch naming, Conventional Commits, and PR expectations.
- **[Development setup](https://michaeldcanady.github.io/servicenow-sdk-go/contributing/setup)** —
  dev container or local environment.
- **[Architecture](https://michaeldcanady.github.io/servicenow-sdk-go/contributing/architecture)** —
  how the SDK is put together.
- **[Testing guide](https://michaeldcanady.github.io/servicenow-sdk-go/contributing/testing)** —
  unit, integration, and e2e suites.
- **[Conventions reference](https://michaeldcanady.github.io/servicenow-sdk-go/contributing/conventions)** —
  the dense field guide reviewers hold PRs to.

The two rules most worth knowing before your first PR:

1. Commits and PR titles follow
   [Conventional Commits](https://www.conventionalcommits.org/) — CI enforces
   this, and `release-please` generates `VERSION` and `CHANGELOG.md` from it
   (never edit those files by hand).
2. Changes to exported API surface must update the docs site (`website/`) or
   explain why not.

Found a bug? [Open an issue](https://github.com/michaeldcanady/servicenow-sdk-go/issues)
with reproduction steps and any error output.
