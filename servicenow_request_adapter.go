package servicenowsdkgo

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

// ServiceNowRequestAdapter is the core service used by ServiceNowServiceClient to make requests to ServiceNow APIs.
type ServiceNowRequestAdapter = internal.ServiceNowRequestAdapter

type serviceNowRequestAdapterOption = internal.ServiceNowRequestAdapterOption

var (
	NewServiceNowRequestAdapterBase = internal.NewServiceNowRequestAdapterBase
)
