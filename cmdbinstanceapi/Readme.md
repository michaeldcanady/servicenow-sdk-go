# CMDB Instance API

The CMDB Instance API provides endpoints for managing Configuration Management Database (CMDB) instances, classes, items, and relations.

## \[GET\] /now/v1/cmdb/instance/{className}

Retrieves multiple CI instances of the specified class.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Cmdb().Instance().ByClass("cmdb_ci_server").Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[POST\] /now/v1/cmdb/instance/{className}

Creates a new CI instance of the specified class.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/cmdbinstanceapi"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	body := &cmdbinstanceapi.CmdbInstance{}
	response, err := client.Now().Cmdb().Instance().ByClass("cmdb_ci_server").Post(context.Background(), body, nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[GET\] /now/v1/cmdb/instance/{className}/{sys_id}

Retrieves the specified CI instance.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Cmdb().Instance().
		ByClass("cmdb_ci_server").ByID("sys_id").Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[PUT\] /now/v1/cmdb/instance/{className}/{sys_id}

Replaces the specified CI instance.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/cmdbinstanceapi"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	body := &cmdbinstanceapi.CmdbInstance{}
	response, err := client.Now().Cmdb().Instance().
		ByClass("cmdb_ci_server").ByID("sys_id").Put(context.Background(), body, nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[PATCH\] /now/v1/cmdb/instance/{className}/{sys_id}

Partially updates the specified CI instance.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/cmdbinstanceapi"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	body := &cmdbinstanceapi.CmdbInstance{}
	response, err := client.Now().Cmdb().Instance().
		ByClass("cmdb_ci_server").ByID("sys_id").Patch(context.Background(), body, nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[POST\] /now/v1/cmdb/instance/{className}/{sys_id}/relation

Creates a relationship between CI instances.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/cmdbinstanceapi"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	body := &cmdbinstanceapi.CmdbInstance{}
	response, err := client.Now().Cmdb().Instance().
		ByClass("cmdb_ci_server").ByID("sys_id").Relation().Post(context.Background(), body, nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[DELETE\] /now/v1/cmdb/instance/{className}/{sys_id}/relation/{rel_sys_id}

Deletes a CI relationship.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	err := client.Now().Cmdb().Instance().
		ByClass("cmdb_ci_server").ByID("sys_id").Relation().
		ByID("rel_sys_id").Delete(context.Background(), nil)
	if err != nil {
		panic(err)
	}
}
```
