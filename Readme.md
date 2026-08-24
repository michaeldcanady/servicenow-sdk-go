![ServiceNow SDK for Go](.github/servicenow-sdk-go_logo.png)

# Service-Now SDK for Go

A type-safe, idiomatic Go client for the ServiceNow REST APIs.

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/michaeldcanady/servicenow-sdk-go?style=plastic)
[![GoDoc](https://img.shields.io/static/v1?style=plastic&label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/michaeldcanady/servicenow-sdk-go)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/michaeldcanady/servicenow-sdk-go?style=plastic)
[![Maintainability](https://qlty.sh/badges/e778f295-dfb1-4637-a15e-f179549fcae4/maintainability.svg)](https://qlty.sh/gh/michaeldcanady/projects/servicenow-sdk-go)

## Install
```bash
go get github.com/michaeldcanady/servicenow-sdk-go/v2@latest
```
> Requires **Go 1.25+**

## Quickstart

```go
package main

import (
    "context"
    "fmt"
    "log"

    servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2"
    "github.com/michaeldcanady/servicenow-sdk-go/v2/v2/credentials"
)

func main() {
    authProvider := credentials.NewBasicProvider("username", "password")

    client, err := servicenowsdkgo.NewServiceNowServiceClient(
        servicenowsdkgo.WithInstance("your-instance"),
        servicenowsdkgo.WithAuthenticationProvider(authProvider),
    )
    if err != nil {
        log.Fatal(err)
    }

    resp, err := client.Now().Table("incident").Get(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    records, err := resp.GetResult()
    if err != nil {
        log.Fatal(err)
    }

    for _, record := range records {
        fmt.Println(record)
    }
}
```

## Contributing

Contributions are welcome: Start with the [Contributor Guide](https://michaeldcanady.github.io/servicenow-sdk-go/contributing/)

## License

[MIT](./LICENSE)
