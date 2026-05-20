package appserviceapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// AppServiceRequestBuilderCreateOrUpdateServiceRequestConfiguration represents the configuration for a CreateOrUpdateService request.
type AppServiceRequestBuilderCreateOrUpdateServiceRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// GetContentRequestBuilderGetQueryParameters represents the query parameters for a GetContent request.
type GetContentRequestBuilderGetQueryParameters struct {
	// Mode determines the amount of relationship data returned.
	Mode *string `url:"Mode,omitempty"`
}

// GetContentRequestBuilderGetRequestConfiguration represents the configuration for a Get request on the getContent endpoint.
type GetContentRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[GetContentRequestBuilderGetQueryParameters]
