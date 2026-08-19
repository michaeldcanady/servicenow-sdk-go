# Batch API

The Batch API provides endpoints for executing multiple API requests in a single batch call.

## \[HEAD\] /now/v1/batch

Checks the availability of the batch endpoint.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	headers, err := client.Now().Batch().Head(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = headers
}
```

## \[POST\] /now/v1/batch

Executes a batch of API requests.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/batchapi"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	body := &batchapi.BatchRequest{}
	response, err := client.Now().Batch().Post(context.Background(), body, nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```
