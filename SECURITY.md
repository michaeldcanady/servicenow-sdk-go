# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 2.x     | :white_check_mark: |
| 1.x     | Critical fixes only |

## Reporting a Vulnerability

Please **do not** open a public issue for security vulnerabilities.

Instead, report it privately via
[GitHub's private vulnerability reporting](https://github.com/michaeldcanady/servicenow-sdk-go/security/advisories/new)
for this repository.

Include, where possible:

- A description of the issue and its impact
- Steps or a proof-of-concept to reproduce it
- Affected package(s) and version(s)

You can expect an acknowledgement within 7 days. Once a fix is available, the
vulnerability will be disclosed via a GitHub security advisory and, where
applicable, submitted to the [Go vulnerability database](https://vuln.go.dev/).

## Scope notes

This SDK handles ServiceNow credentials (basic auth and OAuth2 flows under
`credentials/`). Issues involving credential leakage, token handling, TLS
behavior, or request forgery are especially in scope.
