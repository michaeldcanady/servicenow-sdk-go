# Getting started

## Requirements

- Installed [golang 1.23+](https://go.dev/doc/install).
- Active Service-Now instance or [developer instance](https://developer.servicenow.com/dev.do).

<!-- vale Microsoft.Headings = NO -->
<!-- vale Microsoft.Headings = NO -->
## 1. Install Service-Now SDK for Go
<!-- vale Microsoft.Headings = YES -->

=== "Install the latest"

    ``` bash
    go get github.com/michaeldcanady/servicenow-sdk-go@latest
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

With the `client` object initialized, implementation of the full capabilities offered by the Service-Now SDK for Go can begin. Refer to [apis](/apis/index.md) for details on implementing specific APIs.
