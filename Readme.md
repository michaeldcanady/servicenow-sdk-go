# ServiceNow SDK for Go

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/michaeldcanady/servicenow-sdk-go?style=plastic)
[![GoDoc](https://img.shields.io/static/v1?style=plastic&label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/michaeldcanady/servicenow-sdk-go)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/michaeldcanady/servicenow-sdk-go?style=plastic)
![GitHub issues](https://img.shields.io/github/issues/michaeldcanady/servicenow-sdk-go?style=plastic)
![GitHub](https://img.shields.io/github/license/michaeldcanady/servicenow-sdk-go?style=plastic)
[![Code Climate](https://codeclimate.com/github/michaeldcanady/servicenow-sdk-go.svg)](https://codeclimate.com/github/michaeldcanady/servicenow-sdk-go)
[![codecov](https://codecov.io/gh/michaeldcanady/servicenow-sdk-go/graph/badge.svg?token=MJPM1UAI78)](https://codecov.io/gh/michaeldcanady/servicenow-sdk-go)

A Service-Now API client enabling Go programs to interact with Service-Now in a simple and uniform way

![servicenow-sdk-go](.github/servicenow-sdk-go_logo.png)

## Supported Service-Now APIs

| API                                                                                        | Status | Issues                                                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------ | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Account                                                                                    | ✖️      |                                                                                                                                                                                                               |
| ActivitySubscriptions                                                                      | ✖️      |                                                                                                                                                                                                               |
| Agent Client Collector                                                                     | ✖️      |                                                                                                                                                                                                               |
| Aggregate                                                                                  | ✖️      |                                                                                                                                                                                                               |
| AI Search External User Mapping                                                            | ✖️      |                                                                                                                                                                                                               |
| Alarm Management Open                                                                      | ✖️      |                                                                                                                                                                                                               |
| Application Service                                                                        | ✖️      |                                                                                                                                                                                                               |
| Appointment                                                                                | ✖️      |                                                                                                                                                                                                               |
| [Attachment](https://github.com/michaeldcanady/servicenow-sdk-go/tree/main/attachment-api) | ♻️      | [![Attachment API Issues](https://img.shields.io/github/issues-raw/michaeldcanady/servicenow-sdk-go/attachment%20api?label=%20)](https://github.com/michaeldcanady/servicenow-sdk-go/labels/attachment%20api) |
| Advanced Work Assignment (AWA) Agent                                                       | ✖️      |                                                                                                                                                                                                               |
| AWA Assignment                                                                             | ✖️      |                                                                                                                                                                                                               |
| AWA Routing                                                                                | ✖️      |                                                                                                                                                                                                               |
| Batch                                                                                      | ✖️      |                                                                                                                                                                                                               |
| Case                                                                                       | ✖️      |                                                                                                                                                                                                               |
| Custom Chat Chatbot Interoperability Framework (CCCIF) Media Resource                      | ✖️      |                                                                                                                                                                                                               |
| CdmApplicationsApi                                                                         | ✖️      |                                                                                                                                                                                                               |
| CdmChangesetsApi                                                                           | ✖️      |                                                                                                                                                                                                               |
| CdmEditorApi                                                                               | ✖️      |                                                                                                                                                                                                               |
| CdmPoliciesApi                                                                             | ✖️      |                                                                                                                                                                                                               |
| [Tables](https://github.com/michaeldcanady/servicenow-sdk-go/tree/main/table-api)          | ✔️      | [![Table API Issues](https://img.shields.io/github/issues-raw/michaeldcanady/servicenow-sdk-go/table%20api?label=%20)](https://github.com/michaeldcanady/servicenow-sdk-go/labels/table%20api)                |
---

| Emoji | Meaning       |
| ----- | ------------- |
| ✔️     | Supported     |
| 🆕     | Preview       |
| ♻️     | In progress   |
| ✖️     | Not supported |

## 1. Installation

```Shell
go get github.com/michaeldcanady/servicenow-sdk-go
```

## 2. Getting Started

### 2.1 Create an AuthenticationProvider object

Create a credential object.

```golang
import (
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

cred := credentials.NewUsernamePasswordCredential("username", "password")
```

### 2.2 Get a ServiceNow Client and Adapter object

Create a client object.

```golang
import (
    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

client := servicenowsdkgo.NewServiceNowClient(cred, "instance")
```

### 2.3 Get query parameters

Build query parameter object.

```golang
import (
    "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

params := &tableapi.TableRequestBuilderGetQueryParameters{
        Limit: int32(1),
    }
```

### 2.4 Build request for table

Formulate request and provide parameters.

```golang

records, err := client.Now().Table("table name").Get(params)
if err != nil {
    panic(err)
}
```
