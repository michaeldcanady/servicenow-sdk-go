# ServiceNow SDK for Go

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/michaeldcanady/servicenow-sdk-go?style=plastic)
![example workflow](https://github.com/michaeldcanady/servicenow-sdk-go/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://img.shields.io/static/v1?style=plastic&label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/michaeldcanady/servicenow-sdk-go)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/michaeldcanady/servicenow-sdk-go?style=plastic)
![GitHub issues](https://img.shields.io/github/issues/michaeldcanady/servicenow-sdk-go?style=plastic)
![GitHub](https://img.shields.io/github/license/michaeldcanady/servicenow-sdk-go?style=plastic)
![GitHub all releases](https://img.shields.io/github/downloads/michaeldcanady/servicenow-sdk-go/total?style=plastic)
[![Code Climate](https://codeclimate.com/github/michaeldcanady/servicenow-sdk-go.svg)](https://codeclimate.com/github/michaeldcanady/servicenow-sdk-go)

A Service-Now API client enabling Go programs to interact with Service-Now in a simple and uniform way

![servicenow-sdk-go](.github/servicenow-sdk-go_logo.png)

## Supported Service-Now APIs

| API     | Status |
| ------- | ------ |
| Tables  | ✔️      |
| Another | ♻️      |
| Another | ✖️      |
| Another | 🆕      |

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
    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

cred := servicenowsdkgo.NewUsernamePasswordCredential("username", "password")
```

### 2.2 Get a ServiceNow Client and Adapter object

```golang
import (
    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

client := servicenowsdkgo.NewClient(cred, "instance")
```

### 2.3 Get query parameters

```golang
import (
    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

params := &servicenowsdkgo.TableRequestBuilderGetQueryParameters{
		Limit: int32(1),
	}
```

### 2.4 Build request for table

```golang
import (
    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

records, err := client.Now().Table("table name").Get(params)
if err != nil {
    panic(err)
}
```
