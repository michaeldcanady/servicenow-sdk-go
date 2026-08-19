# Case API

The Case API provides endpoints for managing Customer Service cases, including CRUD operations, activities, and field value lookups.

## \[GET\] /sn_customerservice/v1/case

Retrieves multiple cases.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().CustomerService().Case().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[GET\] /sn_customerservice/v1/case/{id}

Retrieves the specified case.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().CustomerService().Case().ByID("case_id").Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[POST\] /sn_customerservice/v1/case

Creates a new case.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/caseapi"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	body := &caseapi.CaseResult{}
	response, err := client.Now().CustomerService().Case().Post(context.Background(), body, nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[PUT\] /sn_customerservice/v1/case/{id}

Updates the specified case.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/caseapi"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	body := &caseapi.CaseResult{}
	response, err := client.Now().CustomerService().Case().ByID("case_id").Put(context.Background(), body, nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[GET\] /sn_customerservice/v1/case/{id}/activities

Retrieves activities for the specified case.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().CustomerService().Case().ByID("case_id").Activities().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```
