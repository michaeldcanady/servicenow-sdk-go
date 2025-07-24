# Getting started

## Requirements

- Installed [golang 1.23+](https://go.dev/doc/install).
- Active Service-Now instance or [developer instance](https://developer.servicenow.com/dev.do).

## 1. Install Service-Now SDK for Go

<details>
    <summary>Install the latest</summary>
    <code lang="bash">
    go get github.com/michaeldcanady/servicenow-sdk-go
    </code>
</details>
<details>
    <summary>Install a specific version</summary>
    <code lang="bash">
    go get github.com/michaeldcanady/servicenow-sdk-go@{version}
    </code>
</details>

## 2 Create a credential

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

## 2 Create a Service-Now client

```golang
    ...
    client := servicenowsdkgo.NewServiceNowClient(cred, "{instance}")
    ...
```

## 3. Review specific API documentation

Once you have your `client` object, you're good to get started implementing all that the Service-Now SDK for Go has to offer! See [apis](/apis) for information on implementing a specific API!
