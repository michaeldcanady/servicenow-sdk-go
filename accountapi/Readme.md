# Account API

The Account API provides endpoints for managing accounts in ServiceNow.

## \[GET\] /now/v1/account

Retrieves multiple accounts.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Account().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[GET\] /now/v1/account/{account_id}

Retrieves the specified account.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Account().ByID("account_sys_id").Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```
