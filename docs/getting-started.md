# Getting started

## Requirements

- Installed [golang 1.23+](https://go.dev/doc/install).
- Active Service-Now instance or [developer instance](https://developer.servicenow.com/dev.do).

<!-- vale Microsoft.Headings = NO -->
## 1. Install Service-Now SDK for Go
<!-- vale Microsoft.Headings = YES -->

=== "Install the latest"

    ``` bash
    go get github.com/michaeldcanady/servicenow-sdk-go@latest
    ```
=== "Install a specific version"

    ``` bash
    go get github.com/michaeldcanady/servicenow-sdk-go@xSDK_VERSIONx
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
    cred := credentials.NewUsernamePasswordCredential("xSDK_USERNAMEx", "xSDK_PASSWORDx")
    ...
}
```

## 3. Create a Service-Now client

=== "Specify instance"
    ```golang
        client := servicenowsdkgo.NewServiceNowClient(cred, "xSDK_SN_INSTANCEx.service-now.com")
    ```
<!-- vale Microsoft.Headings = NO -->
=== "Specify url"
<!-- vale Microsoft.Headings = YES -->
    ```golang
        client := servicenowsdkgo.NewServiceNowClient(cred, "xSDK_SN_URLx")
    ```

## 4. Review specific api documentation

With the `client` object initialized, implementation of the full capabilities offered by the Service-Now SDK for Go can begin. Refer to [apis](/apis/index.md) for details on implementing specific APIs.
