# ServiceNow SDK for Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/michaeldcanady/servicenow-sdk-go)](https://pkg.go.dev/github.com/michaeldcanady/servicenow-sdk-go)

## 1. Installation

```Shell
go get github.com/michaeldcanady/servicenow-sdk-go
```

## 2. Getting Started

### 2.1 Create an AuthenticationProvider object

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
