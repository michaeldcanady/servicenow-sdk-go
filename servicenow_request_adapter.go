package servicenowsdkgo

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

// ServiceNowRequestAdapter is the core service used by ServiceNowServiceClient to make requests to ServiceNow APIs
type ServiceNowRequestAdapter = internal.ServiceNowRequestAdapter

// serviceNowRequestAdapterOption represents options for the ServiceNowServiceClient
type serviceNowRequestAdapterOption = internal.ServiceNowRequestAdapterOption

var (
	// NewServiceNowRequestAdapterBase creates a new ServiceNowServiceClient using the provided options
	NewServiceNowRequestAdapterBase = internal.NewServiceNowRequestAdapterBase
)
