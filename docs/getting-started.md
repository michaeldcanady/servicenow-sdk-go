# Getting started

## Requirements

- Installed [golang 1.23+](https://go.dev/doc/install).
- Active Service-Now instance or [developer instance](https://developer.servicenow.com/dev.do).

## 1. Install Service-Now SDK for Go

=== "Install the latest"

    ``` bash
    go get github.com/michaeldcanady/servicenow-sdk-go
    ```
=== "Install a specific version"
    
    ``` bash
    go get github.com/michaeldcanady/servicenow-sdk-go@{version}
    ```

## 2. Create a credential

```golang
package main

import (
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
    // instantiates a basic authentication username/password credential but you can use any from the `credentials` submodule or implement your own!
    cred := credentials.NewUsernamePasswordCredential("{username}", "{password}")
    ...
}
```

## 3. Create a Service-Now client

```golang
    ...
    client := servicenowsdkgo.NewServiceNowClient(cred, "{instance}")
    ...
```

## 4. Review specific api documentation

Once you have your `client` object, you're good to get started implementing all that the Service-Now SDK for Go has to offer! See [apis](/apis) for information on implementing a specific api!
