# App Service API

The App Service API provides endpoints for managing ServiceNow Application Services, including CRUD operations, CSDM integration, and service tree actions.

## \[GET\] /now/cmdb/app_service

Retrieves multiple application services.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().AppService().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[GET\] /now/cmdb/app_service/{sys_id}

Retrieves the specified application service.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().AppService().ByID("sys_id").Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[GET\] /now/cmdb/app_service/{sys_id}/getContent

Retrieves the content tree for an application service.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().AppService().ByID("sys_id").GetContent().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[POST\] /now/cmdb/app_service/create

Creates a new application service.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/appserviceapi"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	body := &appserviceapi.CreateServiceRequest{}
	response, err := client.Now().AppService().Create().Post(context.Background(), body, nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```
