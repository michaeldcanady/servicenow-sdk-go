# Activity Subscriptions API

The Activity Subscriptions API provides endpoints for querying activity subscription data and facets.

## \[GET\] /now/v1/actsub/activities

Retrieves activity subscriptions with optional filtering by context, date range, and record.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().ActSub().Activities().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[GET\] /now/v1/actsub/facets/{activity_context}/{context_instance}

Retrieves activity subscription facets for a specific context and instance.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().ActSub().Facets().
		ByContext("incident").
		ByInstance("instance_id").
		Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```
