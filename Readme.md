# ServiceNow SDK for Go

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/michaeldcanady/servicenow-sdk-go?style=plastic)
[![GoDoc](https://img.shields.io/static/v1?style=plastic&label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/michaeldcanady/servicenow-sdk-go)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/michaeldcanady/servicenow-sdk-go?style=plastic)
![GitHub issues](https://img.shields.io/github/issues/michaeldcanady/servicenow-sdk-go?style=plastic)
![GitHub](https://img.shields.io/github/license/michaeldcanady/servicenow-sdk-go?style=plastic)
[![Maintainability](https://qlty.sh/badges/e778f295-dfb1-4637-a15e-f179549fcae4/maintainability.svg)](https://qlty.sh/gh/michaeldcanady/projects/servicenow-sdk-go)
[![codecov](https://codecov.io/gh/michaeldcanady/servicenow-sdk-go/graph/badge.svg?token=MJPM1UAI78)](https://codecov.io/gh/michaeldcanady/servicenow-sdk-go)

A Service-Now API client enabling Go programs to interact with Service-Now in a simple and uniform way

![servicenow-sdk-go](.github/servicenow-sdk-go_logo.png)

## Supported Service-Now APIs

| API                                                                                        | Status | Issues                                                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------ | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Account                                                                                    | ✖️      |                                                                                                                                                                                                                                                     |
| ActivitySubscriptions                                                                      | ✖️      |                                                                                                                                                                                                                                                     |
| Agent Client Collector                                                                     | ✖️      |                                                                                                                                                                                                                                                     |
| Aggregate                                                                                  | ✖️      |                                                                                                                                                                                                                                                     |
| AI Search External User Mapping                                                            | ✖️      |                                                                                                                                                                                                                                                     |
| Alarm Management Open                                                                      | ✖️      |                                                                                                                                                                                                                                                     |
| Application Service                                                                        | ✖️      |                                                                                                                                                                                                                                                     |
| Appointment                                                                                | ✖️      |                                                                                                                                                                                                                                                     |
| [Attachment](https://github.com/michaeldcanady/servicenow-sdk-go/tree/main/attachment-api) | ✔️      | [![Attachment API Issues](https://img.shields.io/github/issues-raw/michaeldcanady/servicenow-sdk-go/attachment%20api?label=%20)](https://github.com/michaeldcanady/servicenow-sdk-go/labels/attachment%20api)                                       |
| Advanced Work Assignment (AWA) Agent                                                       | ✖️      |                                                                                                                                                                                                                                                     |
| AWA Assignment                                                                             | ✖️      |                                                                                                                                                                                                                                                     |
| AWA Routing                                                                                | ✖️      |                                                                                                                                                                                                                                                     |
| [Batch](https://github.com/michaeldcanady/servicenow-sdk-go/tree/main/batch-api)           | ✔️      |                                                                                                                                                                                                                                                     |
| Case                                                                                       | ✖️      |                                                                                                                                                                                                                                                     |
| Custom Chat Chatbot Interoperability Framework (CCCIF) Media Resource                      | ✖️      |                                                                                                                                                                                                                                                     |
| CdmApplicationsApi                                                                         | ✖️      |                                                                                                                                                                                                                                                     |
| CdmChangesetsApi                                                                           | ✖️      |                                                                                                                                                                                                                                                     |
| CdmEditorApi                                                                               | ✖️      |                                                                                                                                                                                                                                                     |
| CdmPoliciesApi                                                                             | ✖️      |                                                                                                                                                                                                                                                     |
| [Tables](https://github.com/michaeldcanady/servicenow-sdk-go/tree/main/table-api)          | ✔️      | [![Table API Issues](https://img.shields.io/github/issues-raw/michaeldcanady/servicenow-sdk-go/module%3A+table-api?label=%20)](https://github.com/michaeldcanady/servicenow-sdk-go/issues?q=is%3Aissue+is%3Aopen+label%3A%22module%3A+table-api%22) |
---

| Emoji | Meaning       |
| ----- | ------------- |
| ✔️     | Supported     |
| 🆕     | Preview       |
| ♻️     | In progress   |
| ✖️     | Not supported |

## Documentation

For detailed information on how to get started and use this SDK, please visit
the [official documentation](https://michaeldcanady.github.io/servicenow-sdk-go/).

## Contributing

We welcome contributions from the community! Whether you're fixing a bug, adding
a new feature, or improving documentation, please review our
[Contributor Guide](https://michaeldcanady.github.io/servicenow-sdk-go/contributing/)
to get started.
