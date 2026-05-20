package caseapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// CaseRequestBuilderGetRequestConfiguration represents configuration for GET /case.
type CaseRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[CaseRequestBuilderGetQueryParameters]

// CaseRequestBuilderPostRequestConfiguration represents configuration for POST /case.
type CaseRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// CaseItemRequestBuilderGetRequestConfiguration represents configuration for GET /case/{id}.
type CaseItemRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// CaseItemRequestBuilderPutRequestConfiguration represents configuration for PUT /case/{id}.
type CaseItemRequestBuilderPutRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// CaseActivitiesRequestBuilderGetRequestConfiguration represents configuration for GET /case/{id}/activities.
type CaseActivitiesRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// CaseFieldValuesRequestBuilderGetRequestConfiguration represents configuration for GET /case/field_values/{field_name}.
type CaseFieldValuesRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[CaseFieldValuesRequestBuilderGetQueryParameters]
