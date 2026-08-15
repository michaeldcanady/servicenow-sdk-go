package caseapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// CaseRequestBuilderGetRequestConfiguration represents configuration for GET /case.
type CaseRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[CaseRequestBuilderGetQueryParameters]

// CaseRequestBuilderPostRequestConfiguration represents configuration for POST /case.
type CaseRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// CaseItemRequestBuilderPutRequestConfiguration represents configuration for PUT /case/{id}.
type CaseItemRequestBuilderPutRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
