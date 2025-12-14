# Overview

<!-- vale Microsoft = NO -->

## What's Service-Now SDK for Go?

Service-Now SDK for Go is a thin golang sdk for the Service-Now REST apis. It leverages the functionality and flexibility of [Kiota modules](https://github.com/orgs/microsoft/repositories?q=kiota-*-go) (we don't use the Kiota CLI as Service-Now's OpenAPI spec, if it exists at all, isn't public).

## Why Service-Now SDK for Go?

Have to write your own SDK - or worse having to manage REST calls by hand - can be tedious and cumbersome. The goal of this project is to make that process easy and intuitive so you can get back to what you enjoy **working on your projects**!

## How to use

This SDK has two modalities of usage: `fluent` and `standard`.

The following block is the base you'll need for **all** implementation methods:
```golang
import (
    "github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

cred := credentials.NewUsernamePasswordCredential("username", "password")

client := servicenowsdkgo.NewServiceNowClient2(cred, "instance")
```

You can make API requests in two ways: through the fluent or standard implementation.
> We recommend the fluent implementation because it emphasizes ease of use and simplicity.

=== "Fluent"

    ``` golang {title="Table api"}
    client.Now2().Table2("table_name")
    ```

    ``` golang {title="Attachment api"}
    client.Now2().Attachment2()
    ```

    ``` golang {title="Batch api"}
    client.Now2().Batch()
    ```

=== "Standard"

    ``` golang {title="Table api"}
    pathParameters := map[string]string{
        "baseurl": "https://www.{instance}.service-now.com/api/now",
        "table":   "incident",
    }

    requestBuilder := tableapi.NewTableRequestBuilder2(client, pathParameters)
    ```

## Development status

Service-Now SDK for Go is being actively developed by the community.
