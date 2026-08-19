# Policy API

The Policy API provides endpoints for managing CDM policies and policy mappings.

## \[POST\] /sn_cdm/v1/policies/mappings

Creates or updates policy mappings.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Cdm().Policies().Mappings().Post(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[DELETE\] /sn_cdm/v1/policies/mappings

Deletes policy mappings.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	err := client.Now().Cdm().Policies().Mappings().Delete(context.Background(), nil)
	if err != nil {
		panic(err)
	}
}
```
