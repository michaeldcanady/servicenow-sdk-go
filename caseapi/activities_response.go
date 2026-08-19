package caseapi

import "github.com/michaeldcanady/servicenow-sdk-go/v2/core"

// ActivitiesResponse represents a single activities response.
type ActivitiesResponse = core.ServiceNowItemResponse[*ActivitiesResultModel]
