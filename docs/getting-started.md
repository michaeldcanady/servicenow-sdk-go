# Getting started
## 1. Installation

Install latest
```Shell
go get github.com/michaeldcanady/servicenow-sdk-go
```
or

Install specific version
```Shell
go get github.com/michaeldcanady/servicenow-sdk-go@version
```

## 2. Getting started

### 2.1 Create a credential

```golang
package main

import (
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
    // instantiates a basic authentication username/password credential but you can use any from the `credentials` submodule or implement your own!
    cred := credentials.NewUsernamePasswordCredential("username", "password")
    ...
```

### 2.2 Create a Service-Now client

```golang
    ...
    client := servicenowsdkgo.NewServiceNowClient(cred, "instance")
    ...
```

## 3. Review API documentation

Once you have your `client` object, you're good to get started implementing all that the Service-Now SDK for Go has to offer! See [apis](/apis) for information on implementing a specific API!
